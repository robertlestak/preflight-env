# preflight-env

a preflight check to ensure that environment variables are set as expected.

## Build

```bash
make
```

## Install

NOTE: you will need `curl`, `bash`, and `jq` installed for the install script to work. It will attempt to install the binary in `/usr/local/bin` and will require `sudo` access. You can override the install directory by setting the `INSTALL_DIR` environment variable.

```bash
curl -sSL https://raw.githubusercontent.com/robertlestak/preflight-env/main/scripts/install.sh | bash
```

## Usage

```bash
Usage of preflight-env:
  -e value
        enviornment variable to check in the form of KEY=VALUE. if VALUE is omitted, only checks if KEY is set.
  -log-level string
        log level (default "info")
```

## Example

```bash
preflight-env -e FOO=bar -e BAZ
```
