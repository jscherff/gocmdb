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
	ID() (string)
	Type() (string)
	Bare() ([]byte)
	JSON() ([]byte, error)
	XML() ([]byte, error)
	CSV() ([]byte, error)
	NVP() ([]byte, error)
	Filename() (string)
}

type Auditable interface {
	ID() (string)
	Type() (string)
	Save(string)
	Restore(string)
	Matches(interface{})
	Compare(string) ([][]string, error)
	Audit(string) (error)
}

type Registerable interface {
	ID() (string)
	Type() (string)
	JSON() ([]byte, error)
}

type Configurable interface {
	GetDeviceSN(string) (error)
	SetDeviceSN(string) (error)
	EraseDeviceSN() (error)
	SetFactorySN(string) (error)
	CopyFactorySN(int) (error)
	Refresh() (error)
	Reset() (error)
}

type Resettable interface {
	Reset() (error)
}

type GenericUSB interface {
	ID() (string)
	Type() (string)
	Save(fn string)
	Restore(string)
	Matches(interface{})
	Compare(string) ([][]string, error)
	Audit(string) (error)

	Refresh() (error)
	Reset() (error)

	Bare() ([]byte)
	JSON() ([]byte, error)
	XML() ([]byte, error)
	CSV() ([]byte, error)
	NVP() ([]byte, error)
	Filename() (string)
}

type MagtekUSB interface {
	ID() (string)
	Type() (string)
	Save(fn string)
	Restore(string)
	Matches(interface{})
	Compare(string) ([][]string, error)
	Audit(string) (error)

	Refresh() (error)
	Reset() (error)

	Bare() ([]byte)
	JSON() ([]byte, error)
	XML() ([]byte, error)
	CSV() ([]byte, error)
	NVP() ([]byte, error)
	Filename() (string)

	GetDeviceSN(string) (error)
	SetDeviceSN(string) (error)
	EraseDeviceSN() (error)
	SetFactorySN(string) (error)
	CopyFactorySN(int) (error)
}

