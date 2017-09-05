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
	"github.com/google/gousb"
	"github.com/jscherff/gocmdb"
	"github.com/jscherff/gocmdb/usbci"
	"reflect"
	"strconv"
	"math"
	"time"
	"fmt"
)

// Device represents a USB device. The Device struct Desc field contains all
// information about the device. It includes the raw device descriptor, the
// config descriptor of the active config, and the size of the data buffer
// required by the device for vendor commands sent via control transfer.
type Device struct {
	*usbci.Device
	DeviceSN string
	FactorySN string
	SoftwareID string
	ProductVer string
	BufferSize int
}

// NewDevice constructs a new Device.
func NewDevice(gd *gousb.Device) (this *Device, err error) {

	ud, err := usbci.NewDevice(gd)

	if err != nil {
		return ud, err
	}

	this = &Device{ud, 0}

	errs := []string

	if this.DeviceSN, e = this.DeviceSN(); e != nil {errs = append(errs, "DeviceSN")}
	if this.FactorySN, e = this.FactorySN(); e != nil {errs = append(errs, "FactorySN")}
	if this.SoftwareID, e = this.SoftwareID(); e != nil {errs = append(errs, "SoftwareID")}
	if this.ProductVer, e = this.ProductVer(); e != nil {errs = append(errs, "ProductVer")}
	if this.BufferSize, e = this.findBufferSize(); e != nil {errs = append(errs, "BufferSize")}

	this.Info.SerialNumber = this.DeviceSN
	this.Info.Vendor0 = this.FactorySN
	this.Info.Vendor1 = this.SoftwareID
	this.Info.Vendor2 = this.ProductVer

	if len(errs) > 0 {
		err = errors.New("getter errors: " + strings.Join(errs, ","))
	}

	return this, err
}

// Convenience method to retrieve device serial number.
func (this *Device) ID() (string, error) {
	return this.DeviceSN()
}

// Convenience method to help identify object type to other apps.
func (this *Device) Type() (string) {
	return reflect.TypeOf(this).String()
}

// DeviceSN retrieves the device configurable serial number from NVRAM.
func (this *Device) DeviceSN() (string, error) {
	return this.getProperty(PropDeviceSN)
}

// FactorySN retrieves the device factory serial number from NVRAM.
func (this *Device) FactorySN() (string, error) {
	val, err = this.getProperty(PropFactorySN)
	if len(val) <= 1 {val = ""}
	return val, err
}

// SoftwareID retrieves the software ID of the device from NVRAM.
func (this *Device) SoftwareID() (string, error) {
	return this.getProperty(PropSoftwareID)
}

// ProductVer retrieves the product version of the device from NVRAM.
func (this *Device) ProductVer() (string, error) {
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

	val, err := this.FactorySN()

	if err != nil {
		return fmt.Errorf("%s: %v", gocmdb.FunctionInfo(), err)
	}

	if len(val) == 0 {
		return fmt.Errorf("%s: no factory serial number", gocmdb.FunctionInfo())
	}

	n = int(math.Min(float64(n), float64(len(val))))
	err = this.SetDeviceSN(s[:n])

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
		err = fmt.Errorf("%s: %v)", gocmdb.FunctionInfo(), err)
	}

	data = make([]byte, this.BufferSize)

	_, err = this.Control(
		usbci.RequestDirectionIn + usbci.RequestTypeClass + usbci.RequestRecipientDevice,
		usbci.RequestGetReport,
		usbci.TypeFeatureReport,
		usbci.ControlInterface,
		data)

	if err != nil {
		return fmt.Errorf("%s: %v", gocmdb.FunctionInfo(), err)
	}

	if data[0] > 0x00 {
		err = fmt.Errorf("%s: command error: %d", gocmdb.FunctionInfo(), int(data[0]))
	}

	time.Sleep(5 * time.Second)

	return err
}

// findBufferSize uses trial and error to find the control transfer data
// buffer size of the device. Failure to use the correct size for control
// transfers carrying vendor commands will result in a LIBUSB_ERROR_PIPE
// error.
func (this *Device) findBufferSize() (error) {

	var rc, size int
	var err error

	for _, size = range BufferSizes {

		data := make([]byte, size)
		copy(data, []byte{CommandGetProp, 0x01, PropSoftwareID})

		rc, err = this.Control(
			usbci.RequestDirectionOut + usbci.RequestTypeClass + usbci.RequestRecipientDevice,
			usbci.RequestSetReport,
			usbci.TypeFeatureReport,
			usbci.ControlInterface,
			data)

		data = make([]byte, size)

		rc, err = this.Control(
			usbci.RequestDirectionIn + usbci.RequestTypeClass + usbci.RequestRecipientDevice,
			usbci.RequestGetReport,
			usbci.TypeFeatureReport,
			usbci.ControlInterface,
			data)

		if rc == size {
			this.BufferSize = size
			break
		}
	}

	if err != nil {
		err = fmt.Errorf("%s: unsupported device", gocmdb.FunctionInfo())
	}

	return err
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
		return val, fmt.Errorf("%s: %v", gocmdb.FunctionInfo(), err)
	}

	data = make([]byte, this.BufferSize)

	_, err = this.Control(
		usbci.RequestDirectionIn + usbci.RequestTypeClass + usbci.RequestRecipientDevice,
		usbci.RequestGetReport,
		usbci.TypeFeatureReport,
		usbci.ControlInterface,
		data)

	if err != nil {
		return val, fmt.Errorf("%s: %v", gocmdb.FunctionInfo(), err)
	}

	if data[0] > 0x00 {
		return val, fmt.Errorf("%s: command error: %d", gocmdb.FunctionInfo(), int(data[0]))
	}

	if data[1] > 0x00 {
		val= string(data[2:int(data[1])+2])
	}

	return val, err
}

// setProperty configures a property in device NVRAM using low-level commands.
func (this *Device) setProperty(id uint8, val string) (err error) {

	if len(val) > this.BufferSize - 3 {
		return fmt.Errorf("%s: property length > data buffer", gocmdb.FunctionInfo())
	}

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
		return fmt.Errorf("%s: %v", gocmdb.FunctionInfo(), err)
	}

	data = make([]byte, this.BufferSize)

	_, err = this.Control(
		usbci.RequestDirectionIn + usbci.RequestTypeClass + usbci.RequestRecipientDevice,
		usbci.RequestGetReport,
		usbci.TypeFeatureReport,
		usbci.ControlInterface,
		data)

	if err != nil {
		return fmt.Errorf("%s: %v", gocmdb.FunctionInfo(), err)
	}

	if data[0] > 0x00 {
		err = fmt.Errorf("%s: command error: %d", gocmdb.FunctionInfo(), int(data[0]))
	}

	return err
}

