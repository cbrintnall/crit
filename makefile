FILES=src/*

all:
	go build -o crit $(FILES)

install: all
	mv crit /usr/local/bin