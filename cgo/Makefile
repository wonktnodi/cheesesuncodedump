include $(GOROOT)/src/Make.inc

TARG=gocurses
CGOFILES=gocurses.go
CGO_LDFLAGS=-lncurses

include $(GOROOT)/src/Make.pkg

CLEANFILES+=main

main: install main.go
		$(GC) main.go
		$(LD) -o $@ main.$O

