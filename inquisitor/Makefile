.PHONY: build build-env run 

GO 	  ?= go
BUILD ?= $(GO) build
RUN   ?= $(GO) run
COMPOSE ?= docker compose

MAIN ?= ./src/inquisitor.go

run:
	$(RUN) $(MAIN)

build:
	$(BUILD) $(MAIN)

build-env:
	$(COMPOSE) up --build -d

.DEFAULT:
	@echo "Inquisitor usage:"
	@echo ""
	@echo "build:     Build inquisitor"
	@echo "build-env: Set up PoC environment"
	@echo "run:       Run without build"
