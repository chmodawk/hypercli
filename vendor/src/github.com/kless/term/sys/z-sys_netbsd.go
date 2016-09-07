// go.mkdef -w sys_unix.go sys_bsd.go
// MACHINE GENERATED BY go.mkdef (github.com/kless/gotool/go.mkdef); DO NOT EDIT
//
// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs _z-sys_netbsd_amd64.go

package sys

const (
	TCSADRAIN = 0x1
	TCSAFLUSH = 0x2
	TCSANOW   = 0x0
)

const TIOCGWINSZ = 0x40087468

const (
	VDISCARD = 0xf
	VEOF     = 0x0
	VEOL     = 0x1
	VEOL2    = 0x2
	VERASE   = 0x3
	VINTR    = 0x8
	VKILL    = 0x5
	VLNEXT   = 0xe
	VMIN     = 0x10
	VQUIT    = 0x9
	VREPRINT = 0x6
	VSTART   = 0xc
	VSTOP    = 0xd
	VSUSP    = 0xa
	VTIME    = 0x11
	VWERASE  = 0x4
)

const (
	BRKINT  = 0x2
	ICRNL   = 0x100
	IGNBRK  = 0x1
	IGNCR   = 0x80
	IGNPAR  = 0x4
	IMAXBEL = 0x2000
	INLCR   = 0x40
	INPCK   = 0x10
	ISTRIP  = 0x20
	IXANY   = 0x800
	IXOFF   = 0x400
	IXON    = 0x200
	PARMRK  = 0x8
)

const (
	BS0    = 0x0
	BS1    = 0x8000
	CR0    = 0x0
	CR1    = 0x1000
	CR2    = 0x2000
	CR3    = 0x3000
	FF0    = 0x0
	FF1    = 0x4000
	NL0    = 0x0
	NL1    = 0x100
	OCRNL  = 0x10
	ONLCR  = 0x2
	ONLRET = 0x40
	ONOCR  = 0x20
	OPOST  = 0x1
	TAB0   = 0x0
	TAB1   = 0x400
	TAB2   = 0x800
	XTABS  = 0xc00
)

const (
	B0      = 0x0
	B110    = 0x6e
	B115200 = 0x1c200
	B1200   = 0x4b0
	B134    = 0x86
	B150    = 0x96
	B1800   = 0x708
	B19200  = 0x4b00
	B200    = 0xc8
	B230400 = 0x38400
	B2400   = 0x960
	B300    = 0x12c
	B38400  = 0x9600
	B4800   = 0x12c0
	B50     = 0x32
	B57600  = 0xe100
	B600    = 0x258
	B75     = 0x4b
	B9600   = 0x2580
	CLOCAL  = 0x8000
	CREAD   = 0x800
	CRTSCTS = 0x10000
	CS5     = 0x0
	CS6     = 0x100
	CS7     = 0x200
	CS8     = 0x300
	CSIZE   = 0x300
	CSTOPB  = 0x400
	EXTA    = 0x4b00
	EXTB    = 0x9600
	HUPCL   = 0x4000
	PARENB  = 0x1000
	PARODD  = 0x2000
)

const (
	ECHO    = 0x8
	ECHOCTL = 0x40
	ECHOE   = 0x2
	ECHOK   = 0x4
	ECHOKE  = 0x1
	ECHONL  = 0x10
	ECHOPRT = 0x20
	EXTPROC = 0x800
	FLUSHO  = 0x800000
	ICANON  = 0x100
	IEXTEN  = 0x400
	ISIG    = 0x80
	NOFLSH  = 0x80000000
	PENDIN  = 0x20000000
	TOSTOP  = 0x400000
)

const TCGETS = 0x402c7413

const TCSETS = 0x802c7414

const TCSETSW = 0x802c7415

const TCSETSF = 0x802c7416

type Termios struct {
	Iflag  uint32
	Oflag  uint32
	Cflag  uint32
	Lflag  uint32
	Cc     [20]uint8
	Ispeed int32
	Ospeed int32
}

type Winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}
