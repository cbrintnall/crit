FILES=src/main.go src/secrets.go

all:
	go build $(FILES)
	mv main crit

install: all
	mv crit /usr/local/bin