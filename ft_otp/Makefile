SHELL := /bin/bash
.SHELLFLAGS := -eu -o pipefail -c

BUILD_DIR ?= ./bin
TARGET ?= $(addprefix $(BUILD_DIR)/, ft_otp)
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

$(BUILD_DIR):
	mkdir -pv $(BUILD_DIR)

.PHONY: clean
clean: $(TARGET)
	rm $@

.DEFAULT:
	@echo "$(MAKECMDGOALS): unknown rule"
	@echo "usage:"
	@echo ""
	@echo "build:	Generate binary"
	@echo "run: 	Run program"
	@echo "clean:   Remove binary"