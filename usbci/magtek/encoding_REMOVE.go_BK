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
	"encoding/json"
	"encoding/xml"
)

func NewDeviceInfoFromXML(x []byte) (*DeviceInfo, error) {
	i := new(DeviceInfo)
	err := xml.Unmarshal(x, i)
	return i, err
}

func NewDeviceInfoFromJSON(j []byte) (*DeviceInfo, error) {
	i := new(DeviceInfo)
	err := json.Unmarshal(j, i)
	return i, err
}

func (i *DeviceInfo) Prune(fields []string) (error) {
	// Pass a list of fields desired. Iterate through struct fields.
	// If struct field is not in the list, set its value to an empty
	// string so it will not be included in JSON/XML/CSV export. See
	// stackoverflow.com/questions/18926303/iterate-through-a-struct-in-go
	// for hints.
	return nil
}

func (i *DeviceInfo) JSON(min bool) ([]byte, error) {
	if min {return json.Marshal(DeviceInfoMin(*i))}
	return json.Marshal(i)
}

func (i *DeviceInfo) XML(min bool) ([]byte, error) {
	if min {return xml.Marshal(DeviceInfoMin(*i))}
	return xml.Marshal(i)
}

func (i *DeviceInfo) FXML(min bool) ([]byte, error) {
	if min {return xml.MarshalIndent(DeviceInfoMin(*i), "", "\t")}
	return xml.MarshalIndent(i, "", "\t")
}
