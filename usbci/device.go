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
	"github.com/google/gousb"
	"github.com/jscherff/gocmdb"
	"strconv"
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

// NewDevice constructs a new Device.
func NewDevice(d *gousb.Device) (nd *Device, err error) {

	nd = &Device{d, 0, 0, 0}

	dd, err := NewDeviceDescriptor(nd)

	if err == nil {
		nd.ManufacturerIndex = int(dd.ManufacturerIndex)
		nd.ProductIndex = int(dd.ProductIndex)
		nd.SerialIndex = int(dd.SerialIndex)
	}

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

	if d.ManufacturerIndex > 0 {
		value, err = d.GetStringDescriptor(d.ManufacturerIndex)
	}

	if err != nil {
		err = fmt.Errorf("%s: %v", gocmdb.GetFunctionInfo(), err)
	}

	return value, err
}

// GetProductName retrieves the product name from device descriptor.
func (d *Device) GetProductName() (value string, err error) {

	if d.ProductIndex > 0 {
		value, err = d.GetStringDescriptor(d.ProductIndex)
	}

	if err != nil {
		err = fmt.Errorf("%s: %v", gocmdb.GetFunctionInfo(), err)
	}

	return value, err
}

// GetDescriptorSN retrieves the serial number of the device from the
// device descriptor. Changes made to the serial number on the device using a
// control transfer are not reflected in the device descriptor until the device
// is power-cycled or performs a device reset.
func (d *Device) GetDescriptSN() (value string, err error) {

	if d.SerialIndex > 0 {
		value, err = d.GetStringDescriptor(d.SerialIndex)
	}

	if err != nil {
		err = fmt.Errorf("%s: %v", gocmdb.GetFunctionInfo(), err)
	}

	return value, err
}
