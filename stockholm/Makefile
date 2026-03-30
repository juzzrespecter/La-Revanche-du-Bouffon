GO ?= go
BUILD ?= $(GO) build
RUN ?= $(GO) run
BUILD_DIR ?= ./bin/
STOCKHOLM_SRC ?= stockholm.go

.PHONY: build
build: $(STOCKHOLM_SRC) | $(BUILD_DIR)


run:
	# capar if not containerized