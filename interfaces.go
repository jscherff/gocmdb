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

type Reportable interface {
	Bare() ([]byte)
	JSON() ([]byte, error)
	XML() ([]byte, error)
	CSV() ([]byte, error)
	NVP() ([]byte, error)
}

type Persistable interface {
	Save(string) (error)
	Restore(string) (error)
}

type Comparable interface {
	Matches(interface{}) (bool)
}

type Configurable interface {
	DeviceSN() (string, error)
	SetDeviceSN(string) (error)
	SetFactorySN(string) (error)
	CopyFactorySN(int) (error)
	EraseDeviceSN() (error)
}

type Resettable interface {
	Reset() (error)
}
