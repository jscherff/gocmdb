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
	"bytes"
	"fmt"
	"encoding/csv"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

const (
	NameIx int = 0
	ValueIx int = 1
)

var ReportFormats = [][]string {
	[]string {"csv",  "Comma-Separated Value format"},
	[]string {"nvp",  "Name-Value Pair format"},
	[]string {"xml",  "eXtensible Markup Language"},
	[]string {"json", "JavaScript Object Notation"},
}

// ErrorDecorator prepends function filename, line number, and function name
// to error messages.
func ErrorDecorator(err error) (error) {

	var msg string

	pc, file, line, success := runtime.Caller(1)
	function := runtime.FuncForPC(pc)

	if success {
		msg = fmt.Sprintf("%s:%d: %s()", filepath.Base(file), line, function.Name())
	} else {
		msg = "unknown goroutine"
	}

	return fmt.Errorf("%s: %v", msg, err)
}

// ToCSV converts a single-tier struct to a string suitable for writing
// to a CSV file. For the csv package, we need to rearranage the elements from
// an ordered list of {{name, value}, {name, value}, ...} to an ordered list
// of {{name, name, ...}, {value, value, ...}}.
func ToCSV (t interface{}) (b []byte, err error) {

	var ssi [][]string

	if ssi, err = ToSlice(t, "csv"); err == nil {

		ss := make([][]string, 2)

		for _, si := range ssi {
			 ss[NameIx] = append(ss[NameIx], si[NameIx])
			 ss[ValueIx] = append(ss[ValueIx], si[ValueIx])
		}

		bb := new(bytes.Buffer)
		cw := csv.NewWriter(bb)
		cw.WriteAll(ss)

		if b, err = bb.Bytes(), cw.Error(); err != nil {
			err = ErrorDecorator(err)
		}

	}

	return b, err
}

// ToNVP converts a single-tier struct to a string containing name-
// value pairs separated by newlines.
func ToNVP (t interface{}) (b []byte, err error) {

	var ssi [][]string

	if ssi, err = ToSlice(t, "nvp"); err == nil {

		var s string

		for _, si := range ssi {
			s += fmt.Sprintf("%s:%s\n", si[NameIx], si[ValueIx])
		}

		b = []byte(s)
	}

	return b, err
}

// Save persists an object to a JSON file.
func Save(t interface{}, fn string) (err error) {

	fh, err := os.Create(fn)
	defer fh.Close()

	if err == nil {
		je := json.NewEncoder(fh)
		err = je.Encode(&t)
	}

	if err != nil {
		err = ErrorDecorator(err)
	}

	return err
}

// Restore restores an object from a JSON file.
func Restore(fn string, t interface{}) (err error) {

	fh, err := os.Open(fn)
	defer fh.Close()

	if err == nil {
		jd := json.NewDecoder(fh)
		err = jd.Decode(&t)
	}

	if err != nil {
		err = ErrorDecorator(err)
	}

	return err
}

// Compare compares the field count, order, names, and values of two structs.
// If the field count or order is different, the structs are not comparable
// and the function returns an error. If the structs differ only in field
// values, the function returns a list of differences.
func Compare(t1 interface{}, t2 interface{}) (ss[][]string, err error) {

	if reflect.DeepEqual(t1, t2) {
		return ss, err
	}

	var (
		st1, st2 [][]string
		lt1, lt2 int
	)

	if st1, err = ToSlice(t1, "cmp"); err != nil {
		return ss, err
	}

	if st2, err = ToSlice(t2, "cmp"); err != nil {
		return ss, err
	}

	if lt1, lt2 = len(st1), len(st2); lt1 != lt2 {
		err = fmt.Errorf("field count: %d != %d", lt1, lt2)
		return ss, ErrorDecorator(err)
	}

	for i := 0; i < lt1; i++ {

		if st1[i][NameIx] != st2[i][NameIx] {
			err = fmt.Errorf("field name %d: %q != %q", i, st1[i][NameIx], st2[i][NameIx])
			return ss, ErrorDecorator(err)
		}

		if st1[i][ValueIx] != st2[i][ValueIx] {
			ss = append(ss, []string{st1[i][NameIx], st1[i][ValueIx], st2[i][ValueIx]})
		}
	}

	return ss, err
}

// ToSlice converts a single-tier struct into a slice of slices in the
// form {{name, value}, {name, value}, ...} for consumption by other methods.
// The outer slice maintains the fields in the same order as the struct. The
// tag parameter is the name of the struct tag to use for special processing.
// The primary purpose of this function is to offload tag processing for other
// functions.
func ToSlice(t interface{}, tag string) (ss[][]string, err error) {

	v := reflect.ValueOf(t)

	if v.Type().Kind() != reflect.Struct {
		err = fmt.Errorf("kind %q is not %q", v.Type().Kind().String(), "struct")
		return ss, ErrorDecorator(err)
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

	return ss, err
}
