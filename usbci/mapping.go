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

var ImportMap = map[string]string {
	"host_name":	"HostName",
	"vendor_id":	"VendorID",
	"product_id":	"ProductID",
	"vendor_name":	"VendorName",
	"product_name":	"ProductName",
	"descript_sn":	"DescriptSN",
	"bus_number":	"BusNumber",
	"bus_address":	"BusAddress",
	"usb_spec":	"USBSpec",
	"usb_class":	"USBClass",
	"usb_subclass":	"USBSubclass",
	"usb_protocol":	"USBProtocol",
	"device_speed":	"DeviceSpeed",
	"device_ver":	"DeviceVer",
	"max_pkt_size":	"MaxPktSize"}

var ExportMap = map[string]string {
	"HostName":	"host_name",
	"VendorID":	"vendor_id",
	"ProductID":	"product_id",
	"VendorName":	"vendor_name",
	"ProductName":	"product_name",
	"DescriptSN":	"descript_sn",
	"BusNumber":	"bus_number",
	"BusAddress":	"bus_address",
	"USBSpec":	"usb_spec",
	"USBClass":	"usb_class",
	"USBSubclass":	"usb_subclass",
	"USBProtocol":	"usb_protocol",
	"DeviceSpeed":	"device_speed",
	"DeviceVer":	"device_ver",
	"MaxPktSize":	"max_pkt_size"}
