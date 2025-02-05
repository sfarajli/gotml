.POSIX:

PREFIX ?= /usr/local

gotml: main.go
	go build

test: gotml
	./gotml test.md

install: gotml
	mkdir -p ${DESTDIR}${PREFIX}/bin
	cp -f $< ${DESTDIR}${PREFIX}/bin

uninstall:
	rm -f ${DESTDIR}${PREFIX}/bin/gotml

clean:
	rm -f gotml

.PHONY: test install uninstall clean
