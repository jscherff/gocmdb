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

	MarshalPrefix string = ""
	MarshalIndent string = "\t"
)

// Generic decorates a gousb Device with Generic Properties and API.
type Generic struct {

	*gousb.Device		`json:"-" xml:"-" csv:"-" nvp:"-" cmp:"-"`

	HostName	string
	VendorID	string
	ProductID	string
	VendorName	string
	ProductName	string
	SerialNum	string
	SoftwareID	string
	ProductVer	string
	BufferSize	int

	BusNumber	string	`json:"-" xml:"-" csv:"-" nvp:"-" cmp:"-"`
	BusAddress	string	`json:"-" xml:"-" csv:"-" nvp:"-" cmp:"-"`
	USBSpec		string	`csv:"-" nvp:"-"`
	USBClass	string	`csv:"-" nvp:"-"`
	USBSubclass	string	`csv:"-" nvp:"-"`
	USBProtocol	string	`csv:"-" nvp:"-"`
	DeviceSpeed	string	`csv:"-" nvp:"-"`
	DeviceVer	string	`csv:"-" nvp:"-"`
	MaxPktSize	string	`csv:"-" nvp:"-"`
	ObjectType	string	`csv:"-" nvp:"-"`

	Vendor		map[string]string `json:",omitempty" xml:",omitempty" csv:"-" nvp:"-" cmp:"-"`

	Changes		[][]string `csv:"-" nvp:"-" cmp:"-"`
}

// NewGeneric instantiates a Generic wrapper for an existing gousb Device.
func NewGeneric(gd *gousb.Device) (*Generic, error) {

	vm := make(map[string]string)

	if gd == nil {
		return &Generic{Device: &gousb.Device{}, Vendor: vm}, nil
	}

	var err error

	this := &Generic{Device: gd, Vendor: vm}
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

// ID is a convenience method to retrieve device serial number.
func (this *Generic) ID() (string) {
	return this.SerialNum
}

// Type is a convenience method to help identify object type to other apps.
func (this *Generic) Type() (string) {
	return reflect.TypeOf(this).String()
}

// Save saves the object to a JSON file.
func (this *Generic) Save(fn string) (error) {
	return gocmdb.Save(*this, fn)
}

// Restore restores the object from a JSON file.
func (this *Generic) Restore(fn string) (error) {
	return gocmdb.Restore(fn, this)
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
	return gocmdb.Compare(*this, *di)
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

// Legacy reports the hostname and serial number in CSV format.
func (this *Generic) Legacy() ([]byte) {
	return []byte(this.HostName + "," + this.SerialNum)
}

// JSON reports all unfiltered fields in JSON format.
func (this *Generic) JSON() ([]byte, error) {
	return json.Marshal(*this)
}

// XML reports all unfiltered fields in XML format.
func (this *Generic) XML() ([]byte, error) {
	return xml.Marshal(*this)
}

// CSV reports all unfiltered fields in CSV format.
func (this *Generic) CSV() ([]byte, error) {
	return gocmdb.ToCSV(*this)
}

// NVP reports all unfiltered fields as name-value pairs.
func (this *Generic) NVP() ([]byte, error) {
	return gocmdb.ToNVP(*this)
}

// PrettyJSON reports all unfiltered fields in formatted JSON format.
func (this *Generic) PrettyJSON() ([]byte, error) {
	return json.MarshalIndent(*this, MarshalPrefix, MarshalIndent)
}

// PrettyXML reports all unfiltered fields in formatted XML format.
func (this *Generic) PrettyXML() ([]byte, error) {
	return xml.MarshalIndent(*this, MarshalPrefix, MarshalIndent)
}
