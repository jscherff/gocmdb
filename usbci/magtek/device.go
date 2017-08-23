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
	BufferSize int
}

var (
	bufferSizes = []int {24, 60}
)

// NewDevice constructs a new Device.
func NewDevice(d *gousb.Device) (nd *Device, err error) {

	pd, err := usbci.NewDevice(d)
	nd = &Device{pd, 0}
	err = nd.findBufferSize()

	return nd, err
}

// GetBufferSize retrieves the size of the device data buffer.
func (d *Device) GetBufferSize() (string, error) {
	return strconv.Itoa(d.BufferSize), nil
}

// GetSoftwareID retrieves the software ID of the device from NVRAM.
func (d *Device) GetSoftwareID() (string, error) {
	return d.getProperty(PropSoftwareID)
}

// GetProductVer retrieves the product version of the device from NVRAM.
func (d *Device) GetProductVer() (value string, err error) {
	value, err = d.getProperty(PropProductVer)
	if len(value) <= 1 {value = ""}
	return value, err
}

// GetDeviceSN retrieves the device configurable serial number from NVRAM.
func (d *Device) GetDeviceSN() (string, error) {
	return d.getProperty(PropDeviceSN)
}

// SetDeviceSN sets the device configurable serial number in NVRAM.
func (d *Device) SetDeviceSN(value string) (error) {
	return d.setProperty(PropDeviceSN, value)
}

// EraseDeviceSN removes the device configurable serial number from NVRAM.
func (d *Device) EraseDeviceSN() (error) {
	return d.setProperty(PropDeviceSN, "")
}

// GetFactorySN retrieves the device factory serial number from NVRAM.
func (d *Device) GetFactorySN() (value string, err error) {
	value, err = d.getProperty(PropFactorySN)
	if len(value) <= 1 {value = ""}
	return value, err
}

// SetFactorySN sets the device factory device serial number in NVRAM.
// This command will fail with result code 07 if the serial number is
// already configured.
func (d *Device) SetFactorySN(value string) (error) {
	return d.setProperty(PropFactorySN, value)
}

// CopyFactorySN copies 'length' characters from the device factory
// serial number to the configurable serial number in NVRAM.
func (d *Device) CopyFactorySN(length int) (error) {

	fs, err := d.GetFactorySN()

	if err != nil {
		return fmt.Errorf("%s: %v", gocmdb.GetFunctionInfo(), err)
	}

	if len(fs) == 0 {
		return fmt.Errorf("%s: factory serial number not present", gocmdb.GetFunctionInfo())
	}

	limit := int(math.Min(float64(length), float64(len(fs))))
	err = d.SetDeviceSN(fs[:limit])

	return err
}

// DeviceReset resets the device using low-level vendor commands.
func (d *Device) DeviceReset() (err error) {

	data := make([]byte, d.BufferSize)
	data[0] = CommandResetDevice

	_, err = d.Control(
		usbci.RequestDirectionOut + usbci.RequestTypeClass + usbci.RequestRecipientDevice,
		usbci.RequestSetReport,
		usbci.TypeFeatureReport,
		usbci.ControlInterface,
		data)

	if err != nil {
		err = fmt.Errorf("%s: %v)", gocmdb.GetFunctionInfo(), err)
	}

	data = make([]byte, d.BufferSize)

	_, err = d.Control(
		usbci.RequestDirectionIn + usbci.RequestTypeClass + usbci.RequestRecipientDevice,
		usbci.RequestGetReport,
		usbci.TypeFeatureReport,
		usbci.ControlInterface,
		data)

	if err != nil {
		return fmt.Errorf("%s: %v", gocmdb.GetFunctionInfo(), err)
	}

	if data[0] > 0x00 {
		err = fmt.Errorf("%s: command error: %d",
			gocmdb.GetFunctionInfo(), int(data[0]))
	}

	time.Sleep(5 * time.Second)

	return err
}

// findBufferSize uses trial and error to find the control transfer data
// buffer size of the device. Failure to use the correct size for control
// transfers carrying vendor commands will result in a LIBUSB_ERROR_PIPE
// error.
func (d *Device) findBufferSize() (err error) {

	var rc, size int

	for _, size = range bufferSizes {

		data := make([]byte, size)
		copy(data, []byte{CommandGetProp, 0x01, PropSoftwareID})

		rc, err = d.Control(
			usbci.RequestDirectionOut + usbci.RequestTypeClass + usbci.RequestRecipientDevice,
			usbci.RequestSetReport,
			usbci.TypeFeatureReport,
			usbci.ControlInterface,
			data)

		data = make([]byte, size)

		rc, err = d.Control(
			usbci.RequestDirectionIn + usbci.RequestTypeClass + usbci.RequestRecipientDevice,
			usbci.RequestGetReport,
			usbci.TypeFeatureReport,
			usbci.ControlInterface,
			data)

		if rc == size {
			d.BufferSize = size
			break
		}
	}

	if err != nil {
		err = fmt.Errorf("%s: unsupported device", gocmdb.GetFunctionInfo())
	}

	return err
}

// getProperty retrieves a property from device NVRAM using low-level commands.
func (d *Device) getProperty(id uint8) (value string, err error) {

	data := make([]byte, d.BufferSize)
	copy(data, []byte{CommandGetProp, 0x01, id})

	_, err = d.Control(
		usbci.RequestDirectionOut + usbci.RequestTypeClass + usbci.RequestRecipientDevice,
		usbci.RequestSetReport,
		usbci.TypeFeatureReport,
		usbci.ControlInterface,
		data)

	if err != nil {
		return value, fmt.Errorf("%s: %v", gocmdb.GetFunctionInfo(), err)
	}

	data = make([]byte, d.BufferSize)

	_, err = d.Control(
		usbci.RequestDirectionIn + usbci.RequestTypeClass + usbci.RequestRecipientDevice,
		usbci.RequestGetReport,
		usbci.TypeFeatureReport,
		usbci.ControlInterface,
		data)

	if err != nil {
		return value, fmt.Errorf("%s: %v", gocmdb.GetFunctionInfo(), err)
	}

	if data[0] > 0x00 {
		return value, fmt.Errorf("%s: command error: %d",
			gocmdb.GetFunctionInfo(), int(data[0]))
	}

	if data[1] > 0x00 {
		value = string(data[2:int(data[1])+2])
	}

	return value, err
}

// setProperty configures a property in device NVRAM using low-level commands.
func (d *Device) setProperty(id uint8, value string) (err error) {

	if len(value) > d.BufferSize - 3 {
		return fmt.Errorf("%s: property length > data buffer", gocmdb.GetFunctionInfo())
	}

	data := make([]byte, d.BufferSize)
	copy(data[0:], []byte{CommandSetProp, uint8(len(value)+1), id})
	copy(data[3:], value)

	_, err = d.Control(
		usbci.RequestDirectionOut + usbci.RequestTypeClass + usbci.RequestRecipientDevice,
		usbci.RequestSetReport,
		usbci.TypeFeatureReport,
		usbci.ControlInterface,
		data)

	if err != nil {
		return fmt.Errorf("%s: %v", gocmdb.GetFunctionInfo(), err)
	}

	data = make([]byte, d.BufferSize)

	_, err = d.Control(
		usbci.RequestDirectionIn + usbci.RequestTypeClass + usbci.RequestRecipientDevice,
		usbci.RequestGetReport,
		usbci.TypeFeatureReport,
		usbci.ControlInterface,
		data)

	if err != nil {
		return fmt.Errorf("%s: %v", gocmdb.GetFunctionInfo(), err)
	}

	if data[0] > 0x00 {
		err = fmt.Errorf("%s: command error: %d",
			gocmdb.GetFunctionInfo(), int(data[0]))
	}

	return err
}
