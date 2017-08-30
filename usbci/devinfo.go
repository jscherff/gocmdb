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
	DescriptorSN	string		`json:"descriptor_sn" csv:"-" nvp:"-"`
	Deltas		[][]string	`json:"deltas" csv:"-" nvp:"-"`
}

func NewDeviceInfo(d *Device) (di *DeviceInfo, e error) {

	ge := new(gocmdb.GetterError)

	di = &DeviceInfo {
		VendorID:	d.VendorID(),
		ProductID:	d.ProductID(),
	}

	if di.HostName, e = os.Hostname(); e != nil {ge.Add("HostName", e)}
	if di.VendorName, e = d.VendorName(); e != nil {ge.Add("VendorName", e)}
	if di.ProductName, e = d.ProductName();	e != nil {ge.Add("ProductName", e)}
	if di.DescriptorSN, e = d.DescriptorSN(); e != nil {ge.Add("DescriptorSN", e)}

	di.SerialNum = di.DescriptorSN

	return di, ge
}

func (di *DeviceInfo) Save(fn string) (error) {
	return gocmdb.SaveObject(*di, fn)
}

func (di *DeviceInfo) Restore(fn string) (error) {
	return gocmdb.RestoreObject(fn, di)
}

func (di *DeviceInfo) Matches(c *gocmdb.Comparable) (bool) {
	return reflect.DeepEqual(di, c)
}

func (di *DeviceInfo) Bare() ([]byte) {
	return []byte(di.HostName + "," + di.SerialNum)
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
