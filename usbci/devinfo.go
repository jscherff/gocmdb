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
	"encoding/json"
	"encoding/xml"
	"reflect"
	"fmt"

	"github.com/jscherff/gocmdb"
)

// DevInfo is a collection of useful information about a USB device.
type DevInfo struct {

	HostName	string	`json:"host_name"`
	VendorID	string	`json:"vendor_id"`
	ProductID	string	`json:"product_id"`
	VendorName	string	`json:"vendor_name"`
	ProductName	string	`json:"product_name"`
	SerialNum	string	`json:"serial_num"`
	ProductVer	string	`json:"product_ver"`
	SoftwareID	string	`json:"software_id"`

	BusNumber	string	`json:"-" xml:"-" csv:"-" nvp:"-" compare:"-"`
	BusAddress	string	`json:"-" xml:"-" csv:"-" nvp:"-" compare:"-"`
	USBSpec		string	`json:"usb_spec" csv:"-" nvp:"-"`
	USBClass	string	`json:"usb_class" csv:"-" nvp:"-"`
	USBSubclass	string	`json:"usb_subclass" csv:"-" nvp:"-"`
	USBProtocol	string	`json:"usb_protocol" csv:"-" nvp:"-"`
	DeviceSpeed	string	`json:"device_speed" csv:"-" nvp:"-"`
	DeviceVer	string	`json:"device_ver" csv:"-" nvp:"-"`
	MaxPktSize	string	`json:"max_pkt_size" csv:"-" nvp:"-"`
	ObjectType	string	`json:"object_type" csv:"-" nvp:"-"`

	Vendor0		string	`json:"vendor0,omitempty" xml:",omitempty" csv:",omitempty" nvp:"omitempty"`
	Vendor1		string	`json:"vendor1,omitempty" xml:",omitempty" csv:",omitempty" nvp:"omitempty"`
	Vendor2		string	`json:"vendor2,omitempty" xml:",omitempty" csv:",omitempty" nvp:"omitempty"`
	Vendor3		string	`json:"vendor3,omitempty" xml:",omitempty" csv:",omitempty" nvp:"omitempty"`
	Vendor4		string	`json:"vendor4,omitempty" xml:",omitempty" csv:",omitempty" nvp:"omitempty"`
	Vendor5		string	`json:"vendor5,omitempty" xml:",omitempty" csv:",omitempty" nvp:"omitempty"`
	Vendor6		string	`json:"vendor6,omitempty" xml:",omitempty" csv:",omitempty" nvp:"omitempty"`
	Vendor7		string	`json:"vendor7,omitempty" xml:",omitempty" csv:",omitempty" nvp:"omitempty"`
	Vendor8		string	`json:"vendor8,omitempty" xml:",omitempty" csv:",omitempty" nvp:"omitempty"`
	Vendor9		string	`json:"vendor9,omitempty" xml:",omitempty" csv:",omitempty" nvp:"omitempty"`

	Changes		[][]string `json:"audit" csv:"-" nvp:"-" compare"-"`
}

// Save saves the object to a JSON file.
func (this *DevInfo) Save(fn string) (error) {
	return gocmdb.SaveObject(*this, fn)
}

// Restore restores the object from a JSON file.
func (this *DevInfo) Restore(fn string) (error) {
	return gocmdb.RestoreObject(fn, this)
}

// Matches returns true if the objects and their properties are identical.
func (this *DevInfo) Matches(i interface{}) (bool) {
	return reflect.DeepEqual(this, i)
}

// Compare compares fields and properties and returns an array of differences.
func (this *DevInfo) Compare(fn string) (ss [][]string, err error) {
	di := new(DevInfo)
	if err = di.Restore(fn); err != nil {
		return ss, err
	}
	return gocmdb.CompareObjects(*this, *di)
}

// Audit calls Compare and places the results in the Changes field.
func (this *DevInfo) Audit(fn string) (err error) {
	this.Changes, err = this.Compare(fn)
	return err
}

// SetSerialNum sets the serial number value. Does not change the serial 
// number on the device.
func (this *DevInfo) SetSerial(val string) {
	this.SerialNum = val
}

// Filename constructs a convenient filename from the bus number, bus address,
// vendor ID, and product ID. Filenames guaranteed unique on a single computer.
func (this *DevInfo) Filename() (string) {
	return fmt.Sprintf("%s-%s-%s-%s", this.BusNumber, this.BusAddress, this.VendorID, this.ProductID)
}

// Reports the hostname and serial number in CSV format.
func (this *DevInfo) Bare() ([]byte) {
	return []byte(this.HostName + "," + this.SerialNum)
}

// Reports all unfiltered fields in JSON format.
func (this *DevInfo) JSON() ([]byte, error) {
	return json.Marshal(*this)
}

// Reports all unfiltered fields in XML format.
func (this *DevInfo) XML() ([]byte, error) {
	return xml.Marshal(*this)
}

// Reports all unfiltered fields in CSV format.
func (this *DevInfo) CSV() ([]byte, error) {
	return gocmdb.ObjectToCSV(*this)
}

// Reports all unfiltered fields as name-value pairs.
func (this *DevInfo) NVP() ([]byte, error) {
	return gocmdb.ObjectToNVP(*this)
}
