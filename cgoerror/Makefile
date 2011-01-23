include $(GOROOT)/src/Make.inc

TARG=cgotest
CGOFILES=cgotest.go
CGO_CFLAGS=-static
#CGO_LDFLAGS=-L. -ltestlib
#CGO_OFILES=$(shell ar t libtestlib.a)
CGO_OFILES=testlib.o

include $(GOROOT)/src/Make.pkg

CLEANFILES+=main

testlib: testlib.c testlib.h
	gcc -c -fPIC testlib.c -o testlib.o
	ar rcs libtestlib.a testlib.o

extract: libtestlib.a
	mkdir _objs
	cd _objs
	ar x ../libtestlib.a
	cd ..

main: testlib install main.go
	$(GC) main.go
	$(LD) -o $@ main.$O

