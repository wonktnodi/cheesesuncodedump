package main

import (
    "fmt"
    "os"
    "syscall"
    "unsafe"
)

const NCCS = 32

type (
    cc_t     byte
    speed_t  uint
    tcflag_t uint

    termios struct {
        c_iflag, c_oflag, c_cflag, c_lflag tcflag_t
        c_line                             cc_t
        c_cc                               [NCCS]cc_t
        c_ispeed, c_ospeed                 speed_t
    }

    winsize struct {
        ws_row, ws_col       uint16
        ws_xpixel, ws_ypixel uint16
    }
)

// termios constants
const (
    t_IGNBRK = tcflag_t(0000001)
    t_BRKINT = tcflag_t(0000002)
    t_PARMRK = tcflag_t(0000010)
    t_ISTRIP = tcflag_t(0000040)
    t_INLCR  = tcflag_t(0000100)
    t_IGNCR  = tcflag_t(0000200)
    t_ICRNL  = tcflag_t(0000400)
    t_IXON   = tcflag_t(0002000)
    t_OPOST  = tcflag_t(0000001)
    t_ECHO   = tcflag_t(0000010)
    t_ECHONL = tcflag_t(0000100)
    t_ICANON = tcflag_t(0000002)
    t_ISIG   = tcflag_t(0000001)
    t_IEXTEN = tcflag_t(0100000)
    t_CSIZE  = tcflag_t(0000060)
    t_CS8    = tcflag_t(0000060)
    t_PARENB = tcflag_t(0000400)
    t_VTIME  = 5
    t_VMIN   = 6
)

// ioctl constants
const (
    i_TCGETS     = 0x5401
    i_TCSETS     = 0x5402
    i_TIOCGWINSZ = 0x5413
    i_TIOCSWINSZ = 0x5414
)

var defaultTermios *termios

func init() {
    var err os.Error
    if defaultTermios, err = getTermios(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func getTermios() (*termios, os.Error) {
    term := new(termios)
    if err := term.get(); err != nil {
        return nil, err
    }

    return term, nil
}

func (self *termios) get() os.Error {
    r1, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
        uintptr(0), uintptr(i_TCGETS),
        uintptr(unsafe.Pointer(self)))
    if err := os.NewSyscallError("SYS_IOCTL", int(errno)); err != nil {
        return err
    }

    if err := os.NewSyscallError("SYS_IOCTL", int(r1)); err != nil {
        return err
    }

    return nil
}

func (self *termios) set() os.Error {
    r1, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
        uintptr(0), uintptr(i_TCSETS),
        uintptr(unsafe.Pointer(self)))
    if err := os.NewSyscallError("SYS_IOCTL", int(errno)); err != nil {
        return err
    }

    if err := os.NewSyscallError("SYS_IOCTL", int(r1)); err != nil {
        return err
    }

    return nil
}

func (self *termios) setRaw() os.Error {
    self.c_iflag &= ^(t_IGNBRK | t_BRKINT | t_PARMRK | t_ISTRIP |
        t_INLCR | t_IGNCR | t_ICRNL | t_IXON)
    self.c_oflag &= ^t_OPOST
    self.c_lflag &= ^(t_ECHO | t_ECHONL | t_ICANON | t_ISIG | t_IEXTEN)
    self.c_cflag &= ^(t_CSIZE | t_PARENB)
    self.c_cflag |= t_CS8

    self.c_cc[t_VMIN] = 1
    self.c_cc[t_VTIME] = 0

    return self.set()
}

func getWinsize() (*winsize, os.Error) {
    size := new(winsize)
    if err := size.get(); err != nil {
        return nil, err
    }

    return size, nil
}

func (self *winsize) get() os.Error {
    r1, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
        uintptr(0), uintptr(i_TIOCGWINSZ),
        uintptr(unsafe.Pointer(self)))
    if err := os.NewSyscallError("SYS_IOCTL", int(errno)); err != nil {
        return err
    }
    if err := os.NewSyscallError("SYS_IOCTL", int(r1)); err != nil {
        return err
    }

    return nil
}

func (self *winsize) set() os.Error {
    r1, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
        uintptr(0), uintptr(i_TIOCSWINSZ),
        uintptr(unsafe.Pointer(self)))
    if err := os.NewSyscallError("SYS_IOCTL", int(errno)); err != nil {
        return err
    }
    if err := os.NewSyscallError("SYS_IOCTL", int(r1)); err != nil {
        return err
    }

    return nil
}

