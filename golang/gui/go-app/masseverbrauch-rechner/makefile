clean:
	@git clean -f -x -d

build:
	@GOARCH=wasm GOOS=js go build -o web/app.wasm && go build -o bin/masseverbrauch-rechner .

run:
	@GOARCH=wasm GOOS=js go build -o web/app.wasm && go run .
