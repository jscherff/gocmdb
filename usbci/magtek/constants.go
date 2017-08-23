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

const (
	MagtekVendorID uint16 = 0x0801

	SureswipeKbPID uint16 = 0x0001
	SureswipeHidPID uint16 = 0x0002
	MagnesafeSwipeKbPID uint16 = 0x0001
	MagnesafeInsertKbPID uint16 = 0x0001
	MagnesafeSwipeHidPID uint16 = 0x0011
	MagnesafeInsertHidPID uint16 = 0x0013
	MagnesafeWirelessHidPID uint16 = 0x0014

	BufferSizeSureswipe int = 24
	BufferSizeMagnesafe int = 60

	CommandGetProp uint8 = 0x00
	CommandSetProp uint8 = 0x01
	CommandResetDevice uint8 = 0x02

	ResultCodeSuccess uint8 = 0x00
	ResultCodeFailure uint8 = 0x01
	ResultCodeBadParam uint8 = 0x02

	PropSoftwareID uint8 = 0x00
	PropDeviceSN uint8 = 0x01
	PropFactorySN uint8 = 0x03
	PropProductVer uint8 = 0x04

	DefaultSNLength int = 7
)
