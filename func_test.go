package gocmdb

import (
	`crypto/sha256`
	`fmt`
	`os`
	`path/filepath`
	`reflect`
	`testing`
	`github.com/google/gousb`
	`github.com/jscherff/gocmdb/usbci`
	`github.com/jscherff/gotest`
)

func TestGetterMethods(t *testing.T) {

	gotest.Assert(t, mag[`mag1`].ID() == mag[`mag1`].SerialNum, `ID() does not match (device).SerialNum`)
	gotest.Assert(t, mag[`mag1`].VID() == mag[`mag1`].VendorID, `VID() does not match (device).VenndorID`)
	gotest.Assert(t, mag[`mag1`].PID() == mag[`mag1`].ProductID, `PID() does not match (device).ProductID`)
	gotest.Assert(t, mag[`mag1`].Type() == fmt.Sprintf(`%T`, mag[`mag1`]), `Type does not match TypeOf(device)`)

	if hostName, err := os.Hostname(); err != nil {
		return
	} else {
		gotest.Assert(t, mag[`mag1`].Host() == hostName, `Host() does not match os.Hostname()`)
	}
}

func TestFilenameMethod(t *testing.T) {

	fileName := fmt.Sprintf(`%03d-%03d-%s-%s`,
                mag[`mag1`].PortNumber,
                mag[`mag1`].BusNumber,
                mag[`mag1`].VendorID,
                mag[`mag1`].ProductID,
        )

	gotest.Assert(t, mag[`mag1`].Filename() == fileName, `(device).Filename() string incorrect`)
}

func TestReportMethods(t *testing.T) {

	var (
		b []byte
		err error
	)

	b, err = mag[`mag1`].PrettyJSON()
	gotest.Ok(t, err)
	gotest.Assert(t, sha256.Sum256(b) == sigPrettyJSON[`mag1`], `unexpected hash signature of JSON output`)

	b, err = mag[`mag1`].JSON()
	gotest.Ok(t, err)
	gotest.Assert(t, sha256.Sum256(b) == sigJSON[`mag1`], `unexpected hash signature of NVP output`)

	b, err = mag[`mag1`].PrettyXML()
	gotest.Ok(t, err)
	gotest.Assert(t, sha256.Sum256(b) == sigPrettyXML[`mag1`], `unexpected hash signature of XML output`)

	b, err = mag[`mag1`].XML()
	gotest.Ok(t, err)
	gotest.Assert(t, sha256.Sum256(b) == sigXML[`mag1`], `unexpected hash signature of NVP output`)

	b, err = mag[`mag1`].CSV()
	gotest.Ok(t, err)
	gotest.Assert(t, sha256.Sum256(b) == sigCSV[`mag1`], `unexpected hash signature of CSV output`)

	b, err = mag[`mag1`].NVP()
	gotest.Ok(t, err)
	gotest.Assert(t, sha256.Sum256(b) == sigNVP[`mag1`], `unexpected hash signature of NVP output`)

	b = mag[`mag1`].Legacy()
	gotest.Ok(t, err)
	gotest.Assert(t, sha256.Sum256(b) == sigLegacy[`mag1`], `unexpected hash signature of NVP output`)
}

func TestPersistenceMethods(t *testing.T) {

	var err error

	t.Run("Save() and RestoreFile()", func(t *testing.T) {

		fn := filepath.Join(os.Getenv(`TEMP`), `mag1.json`)

		err = mag[`mag1`].Save(fn)
		gotest.Ok(t, err)

		mag3, err := usbci.NewMagtek(nil)
		gotest.Ok(t, err)

		err = mag3.RestoreFile(fn)
		gotest.Ok(t, err)

		gotest.Assert(t, reflect.DeepEqual(mag[`mag1`], mag3), `restored device not identical to saved device`)
	})

	t.Run("JSON() and RestoreJSON()", func(t *testing.T) {

		j, err := mag[`mag1`].JSON()
		gotest.Ok(t, err)

		mag3, err := usbci.NewMagtek(nil)
		gotest.Ok(t, err)

		err = mag3.RestoreJSON(j)
		gotest.Ok(t, err)

		gotest.Assert(t, reflect.DeepEqual(mag[`mag1`], mag3), `restored device not identical to saved device`)
	})
}

func TestCompareMethods(t *testing.T) {

	t.Run("Save() and CompareFile()", func(t *testing.T) {

		mag3, err := usbci.NewMagtek(nil)
		gotest.Ok(t, err)

		err = mag3.RestoreJSON(magJSON[`mag1`])
		gotest.Ok(t, err)

		fn1 := filepath.Join(os.Getenv(`TEMP`), `mag1.json`)
		fn2 := filepath.Join(os.Getenv(`TEMP`), `mag2.json`)

		err = mag[`mag1`].Save(fn1)
		gotest.Ok(t, err)

		err = mag[`mag2`].Save(fn2)
		gotest.Ok(t, err)

		ss1, err := mag3.CompareFile(fn1)
		gotest.Ok(t, err)
		gotest.Assert(t, len(ss1) == 0, `device should match its clone`)

		ss2, err := mag3.CompareFile(fn2)
		gotest.Ok(t, err)
		gotest.Assert(t, len(ss2) == 2, `modified device should not match original`)
	})

	t.Run("JSON() and CompareJSON()", func(t *testing.T) {

		mag3, err := usbci.NewMagtek(nil)
		gotest.Ok(t, err)

		err = mag3.RestoreJSON(magJSON[`mag1`])
		gotest.Ok(t, err)

		j1, err := mag[`mag1`].JSON()
		gotest.Ok(t, err)

		j2, err := mag[`mag2`].JSON()
		gotest.Ok(t, err)

		ss1, err := mag3.CompareJSON(j1)
		gotest.Ok(t, err)
		gotest.Assert(t, len(ss1) == 0, `device should match its clone`)

		ss2, err := mag3.CompareJSON(j2)
		gotest.Ok(t, err)
		gotest.Assert(t, len(ss2) == 2, `modified device should not match original`)
	})
}

func TestAuditMethods(t *testing.T) {

	t.Run("Save() and AuditFile()", func(t *testing.T) {

		mag3, err := usbci.NewMagtek(nil)
		gotest.Ok(t, err)

		err = mag3.RestoreJSON(magJSON[`mag2`])
		gotest.Ok(t, err)

		fn1 := filepath.Join(os.Getenv(`TEMP`), `mag1.json`)
		fn2 := filepath.Join(os.Getenv(`TEMP`), `mag2.json`)

		err = mag[`mag1`].Save(fn1)
		gotest.Ok(t, err)

		err = mag[`mag2`].Save(fn2)
		gotest.Ok(t, err)

		err = mag3.AuditFile(fn2)
		gotest.Ok(t, err)
		gotest.Assert(t, len(mag3.Changes) == 0, `device should match its clone`)

		err = mag3.AuditFile(fn1)
		gotest.Ok(t, err)
		gotest.Assert(t, len(mag3.Changes) == 2, `modified device should not match original`)

		if len(mag3.Changes) < 2 { return }

		gotest.Assert(t, reflect.DeepEqual(mag3.Changes, magChanges),
			`(device).Changes contains bad data`)
		gotest.Assert(t, reflect.DeepEqual(mag3.GetChanges(), magChanges),
			`(device).GetChanges() returns bad data`)
	})

	t.Run("JSON() and AuditJSON()", func(t *testing.T) {

		mag3, err := usbci.NewMagtek(nil)
		gotest.Ok(t, err)

		err = mag3.RestoreJSON(magJSON[`mag2`])
		gotest.Ok(t, err)

		j1, err := mag[`mag1`].JSON()
		gotest.Ok(t, err)

		j2, err := mag[`mag2`].JSON()
		gotest.Ok(t, err)

		err = mag3.AuditJSON(j2)
		gotest.Ok(t, err)
		gotest.Assert(t, len(mag3.Changes) == 0, `device should match its clone`)

		err = mag3.AuditJSON(j1)
		gotest.Ok(t, err)
		gotest.Assert(t, len(mag3.Changes) == 2, `modified device should not match original`)

		if len(mag3.Changes) < 2 { return }

		gotest.Assert(t, reflect.DeepEqual(mag3.Changes, magChanges),
			`(device).Changes contains bad data`)
		gotest.Assert(t, reflect.DeepEqual(mag3.GetChanges(), magChanges),
			`(device).GetChanges() returns bad data`)
	})
}

func TestChangeMethods(t *testing.T) {

	t.Run("AddChange() and GetChanges()", func(t *testing.T) {

		var changes = []string{`SoftwareID`, `21042818B01`, `21042818B03`}

		mag3, err := usbci.NewMagtek(nil)
		gotest.Ok(t, err)

		err = mag3.RestoreJSON(magJSON[`mag1`])
		gotest.Ok(t, err)

		mag3.AddChange(`SoftwareID`, `21042818B01`, `21042818B03`)
		gotest.Assert(t, len(mag3.Changes) == 1, `(device).Changes should contain one change`)
		gotest.Assert(t, len(mag3.GetChanges()) == 1, `(device).GetChanges() should contain one change`)

		if len(mag3.Changes) < 1 { return }

		gotest.Assert(t, reflect.DeepEqual(mag3.GetChanges()[0], changes),
			`(device).GetChanges() returns bad data`)
		gotest.Assert(t, reflect.DeepEqual(mag3.Changes[0], changes),
			`(device).Changes contains bad data`)
	})

	t.Run("SetChanges() and GetChanges()", func(t *testing.T) {

		mag3, err := usbci.NewMagtek(nil)
		gotest.Ok(t, err)

		err = mag3.RestoreJSON(magJSON[`mag2`])
		gotest.Ok(t, err)

		ss, err := mag3.CompareJSON(magJSON[`mag1`])
		gotest.Ok(t, err)

		mag3.SetChanges(ss)
		gotest.Assert(t, len(mag3.Changes) == 2, `(device).Changes should contain two changes`)
		gotest.Assert(t, len(mag3.GetChanges()) == 2, `(device).GetChanges() should contain two changes`)

		if len(mag3.Changes) < 2 { return }

		gotest.Assert(t, reflect.DeepEqual(mag3.Changes, magChanges),
			`(device).Changes contains bad data`)
		gotest.Assert(t, reflect.DeepEqual(mag3.GetChanges(), magChanges),
			`(device).GetChanges() returns bad data`)
	})
}

func TestSerialMethods(t *testing.T) {

	t.Run("magtek Sureswipe Card Reader", func(t *testing.T) {

		ctx := gousb.NewContext()
		defer ctx.Close()

		dev, err := ctx.OpenDeviceWithVIDPID(0x0801, 0x0001)

		if err == gousb.ErrorNoDevice || err == gousb.ErrorNotFound {
			t.Skip(`no compatible devices found`)
		}

		gotest.Ok(t, err)
		defer dev.Close()

		// Set device SN

		mdev, err := usbci.NewMagtek(dev)
		gotest.Ok(t, err)

		oldSn, err := mdev.GetDeviceSN()
		gotest.Ok(t, err)

		err = mdev.SetDeviceSN(`TESTING`)
		gotest.Ok(t, err)

		newSn, err := mdev.GetDeviceSN()
		gotest.Ok(t, err)
		gotest.Assert(t, newSn == `TESTING`, `setting device SN to new value unsuccessful`)

		errs := mdev.Refresh()
		gotest.Assert(t, len(errs) == 0, `(device).Refresh() produced setter errors`)

		// Erase device SN

		err = mdev.EraseDeviceSN()
		gotest.Ok(t, err)

		newSn, err = mdev.GetDeviceSN()
		gotest.Ok(t, err)
		gotest.Assert(t, newSn == ``, `erasing device SN was unsuccessful`)

		// Restore device SN

		err = mdev.SetDeviceSN(oldSn)
		gotest.Ok(t, err)
		newSn, err = mdev.GetDeviceSN()
		gotest.Ok(t, err)
		gotest.Assert(t, newSn == oldSn, `restoring device SN to previous value unsuccessful`)

		err = mdev.Reset()
		gotest.Ok(t, err)
	})

	t.Run("magtek Dynamag Card Reader", func(t *testing.T) {

		ctx := gousb.NewContext()
		defer ctx.Close()

		dev, err := ctx.OpenDeviceWithVIDPID(0x0801, 0x0001)

		if err == gousb.ErrorNoDevice || err == gousb.ErrorNotFound {
			t.Skip(`no compatible devices found`)
		}

		gotest.Ok(t, err)
		defer dev.Close()

		mdev, err := usbci.NewMagtek(dev)
		gotest.Ok(t, err)
		gotest.Assert(t, mdev.FactorySN != ``, `device does not have a factory SN`)

		oldSn, err := mdev.GetDeviceSN()
		gotest.Ok(t, err)

		err = mdev.CopyFactorySN(7)
		gotest.Ok(t, err)
		gotest.Assert(t, mdev.DeviceSN == mdev.FactorySN[:7], `copying factory SN to device SN unsuccessful`)

		err = mdev.SetDeviceSN(oldSn)
		gotest.Ok(t, err)
		newSn, err := mdev.GetDeviceSN()
		gotest.Ok(t, err)
		gotest.Assert(t, newSn == oldSn, `restoring device SN to previous value unsuccessful`)

		err = mdev.Reset()
		gotest.Ok(t, err)
	})
}
