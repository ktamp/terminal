// go.mkdef -w -pub=false terminal_unix.go sys_unix.go def_bsd.go
// MACHINE GENERATED BY go.mkdef (github.com/kless/cutil); DO NOT EDIT
//
// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs _zdef_openbsd_amd64.go

package terminal

const (
	_BRKINT     = 0x2
	_CS8        = 0x300
	_CSIZE      = 0x300
	_ECHO       = 0x8
	_ECHONL     = 0x10
	_ICANON     = 0x100
	_ICRNL      = 0x100
	_IEXTEN     = 0x400
	_IGNBRK     = 0x1
	_IGNCR      = 0x80
	_INLCR      = 0x40
	_ISIG       = 0x80
	_ISTRIP     = 0x20
	_IXON       = 0x200
	_OPOST      = 0x1
	_PARENB     = 0x1000
	_PARMRK     = 0x8
	_TCSADRAIN  = 0x1
	_TCSAFLUSH  = 0x2
	_TCSANOW    = 0x0
	_TCGETS     = 0x402c7413
	_TIOCGWINSZ = 0x40087468
	_TCSETS     = 0x802c7414
	_TCSETSF    = 0x802c7416
	_TCSETSW    = 0x802c7415
	_VMIN       = 0x10
	_VTIME      = 0x11
)

type termios struct {
	Iflag  uint32
	Oflag  uint32
	Cflag  uint32
	Lflag  uint32
	Cc     [20]uint8
	Ispeed int32
	Ospeed int32
}
type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}
