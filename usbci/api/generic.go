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

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/google/gousb"
	"github.com/jscherff/gocmdb"
)

const (
	RequestDirectionOut uint8 = 0x00
	RequestDirectionIn uint8 = 0x80
	RequestTypeStandard uint8 = 0x00
	RequestTypeClass uint8 = 0x20
	RequestTypeVendor uint8 = 0x40
	RequestRecipientDevice uint8 = 0x00
	RequestRecipientInterface uint8 = 0x01
	RequestRecipientEndpoint uint8 = 0x02
	RequestRecipientOther uint8 = 0x03

	RequestGetReport uint8 = 0x01
	RequestSetReport uint8 = 0x09
	RequestGetDescriptor uint8 = 0x06

	TypeDeviceDescriptor uint16 = 0x0100
	TypeConfigDescriptor uint16 = 0x0200
	TypeHidDescriptor uint16 = 0x2200
	TypeFeatureReport uint16 = 0x0300

	ControlInterface uint16 = 0x0000

	BufferSizeDeviceDescriptor int = 18
	BufferSizeConfigDescriptor int = 9
)

// Generic decorates a gousb Device with Generic Properties and API.
type Generic struct {

	*gousb.Device

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

// NewGeneric instantiates a Generic wrapper for an existing gousb Device.
func NewGeneric(gd *gousb.Device) (*Generic, error) {

	var err error

	this := &Generic{Device: gd}
	errs := this.Init()

	if len(errs) > 0 {

		emsg := "initialization failures:"

		for k, _ := range errs {
			emsg = fmt.Sprintf("%s %s", emsg, k)
		}

		err = gocmdb.ErrorDecorator(errors.New(emsg))
	}

	return this, err
}


// Init initializes API properties.
func (this *Generic) Init() (errs map[string]bool) {

	var err error

	if this.HostName, err = os.Hostname(); err != nil {
		errs["HostName"] = true
	}
	if this.VendorName, err = this.Manufacturer(); err != nil {
		errs["VendorName"] = true
	}
	if this.ProductName, err = this.Product(); err != nil {
		errs["ProductName"] = true
	}
	if this.SerialNum, err = this.SerialNumber(); err != nil {
		errs["SerialNum"] = true
	}

	this.VendorID = this.Desc.Vendor.String()
	this.ProductID = this.Desc.Product.String()
	this.BusNumber = fmt.Sprintf("%03d", this.Desc.Bus)
	this.BusAddress = fmt.Sprintf("%03d", this.Desc.Address)
	this.USBSpec = this.Desc.Spec.String()
	this.USBClass = this.Desc.Class.String()
	this.USBSubclass = this.Desc.SubClass.String()
	this.USBProtocol = this.Desc.Protocol.String()
	this.DeviceSpeed = this.Desc.Speed.String()
	this.DeviceVer = this.Desc.Device.String()
	this.MaxPktSize = strconv.Itoa(this.Desc.MaxControlPacketSize)
	this.ObjectType = this.Type()

	return errs
}

// Refresh updates properties whose underlying values may have changed.
func (this *Generic) Refresh() (errs map[string]bool) {

	var err error

	if this.SerialNum, err = this.SerialNumber(); err != nil {
		errs["SerialNum"] = true
	}

	return errs
}

// Convenience method to retrieve device serial number.
func (this *Generic) ID() (string) {
	return this.SerialNum
}

// Convenience method to help identify object type to other apps.
func (this *Generic) Type() (string) {
	return reflect.TypeOf(this).String()
}

// Save saves the object to a JSON file.
func (this *Generic) Save(fn string) (error) {
	return gocmdb.SaveObject(*this, fn)
}

// Restore restores the object from a JSON file.
func (this *Generic) Restore(fn string) (error) {
	return gocmdb.RestoreObject(fn, this)
}

// Matches returns true if the objects and their properties are identical.
func (this *Generic) Matches(i interface{}) (bool) {
	return reflect.DeepEqual(this, i)
}

// Compare compares fields and properties and returns an array of differences.
func (this *Generic) Compare(fn string) (ss [][]string, err error) {
	di := new(Generic)
	if err = di.Restore(fn); err != nil {
		return ss, err
	}
	return gocmdb.CompareObjects(*this, *di)
}

// Audit calls Compare and places the results in the Changes field.
func (this *Generic) Audit(fn string) (err error) {
	this.Changes, err = this.Compare(fn)
	return err
}

// SetSerialNum sets the serial number property only. Does not change the
// underlying serial number in the descriptor or on the device.
func (this *Generic) SetSerial(val string) {
	this.SerialNum = val
}

// Filename constructs a convenient filename from the bus number, bus address,
// vendor ID, and product ID. Filenames guaranteed unique on a single computer.
func (this *Generic) Filename() (string) {
	return fmt.Sprintf("%s-%s-%s-%s", this.BusNumber, this.BusAddress, this.VendorID, this.ProductID)
}

// Reports the hostname and serial number in CSV format.
func (this *Generic) Bare() ([]byte) {
	return []byte(this.HostName + "," + this.SerialNum)
}

// Reports all unfiltered fields in JSON format.
func (this *Generic) JSON() ([]byte, error) {
	return json.Marshal(*this)
}

// Reports all unfiltered fields in XML format.
func (this *Generic) XML() ([]byte, error) {
	return xml.Marshal(*this)
}

// Reports all unfiltered fields in CSV format.
func (this *Generic) CSV() ([]byte, error) {
	return gocmdb.ObjectToCSV(*this)
}

// Reports all unfiltered fields as name-value pairs.
func (this *Generic) NVP() ([]byte, error) {
	return gocmdb.ObjectToNVP(*this)
}
