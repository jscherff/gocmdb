package gocmdb

import (
	`log`
	`github.com/jscherff/gocmdb/usbci`
)

var (
	mag1, mag2 *usbci.Magtek
	gen1, gen2 *usbci.Generic

	magChanges = make([][]string, 2)
)

func init() {

	magChanges[0] = []string{`SoftwareID`, `21042840G01`, `21042840G02`}
	magChanges[1] = []string{`USBSpec`, `1.10`, `2.00`}

	var errM1, errM2, errG1, errG2 error

	if mag1, errM1 = usbci.NewMagtek(nil); errM1 == nil {
		errM1 = mag1.RestoreJSON(mag1JSON)
	}

	if mag2, errM2 = usbci.NewMagtek(nil); errM2 == nil {
		errM2 = mag2.RestoreJSON(mag2JSON)
	}

	if gen1, errG1 = usbci.NewGeneric(nil); errG1 == nil {
		errG1 = gen1.RestoreJSON(gen1JSON)
	}

	if gen2, errG2 = usbci.NewGeneric(nil); errG2 == nil {
		errG2 = gen2.RestoreJSON(gen2JSON)
	}

	if errM1 != nil || errM2 != nil || errG1 != nil || errG2 != nil {
		log.Fatalln(`Testing setup failed: could not restore devices.`)
	}
}
