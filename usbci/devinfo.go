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
	DescriptSN	string		`json:"descript_sn" csv:"-" nvp:"-"`
	Deltas		[][]string	`json:"deltas" csv:"-" nvp:"-"`
}

func NewDeviceInfo(d *Device) (di *DeviceInfo, errs []error) {

	var e error

	di = &DeviceInfo {
		VendorID:	d.VendorID(),
		ProductID:	d.ProductID(),
	}

	if di.HostName, e = os.Hostname();	e != nil {errs = append(errs, e)}
	if di.VendorName, e = d.VendorName();	e != nil {errs = append(errs, e)}
	if di.ProductName, e = d.ProductName();	e != nil {errs = append(errs, e)}
	if di.DescriptSN, e = d.DescriptSN();	e != nil {errs = append(errs, e)}

	di.SerialNum = di.DescriptSN

	return di, errs
}

func GetDeviceInfo(fn string) (*DeviceInfo, error) {
	di := new(DeviceInfo)
	e := gocmdb.RestoreObject(di, fn)
	return di, e
}

func (di *DeviceInfo) JSON() ([]byte, error) {
	return json.Marshal(*di)
}

func (di *DeviceInfo) XML() ([]byte, error) {
	return xml.Marshal(*di)
}

func (di *DeviceInfo) CSV() ([]byte, error) {
	return gocmdb.ObjectToCSV(*di)
}

func (di *DeviceInfo) NVP() ([]byte, error) {
	return gocmdb.ObjectToNVP(*di)
}

func (di *DeviceInfo) Save(fn string) (error) {
	return gocmdb.SaveObject(*di, fn)
}

func (di *DeviceInfo) Matches(t interface{}) (bool) {
	return reflect.DeepEqual(di, t)
}
