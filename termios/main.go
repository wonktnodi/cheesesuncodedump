package main

import (
    "fmt"
    "os"
)

func main() {
    var (
        s *winsize
        t *termios
        err os.Error
    )

    defer func() {
        if err != nil { fmt.Println(err) }
        if err = defaultTermios.set(); err != nil { fmt.Println(err) }
        fmt.Print("\n")
    }()

    if s, err = getWinsize(); err != nil { return }
    fmt.Printf("Window Size:\n\tLines: %d\n\tColumns: %d\n", s.ws_row, s.ws_col)

    fmt.Println("Entering Raw mode. . .")
    fmt.Println("Type some characters!  Press 'q' to quit!")

    if t, err = getTermios(); err != nil { return }
    if err = t.setRaw(); err != nil { return }

    var buff [1]byte
    for buff[0] != 'q' {
        if _, err = os.Stdin.Read(buff[:]); err != nil { return }
        fmt.Printf("%c", buff[0])
    }
}
