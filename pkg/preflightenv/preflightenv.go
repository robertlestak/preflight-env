package preflightenv

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

var (
	Logger *log.Logger
)

func init() {
	if Logger == nil {
		Logger = log.New()
		Logger.SetOutput(os.Stdout)
		Logger.SetLevel(log.InfoLevel)
	}
}

type PreflightEnv struct {
	Equiv   bool              `json:"equivalent" yaml:"equivalent"`
	EnvVars map[string]string `json:"envVars" yaml:"envVars"`
}

func LoadConfig(filepath string) (*PreflightEnv, error) {
	l := Logger.WithFields(log.Fields{
		"fn": "LoadConfig",
	})
	l.Debug("loading config")
	var err error
	pf := &PreflightEnv{}
	bd, err := os.ReadFile(filepath)
	if err != nil {
		l.WithError(err).Error("error reading file")
		return pf, err
	}
	if err := yaml.Unmarshal(bd, pf); err != nil {
		// try with json
		if err := json.Unmarshal(bd, pf); err != nil {
			l.WithError(err).Error("error unmarshalling config")
			return pf, err
		}
	}
	return pf, err
}

func (pf *PreflightEnv) Equivalent() {
	l := Logger
	l.Debug("printing equivalent command")
	var cmd string
	for k, v := range pf.EnvVars {
		if v == "" {
			cmd += `if [ -z "$` + k + `" ]; then echo "expecting ` + k + ` to exist" && exit 1; fi; `
		} else {
			cmd += `if [ "$` + k + `" != "` + v + `" ]; then echo "failed - expected: ` + v + `, got: $` + k + `" && exit 1; fi; `
		}
	}
	cmd = strings.TrimSpace(cmd)
	cmd = fmt.Sprintf("sh -c '%s'", cmd)
	fmt.Println(cmd)
}

func (pf *PreflightEnv) Run() error {
	l := Logger.WithFields(log.Fields{
		"preflight": "env",
	})
	l.Debug("starting preflight-env")
	if pf.Equiv {
		pf.Equivalent()
		return nil
	}
	for k, v := range pf.EnvVars {
		if v == "" {
			// checking if env var exists
			ev := os.Getenv(k)
			if ev == "" {
				failStr := fmt.Sprintf("failed - expected: %s to exit, got: nil", k)
				l.Error(failStr)
				return errors.New(failStr)
			}
		} else {
			// checking if env var is set to correct value
			ev := os.Getenv(k)
			if ev != v {
				failStr := fmt.Sprintf("failed - expected: %s, got: %s", v, ev)
				l.Error(failStr)
				return errors.New(failStr)
			}
		}
	}
	l.Info("passed")
	return nil
}
