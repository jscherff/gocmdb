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

import "os"

type DeviceInfo struct {
	HostName	string
	VendorID	string	`json:",omitempty" xml:",omitempty" csv:",omitempty"`
	ProductID	string	`json:",omitempty" xml:",omitempty" csv:",omitempty"`
	VendorName	string	`json:",omitempty" xml:",omitempty" csv:",omitempty"`
	ProductName	string	`json:",omitempty" xml:",omitempty" csv:",omitempty"`
	DescriptSN	string	`json:",omitempty" xml:",omitempty" csv:",omitempty"`
	BusNumber	string	`json:",omitempty" xml:",omitempty" csv:",omitempty"`
	BusAddress	string	`json:",omitempty" xml:",omitempty" csv:",omitempty"`
	USBSpec		string	`json:",omitempty" xml:",omitempty" csv:",omitempty"`
	USBClass	string	`json:",omitempty" xml:",omitempty" csv:",omitempty"`
	USBSubclass	string	`json:",omitempty" xml:",omitempty" csv:",omitempty"`
	USBProtocol	string	`json:",omitempty" xml:",omitempty" csv:",omitempty"`
	DeviceSpeed	string	`json:",omitempty" xml:",omitempty" csv:",omitempty"`
	DeviceVer	string	`json:",omitempty" xml:",omitempty" csv:",omitempty"`
	MaxPktSize	string	`json:",omitempty" xml:",omitempty" csv:",omitempty"`
}

type DeviceInfoMin struct {
	HostName	string
	VendorID	string	`json:",omitempty" xml:",omitempty" csv:",omitempty"`
	ProductID	string	`json:",omitempty" xml:",omitempty" csv:",omitempty"`
	VendorName	string	`json:"-" xml:"-" csv:"-"`
	ProductName	string	`json:"-" xml:"-" csv:"-"`
	DescriptSN	string	`json:"-" xml:"-" csv:"-"`
	BusNumber	string	`json:"-" xml:"-" csv:"-"`
	BusAddress	string	`json:"-" xml:"-" csv:"-"`
	USBSpec		string	`json:"-" xml:"-" csv:"-"`
	USBClass	string	`json:"-" xml:"-" csv:"-"`
	USBSubclass	string	`json:"-" xml:"-" csv:"-"`
	USBProtocol	string	`json:"-" xml:"-" csv:"-"`
	DeviceSpeed	string	`json:"-" xml:"-" csv:"-"`
	DeviceVer	string	`json:"-" xml:"-" csv:"-"`
	MaxPktSize	string	`json:"-" xml:"-" csv:"-"`
}

func NewDeviceInfo(d *Device) (ni *DeviceInfo, errs []error) {

	var e error

	ni = new(DeviceInfo)

	ni.VendorID = d.GetVendorID()
	ni.ProductID = d.GetProductID()
	ni.BusNumber = d.GetBusNumber()
	ni.BusAddress = d.GetBusAddress()
	ni.USBSpec = d.GetUSBSpec()
	ni.USBClass = d.GetUSBClass()
	ni.USBSubclass = d.GetUSBSubclass()
	ni.USBProtocol = d.GetUSBProtocol()
	ni.DeviceSpeed = d.GetDeviceSpeed()
	ni.DeviceVer = d.GetDeviceVer()
	ni.MaxPktSize = d.GetMaxPktSize()
/*
	ni = &DeviceInfo {
		VendorID:	d.GetVendorID(),
		ProductID:	d.GetProductID(),
		BusNumber:	d.GetBusNumber(),
		BusAddress:	d.GetBusAddress(),
		USBSpec:	d.GetUSBSpec(),
		USBClass:	d.GetUSBClass(),
		USBSubclass:	d.GetUSBSubclass(),
		USBProtocol:	d.GetUSBProtocol(),
		DeviceSpeed:	d.GetDeviceSpeed(),
		DeviceVer:	d.GetDeviceVer(),
		MaxPktSize:	d.GetMaxPktSize()}
*/

	if ni.HostName, e = os.Hostname(); e != nil {errs = append(errs, e)}
	if ni.VendorName, e = d.GetVendorName(); e != nil {errs = append(errs, e)}
	if ni.ProductName, e = d.GetProductName(); e != nil {errs = append(errs, e)}
	if ni.DescriptSN, e = d.GetDescriptSN(); e != nil {errs = append(errs, e)}

	return ni, errs
}
