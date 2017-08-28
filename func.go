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
	"strings"
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

// StructToNVP converts a single-tier struct to a string containing name-
// value pairs separated by newlines.
func StructToNVP (t interface{}) (s string, e error) {

	if ssnvp, e := StructToSlice(t, "nvp"); e == nil {
		for _, snvp := range ssnvp {
			s += fmt.Sprintf("%s: %s\n", snvp[0], snvp[1])
		}
	}

	return s, e
}

// StructToCSV converts a single-tier struct to a string suitable for writing
// to a CSV file.
func StructToCSV (t interface{}) (s string, e error) {

	ssnvp, e := StructToSlice(t, "csv")
	if e != nil {return s, e}

	// For the encoding/csv package, we need to rearranage the elements
	// from an ordered list of {{name, value}, {name, value}, ...} to an
	// ordered list of {{name, name, ...}, {value, value, ...}}.

	ss := make([][]string, 2)

	for _, snvp := range ssnvp {
		for i, v := range snvp {ss[i] = append(ss[i], v)}
	}

	b := new(bytes.Buffer)
	w := csv.NewWriter(b)
	w.WriteAll(ss)

	return b.String(), w.Error()
}

// StructCompare compares the field count, order, names, and values of two
// structs. If the field count or order is different, the structs are not
// comparable and the function returns an error. If the structs differ only
// in field values, the function returns a list of differences.
func StructCompare(a interface{}, b interface{}) (ss[][]string, e error) {

	var iName, iValue int = 0, 1

	if reflect.DeepEqual(a, b) {return ss, e}

	as, e := StructToSlice(a, "")
	if e != nil {return ss, e}

	bs, e := StructToSlice(b, "")
	if e != nil {return ss, e}

	if al, bl := len(as), len(bs); al != bl {
		return ss, fmt.Errorf("field count mismatch: %d != %d", al, bl)
	}

	for i := 0; i < len(as); i++ {

		if as[i][iName] != bs[i][iName] {
			return ss, fmt.Errorf("field name mismatch: %q != %q", i, as[i][iName], bs[i][iName])
		}

		if as[i][iValue] != bs[i][iValue] {
			ss = append(ss, []string{as[i][iName], as[i][iValue], bs[i][iValue]})
		}
	}

	return ss, e
}

// StructToSlice converts a single-tier struct into a slice of slices in the
// form {{name, value}, {name, value}, ...} for consumption by other methods.
// The outer slice maintains the fields in the same order as the struct. The
// tag parameter is the name of the struct tag to use for special processing.
// The primary purpose of this function is to offload tag processing for other
// functions.
func StructToSlice(t interface{}, tag string) (ss[][]string, e error) {

	v := reflect.ValueOf(t)

	if v.Type().Kind() != reflect.Struct {
		return ss, fmt.Errorf("%s: kind is not 'struct'", GetFunctionInfo())
	}

	OUTER:
	for i := 0; i < v.NumField(); i++ {

		fname := v.Type().Field(i).Name
		fval := fmt.Sprintf("%v", v.Field(i).Interface())

		// Process field tags. Function follows the same tag
		// rules as encoding/xml and encoding/json, but only
		// support modified field names, the '-' option, and
		// the 'omitempty' option.

		if tval, ok := v.Type().Field(i).Tag.Lookup(tag); ok {

			topt := strings.Split(tval, ",")

			// Look for unconditional skip or field alias.

			switch {
			case topt[0] == "":
				break
			case topt[0] == "-":
				continue OUTER
			default:
				fname = topt[0]
			}

			// Process tag options. Curreonly only supports
			// the 'omitempty' option. Can be expanded with
			// additional cases.

			for j := 1; j < len(topt); j++ {
				switch topt[j] {
				case "omitempty":
					if fval == "" {continue OUTER}
				}
			}
		}

		ss = append(ss, []string{fname, fval})
	}

	return ss, e
}
