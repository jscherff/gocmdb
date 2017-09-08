// Copyright 2017 John Scherff
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package usbci

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"time"

	"github.com/google/gousb"
	"github.com/jscherff/gocmdb"
)

const (
	MagtekVendorID uint16 = 0x0801

	SureswipeKbPID uint16 = 0x0001
	SureswipeHidPID uint16 = 0x0002
	MagnesafeSwipeKbPID uint16 = 0x0001
	MagnesafeInsertKbPID uint16 = 0x0001
	MagnesafeSwipeHidPID uint16 = 0x0011
	MagnesafeInsertHidPID uint16 = 0x0013
	MagnesafeWirelessHidPID uint16 = 0x0014

	BufferSizeSureswipe int = 24
	BufferSizeMagnesafe int = 60

	CommandGetProp uint8 = 0x00
	CommandSetProp uint8 = 0x01
	CommandResetDevice uint8 = 0x02

	ResultCodeSuccess uint8 = 0x00
	ResultCodeFailure uint8 = 0x01
	ResultCodeBadParam uint8 = 0x02

	PropSoftwareID uint8 = 0x00
	PropDeviceSN uint8 = 0x01
	PropFactorySN uint8 = 0x03
	PropProductVer uint8 = 0x04

	DefaultSNLength int = 7
)

var (
	BufferSizes = []int{24, 60}
)

// Magtek decorates a gousb Device with Generic and Magtek Properties and API.
type Magtek struct {
	*Generic
}

// NewMagtek instantiates a Magtek wrapper for an existing gousb Device.
func NewMagtek(gd *gousb.Device) (*Magtek, error) {

	vm := make(map[string]string)

	if gd == nil {
		return &Magtek{&Generic{Device: &gousb.Device{}, Vendor: vm}}, nil
	}

	var err error

	this := &Magtek{&Generic{Device: gd, Vendor: vm}}
	errs := this.Init()

	if errs["BufferSize"] {this = nil}

	if len(errs) > 0 {

		emsg := "initialization failures:"

		for k, _ := range errs {
			emsg = fmt.Sprintf("%s %s", emsg, k)
		}

		err = gocmdb.ErrorDecorator(errors.New(emsg))
	}

	return this, err
}

// Init initializes API properties.
func (this *Magtek) Init() (errs map[string]bool) {

	errs = this.Generic.Init()

	var err error

	if this.BufferSize, err = this.GetBufferSize(); err != nil {
		errs["BufferSize"] = true
		return errs
	}
	if this.SoftwareID, err = this.GetSoftwareID(); err != nil {
		errs["SoftwareID"] = true
	}
	if this.ProductVer, err = this.GetProductVer(); err != nil {
		errs["ProductVer"] = true
	}
	if this.DeviceSN, err = this.GetDeviceSN(); err != nil {
		errs["DeviceSN"] = true
	}
	if this.FactorySN, err = this.GetFactorySN(); err != nil {
		errs["FactorySN"] = true
	}
	if this.DescriptorSN, err = this.SerialNumber(); err != nil {
		errs["DescriptorSN"] = true
	}

	this.SerialNum = this.DeviceSN
	this.ObjectType = this.Type()

	return errs
}

// Refresh updates API properties whose values may have changed.
func (this *Magtek) Refresh() (errs map[string]bool) {

	errs = this.Generic.Refresh()

	var err error

	if this.DeviceSN, err = this.GetDeviceSN(); err != nil {
		errs["DeviceSN"] = true
	}
	if this.FactorySN, err = this.GetFactorySN(); err != nil {
		errs["FactorySN"] = true
	}
	if this.DescriptorSN, err = this.SerialNumber(); err != nil {
		errs["DescriptorSN"] = true
	}

	this.SerialNum = this.DeviceSN

	return errs
}

// Type is a convenience method to help identify object type to other apps.
func (this *Magtek) Type() (string) {
	return reflect.TypeOf(this).String()
}

// DeviceSN retrieves the device configurable serial number from NVRAM.
func (this *Magtek) GetDeviceSN() (string, error) {
	return this.getProperty(PropDeviceSN)
}

// FactorySN retrieves the device factory serial number from NVRAM.
func (this *Magtek) GetFactorySN() (string, error) {
	val, err := this.getProperty(PropFactorySN)
	if len(val) <= 1 {val = ""}
	return val, err
}

// SoftwareID retrieves the software ID of the device from NVRAM.
func (this *Magtek) GetSoftwareID() (string, error) {
	return this.getProperty(PropSoftwareID)
}

// ProductVer retrieves the product version of the device from NVRAM.
func (this *Magtek) GetProductVer() (string, error) {
	val, err := this.getProperty(PropProductVer)
	if len(val) <= 1 {val = ""}
	return val, err
}

// SetDeviceSN sets the device configurable serial number in NVRAM.
func (this *Magtek) SetDeviceSN(val string) (error) {
	return this.setProperty(PropDeviceSN, val)
}

// EraseDeviceSN removes the device configurable serial number from NVRAM.
func (this *Magtek) EraseDeviceSN() (error) {
	return this.setProperty(PropDeviceSN, "")
}

// SetFactorySN sets the device factory device serial number in NVRAM.
// This will fail with result code 07 if serial number is already set.
func (this *Magtek) SetFactorySN(val string) (error) {
	return this.setProperty(PropFactorySN, val)
}

// CopyFactorySN copies 'length' characters from the device factory
// serial number to the configurable serial number in NVRAM.
func (this *Magtek) CopyFactorySN(n int) (error) {

	val, err := this.GetFactorySN()

	if err != nil {
		return gocmdb.ErrorDecorator(err)
	}

	if len(val) == 0 {
		return gocmdb.ErrorDecorator(fmt.Errorf("no factory serial number"))
	}

	n = int(math.Min(float64(n), float64(len(val))))
	err = this.SetDeviceSN(val[:n])

	return err
}

// Reset overides inherited Reset method with a low-level vendor reset.
func (this *Magtek) Reset() (error) {

	data := make([]byte, this.BufferSize)
	data[0] = CommandResetDevice

	_, err := this.Control(
		RequestDirectionOut + RequestTypeClass + RequestRecipientDevice,
		RequestSetReport,
		TypeFeatureReport,
		ControlInterface,
		data)

	if err != nil {
		err = gocmdb.ErrorDecorator(err)
	}

	data = make([]byte, this.BufferSize)

	_, err = this.Control(
		RequestDirectionIn + RequestTypeClass + RequestRecipientDevice,
		RequestGetReport,
		TypeFeatureReport,
		ControlInterface,
		data)

	if err != nil {
		return gocmdb.ErrorDecorator(err)
	}

	if data[0] > 0x00 {
		err = gocmdb.ErrorDecorator(fmt.Errorf("command error: %d", int(data[0])))
	}

	time.Sleep(5 * time.Second)

	return err
}

// GetBufferSize uses trial and error to find the control transfer data
// buffer size of the device. Failure to use the correct size for control
// transfers carrying vendor commands will result in a LIBUSB_ERROR_PIPE
// error.
func (this *Magtek) GetBufferSize() (size int, err error) {

	var rc int

	for _, size = range BufferSizes {

		data := make([]byte, size)
		copy(data, []byte{CommandGetProp, 0x01, PropSoftwareID})

		rc, err = this.Control(
			RequestDirectionOut + RequestTypeClass + RequestRecipientDevice,
			RequestSetReport,
			TypeFeatureReport,
			ControlInterface,
			data)

		if err != nil {continue}

		rc, err = this.Control(
			RequestDirectionIn + RequestTypeClass + RequestRecipientDevice,
			RequestGetReport,
			TypeFeatureReport,
			ControlInterface,
			data)

		if err != nil {continue}

		if rc == size {
			break
		}
	}

	if err != nil {
		err = gocmdb.ErrorDecorator(fmt.Errorf("unsupported device"))
	}

	return size, err
}

// getProperty retrieves a property from device NVRAM using low-level commands.
func (this *Magtek) getProperty(id uint8) (val string, err error) {

	data := make([]byte, this.BufferSize)
	copy(data, []byte{CommandGetProp, 0x01, id})

	_, err = this.Control(
		RequestDirectionOut + RequestTypeClass + RequestRecipientDevice,
		RequestSetReport,
		TypeFeatureReport,
		ControlInterface,
		data)

	if err != nil {
		return val, gocmdb.ErrorDecorator(err)
	}

	data = make([]byte, this.BufferSize)

	_, err = this.Control(
		RequestDirectionIn + RequestTypeClass + RequestRecipientDevice,
		RequestGetReport,
		TypeFeatureReport,
		ControlInterface,
		data)

	if err != nil {
		return val, gocmdb.ErrorDecorator(err)
	}

	if data[0] > 0x00 {
		return val, gocmdb.ErrorDecorator(fmt.Errorf("command error: %d", int(data[0])))
	}

	if data[1] > 0x00 {
		val= string(data[2:int(data[1])+2])
	}

	return val, err
}

// setProperty configures a property in device NVRAM using low-level commands.
func (this *Magtek) setProperty(id uint8, val string) (err error) {

	data := make([]byte, this.BufferSize)
	copy(data[0:], []byte{CommandSetProp, uint8(len(val)+1), id})
	copy(data[3:], val)

	_, err = this.Control(
		RequestDirectionOut + RequestTypeClass + RequestRecipientDevice,
		RequestSetReport,
		TypeFeatureReport,
		ControlInterface,
		data)

	if err != nil {
		return gocmdb.ErrorDecorator(err)
	}

	data = make([]byte, this.BufferSize)

	_, err = this.Control(
		RequestDirectionIn + RequestTypeClass + RequestRecipientDevice,
		RequestGetReport,
		TypeFeatureReport,
		ControlInterface,
		data)

	if err != nil {
		return gocmdb.ErrorDecorator(err)
	}

	if data[0] > 0x00 {
		err = gocmdb.ErrorDecorator(fmt.Errorf("command error: %d", int(data[0])))
	}

	if err == nil {
		this.Refresh()
	}

	return err
}
