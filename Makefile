APP_NAME := myapp
BUILD_DIR := build

GOOS := linux
GOARCH := arm
GOARM := 7

##################################################################
# Change these EVNs according to your network and board.
# You only need to set RASPBERRY_PI_IP.
# The next two ENVs are used to find RASPBERRY_PI_IP.
RASPBERRY_PI_IP := 192.168.189.103
NETWORK_RANGE := 192.168.189.0/24
RASPBERRY_PI_MAC := b8:27:eb:82:33:e8
##################################################################


.PHONY: all clean run

all: $(BUILD_DIR)/$(APP_NAME)

$(BUILD_DIR)/$(APP_NAME): *.go
	GOOS=$(GOOS) GOARCH=$(GOARCH) GOARM=$(GOARM) go build -o $(BUILD_DIR)/$(APP_NAME) .

clean:
	rm -rf $(BUILD_DIR)

run: all
	$(BUILD_DIR)/$(APP_NAME)

test:
	go build -o $(BUILD_DIR)/$(APP_NAME)_test && sudo ./$(BUILD_DIR)/$(APP_NAME)_test --port /dev/tty0

find-ip:
	nmap -sn $(NETWORK_RANGE)
	arp -an | awk '/$(RASPBERRY_PI_MAC)/ {print $$2}' | tr -d '()'

transfer:
	make
	ssh root@$(RASPBERRY_PI_IP) 'rm -f myapp'
	scp build/myapp root@$(RASPBERRY_PI_IP):~

deploy:
	make transfer
	ssh root@$(RASPBERRY_PI_IP) './myapp'

errors:
	ssh root@$(RASPBERRY_PI_IP) 'cat errors.log'
