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

package magtek

import (
	"reflect"
	"strconv"
	"strings"
	"math"
	"time"
	"fmt"

	"github.com/google/gousb"
	"github.com/jscherff/gocmdb"
	"github.com/jscherff/gocmdb/usbci"
)

// Device represents a USB device. The Device struct Desc field contains all
// information about the device. It includes the raw device descriptor, the
// config descriptor of the active config, and the size of the data buffer
// required by the device for vendor commands sent via control transfer.
type Device struct {
	*usbci.Device
	BufferSize int		`json:"bufer_size" csv:"-" nvp:"-"`
	DeviceSN string		`json:"device_sn" csv:"-" nvp:"-"`
	FactorySN string	`json:"factory_sn" csv:"-" nvp:"-"`
	DescriptorSN string	`json:"descriptor_sn" csv:"-" nvp:"-"`
}

var BufferSizes = []int{24, 60}

// NewDevice converts an existing gousb device to a magtek Device.
func NewDevice(gd *gousb.Device) (*Device, error) {

	var err error
	var errs []string

	ud, err := usbci.NewDevice(gd)

	if err != nil {
		return nil, gocmdb.ErrorDecorator(err)
	}

	this := &Device{ud, 0, "", "", ""}

	if this.BufferSize, err = this.GetBufferSize(); err != nil {
		return nil, gocmdb.ErrorDecorator(err)
	}
	if this.SoftwareID, err = this.GetSoftwareID(); err != nil {
		errs = append(errs, "SoftwareID")
	}
	if this.ProductVer, err = this.GetProductVer(); err != nil {
		errs = append(errs, "ProductVer")
	}
	if this.DeviceSN, err = this.GetDeviceSN(); err != nil {
		errs = append(errs, "DeviceSN")
	}
	if this.FactorySN, err = this.GetFactorySN(); err != nil {
		errs = append(errs, "FactorySN")
	}
	if this.DescriptorSN, err = this.SerialNumber(); err != nil {
		errs = append(errs, "DescriptorSN")
	}

	this.SerialNum = this.DeviceSN
	this.ObjectType = this.Type()

	this.Vendor0 = "BufferSize:" + strconv.Itoa(this.BufferSize)
	this.Vendor1 = "DeviceSN:" + this.DeviceSN
	this.Vendor2 = "FactorySN:" + this.FactorySN
	this.Vendor3 = "DescriptorSN:" + this.DescriptorSN

	if len(errs) > 0 {
		err = fmt.Errorf("getter errors: ", strings.Join(errs, ","))
		err = gocmdb.ErrorDecorator(err)
	}

	return this, err
}

// Refresh items whose underlying values may have chanegd.
func (this *Device) Refresh() (error) {

	var err error
	var errs []string

	if this.DeviceSN, err = this.GetDeviceSN(); err != nil {
		errs = append(errs, "DeviceSN")
	}
	if this.FactorySN, err = this.GetFactorySN(); err != nil {
		errs = append(errs, "FactorySN")
	}
	if this.DescriptorSN, err = this.SerialNumber(); err != nil {
		errs = append(errs, "DescriptorSN")
	}

	this.SerialNum = this.DeviceSN

	if len(errs) > 0 {
		err = fmt.Errorf("getter errors: ", strings.Join(errs, ","))
		err = gocmdb.ErrorDecorator(err)
	}

	return err
}

// Convenience method to retrieve device serial number.
func (this *Device) ID() (string) {
	return this.SerialNum
}

// Convenience method to help identify object type to other apps.
func (this *Device) Type() (string) {
	return reflect.TypeOf(this).String()
}

// DeviceSN retrieves the device configurable serial number from NVRAM.
func (this *Device) GetDeviceSN() (string, error) {
	return this.getProperty(PropDeviceSN)
}

// FactorySN retrieves the device factory serial number from NVRAM.
func (this *Device) GetFactorySN() (string, error) {
	val, err := this.getProperty(PropFactorySN)
	if len(val) <= 1 {val = ""}
	return val, err
}

// SoftwareID retrieves the software ID of the device from NVRAM.
func (this *Device) GetSoftwareID() (string, error) {
	return this.getProperty(PropSoftwareID)
}

// ProductVer retrieves the product version of the device from NVRAM.
func (this *Device) GetProductVer() (string, error) {
	val, err := this.getProperty(PropProductVer)
	if len(val) <= 1 {val = ""}
	return val, err
}

// SetDeviceSN sets the device configurable serial number in NVRAM.
func (this *Device) SetDeviceSN(val string) (error) {
	return this.setProperty(PropDeviceSN, val)
}

// EraseDeviceSN removes the device configurable serial number from NVRAM.
func (this *Device) EraseDeviceSN() (error) {
	return this.setProperty(PropDeviceSN, "")
}

// SetFactorySN sets the device factory device serial number in NVRAM.
// This will fail with result code 07 if serial number is already set.
func (this *Device) SetFactorySN(val string) (error) {
	return this.setProperty(PropFactorySN, val)
}

// CopyFactorySN copies 'length' characters from the device factory
// serial number to the configurable serial number in NVRAM.
func (this *Device) CopyFactorySN(n int) (error) {

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
func (this *Device) Reset() (error) {

	data := make([]byte, this.BufferSize)
	data[0] = CommandResetDevice

	_, err := this.Control(
		usbci.RequestDirectionOut + usbci.RequestTypeClass + usbci.RequestRecipientDevice,
		usbci.RequestSetReport,
		usbci.TypeFeatureReport,
		usbci.ControlInterface,
		data)

	if err != nil {
		err = gocmdb.ErrorDecorator(err)
	}

	data = make([]byte, this.BufferSize)

	_, err = this.Control(
		usbci.RequestDirectionIn + usbci.RequestTypeClass + usbci.RequestRecipientDevice,
		usbci.RequestGetReport,
		usbci.TypeFeatureReport,
		usbci.ControlInterface,
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
func (this *Device) GetBufferSize() (size int, err error) {

	var rc int

	for _, size = range BufferSizes {

		data := make([]byte, size)
		copy(data, []byte{CommandGetProp, 0x01, PropSoftwareID})

		rc, err = this.Control(
			usbci.RequestDirectionOut + usbci.RequestTypeClass + usbci.RequestRecipientDevice,
			usbci.RequestSetReport,
			usbci.TypeFeatureReport,
			usbci.ControlInterface,
			data)

		if err != nil {continue}

		rc, err = this.Control(
			usbci.RequestDirectionIn + usbci.RequestTypeClass + usbci.RequestRecipientDevice,
			usbci.RequestGetReport,
			usbci.TypeFeatureReport,
			usbci.ControlInterface,
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
func (this *Device) getProperty(id uint8) (val string, err error) {

	data := make([]byte, this.BufferSize)
	copy(data, []byte{CommandGetProp, 0x01, id})

	_, err = this.Control(
		usbci.RequestDirectionOut + usbci.RequestTypeClass + usbci.RequestRecipientDevice,
		usbci.RequestSetReport,
		usbci.TypeFeatureReport,
		usbci.ControlInterface,
		data)

	if err != nil {
		return val, gocmdb.ErrorDecorator(err)
	}

	data = make([]byte, this.BufferSize)

	_, err = this.Control(
		usbci.RequestDirectionIn + usbci.RequestTypeClass + usbci.RequestRecipientDevice,
		usbci.RequestGetReport,
		usbci.TypeFeatureReport,
		usbci.ControlInterface,
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
func (this *Device) setProperty(id uint8, val string) (err error) {

	data := make([]byte, this.BufferSize)
	copy(data[0:], []byte{CommandSetProp, uint8(len(val)+1), id})
	copy(data[3:], val)

	_, err = this.Control(
		usbci.RequestDirectionOut + usbci.RequestTypeClass + usbci.RequestRecipientDevice,
		usbci.RequestSetReport,
		usbci.TypeFeatureReport,
		usbci.ControlInterface,
		data)

	if err != nil {
		return gocmdb.ErrorDecorator(err)
	}

	data = make([]byte, this.BufferSize)

	_, err = this.Control(
		usbci.RequestDirectionIn + usbci.RequestTypeClass + usbci.RequestRecipientDevice,
		usbci.RequestGetReport,
		usbci.TypeFeatureReport,
		usbci.ControlInterface,
		data)

	if err != nil {
		return gocmdb.ErrorDecorator(err)
	}

	if data[0] > 0x00 {
		err = gocmdb.ErrorDecorator(fmt.Errorf("command error: %d", int(data[0])))
	}

	if err == nil {
		err = this.Refresh()
	}

	return err
}
