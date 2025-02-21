OUTPUT_EXE=bin/masseverbrauch-rechner

clean:
	@git clean -f -x -d \
		&& rm -rf dist # ui-*.css file won't be removed

pages:
	@rm -rf dist \
		&& GOARCH=wasm GOOS=js go build -o web/app.wasm \
		&& go run . --generate-pages \
		&& cp -r web dist/

build:
	@GOARCH=wasm GOOS=js go build -o web/app.wasm \
		&& go build -o ${OUTPUT_EXE} .

run:
	@GOARCH=wasm GOOS=js go build -o web/app.wasm \
		&& go run .
