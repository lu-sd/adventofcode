
.PHONY: all
COOKIE_FILE=cmd/cookie.txt

all: build

build: $(COOKIE_FILE)
	@echo "build generator"
	@go build -o generator cmd/main.go
