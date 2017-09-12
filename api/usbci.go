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

package api

// usbci is a collection of getters for use by the gocmdbd webserver
type usbci struct {
	HostName     string		`json:"host_name"     csv:"host_name"`
	VendorID     string		`json:"vendor_id"     csv:"vendor_id"`
	ProductID    string		`json:"product_id"    csv:"product_id"`
	VendorName   string		`json:"vendor_name"   csv:"vendor_name"`
	ProductName  string		`json:"product_name"  csv:"product_name"`
	SerialNum    string		`json:"serial_num"    csv:"serial_num"`
	SoftwareID   string		`json:"software_id"   csv:"software_id"`
	ProductVer   string		`json:"product_ver"   csv:"product_ver"`

	BusNumber    int		`json:"bus_number"    csv:"-" nvp:"-" cmp:"-"`
	BusAddress   int		`json:"bus_address"   csv:"-" nvp:"-" cmp:"-"`
	PortNumber   int		`json:"port_number"   csv:"-" nvp:"-" cmp:"-"`
	BufferSize   int		`json:"buffer_size"   csv:"-" nvp:"-"`
	MaxPktSize   int		`json:"max_pkt_size"  csv:"-" nvp:"-"`
	USBSpec	     string		`json:"usb_spec"      csv:"-" nvp:"-"`
	USBClass     string		`json:"usb_class"     csv:"-" nvp:"-"`
	USBSubclass  string		`json:"usb_subclass"  csv:"-" nvp:"-"`
	USBProtocol  string		`json:"usb_protocol"  csv:"-" nvp:"-"`
	DeviceSpeed  string		`json:"device_speed"  csv:"-" nvp:"-"`
	DeviceVer    string		`json:"device_ver"    csv:"-" nvp:"-"`
	ObjectType   string		`json:"object_type"   csv:"-" nvp:"-"`

	DeviceSN     string		`json:"device_sn"     csv:"-" nvp:"-"`
	FactorySN    string		`json:"factory_sn"    csv:"-" nvp:"-"`
	DescriptorSN string		`json:"descriptor_sn" csv:"-" nvp:"-"`

	Changes	     [][]string		`json:"changes"       csv:"-" nvp:"-" cmp:"-"`
	Vendor       map[string]string	`json:"vendor"        csv:"-" nvp:"-" cmp:"-"`

}

func Newusbci() (*usbci) {
	return &usbci{Vendor: make(map[string]string)}
}

func (this *usbci) ID() (string) { return this.SerialNum }
func (this *usbci) VID() (string) { return this.VendorID }
func (this *usbci) PID() (string) { return this.ProductID }
func (this *usbci) GetHostName() (string) { return this.HostName }
func (this *usbci) GetVendorID() (string) { return this.VendorID }
func (this *usbci) GetProductID() (string) { return this.ProductID }
func (this *usbci) GetVendorName() (string) { return this.VendorName }
func (this *usbci) GetProductName() (string)	{ return this.ProductName }
func (this *usbci) GetSerialNum() (string) { return this.SerialNum }
func (this *usbci) GetSoftwareID() (string) { return this.SoftwareID }
func (this *usbci) GetProductVer() (string) { return this.ProductVer }
func (this *usbci) GetBufferSize() (int) { return this.BufferSize }
func (this *usbci) GetBusNumber() (int) { return this.BusNumber }
func (this *usbci) GetBusAddress() (int) { return this.BusAddress }
func (this *usbci) GetPortNumber() (int) { return this.PortNumber }
func (this *usbci) GetUSBSpec() (string) { return this.USBSpec }
func (this *usbci) GetUSBClass() (string) { return this.USBClass }
func (this *usbci) GetUSBSubclass() (string) { return this.USBSubclass }
func (this *usbci) GetUSBProtocol() (string) { return this.USBProtocol }
func (this *usbci) GetDeviceSpeed() (string) { return this.DeviceSpeed }
func (this *usbci) GetDeviceVer() (string) { return this.DeviceVer }
func (this *usbci) GetMaxPktSize() (int) { return this.MaxPktSize }
func (this *usbci) GetObjectType() (string) { return this.ObjectType }
func (this *usbci) GetDeviceSN() (string) { return this.DeviceSN }
func (this *usbci) GetFactorySN() (string) { return this.FactorySN }
func (this *usbci) GetDescriptorSN() (string) { return this.DescriptorSN }
func (this *usbci) GetVendor() (map[string]string) { return this.Vendor }
func (this *usbci) GetChanges() ([][]string) { return this.Changes }
