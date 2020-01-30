FILES=src/*
BUILDS=builds

all: clean build-mac build-linux build-windows

clean:
	rm -rf ./builds

build-mac: $(FILES)
	GOOS=darwin GOARCH=386 go build -o $(BUILDS)/crit_darwin_386 $(FILES)
	GOOS=darwin GOARCH=amd64 go build -o $(BUILDS)/crit_darwin_amd64 $(FILES)

build-linux: $(FILES)
	GOOS=linux GOARCH=386 go build -o $(BUILDS)/crit_linux_386 $(FILES)
	GOOS=linux GOARCH=amd64 go build -o $(BUILDS)/crit_linux_amd64 $(FILES)

build-windows: $(FILES)
	GOOS=windows GOARCH=386 go build -o $(BUILDS)/crit_windows_386 $(FILES)
	GOOS=windows GOARCH=amd64 go build -o $(BUILDS)/crit_windows_amd64 $(FILES)