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

package gocmdb

import (
	`encoding/json`
	`log`
	`os`
	`github.com/jscherff/gocmdb/usbci`
)

type TestData struct {
	Jsn map[string][]byte
	Mag map[string]*usbci.Magtek
	Gen map[string]*usbci.Generic
	Sig map[string]map[string][32]byte
	Chg [][]string
	Clg []string
}

var td *TestData

func init() {

	td = new(TestData)

	if err := createObjects(); err != nil {
		log.Fatal(err)
	}
}

func createObjects() error {

	if fh, err := os.Open(`testdata.json`); err != nil {
		return err
	} else {
		defer fh.Close()
		return json.NewDecoder(fh).Decode(&td)
	}
}
