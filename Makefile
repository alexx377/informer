.PHONY: all clean pack

all: informer

pack: informer
	tar -czvf ./pack/informer.tar.gz --directory=bin/ informer informer.exe informer.json

clean:
	rm -rf ./bin/*
	rm -rf ./pack/*

informer: informer.go informer.json
	GOOS=linux GOARCH=386 go build -o ./bin/informer
	GOOS=windows GOARCH=386	go build -o ./bin/informer.exe
	cp -f ./informer.json ./bin/informer.json
