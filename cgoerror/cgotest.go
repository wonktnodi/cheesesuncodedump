package cgotest

// #include "testlib.h"
import "C"

func GetStuff() int {
    return int(C.GetStuff())
}

