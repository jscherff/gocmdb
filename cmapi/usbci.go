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

package cmapi

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	FieldNameIx int = 0
	OldValueIx int = 1
	NewValueIx int = 2
)

// UsbCi is a collection of getters for use by the gocmdbd webserver
type UsbCi struct {
	HostName     string		`json:"host_name"     db:"host_name"`
	VendorID     string		`json:"vendor_id"     db:"vendor_id"`
	ProductID    string		`json:"product_id"    db:"product_id"`
	SerialNum    string		`json:"serial_num"    db:"serial_num"`
	VendorName   string		`json:"vendor_name"   db:"vendor_name"`
	ProductName  string		`json:"product_name"  db:"product_name"`
	ProductVer   string		`json:"product_ver"   db:"product_ver"`
	SoftwareID   string		`json:"software_id"   db:"software_id"`

	BusNumber    int		`json:"bus_number"    db:"bus_number"`
	BusAddress   int		`json:"bus_address"   db:"bus_address"`
	PortNumber   int		`json:"port_number"   db:"port_number"`
	BufferSize   int		`json:"buffer_size"   db:"buffer_size"`
	MaxPktSize   int		`json:"max_pkt_size"  db:"max_pkt_size"`

	USBSpec	     string		`json:"usb_spec"      db:"usb_spec"`
	USBClass     string		`json:"usb_class"     db:"usb_class"`
	USBSubclass  string		`json:"usb_subclass"  db:"usb_subclass"`
	USBProtocol  string		`json:"usb_protocol"  db:"usb_protocol"`
	DeviceSpeed  string		`json:"device_speed"  db:"device_speed"`
	DeviceVer    string		`json:"device_ver"    db:"device_ver"`
	DeviceSN     string		`json:"device_sn"     db:"device_sn"`
	FactorySN    string		`json:"factory_sn"    db:"factory_sn"`
	DescriptorSN string		`json:"descriptor_sn" db:"descriptor_sn"`
	ObjectType   string		`json:"object_type"   db:"object_type"`

	Vendor       map[string]string	`json:"vendor"        db:"-"`
	Changes	     [][]string		`json:"changes"       db:"-"`
}

func NewUsbCi() (*UsbCi) {
	return &UsbCi{Vendor: make(map[string]string)}
}

func (this *UsbCi) ID() (string) { return this.SerialNum }
func (this *UsbCi) VID() (string) { return this.VendorID }
func (this *UsbCi) PID() (string) { return this.ProductID }
func (this *UsbCi) GetHostName() (string) { return this.HostName }
func (this *UsbCi) GetVendorID() (string) { return this.VendorID }
func (this *UsbCi) GetProductID() (string) { return this.ProductID }
func (this *UsbCi) GetSerialNum() (string) { return this.SerialNum }
func (this *UsbCi) GetVendorName() (string) { return this.VendorName }
func (this *UsbCi) GetProductName() (string)	{ return this.ProductName }
func (this *UsbCi) GetSoftwareID() (string) { return this.SoftwareID }
func (this *UsbCi) GetProductVer() (string) { return this.ProductVer }
func (this *UsbCi) GetBusNumber() (int) { return this.BusNumber }
func (this *UsbCi) GetBusAddress() (int) { return this.BusAddress }
func (this *UsbCi) GetPortNumber() (int) { return this.PortNumber }
func (this *UsbCi) GetBufferSize() (int) { return this.BufferSize }
func (this *UsbCi) GetMaxPktSize() (int) { return this.MaxPktSize }
func (this *UsbCi) GetUSBSpec() (string) { return this.USBSpec }
func (this *UsbCi) GetUSBClass() (string) { return this.USBClass }
func (this *UsbCi) GetUSBSubclass() (string) { return this.USBSubclass }
func (this *UsbCi) GetUSBProtocol() (string) { return this.USBProtocol }
func (this *UsbCi) GetDeviceSpeed() (string) { return this.DeviceSpeed }
func (this *UsbCi) GetDeviceVer() (string) { return this.DeviceVer }
func (this *UsbCi) GetDeviceSN() (string) { return this.DeviceSN }
func (this *UsbCi) GetFactorySN() (string) { return this.FactorySN }
func (this *UsbCi) GetDescriptorSN() (string) { return this.DescriptorSN }
func (this *UsbCi) GetObjectType() (string) { return this.ObjectType }
func (this *UsbCi) GetVendor() (map[string]string) { return this.Vendor }
func (this *UsbCi) GetChanges() ([][]string) { return this.Changes }
