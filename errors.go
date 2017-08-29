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
	"strings"
	"fmt"
)

// GetterError holds a collection of errors encountered when calling several
// getter methods while populating data structures.
type GetterError struct {
	Getters	[]string
	Errors	[]error
}

// Add adds a new getter error to GetterError.
func (e *GetterError) Add(gs string, ge error) {
	e.Getters = append(e.Getters, gs)
	e.Errors = append(e.Errors, ge)
}

// Error implements the Error method of the error interface.
func (e *GetterError) Error() (s string) {
	return fmt.Sprintf("getter errors: %s\n", strings.Join(e.Getters, ", "))
}
