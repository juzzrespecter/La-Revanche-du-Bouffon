.PHONY: build run

GO ?= go
BUILD ?= $(GO) build
RUN ?= $(GO) run

MAIN ?= ft_otp.go

build:
	$(BUILD) $(MAIN)

run:
	$(RUN) $(MAIN)

.DEFAULT:
	@echo "One-time password generator"
	@echo ""
	@echo "build: Generate binary"
	@echo "run: run program"