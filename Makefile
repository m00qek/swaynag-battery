all: run

clean:
	@rm bin/ -rf

run:
	@go run *.go

build: clean
	@go build -o bin/sway-battery-nagbar
