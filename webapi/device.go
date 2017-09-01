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

package webapi

import (
	"encoding/json"
	"encoding/xml"
)

type Device struct {
	HostName	string	`json:"hostname,omitempty"`
	VendorID	string	`json:"vendor_id,omitempty"`
	ProductID	string	`json:"product_id,omitempty"`
	SerialNum	string	`json:"serial_num,omitempty"`
	VendorName	string	`json:"vendor_name,omitempty"`
	ProductName	string	`json:"product_name,omitempty"`
	ProductVer	string	`json:"product_ver,omitempty"`
	SoftwareID	string	`json:"software_id,omitempty"`
	ObjectType	string	`json:"object_type,omitempty"`
}

func NewDevice(j []byte) (*Device, error) {
	di := new(Device)
	e := json.Unmarshal(j, &di)
	return di, e
}

func (this *Device) ID() (string) {
	return this.SerialNum
}

func (this *Device) JSON() ([]byte, error) {
	return json.Marshal(*this)
}

func (this *Device) XML() ([]byte, error) {
	return xml.Marshal(*this)
}
