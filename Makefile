APP_NAME := myapp
BUILD_DIR := build

GOOS := linux
GOARCH := arm
GOARM := 7

.PHONY: all clean run

all: $(BUILD_DIR)/$(APP_NAME)

$(BUILD_DIR)/$(APP_NAME): *.go
	GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM) go build -o $(BUILD_DIR)/$(APP_NAME) .

clean:
	rm -rf $(BUILD_DIR)

run: all
	$(BUILD_DIR)/$(APP_NAME)

test:
	go build -o $(BUILD_DIR)/$(APP_NAME)_test && sudo ./$(BUILD_DIR)/$(APP_NAME)_test