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
	HostName	string	`csv:"hostname"`
	DeviceSN	string	`csv:"device_sn"`
	VendorID	string	`csv:"vendor_id"`
	ProductID	string	`csv:"product_id"`
	SoftwareID	string	`csv:"software_id"`
	VendorName	string	`csv:"vendor_name"`
	ProductName	string	`csv:"product_name"`
	ProductVer	string	`csv:"product_ver"`
	FactorySN	string	`csv:"factory_sn"`
	DescriptSN	string	`csv:"descript_sn"`
	BusNumber	string	`csv:"bus_number"`
	BusAddress	string	`csv:"bus_address"`
	USBSpec		string	`csv:"usb_spec"`
	USBClass	string	`csv:"usb_class"`
	USBSubclass	string	`csv:"usb_subclass"`
	USBProtocol	string	`csv:"usb_protocol"`
	DeviceSpeed	string	`csv:"device_speed"`
	DeviceVer	string	`csv:"device_ver"`
	MaxPktSize	string	`csv:"max_pkt_size"`
	BufferSize	string	`csv:"buffer_size"`
}

type DeviceInfoMin struct {
	HostName	string	`csv:"hostname"`
	DeviceSN	string	`csv:"device_sn"`
	VendorID	string	`csv:"vendor_id"`
	ProductID	string	`csv:"product_id"`
	SoftwareID	string	`csv:"software_id,omitempty"`
	VendorName	string	`json:"-" xml:"-" csv:"-" nvp:"-"`
	ProductName	string	`json:"-" xml:"-" csv:"-" nvp:"-"`
	ProductVer	string	`json:"-" xml:"-" csv:"-" nvp:"-"`
	FactorySN	string	`json:"-" xml:"-" csv:"-" nvp:"-"`
	DescriptSN	string	`json:"-" xml:"-" csv:"-" nvp:"-"`
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
	if ni.DeviceSN, e = d.GetDeviceSN(); e != nil {errs = append(errs, e)}
	if ni.SoftwareID, e = d.GetSoftwareID(); e != nil {errs = append(errs, e)}
	if ni.VendorName, e = d.GetVendorName(); e != nil {errs = append(errs, e)}
	if ni.ProductName, e = d.GetProductName(); e != nil {errs = append(errs, e)}
	if ni.ProductVer, e = d.GetProductVer(); e != nil {errs = append(errs, e)}
	if ni.FactorySN, e = d.GetFactorySN(); e != nil {errs = append(errs, e)}
	if ni.DescriptSN, e = d.GetDescriptSN(); e != nil {errs = append(errs, e)}
	if ni.BufferSize, e = d.GetBufferSize(); e != nil {errs = append(errs, e)}

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

func (i *DeviceInfo) CSV(min bool) (string, error) {
	if min {return gocmdb.StructToCSV(DeviceInfoMin(*i))}
	return gocmdb.StructToCSV(*i)
}

func (i *DeviceInfo) NVP(min bool) (string, error) {
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
