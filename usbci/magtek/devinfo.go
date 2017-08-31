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
	"strings"
	"reflect"
	"errors"
	"os"
)

type DeviceInfo struct {
	HostName	string		`json:"hostname"`
	VendorID	string		`json:"vendor_id"`
	ProductID	string		`json:"product_id"`
	VendorName	string		`json:"vendor_name"`
	ProductName	string		`json:"product_name"`
	SerialNum	string		`json:"serial_num"`
	DeviceSN	string		`json:"device_sn" csv:"-" nvp:"-"`
	FactorySN	string		`json:"factory_sn" csv:"-" nvp:"-"`
	DescriptorSN	string		`json:"descriptor_sn" csv:"-" nvp:"-"`
	ProductVer	string		`json:"product_ver"`
	SoftwareID	string		`json:"software_id"`
	Deltas		[][]string	`json:"deltas" csv:"-" nvp:"-"`
}

func NewDeviceInfo(d *Device) (di *DeviceInfo, e error) {

	var es []string

	di = &DeviceInfo {
		VendorID:	d.VendorID(),
		ProductID:	d.ProductID(),
	}

	if di.HostName, e = os.Hostname(); e != nil {es = append(es, "HostName")}
	if di.VendorName, e = d.VendorName(); e != nil {es = append(es, "VendorName")}
	if di.ProductName, e = d.ProductName(); e != nil {es = append(es, "ProductName")}
	if di.DeviceSN, e = d.DeviceSN(); e != nil {es = append(es, "DeviceSN")}
	if di.FactorySN, e = d.FactorySN(); e != nil {es = append(es, "FactorySN")}
	if di.DescriptorSN, e = d.DescriptorSN(); e != nil {es = append(es, "DescriptorSN")}
	if di.ProductVer, e = d.ProductVer(); e != nil {es = append(es, "ProductVer")}
	if di.SoftwareID, e = d.SoftwareID(); e != nil {es = append(es, "SoftwareId")}

	di.SerialNum = di.DeviceSN

	if len(es) > 0 {
		e = errors.New("getter errors: " + strings.Join(es, ","))
	}

	return di, e
}

func (this *DeviceInfo) ID() (string, error) {
	var e error
	if len(this.SerialNum) == 0 {e = errors.New("no unique identifier")}
	return this.SerialNum, e
}

func (this *DeviceInfo) Type() (string) {
	return reflect.TypeOf(*this).String()
}

func (this *DeviceInfo) Save(fn string) (error) {
	return gocmdb.SaveObject(*this, fn)
}

func (this *DeviceInfo) Restore(fn string) (error) {
	return gocmdb.RestoreObject(fn, this)
}

func (this *DeviceInfo) Matches(i interface{}) (bool) {
	return reflect.DeepEqual(this, i)
}

func (this *DeviceInfo) Compare(fn string) (ss [][]string, e error) {
	di := new(DeviceInfo)
	if e = di.Restore(fn); e != nil {return ss, e}
	return gocmdb.CompareObjects(*this, *di)
}

func (this *DeviceInfo) Bare() ([]byte) {
	return []byte(this.HostName + "," + this.SerialNum)
}

func (this *DeviceInfo) JSON() ([]byte, error) {
	return json.Marshal(*this)
}

func (this *DeviceInfo) XML() ([]byte, error) {
	return xml.Marshal(*this)
}

func (this *DeviceInfo) CSV() ([]byte, error) {
	return gocmdb.ObjectToCSV(*this)
}

func (this *DeviceInfo) NVP() ([]byte, error) {
	return gocmdb.ObjectToNVP(*this)
}
