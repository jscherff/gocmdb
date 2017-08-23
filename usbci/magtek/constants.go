package magtek

const (
	MagtekVendorID uint16 = 0x0801

	SureswipeKbPID uint16 = 0x0001
	SureswipeHidPID uint16 = 0x0002
	MagnesafeSwipeKbPID uint16 = 0x0001
	MagnesafeInsertKbPID uint16 = 0x0001
	MagnesafeSwipeHidPID uint16 = 0x0011
	MagnesafeInsertHidPID uint16 = 0x0013
	MagnesafeWirelessHidPID uint16 = 0x0014

	BufferSizeSureswipe int = 24
	BufferSizeMagnesafe int = 60

	CommandGetProp uint8 = 0x00
	CommandSetProp uint8 = 0x01
	CommandResetDevice uint8 = 0x02

	ResultCodeSuccess uint8 = 0x00
	ResultCodeFailure uint8 = 0x01
	ResultCodeBadParam uint8 = 0x02

	PropSoftwareID uint8 = 0x00
	PropDeviceSN uint8 = 0x01
	PropFactorySN uint8 = 0x03
	PropProductVer uint8 = 0x04

	DefaultSNLength int = 7
)
