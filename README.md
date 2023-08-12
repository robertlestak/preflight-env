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
  -config string
        path to config file
  -e value
        enviornment variable to check in the form of KEY=VALUE. if VALUE is omitted, only checks if KEY is set.
  -log-level string
        log level (default "info")
```

## Example

```bash
preflight-env -e FOO=bar -e BAZ
```

### Docker example

```bash
docker run --rm robertlestak/preflight-env \
      -e FOO=bar \
      -e BAZ
```

## Config file

You can also use a config file rather than cli args.

```yaml
envVars:
      HELLO: world
      FOO: bar
      BAZ: # this will check if BAZ is set, but not its value
```

```bash
preflight-env -config config.yaml
```