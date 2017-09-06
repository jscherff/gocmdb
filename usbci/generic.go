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
	"github.com/jscherff/gocmdb/usbci/api"
)

// Device represents a USB device. The Device struct Desc field contains all
// information about the device. It includes the raw device descriptor, the
// config descriptor of the active config, and the size of the data buffer
// required by the device for vendor commands sent via control transfer.
type Generic struct {
	*gousb.Device
	*api.Generic
}

// NewDevice converts an existing gousb device to a usbci Device.
func NewDevice(gd *gousb.Device) (this *Generic, err error) {

	var errs []string

	mapi := &api.Generic{}
	this := &Device{gd,mapi}

	if err = this.Generic.Init(); err != nil {
		errs = append(errs, "Generic")
	}

	if len(errs) > 0 {
		err = fmt.Errorf("API errors: ", strings.Join(errs, ","))
		err = gocmdb.ErrorDecorator(err)
	}

	return this, err
}

// Refresh items whose underlying values may have chanegd.
func (this *Generic) Refresh() (err error) {

	var errs []string

	if err = this.Generic.Refresh(); err != nil {
		errs = append(errs, "Generic")
	}

	if len(errs) > 0 {
		err = fmt.Errorf("API errors: ", strings.Join(errs, ","))
		err = gocmdb.ErrorDecorator(err)
	}

	return err
}

// Convenience method to retrieve device serial number.
func (this *Generic) ID() (string) {
	return this.Generic.ID()
}

// Convenience method to help identify object type to other apps.
func (this *Generic) Type() (string) {
	return this.Generic.Type()
}
