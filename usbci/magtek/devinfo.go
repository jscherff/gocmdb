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
	HostName	string
	DeviceSN	string
	VendorID	string
	ProductID	string
	SoftwareID	string
	VendorName	string
	ProductName	string
	ProductVer	string
	FactorySN	string
	DescriptSN	string
	BusNumber	string
	BusAddress	string
	USBSpec		string
	USBClass	string
	USBSubclass	string
	USBProtocol	string
	DeviceSpeed	string
	DeviceVer	string
	MaxPktSize	string
	BufferSize	string
}

type DeviceInfoMin struct {
	HostName	string
	DeviceSN	string
	VendorID	string	`json:",omitempty" xml:",omitempty" csv:",omitempty"`
	ProductID	string	`json:",omitempty" xml:",omitempty" csv:",omitempty"`
	SoftwareID	string	`json:",omitempty" xml:",omitempty" csv:",omitempty"`
	VendorName	string	`json:"-" xml:"-" csv:"-"`
	ProductName	string	`json:"-" xml:"-" csv:"-"`
	ProductVer	string	`json:"-" xml:"-" csv:"-"`
	FactorySN	string	`json:"-" xml:"-" csv:"-"`
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
	BufferSize	string	`json:"-" xml:"-" csv:"-"`
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


