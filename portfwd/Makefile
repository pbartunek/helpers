CC=go build
LDFLAGS="-s -w"
ODIR=dist

.SILENT: linux win all clean
.PHONY: clean

win32:
	GOOS=windows GOARCH=386 $(CC) -o $(ODIR)/portfwd.exe portfwd.go

win:
	GOOS=windows GOARCH=amd64 $(CC) -o $(ODIR)/portfwd64.exe portfwd.go

linux:
	GOOS=linux GOARCH=amd64 $(CC) -o $(ODIR)/portfwd portfwd.go

macos:
	GOOS=darwin GOARCH=amd64 $(CC) -o $(ODIR)/portfwd_macos portfwd.go

all: win32 win linux macos

clean:
	rm -rf dist
