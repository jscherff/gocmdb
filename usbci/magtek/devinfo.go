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
	"github.com/jscherff/gocmdb"
	"encoding/json"
	"encoding/xml"
	"reflect"
	"os"
)

// See the following for tips on runtime manipulation of struct tags:
// https://beta.golang.org/doc/go1.8#language
// https://play.golang.org/p/QNArOeqy94

type DeviceInfo struct {
	HostName	string	`json:"hostname" csv:"hostname"`
	VendorID	string	`json:"vendor_id" csv:"vendor_id"`
	ProductID	string	`json:"product_id" csv:"product_id"`
	VendorName	string	`json:"vendor_name" csv:"vendor_name"`
	ProductName	string	`json:"product_name" csv:"product_name"`
	SerialNum	string	`json:"serial_num" csv:"serial_num"`
	DeviceSN	string	`json:"device_sn" csv:"device_sn"`
	FactorySN	string	`json:"factory_sn" csv:"factory_sn"`
	DescriptSN	string	`json:"descriptor_sn" csv:"descriptor_sn"`
	ProductVer	string	`json:"product_ver" csv:"product_ver"`
	SoftwareID	string	`json:"software_id" csv:"software_id"`
	BusNumber	string	`json:"bus_number" csv:"bus_number"`
	BusAddress	string	`json:"bus_address" csv:"bus_address"`
	USBSpec		string	`json:"usb_spec" csv:"usb_spec"`
	USBClass	string	`json:"usb_class" csv:"usb_class"`
	USBSubclass	string	`json:"usb_subclass" csv:"usb_subclass"`
	USBProtocol	string	`json:"usb_protocol" csv:"usb_protocol"`
	DeviceSpeed	string	`json:"device_speed" csv:"device_speed"`
	DeviceVer	string	`json:"device_ver" csv:"device_ver"`
	MaxPktSize	string	`json:"max_pkt_size" csv:"max_pkt_size"`
	BufferSize	string	`json:"buffer_size" csv:"buffer_size"`
}

type DeviceInfoMin struct {
	HostName	string	`json:"hostname" csv:"hostname"`
	VendorID	string	`json:"vendor_id" csv:"vendor_id"`
	ProductID	string	`json:"product_id" csv:"product_id"`
	VendorName	string	`json:"vendor_name" csv:"vendor_name"`
	ProductName	string	`json:"product_name" csv:"product_name"`
	SerialNum	string	`json:"serial_num" csv:"serial_num"`
	DeviceSN	string	`json:"-" xml:"-" csv:"-" nvp:"-"`
	FactorySN	string	`json:"-" xml:"-" csv:"-" nvp:"-"`
	DescriptSN	string	`json:"-" xml:"-" csv:"-" nvp:"-"`
	ProductVer	string	`json:"product_ver" csv:"product_ver"`
	SoftwareID	string	`json:"software_id" csv:"software_id`
	BusNumber	string	`json:"-" xml:"-" csv:"-" nvp:"-"`
	BusAddress	string	`json:"-" xml:"-" csv:"-" nvp:"-"`
	USBSpec		string	`json:"-" xml:"-" csv:"-" nvp:"-"`
	USBClass	string	`json:"-" xml:"-" csv:"-" nvp:"-"`
	USBSubclass	string	`json:"-" xml:"-" csv:"-" nvp:"-"`
	USBProtocol	string	`json:"-" xml:"-" csv:"-" nvp:"-"`
	DeviceSpeed	string	`json:"-" xml:"-" csv:"-" nvp:"-"`
	DeviceVer	string	`json:"-" xml:"-" csv:"-" nvp:"-"`
	MaxPktSize	string	`json:"-" xml:"-" csv:"-" nvp:"-"`
	BufferSize	string	`json:"-" xml:"-" csv:"-" nvp:"-"`
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
	if ni.DeviceSN, e = d.GetDeviceSN(); e != nil {errs = append(errs, e)}
	if ni.FactorySN, e = d.GetFactorySN(); e != nil {errs = append(errs, e)}
	if ni.DescriptSN, e = d.GetDescriptSN(); e != nil {errs = append(errs, e)}
	if ni.ProductVer, e = d.GetProductVer(); e != nil {errs = append(errs, e)}
	if ni.SoftwareID, e = d.GetSoftwareID(); e != nil {errs = append(errs, e)}
	if ni.BufferSize, e = d.GetBufferSize(); e != nil {errs = append(errs, e)}

	ni.SerialNum = ni.DeviceSN

	return ni, errs
}

func (i *DeviceInfo) JSON(min bool) ([]byte, error) {
	if min {return json.Marshal(DeviceInfoMin(*i))}
	return json.Marshal(*i)
}

func (i *DeviceInfo) XML(min bool) ([]byte, error) {
	if min {return xml.Marshal(DeviceInfoMin(*i))}
	return xml.Marshal(*i)
}

func (i *DeviceInfo) CSV(min bool) ([]byte, error) {
	if min {return gocmdb.StructToCSV(DeviceInfoMin(*i))}
	return gocmdb.StructToCSV(*i)
}

func (i *DeviceInfo) NVP(min bool) ([]byte, error) {
	if min {return gocmdb.StructToNVP(DeviceInfoMin(*i))}
	return gocmdb.StructToNVP(*i)
}

func (i *DeviceInfo) Matches(t interface{}) (bool) {
	return reflect.DeepEqual(i, t)
}

func (i *DeviceInfo) Copy(min bool) (ni *DeviceInfo, e error) {

	ni = new(DeviceInfo)

	if min {
		b, e := json.Marshal(DeviceInfoMin(*i))
		if e != nil {return nil, e}
		e = json.Unmarshal(b, ni)
		if e != nil {return nil, e}
	} else {
		*ni = *i
	}

	return ni, e
}
