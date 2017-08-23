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
	"github.com/jscherff/cmdb"
	"github.com/google/gousb"
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
	*gousb.Device
	ManufacturerIndex int	// Index of Manufacturer String Descriptor
	ProductIndex int	// Index of Product String Descriptor
	SerialIndex int		// Index of Serial Number String Descriptor
}

var (
	bufferSizes = []int {24, 60}
)

// NewDevice constructs a new Device.
func NewDevice(d *gousb.Device) (nd *Device, err error) {

	nd = &Device{d, 0, new(DeviceDescriptor), new(ConfigDescriptor)}

	err = nd.findBufferSize()

	if err != nil {
		return nd, err
	}

	_ = nd.getDeviceDescriptor()
	_ = nd.getConfigDescriptor()

	return nd, err
}

// GetBusNumber retrieves the USB bus number of the device.
func (d *Device) GetBusNumber() string {
	return strconv.Itoa(d.Desc.Bus)
}

// GetBusAddress retrieves address of the device on the USB bus.
func (d *Device) GetBusAddress() string {
	return strconv.Itoa(d.Desc.Address)
}

// GetDeviceSpeed retrieves the negotiated operating speed of the device.
func (d *Device) GetDeviceSpeed() string {
	return d.Desc.Speed.String()
}

// GetUSBSpec retrieves the USB specification release number of the device.
func (d *Device) GetUSBSpec() string {
	return d.Desc.Spec.String()
}

// GetDeviceVer retrieves the major/minor version number ofthe device.
func (d *Device) GetDeviceVer() string {
	return d.Desc.Device.String()
}

// GetVendorId retrieves the USB vendor ID of the device.
func (d *Device) GetVendorID() string {
	return d.Desc.Vendor.String()
}

// GetProductID retrieves the USB product ID of the device.
func (d *Device) GetProductID() string {
	return d.Desc.Product.String()
}

// GetUSBClass retrieves the USB class of the device.
func (d *Device) GetUSBClass() string {
	return d.Desc.Class.String()
}

// GetUSBSubclass retrieves the USB subclass of the device.
func (d *Device) GetUSBSubclass() string {
	return d.Desc.SubClass.String()
}

// GetUSBProtocol retrieves the USB protocol of the device.
func (d *Device) GetUSBProtocol() string {
	return d.Desc.Protocol.String()
}

// GetMaxPktSize retrieves the maximum size of the control transfer.
func (d *Device) GetMaxPktSize() string {
	return strconv.Itoa(d.Desc.MaxControlPacketSize)
}

// GetVendorName retrieves the manufacturer name from device descriptor.
func (d *Device) GetVendorName() (value string, err error) {

	if d.DeviceDescriptor.ManufacturerIndex > 0 {
		value, err = d.GetStringDescriptor(int(d.DeviceDescriptor.ManufacturerIndex))
	}

	if err != nil {
		err = fmt.Errorf("%s: %v", getFunctionInfo(), err)
	}

	return value, err
}

// GetProductName retrieves the product name from device descriptor.
func (d *Device) GetProductName() (value string, err error) {

	if d.DeviceDescriptor.ProductIndex > 0 {
		value, err = d.GetStringDescriptor(int(d.DeviceDescriptor.ProductIndex))
	}

	if err != nil {
		err = fmt.Errorf("%s: %v", getFunctionInfo(), err)
	}

	return value, err
}

// GetDescriptorSN retrieves the serial number of the device from the
// device descriptor. Changes made to the serial number on the device using a
// control transfer are not reflected in the device descriptor until the device
// is power-cycled or performs a device reset.
func (d *Device) GetDescriptSN() (value string, err error) {

	if d.DeviceDescriptor.SerialNumIndex > 0 {
		value, err = d.GetStringDescriptor(int(d.DeviceDescriptor.SerialNumIndex))
	}

	if err != nil {
		err = fmt.Errorf("%s: %v", getFunctionInfo(), err)
	}

	return value, err
}

// UsbReset performs a USB port reset to reinitialize the device.
func (d *Device) UsbReset() (err error) {
	return d.Reset()
}

// getDeviceDescriptor retrieves the raw device descriptor.
func (d *Device) getDeviceDescriptor() (err error) {

	data := make([]byte, BufferSizeDeviceDescriptor)

	_, err = d.Control(
		RequestDirectionIn + RequestTypeStandard + RequestRecipientDevice,
		RequestGetDescriptor,
		TypeDeviceDescriptor,
		ControlInterface,
		data)

	if err == nil {

		*d.DeviceDescriptor = DeviceDescriptor {
			data[0],
			data[1],
			uint16(data[2]) + (uint16(data[3])<<8),
			data[4],
			data[5],
			data[6],
			data[7],
			uint16(data[8]) + (uint16(data[9])<<8),
			uint16(data[10]) + (uint16(data[11])<<8),
			uint16(data[12]) + (uint16(data[13])<<8),
			data[14],
			data[15],
			data[16],
			data[17]}
	} else {
		err = fmt.Errorf("%s: %v", getFunctionInfo(), err)
	}

	return err
}

// getConfigDescriptor retrieves the raw active config descriptor.
func (d *Device) getConfigDescriptor() (err error) {

	data := make([]byte, BufferSizeConfigDescriptor)

	_, err = d.Control(
		RequestDirectionIn + RequestTypeStandard + RequestRecipientDevice,
		RequestGetDescriptor,
		TypeConfigDescriptor,
		ControlInterface,
		data)

	if err == nil {

		*d.ConfigDescriptor = ConfigDescriptor {
			data[0],
			data[1],
			uint16(data[2]) + (uint16(data[3]) << 8),
			data[4],
			data[5],
			data[6],
			data[7],
			data[8]}
	} else {
		return fmt.Errorf("%s: %v", getFunctionInfo(), err)
	}

	return err
}
