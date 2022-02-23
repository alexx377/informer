.PHONY: all clean pack

all: informer

pack: informer
	zip -9 -j -o ./pack/informer.zip ./bin/informer ./bin/informer.exe ./informer.json

clean:
	rm -rf ./bin/*
	rm -rf ./pack/*

informer: informer.go informer.json
	GOOS=linux GOARCH=386 go build -o ./bin/informer
	GOOS=windows GOARCH=386	go build -o ./bin/informer.exe
	cp -f ./informer.json ./bin/informer.json
