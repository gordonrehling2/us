export TOOLS=./bin

$$TOOLS/hi:
	echo "Error: need to install bash script ./bin/hi"


deps: $$TOOLS/hi
	dep ensure


lint:
	golint $(go list ./... | grep -v /vendor/)

dbuild:
	docker build -t us .

drun:
	docker run -it --rm --name us us

run: lint
	$$TOOLS/hi
	go run main.go