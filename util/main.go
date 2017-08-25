package main

import (
	//"github.com/jscherff/gocmdb/usbci"
	"github.com/jscherff/gocmdb/usbci/magtek"
	"github.com/google/gousb"
	"encoding/json"
	"flag"
	"log"
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "You must specify a mode of operation.\n")
		fsMode.Usage()
		os.Exit(1)
	}

	fsMode.Parse(os.Args[1:2])

	var flagset *flag.FlagSet

	switch {

	case *fModeReport:
		flagset = fsReport

	case *fModeConfig:
		flagset = fsConfig

	case *fModeReset:
		flagset = fsReset
	}

	if flagset.Parse(os.Args[2:]); flagset.NFlag() == 0 {
		fmt.Fprintf(os.Stderr, "You must specify at least one option.\n")
		flagset.Usage()
		os.Exit(1)
	}

	context := gousb.NewContext()
	defer context.Close()

	// Open devices that report a Magtek vendor ID, 0x0801.
	// We omit error checking on OpenDevices() because this
	// function terminates with 'libusb: not found [code -5]'
	// on Windows systems.

	devices, _ := context.OpenDevices(func(desc *gousb.DeviceDesc) bool {
		return uint16(desc.Vendor) == magtek.MagtekVendorID
	})

	if len(devices) == 0 {
		log.Fatalf("No Magtek devices found")
	}

	for _, device := range devices {

		defer device.Close()
		//device, err := usbci.NewDevice(device)
		device, err := magtek.NewDevice(device)
		fmt.Printf("%s\n\n", device)

		if err != nil {
			log.Fatalf("Error: %v", err); continue
		}


		di, errs := magtek.NewDeviceInfo(device)
		fmt.Printf("%v\n\n", di)

		if len(errs) > 0 {
			log.Fatalf("Errors encountered"); continue
		}

/*
		dx := []byte("<DeviceInfo><HostName>John-SurfacePro</HostName><VendorID>0801</VendorID><ProductID>0001</ProductID><VendorName>Mag-Tek</VendorName><ProductName>USB Swipe Reader</ProductName><ProductVer>V05</ProductVer><SoftwareID>21042840G01</SoftwareID><DeviceSN>B164F78</DeviceSN><FactorySN>B164F78022713AA</FactorySN><DescriptSN>B164F78</DescriptSN><BusNumber>1</BusNumber><BusAddress>29</BusAddress><USBSpec>1.10</USBSpec><USBClass>per-interface</USBClass><USBSubclass>per-interface</USBSubclass><USBProtocol>0</USBProtocol><DeviceSpeed>full</DeviceSpeed><DeviceVer>1.00</DeviceVer><MaxPktSize>8</MaxPktSize><BufferSize>60</BufferSize></DeviceInfo>")

		di, err := magtek.NewDeviceInfoFromXML(dx)
		fmt.Println(di)

		di.SoftwareID = ""
*/
/*
		b, _ := di.JSON(true)
		fmt.Println(string(b) + "\n")

		b, _ = di.XML(true)
		fmt.Println(string(b) + "\n")

		s, _ := di.CSV(true)
		fmt.Println(s + "\n")
*/
		j := "{\"HostName\":\"John-SurfacePro\",\"DeviceSN\":\"B164F78\",\"VendorID\":\"0801\",\"ProductID\":\"0001\",\"SoftwareID\":\"21042840G01\",\"VendorName\":\"Mag-Tek\",\"ProductName\":\"USB Swipe Reader\",\"ProductVer\":\"V05\",\"FactorySN\":\"B164F78022713AA\",\"DescriptSN\":\"B164F78\",\"BusNumber\":\"1\",\"BusAddress\":\"4\",\"USBSpec\":\"1.10\",\"USBClass\":\"per-interface\",\"USBSubclass\":\"per-interface\",\"USBProtocol\":\"0\",\"DeviceSpeed\":\"full\",\"DeviceVer\":\"1.00\",\"MaxPktSize\":\"8\",\"BufferSize\":\"60\"}"

		//j := "{\"HostName\":\"John-SurfacePro\",\"DeviceSN\":\"B164F78\",\"VendorID\":\"0801\",\"ProductID\":\"0001\",\"SoftwareID\":\"21042840G01\",\"ProductName\":\"USB Swipe Reader\",\"ProductVer\":\"V05\",\"FactorySN\":\"B164F78022713AA\",\"DescriptSN\":\"B164F78\",\"BusNumber\":\"1\",\"BusAddress\":\"4\",\"USBSpec\":\"1.10\",\"USBClass\":\"per-interface\",\"USBSubclass\":\"per-interface\",\"USBProtocol\":\"0\",\"DeviceSpeed\":\"full\",\"DeviceVer\":\"1.00\",\"MaxPktSize\":\"8\",\"BufferSize\":\"60\"}"

		di2 := new(magtek.DeviceInfo)
		if err = json.Unmarshal([]byte(j), &di2); err != nil {log.Fatalf("%v", err)}

		di3, err := di.Copy(true)
		fmt.Printf("A COPY:\n%s\n\n", di3)

		di4, err := di3.Copy(false)
		fmt.Printf("A COPY OF A COPY:\n%s\n\n", di4)

		fmt.Println(di4.Matches(di3))

		b, _ := di.JSON(false)
		fmt.Println(string(b) + "\n")

		b, _ = di.XML(false)
		fmt.Println(string(b) + "\n")

		s, _ := di.CSV(true)
		fmt.Println(s + "\n")

		s, _ = di.NVP(true)
		fmt.Println(s + "\n")

		fmt.Println(di.Matches(di2))
		os.Exit(0)

		switch {

		case *fModeReport:
			err = report(device)

		case *fModeConfig:
			err = config(device)

		case *fModeReset:
			err = reset(device)
		}
	}
}
