GO ?= go
NAME ?= stockholm
BUILD ?= $(GO) build
RUN ?= $(GO) run
BUILD_DIR ?= ./bin/
TARGET ?= $(addprefix $(BUILD_DIR), $(NAME)) 
STOCKHOLM_SRC ?= $(addsufix $(NAME), .go)

MKDIR ?= mkdir -pv
RM ?= rm -rfvi

.PHONY: build
build: $(STOCKHOLM_SRC) | $(BUILD_DIR)
	$(BUILD) -o $(TARGET)

.PHONY: run
run:
ifeq (,$(wildcard /.dockerenv))
	$(error "no estás virtualizado amigo")
endif
	$(RUN) $(STOCKHOLM_SRC)

.PHONY: clean
	$(RM) $(BUILD_DIR)

$(BUILD_DIR):
	$(MKDIR) $@	

.DEFAULT:
	@echo "[ STOCKHOLM options ]"
	@echo ""
	@echo "make [build] : build stockholm binary"
	@echo "make run     : run without building"
	@echo "make clean   : remove build"