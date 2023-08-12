bin: bin/preflight-env_darwin_amd64 bin/preflight-env_linux_amd64 bin/preflight-env_windows_amd64.exe
bin: bin/preflight-env_darwin_arm64 bin/preflight-env_linux_arm64 bin/preflight-env_windows_arm64.exe

bin/preflight-env_darwin_amd64:
	@mkdir -p bin
	@echo "Compiling preflight-env..."
	GOOS=darwin GOARCH=amd64 go build -o $@ cmd/preflight-env/*.go

bin/preflight-env_darwin_arm64:
	@mkdir -p bin
	@echo "Compiling preflight-env..."
	GOOS=darwin GOARCH=arm64 go build -o $@ cmd/preflight-env/*.go

bin/preflight-env_linux_amd64:
	@mkdir -p bin
	@echo "Compiling preflight-env..."
	GOOS=linux GOARCH=amd64 go build -o $@ cmd/preflight-env/*.go

bin/preflight-env_linux_arm64:
	@mkdir -p bin
	@echo "Compiling preflight-env..."
	GOOS=linux GOARCH=arm64 go build -o $@ cmd/preflight-env/*.go

bin/preflight-env_windows_amd64.exe:
	@mkdir -p bin
	@echo "Compiling preflight-env..."
	GOOS=windows GOARCH=amd64 go build -o $@ cmd/preflight-env/*.go

bin/preflight-env_windows_arm64.exe:
	@mkdir -p bin
	@echo "Compiling preflight-env..."
	GOOS=windows GOARCH=arm64 go build -o $@ cmd/preflight-env/*.go

.PHONY: install
install: bin
	@echo "Installing preflight-env..."
	@scp bin/preflight-env_$$(go env GOOS)_$$(go env GOARCH) /usr/local/bin/preflight-env