# Build directory
BUILD_DIR := build

# Install directory
INSTALL_DIR := /usr/bin

# Go build flags (for release)
GO_BUILD_FLAGS := -ldflags="-s -w"

build: limetool

limetool:
	@mkdir -p $(BUILD_DIR)
	@go build $(GO_BUILD_FLAGS) -o $(BUILD_DIR)/limetool cmd/limetool_cli.go

install: build
	install -m 755 $(BUILD_DIR)/limetool $(INSTALL_DIR)
