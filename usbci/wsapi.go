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

// WSAPI is a collection of getters for use by the gocmdbd webserver
type WSAPI struct {
	*Generic
}

func NewWSAPI() (*WSAPI) {
	g, _ := NewGeneric(nil)
	return &WSAPI{g}
}

func (this *WSAPI) GetHostName() (string) { return this.HostName }
func (this *WSAPI) GetVendorID() (string) { return this.VendorID }
func (this *WSAPI) GetProductID() (string) { return this.ProductID }
func (this *WSAPI) GetVendorName() (string) { return this.VendorName }
func (this *WSAPI) GetProductName() (string) { return this.ProductName }
func (this *WSAPI) GetSerialNum() (string) { return this.SerialNum }
func (this *WSAPI) GetSoftwareID() (string) { return this.SoftwareID }
func (this *WSAPI) GetProductVer() (string) { return this.ProductVer }
func (this *WSAPI) GetBufferSize() (int) { return this.BufferSize }
func (this *WSAPI) GetBusNumber() (int) { return this.BusNumber }
func (this *WSAPI) GetBusAddress() (int) { return this.BusAddress }
func (this *WSAPI) GetPortNumber() (int) { return this.PortNumber }
func (this *WSAPI) GetUSBSpec() (string) { return this.USBSpec }
func (this *WSAPI) GetUSBClass() (string) { return this.USBClass }
func (this *WSAPI) GetUSBSubclass() (string) { return this.USBSubclass }
func (this *WSAPI) GetUSBProtocol() (string) { return this.USBProtocol }
func (this *WSAPI) GetDeviceSpeed() (string) { return this.DeviceSpeed }
func (this *WSAPI) GetDeviceVer() (string) { return this.DeviceVer }
func (this *WSAPI) GetMaxPktSize() (int) { return this.MaxPktSize }
func (this *WSAPI) GetObjectType() (string) { return this.ObjectType }
func (this *WSAPI) GetDeviceSN() (string) { return this.DeviceSN }
func (this *WSAPI) GetFactorySN() (string) { return this.FactorySN }
func (this *WSAPI) GetDescriptorSN() (string) { return this.DescriptorSN }
func (this *WSAPI) GetVendor() (map[string]string) { return this.Vendor }
func (this *WSAPI) GetChanges() ([][]string) { return this.Changes }
