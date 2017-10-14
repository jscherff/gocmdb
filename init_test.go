package gocmdb

import (
	`crypto/sha256`
	`log`
	`testing`
	`github.com/google/gousb`
	`github.com/jscherff/gocmdb/usbci`
)

func init() {

	magChanges[0] = []string{`SoftwareID`, `21042840G01`, `21042840G02`}
	magChanges[1] = []string{`USBSpec`, `1.10`, `2.00`}

	if err := createObjects(); err != nil {
		log.Fatal(err)
	}

	if err := generateSigs(); err != nil {
		log.Fatal(err)
	}
}

func createObjects() (err error) {

	for k, j := range magJSON {

		if mag[k], err = usbci.NewMagtek(nil); err != nil {
			return err
		}
		if err := mag[k].RestoreJSON(j); err != nil {
			return err
		}
	}

	for k, j := range genJSON {

		if gen[k], err = usbci.NewGeneric(nil); err != nil {
			return err
		}
		if err := gen[k].RestoreJSON(j); err != nil {
			return err
		}
	}

	return nil
}

func generateSigs() error {

	for k := range magJSON {

		if b, err := mag[k].CSV(); err != nil {
			return err
		} else {
			sigCSV[k] = sha256.Sum256(b)
		}
		if b, err := mag[k].NVP(); err != nil {
			return err
		} else {
			sigNVP[k] = sha256.Sum256(b)
		}
		if b, err := mag[k].XML(); err != nil {
			return err
		} else {
			sigXML[k] = sha256.Sum256(b)
		}
		if b, err := mag[k].JSON(); err != nil {
			return err
		} else {
			sigJSON[k] = sha256.Sum256(b)
		}
		if b, err := mag[k].PrettyXML(); err != nil {
			return err
		} else {
			sigPrettyXML[k] = sha256.Sum256(b)
		}
		if b, err := mag[k].PrettyJSON(); err != nil {
			return err
		} else {
			sigPrettyJSON[k] = sha256.Sum256(b)
		}

		b := mag[k].Legacy()
		sigLegacy[k] = sha256.Sum256(b)
	}

	for k := range genJSON {

		if b, err := gen[k].CSV(); err != nil {
			return err
		} else {
			sigCSV[k] = sha256.Sum256(b)
		}
		if b, err := gen[k].NVP(); err != nil {
			return err
		} else {
			sigNVP[k] = sha256.Sum256(b)
		}
		if b, err := gen[k].XML(); err != nil {
			return err
		} else {
			sigXML[k] = sha256.Sum256(b)
		}
		if b, err := gen[k].JSON(); err != nil {
			return err
		} else {
			sigJSON[k] = sha256.Sum256(b)
		}
		if b, err := gen[k].PrettyXML(); err != nil {
			return err
		} else {
			sigPrettyXML[k] = sha256.Sum256(b)
		}
		if b, err := gen[k].PrettyJSON(); err != nil {
			return err
		} else {
			sigPrettyJSON[k] = sha256.Sum256(b)
		}

		b := gen[k].Legacy()
		sigLegacy[k] = sha256.Sum256(b)
	}

	return nil
}

func restoreState(tb testing.TB) {

	tb.Helper()

	if err := createObjects(); err != nil {
		tb.Fatal(err)
	}
}

func getMagtekDevice(tb testing.TB, c *gousb.Context) (mdev *usbci.Magtek, err error) {

	tb.Helper()

	dev, err := c.OpenDeviceWithVIDPID(0x0801, 0x0001)

	if dev != nil {
		mdev, err = usbci.NewMagtek(dev)
	}

	return mdev, err
}

