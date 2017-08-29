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
	"encoding/json"
	"encoding/xml"
	"reflect"
	"os"
)

type DeviceInfo struct {
	HostName	string		`json:"hostname" csv:"hostname"`
	VendorID	string		`json:"vendor_id" csv:"vendor_id"`
	ProductID	string		`json:"product_id" csv:"product_id"`
	VendorName	string		`json:"vendor_name" csv: "vendor_name"`
	ProductName	string		`json:"product_name" csv: "product_name"`
	SerialNum	string		`json:"serial_num" csv:"serial_num"`
	DescriptSN	string		`json:"-" xml:"-" csv:"-" nvp:"-"`
	BusNumber	string		`json:"-" xml:"-" csv:"-" nvp:"-"`
	BusAddress	string		`json:"-" xml:"-" csv:"-" nvp:"-"`
	USBSpec		string		`json:"-" xml:"-" csv:"-" nvp:"-"`
	USBClass	string		`json:"-" xml:"-" csv:"-" nvp:"-"`
	USBSubclass	string		`json:"-" xml:"-" csv:"-" nvp:"-"`
	USBProtocol	string		`json:"-" xml:"-" csv:"-" nvp:"-"`
	DeviceSpeed	string		`json:"-" xml:"-" csv:"-" nvp:"-"`
	DeviceVer	string		`json:"-" xml:"-" csv:"-" nvp:"-"`
	MaxPktSize	string		`json:"-" xml:"-" csv:"-" nvp:"-"`
	Deltas		[][]string	`json:"deltas" csv:"-" nvp:"-"`
}

type DeviceInfoAll struct {
	HostName	string		`json:"hostname" csv:"hostname"`
	VendorID	string		`json:"vendor_id" csv:"vendor_id"`
	ProductID	string		`json:"product_id" csv:"product_id"`
	VendorName	string		`json:"vendor_name" csv:"vendor_name"`
	ProductName	string		`json:"product_name" csv:"product_name"`
	SerialNum	string		`json:"serial_num" csv:"serial_num"`
	DescriptSN	string		`json:"descriptor_sn" csv:"descriptor_sn"`
	BusNumber	string		`json:"bus_number" csv:"bus_number"`
	BusAddress	string		`json:"bus_address" csv:"bus_address"`
	USBSpec		string		`json:"usb_spec" csv:"usb_spec"`
	USBClass	string		`json:"usb_class" csv:"usb_class"`
	USBSubclass	string		`json:"usb_subclass" csv:"usb_subclass"`
	USBProtocol	string		`json:"usb_protocol" csv:"usb_protocol"`
	DeviceSpeed	string		`json:"device_speed" csv:"device_speed"`
	DeviceVer	string		`json:"device_ver" csv:"device_ver"`
	MaxPktSize	string		`json:"max_pkt_size" csv:"max_pkt_size"`
	Deltas		[][]string	`json:"deltas" csv:"-" nvp:"-"`
}

func NewDeviceInfo(d *Device) (ni *DeviceInfo, errs []error) {

	var e error

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

	if ni.HostName, e = os.Hostname(); e != nil {errs = append(errs, e)}
	if ni.VendorName, e = d.GetVendorName(); e != nil {errs = append(errs, e)}
	if ni.ProductName, e = d.GetProductName(); e != nil {errs = append(errs, e)}
	if ni.DescriptSN, e = d.GetDescriptSN(); e != nil {errs = append(errs, e)}

	ni.SerialNum = ni.DescriptSN

	return ni, errs
}

func GetDeviceInfo(fn string) (i *DeviceInfo, e error) {
	i = new(DeviceInfo)
	e = gocmdb.RestoreObject(i, fn)
	return i, e
}

func (i *DeviceInfo) JSON(all bool) ([]byte, error) {
	if all {return json.Marshal(DeviceInfoAll(*i))}
	return json.Marshal(*i)
}

func (i *DeviceInfo) XML(all bool) ([]byte, error) {
	if all {return xml.Marshal(DeviceInfoAll(*i))}
	return xml.Marshal(*i)
}

func (i *DeviceInfo) CSV(all bool) ([]byte, error) {
	if all {return gocmdb.ObjectToCSV(DeviceInfoAll(*i))}
	return gocmdb.ObjectToCSV(*i)
}

func (i *DeviceInfo) NVP(all bool) ([]byte, error) {
	if all {return gocmdb.ObjectToNVP(DeviceInfoAll(*i))}
	return gocmdb.ObjectToNVP(*i)
}

func (i *DeviceInfo) Save(fn string, all bool) (error) {
	if all {return gocmdb.SaveObject(DeviceInfoAll(*i), fn)}
	return gocmdb.SaveObject(*i, fn)
}

func (i *DeviceInfo) Matches(t interface{}) (bool) {
	return reflect.DeepEqual(i, t)
}

func (i *DeviceInfo) Copy(all bool) (ni *DeviceInfo, e error) {

	ni = new(DeviceInfo)

	if all {
		b, e := json.Marshal(DeviceInfoAll(*i))
		if e != nil {return nil, e}
		e = json.Unmarshal(b, ni)
		if e != nil {return nil, e}
	} else {
		*ni = *i
	}

	return ni, e
}
