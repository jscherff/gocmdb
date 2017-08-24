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
	"path/filepath"
	"encoding/csv"
	"runtime"
	"reflect"
	"bytes"
	"fmt"
)

// GetFunctionInfo returns function filename, line number, and function name
// for error reporting.
func GetFunctionInfo() string {

	pc, file, line, success := runtime.Caller(1)
	function := runtime.FuncForPC(pc)

	if !success {
		return "Unknown goroutine"
	}

	return fmt.Sprintf("%s:%d: %s()", filepath.Base(file), line, function.Name())
}

// StructToCSV converts a single-tier struct to a CSV-formatted string.
func StructToCSV (t interface{}) (s string, e error) {

	v := reflect.ValueOf(t)

	if v.Type().Kind() != reflect.Struct {
		return s, fmt.Errorf("%s: kind is not 'struct'", gocmdb.GetFunctionInfo())
	}

	var data = make([][]string, 2)

	for i := 0; i < v.NumField(); i++ {
		//TODO: skip if fieled has tag `csv:"-"`
		data[0] = append(data[0], v.Type().Field(i).Name)
		data[1] = append(data[1], fmt.Sprintf("%v", v.Field(i).Interface()))
	}

	b := new(bytes.Buffer)
	w := csv.NewWriter(b)
	w.WriteAll(data)

	e = w.Error()
	s = b.String()

	return s, e
}
