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

type Identifiable interface {
	ID() (string)
	VID() (string)
	PID() (string)
	Host() (string)
	Type() (string)
}

type Reportable interface {
	JSON() ([]byte, error)
	XML() ([]byte, error)
	CSV() ([]byte, error)
	NVP() ([]byte, error)
	Legacy() ([]byte)
	PrettyJSON() ([]byte, error)
	PrettyXML() ([]byte, error)
	Filename() (string)
}

type Comparable interface {
	Save(string) (error)
	RestoreFile(string) (error)
	RestoreJSON([]byte) (error)
	CompareFile(string) ([][]string, error)
	CompareJSON([]byte) ([][]string, error)
	AuditFile(string) (error)
	AuditJSON([]byte) (error)
	AddChange(string, string, string)
	SetChanges([][]string)
	GetChanges() ([][]string)
}

type Resettable interface {
	Refresh() (map[string]bool)
	Reset() (error)
}

type Registerable interface {
	Identifiable
	Reportable
}

type Auditable interface {
	Identifiable
	Reportable
	Comparable
}

type Configurable interface {
	Identifiable
	Reportable
	Comparable
	Resettable
	GetDeviceSN() (string, error)
	SetDeviceSN(string) (error)
	EraseDeviceSN() (error)
	SetFactorySN(string) (error)
	CopyFactorySN(int) (error)
}

type GenericUSB interface {
	Auditable
	Resettable
}

type MagtekUSB interface {
	Configurable
}
