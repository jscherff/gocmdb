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
	ID() (string, error)
	CSV() ([]byte, error)
	NVP() ([]byte, error)
	XML() ([]byte, error)
	JSON() ([]byte, error)
	Bare() ([]byte)
}

type Auditable interface {
	ID() (string, error)
	Save(string) (error)
	Restore(string) (error)
	Compare(string) ([][]string, error)
}

type Registerable interface {
	Type() (string)
	ID() (string, error)
	JSON() ([]byte, error)
}

type Configurable interface {
	DeviceSN() (string, error)
	SetDeviceSN(string) (error)
	SetFactorySN(string) (error)
	CopyFactorySN(int) (error)
	EraseDeviceSN() (error)
	Reset() (error)
}

type Resettable interface {
	Reset() (error)
}
