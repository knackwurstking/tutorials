clean:
	@git clean -f -x -d

build:
	go build -o bin/masseverbrauch-rechner .

dev:
	nodemon -L --signal SIGTERM --exec 'go run .' --ext 'go,mod,sum' --delay 3 --ignore bin
