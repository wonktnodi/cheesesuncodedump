include $(GOROOT)/src/Make.inc

main: termios.go main.go
	$(GC) -o main.$O termios.go main.go
	$(LD) -o main main.$O

clean:
	@rm main main.6

all: clean main

