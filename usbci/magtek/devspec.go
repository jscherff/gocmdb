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
)

type DeviceSpec struct {
	BusNumber	string		`json:"bus_number"`
	BusAddress	string		`json:"bus_address"`
	USBSpec		string		`json:"usb_spec"`
	USBClass	string		`json:"usb_class"`
	USBSubclass	string		`json:"usb_subclass"`
	USBProtocol	string		`json:"usb_protocol"`
	DeviceSpeed	string		`json:"device_speed"`
	DeviceVer	string		`json:"device_ver"`
	MaxPktSize	string		`json:"max_pkt_size"`
	BufferSize	string		`json:"buffer_size"`
	Deltas		[][]string	`json:"deltas" csv:"-" nvp:"-"`
}

func NewDeviceSpec(d *Device) (ds *DeviceSpec, e error) {

	var es []string

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

	if ds.BufferSize, e = d.BufferSize(); e != nil {es = append(es, "BufferSize")}

	if len(es) > 0 {
		e = errors.New("getter errors: " + strings.Join(es, ","))
	}

	return ds, e
}

func (ds *DeviceSpec) Save(fn string) (error) {
	return gocmdb.SaveObject(*ds, fn)
}

func (ds *DeviceSpec) Restore(fn string) (error) {
	return gocmdb.RestoreObject(fn, ds)
}

func (ds *DeviceSpec) Matches(i interface{}) (bool) {
	return reflect.DeepEqual(ds, i)
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
