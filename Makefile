build:
	go build -o app src/*.go

run:
	go build -o app src/*.go && ./app