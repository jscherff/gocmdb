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
