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
	"reflect"
	"strconv"
	"fmt"
)

// Device represents a USB device. The Device struct Desc field contains all
// information about the device. It includes the raw device descriptor, the
// config descriptor of the active config, and the size of the data buffer
// required by the device for vendor commands sent via control transfer.
type Device struct {
	*gousb.Device
	manufacturerIx int	// Index of Manufacturer String Descriptor
	productIx int		// Index of Product String Descriptor
	serialIx int		// Index of Serial Number String Descriptor
}

// NewDevice constructs a new Device.
func NewDevice(d *gousb.Device) (*Device, error) {

	nd := &Device{d, 0, 0, 0}

	dd, err := NewDeviceDescriptor(nd)

	if err == nil {
		nd.manufacturerIx = int(dd.ManufacturerIndex)
		nd.productIx = int(dd.ProductIndex)
		nd.serialIx = int(dd.SerialIndex)
	}

	return nd, err
}

// Convenience method to retrieve device serial number.
func (this *Device) ID() (string, error) {
	return this.DescriptorSN()
}

// Convenience method to help identify object type to other apps.
func (this *Device) Type() (string) {
	return reflect.TypeOf(this).String()
}

// BusNumber retrieves the USB bus number of the device.
func (this *Device) BusNumber() string {
	return strconv.Itoa(this.Desc.Bus)
}

// BusAddress retrieves address of the device on the USB bus.
func (this *Device) BusAddress() string {
	return strconv.Itoa(this.Desc.Address)
}

// DeviceSpeed retrieves the negotiated operating speed of the device.
func (this *Device) DeviceSpeed() string {
	return this.Desc.Speed.String()
}

// USBSpec retrieves the USB specification release number of the device.
func (this *Device) USBSpec() string {
	return this.Desc.Spec.String()
}

// DeviceVer retrieves the major/minor version number ofthe device.
func (this *Device) DeviceVer() string {
	return this.Desc.Device.String()
}

// GetVendorId retrieves the USB vendor ID of the device.
func (this *Device) VendorID() string {
	return this.Desc.Vendor.String()
}

// ProductID retrieves the USB product ID of the device.
func (this *Device) ProductID() string {
	return this.Desc.Product.String()
}

// USBClass retrieves the USB class of the device.
func (this *Device) USBClass() string {
	return this.Desc.Class.String()
}

// USBSubclass retrieves the USB subclass of the device.
func (this *Device) USBSubclass() string {
	return this.Desc.SubClass.String()
}

// USBProtocol retrieves the USB protocol of the device.
func (this *Device) USBProtocol() string {
	return this.Desc.Protocol.String()
}

// MaxPktSize retrieves the maximum size of the control transfer.
func (this *Device) MaxPktSize() string {
	return strconv.Itoa(this.Desc.MaxControlPacketSize)
}

// VendorName retrieves the manufacturer name from device descriptor.
func (this *Device) VendorName() (value string, err error) {

	if this.manufacturerIx > 0 {
		value, err = this.GetStringDescriptor(this.manufacturerIx)
	}

	if err != nil {
		err = fmt.Errorf("%s: %v", gocmdb.FunctionInfo(), err)
	}

	return value, err
}

// ProductName retrieves the product name from device descriptor.
func (this *Device) ProductName() (value string, err error) {

	if this.productIx > 0 {
		value, err = this.GetStringDescriptor(this.productIx)
	}

	if err != nil {
		err = fmt.Errorf("%s: %v", gocmdb.FunctionInfo(), err)
	}

	return value, err
}

// GetDescriptorSN retrieves the serial number of the device from the
// device descriptor. Changes made to the serial number on the device using a
// control transfer are not reflected in the device descriptor until the device
// is power-cycled or performs a device reset.
func (this *Device) DescriptorSN() (value string, err error) {

	if this.serialIx > 0 {
		value, err = this.GetStringDescriptor(this.serialIx)
	}

	if err != nil {
		err = fmt.Errorf("%s: %v", gocmdb.FunctionInfo(), err)
	}

	return value, err
}
