APP_NAME=go-scamalytics
BUILD_DIR=build
SRC_FILES=*.go

# Windows x86-64
build_win64:
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)_win64.exe $(SRC_FILES)

# Windows x86
build_win32:
	GOOS=windows GOARCH=386 go build -o $(BUILD_DIR)/$(APP_NAME)_win32.exe $(SRC_FILES)

# Linux x86-64
build_linux64:
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)_linux64 $(SRC_FILES)

# Linux x86
build_linux32:
	GOOS=linux GOARCH=386 go build -o $(BUILD_DIR)/$(APP_NAME)_linux32 $(SRC_FILES)

# macOS Intel
build_macos_intel:
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME)_macos_intel $(SRC_FILES)

# macOS M1
build_macos_arm:
	GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DIR)/$(APP_NAME)_macos_m1 $(SRC_FILES)


build_all: build_win64 build_win32 build_linux64 build_linux32 build_macos_intel build_macos_arm

clean:
	rm -rf $(BUILD_DIR)
