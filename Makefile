APP=tweety
BIN=$(PWD)/bin/$(APP)
VERSION=1.0.0
GO ?= go

linux: clean
	@cd cmd/tweety GOOS=linux $(GO) build -o $(BIN)

build: clean
	@cd cmd/tweety && $(GO) build -o $(BIN)

clean:
	@echo "[clean] Cleaning files..."
	@rm -f $(BIN)
