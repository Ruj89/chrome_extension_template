APP_NAME = native_app
CHROME_APP_NAME = com.example.nativeapp
BUILD_DIR = build
GO_SOURCE_DIR = native
TOOLS_DIR = tools
EXTENSION_SOURCE_DIR = extension

all: build

build: $(BUILD_DIR)/$(APP_NAME) $(BUILD_DIR)/$(CHROME_APP_NAME).json
	@echo "Build completed!"

$(BUILD_DIR)/$(APP_NAME): $(GO_SOURCE_DIR)/main.go
	@echo "Building the Go binary..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(GO_SOURCE_DIR)/main.go

ifneq ($(BASE_DIR),)
DESTINATION_BUILD_DIR = $(abspath $(BASE_DIR)/$(BUILD_DIR))
MANIFEST_PATH = $(abspath $(BASE_DIR)/$(EXTENSION_SOURCE_DIR))
else
DESTINATION_BUILD_DIR = $(realpath $(BUILD_DIR))
MANIFEST_PATH = $(realpath $(EXTENSION_SOURCE_DIR))
endif

$(BUILD_DIR)/.extension_id: $(EXTENSION_SOURCE_DIR)
	@echo "Calculating extension ID"
	go run $(TOOLS_DIR)/generate_chrome_id.go "$(MANIFEST_PATH)" > $(BUILD_DIR)/.extension_id

$(BUILD_DIR)/$(CHROME_APP_NAME).json: $(BUILD_DIR)/.extension_id $(EXTENSION_SOURCE_DIR)/manifest.json $(GO_SOURCE_DIR)/$(CHROME_APP_NAME).json.template
	@echo "Generating the manifest file"
	@sed \
		-e "s|\$$PATH|$(DESTINATION_BUILD_DIR)/$(APP_NAME)|g" \
		-e "s|\$$EXTENSION_ID|$(shell cat $(BUILD_DIR)/.extension_id)|g" \
		$(GO_SOURCE_DIR)/$(CHROME_APP_NAME).json.template > $(BUILD_DIR)/$(CHROME_APP_NAME).json

install: $(BUILD_DIR)/$(CHROME_APP_NAME).json
	@echo "Installing Chrome application manifest"
	cp $(BUILD_DIR)/$(CHROME_APP_NAME).json $(HOME)/.config/BraveSoftware/Brave-Browser/NativeMessagingHosts

clean:
	@echo "Cleaning build directory..."
	rm -rf $(BUILD_DIR)
