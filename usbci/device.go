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
	"github.com/google/gousb"
	"github.com/jscherff/gocmdb"
	"reflect"
	"strconv"
	"fmt"
)

// Device represents a USB device. The Device struct Desc field contains all
// information about the device. It includes the raw device descriptor, the
// config descriptor of the active config, and the size of the data buffer
// required by the device for vendor commands sent via control transfer.
type Device struct {

	*gousb.Device

	Info struct {
		HostName	string	`json:"host_name"`
		VendorID	string	`json:"vendor_id"`
		ProductID	string	`json:"product_id"`
		VendorName	string	`json:"vendor_name"`
		ProductName	string	`json:"product_name"`
		SerialNumber	string	`json:"serial_number"`
		USBSpec		string	`json:"usb_spec"`
		USBClass	string	`json:"usb_class"`
		USBSubclass	string	`json:"usb_subclass"`
		USBProtocol	string	`json:"usb_protocol"`
		DeviceSpeed	string	`json:"device_speed"`
		DeviceVer	string	`json:"device_ver"`
		MaxPktSize	string	`json:"max_pkt_size"`
		BusNumber	string	`json:"-" xml:"-" csv:"-" nvp:"-"`
		BusAddress	string	`json:"-" xml:"-" csv:"-" nvp:"-"`
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
	}
}

// NewDevice constructs a new Device.
func NewDevice(gd *gousb.Device) (this *Device, err error) {

	this = &Device{gd,
		&Device.Info{
			VendorID: this.Desc.Vendor.String(),
			ProductID: this.Desc.Product.String(),
			USBSpec: this.Desc.Spec.String(),
			USBClass: this.Desc.Class.String(),
			USBSubclass: this.Desc.SubClass.String(),
			USBProtocol: this.Desc.Protocol.String(),
			DeviceSpeed: this.Desc.Speed.String(),
			DeviceVer: this.Desc.Device.String(),
			MaxPktSize: strconv.Itoa(this.Desc.MaxControlPacketSize),
			BusNumber: strconv.Itoa(this.Desc.Bus),
			BusAddress: strconv.Itoa(this.Desc.Address)
		}
	}

	errs := []string

	if this.Info.Hostname, e = os.Hostname(); e != nil {errs = append(errs, "HostName")}
	if this.Info.VendorName, e = this.Manufacturer(); e != nil {errs = append(errs, "VendorName")}
	if this.Info.ProductName, e = this.Product(); e != nil {errs = append(errs, "ProductName")}
	if this.Info.SerialNumber, e = this.SerialNumber(); e != nil {errs = append(errs, "SerialNumber")}

	if len(errs) > 0 {
		err = errors.New("getter errors: " + strings.Join(errs, ","))
	}

	return this, err
}

// Convenience method to retrieve device serial number.
func (this *Device) ID() (sn string, err error) {
	if len(this.Info.SerialNumber) == 0 {
		err = errors.New("no unique identifier")
	}
	return this.Info.SerialNumber, err
}

func (this *Device) Type() (string) {
	return reflect.TypeOf(*this).String()
}

func (this *Device) Save(fn string) (error) {
	return gocmdb.SaveObject(*this, fn)
}

func (this *Device) Restore(fn string) (error) {
	return gocmdb.RestoreObject(fn, this)
}

func (this *Device) Matches(i interface{}) (bool) {
	return reflect.DeepEqual(this.Info, i.Info)
}

func (this *Device) Compare(fn string) (ss [][]string, err error) {
	d := new(Device)
	if err = d.Restore(fn); err != nil {
		return ss, err
	}
	return gocmdb.CompareObjects(*this.Info, *d.Info)
}

func (this *Device) Bare() ([]byte) {
	return []byte(fmt.Sprintf("%s,%s", this.Info.HostName, this.Info.SerialNumber))
}

func (this *Device) JSON() ([]byte, error) {
	return json.Marshal(*this.Info)
}

func (this *Device) XML() ([]byte, error) {
	return xml.Marshal(*this.Info)
}

func (this *Device) CSV() ([]byte, error) {
	return gocmdb.ObjectToCSV(*this.Info)
}

func (this *Device) NVP() ([]byte, error) {
	return gocmdb.ObjectToNVP(*this.Info)
}
