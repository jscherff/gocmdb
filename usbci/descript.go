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
	"github.com/jscherff/gocmdb"
	"fmt"
)

// DeviceDescriptor specifies some basic information about the USB device,
// such as the supported USB version, maximum packet size, vendor and product
// IDs.
type DeviceDescriptor struct {
	Length uint8			// Size of the Descriptor in Bytes
	DescriptorType uint8		// Device Descriptor Type (0x01)
	UsbSpecification uint16		// BCD of Device USB Specification Number
	DeviceClass uint8		// Device Class Code or Vendor Specified
	DeviceSubClass uint8		// Subclass Code Assigned by USB Org
	DeviceProtocol uint8		// Protocol Code Assigned by USB Org
	MaxPktSize uint8		// Maximum Packet Size for Zero Endpoint
	VendorID uint16			// Vendor ID Assigned by USB Org
	ProductID uint16		// Product ID (Assigned by Manufacturer)
	DeviceReleaseNumber uint16	// BCD of Device Release Number
	ManufacturerIndex uint8		// Index Manufacturer String Descriptor
	ProductIndex uint8		// Index of Product String Descriptor
	SerialIndex uint8		// Index of Serial Number String Descriptor
	NumConfigurations uint8		// Number of Possible Configurations
}

// NewDeviceDescriptor constructs a new DeviceDescriptor.
func NewDeviceDescriptor(d *Device) (dd *DeviceDescriptor, err error) {

	dd = new(DeviceDescriptor)
	data := make([]byte, BufferSizeDeviceDescriptor)

	_, err = d.Control(
		RequestDirectionIn + RequestTypeStandard + RequestRecipientDevice,
		RequestGetDescriptor,
		TypeDeviceDescriptor,
		ControlInterface,
		data)

	if err == nil {

		*dd = DeviceDescriptor {
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
		err = fmt.Errorf("%s: %v", gocmdb.GetFunctionInfo(), err)
	}

	return dd, err
}

// ConfigDescriptor represents the active configuration of the USB device.
// A device can have several configurations, though most have only one.
type ConfigDescriptor struct {
	Length uint8			// Size of Descriptor in Bytes
	DescriptorType uint8		// Configuration Descriptor Type (0x02)
	TotalLength uint16		// Total Length of Data Returned
	NumInterfaces uint8		// Number of Interfaces
	ConfigurationValue uint8	// Value to Select This Configuration
	ConfigurationIndex uint8	// Index of String Descriptor for Configuration
	Attributes uint8		// Bitmap of Power Attributes
	MaxPower uint8			// Maximum Power Consumption in 2mA units
}

// NewConfigDescriptor constructs a new ConfigDescriptor.
func NewConfigDescriptor(d *Device) (cd *ConfigDescriptor, err error) {

	cd = new(ConfigDescriptor)
	data := make([]byte, BufferSizeConfigDescriptor)

	_, err = d.Control(
		RequestDirectionIn + RequestTypeStandard + RequestRecipientDevice,
		RequestGetDescriptor,
		TypeConfigDescriptor,
		ControlInterface,
		data)

	if err == nil {

		*cd = ConfigDescriptor {
			data[0],
			data[1],
			uint16(data[2]) + (uint16(data[3]) << 8),
			data[4],
			data[5],
			data[6],
			data[7],
			data[8]}
	} else {
		err = fmt.Errorf("%s: %v", gocmdb.GetFunctionInfo(), err)
	}

	return cd, err
}