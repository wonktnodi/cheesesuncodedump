include $(GOROOT)/src/Make.inc

TARG=callback
CGOFILES=callback.go
LDPATH_linux=-Wl,-R,`pwd`
CGO_LDFLAGS=_cgo_export.o calladd.so funcwrap.so $(LDPATH_linux)
CGO_DEPS=_cgo_export.o calladd.so funcwrap.so

CLEANFILES+=main

include $(GOROOT)/src/Make.pkg

calladd.o: calladd.c
	gcc -g -c -fPIC calladd.c

calladd.so: calladd.o
	gcc -shared -o $@ calladd.o

funcwrap.o: funcwrap.c _cgo_export.h
	gcc $(_CGO_CFLAGS_$(GOARCH)) -g -c -fPIC $(CFLAGS) funcwrap.c

funcwrap.so: funcwrap.o
	gcc $(_CGO_CFLAGS_$(GOARCH)) -o $@ funcwrap.o $(_CGO_LDFLAGS_$(GOOS))

main: calladd.so install main.go
	$(GC) main.go
	$(LD) -o $@ main.$O

