all:
	go mod tidy && go build -o ./bin/fsfs .