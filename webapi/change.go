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

import "encoding/json"

type Change struct {
	FieldName	string	`json:"omitempty"`
	OldValue	string	`json:"omitempty"`
	NewValue	string	`json:"omitempty"`
}

type Changes []Change

func NewChanges(ss [][]string) (c Changes) {

	for _, s := range ss {
		if len(s) != 3 {continue}
		c = append(c, Change{s[0], s[1], s[2]})
	}

	return c
}

func (this *Changes) JSON() ([]byte, error) {
	return json.Marshal(*this)
}
