SHELL := /bin/bash
.SHELLFLAGS := -eu -o pipefail -c

BIN ?= ft_otp
BUILD_DIR ?= ./bin
TARGET ?= $(addprefix $(BUILD_DIR)/, $(BIN))
MAIN ?= ft_otp.go

GO ?= go
BUILD ?= $(GO) build
RUN ?= $(GO) run

.PHONY: build
build: $(MAIN) | $(BUILD_DIR)
	$(BUILD) -o $(TARGET) $(notdir $<)

.PHONY: run
run:
	$(RUN) $(MAIN)

.PHONY: test
test:
	$(GO) mod tidy
	$(GO) test -v ./...

$(BUILD_DIR):
	mkdir -pv $(BUILD_DIR)

.PHONY: clean
clean: $(TARGET)
	rm $@

.DEFAULT:
	@echo "$(MAKECMDGOALS): unknown rule"
	@echo "usage:"
	@echo ""
	@echo "build:	Generate ft_otp binary"
	@echo "run: 	Run ft_otp"
	@echo "clean:   Remove binary"