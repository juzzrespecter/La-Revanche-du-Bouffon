.PHONY: build build-spider build-scorpion 

GO ?= go
BUILD ?= $(GO) build
BUILD_DIR ?= ./bin/
SPIDER_DIR ?= spider/
SCORPION_DIR ?= scorpion/
SPIDER_SRC ?= $(addprefix $(SPIDER_DIR), spider.go)
SCORPION_SRC ?= $(addprefix $(SCORPION_DIR), scorpion.go)
TARGET_SPIDER ?= $(addprefix $(BUILD_DIR), spider)
TARGET_SCORPION ?= $(addprefix $(BUILD_DIR), scorpion)

.PHONY: build
build: build-spider build-scorpion

.PHONY: build-spider
build-spider: $(SPIDER_SRC) $(BUILD_DIR)
	$(BUILD) -o $(TARGET_SPIDER) $<

.PHONY: build-scorpion
build-scorpion: $(SCORPION_SRC) $(BUILD_DIR)
	$(BUILD) -o $(TARGET_SCORPION) $< 

$(BUILD_DIR):
	mkdir -pv $@

.PHONY: clean
clean:
	rm -rfv $(BUILD_DIR)

.DEFAULT:
	@echo "Build project"
	@echo ""
	@echo "build:          Build all binaries"
	@echo "build-spider:   Build crawler"
	@echo "build-scorpion: Build exif tool"