.POSIX:

PREFIX ?= /usr/local

gotml:
	go build

test: gotml
	./gotml test.md

install: gotml
	mkdir -p ${DESTDIR}${PREFIX}/bin
	cp -f $< ${DESTDIR}${PREFIX}/bin

unistall:
	rm -f ${DESTDIR}${PREFIX}/bin/gotml

clean:
	rm -f gotml

.PHONY: test install uninstall clean
