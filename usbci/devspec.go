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
)

type DeviceSpec struct {
	BusNumber	string		`json:"bus_number" csv:"bus_number"`
	BusAddress	string		`json:"bus_address" csv:"bus_address"`
	USBSpec		string		`json:"usb_spec" csv:"usb_spec"`
	USBClass	string		`json:"usb_class" csv:"usb_class"`
	USBSubclass	string		`json:"usb_subclass" csv:"usb_subclass"`
	USBProtocol	string		`json:"usb_protocol" csv:"usb_protocol"`
	DeviceSpeed	string		`json:"device_speed" csv:"device_speed"`
	DeviceVer	string		`json:"device_ver" csv:"device_ver"`
	MaxPktSize	string		`json:"max_pkt_size" csv:"max_pkt_size"`
	Deltas		[][]string	`json:"deltas" csv:"-" nvp:"-"`
}

func NewDeviceSpec(d *Device) (ds *DeviceSpec, e error) {

	ds = &DeviceSpec {
		BusNumber:	d.BusNumber(),
		BusAddress:	d.BusAddress(),
		USBSpec:	d.USBSpec(),
		USBClass:	d.USBClass(),
		USBSubclass:	d.USBSubclass(),
		USBProtocol:	d.USBProtocol(),
		DeviceSpeed:	d.DeviceSpeed(),
		DeviceVer:	d.DeviceVer(),
		MaxPktSize:	d.MaxPktSize(),
	}

	return ds, nil
}

func (ds *DeviceSpec) Save(fn string) (error) {
	return gocmdb.SaveObject(*ds, fn)
}

func (ds *DeviceSpec) Restore(fn string) (error) {
	return gocmdb.RestoreObject(fn, ds)
}

func (di *DeviceSpec) Matches(i interface{}) (bool) {
	return reflect.DeepEqual(di, i)
}

func (ds *DeviceSpec) Bare() ([]byte) {
	return []byte{}
}

func (ds *DeviceSpec) JSON() ([]byte, error) {
	return json.Marshal(*ds)
}

func (ds *DeviceSpec) XML() ([]byte, error) {
	return xml.Marshal(*ds)
}

func (ds *DeviceSpec) CSV() ([]byte, error) {
	return gocmdb.ObjectToCSV(*ds)
}

func (ds *DeviceSpec) NVP() ([]byte, error) {
	return gocmdb.ObjectToNVP(*ds)
}
