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
	"reflect"
	"strconv"
	"strings"
	"fmt"
	"os"

	"github.com/google/gousb"
	"github.com/jscherff/gocmdb"
)

// Device represents a USB device. The Device struct Desc field contains all
// information about the device. It includes the raw device descriptor, the
// config descriptor of the active config, and the size of the data buffer
// required by the device for vendor commands sent via control transfer.
type Device struct {
	*gousb.Device
	*DevInfo
}

// NewDevice converts an existing gousb device to a usbci Device.
func NewDevice(gd *gousb.Device) (*Device, error) {

	var err error
	var errs []string

	di := &DevInfo{}
	this := &Device{gd,di}

	if this.HostName, err = os.Hostname(); err != nil {
		errs = append(errs, "HostName")
	}
	if this.VendorName, err = this.Manufacturer(); err != nil {
		errs = append(errs, "VendorName")
	}
	if this.ProductName, err = this.Product(); err != nil {
		errs = append(errs, "ProductName")
	}
	if this.SerialNum, err = this.SerialNumber(); err != nil {
		errs = append(errs, "SerialNum")
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

	if len(errs) > 0 {
		err = fmt.Errorf("getter errors: ", strings.Join(errs, ","))
		err = gocmdb.ErrorDecorator(err)
	}

	return this, err
}

// Refresh items whose underlying values may have chanegd.
func (this *Device) Refresh() (err error) {
	this.SerialNum, err = this.SerialNumber()
	return err
}

// Convenience method to retrieve device serial number.
func (this *Device) ID() (string) {
	return this.SerialNum
}

// Convenience method to help identify object type to other apps.
func (this *Device) Type() (string) {
	return reflect.TypeOf(this).String()
}
