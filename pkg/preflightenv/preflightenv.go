package preflightenv

import (
	"encoding/json"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type PreflightEnv struct {
	EnvVars map[string]string `json:"envVars" yaml:"envVars"`
}

func LoadConfig(filepath string) (*PreflightEnv, error) {
	l := log.WithFields(log.Fields{
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

func (pf *PreflightEnv) Run() error {
	l := log.WithFields(log.Fields{
		"fn": "Run",
	})
	l.Debug("starting preflight-env")
	for k, v := range pf.EnvVars {
		if v == "" {
			// checking if env var exists
			ev := os.Getenv(k)
			if ev == "" {
				l.Errorf("env var %s not set", k)
				return fmt.Errorf("env var %s not set", k)
			}
		} else {
			// checking if env var is set to correct value
			ev := os.Getenv(k)
			if ev != v {
				l.Errorf("env var %s not set to correct value", k)
				return fmt.Errorf("env var %s not set to correct value", k)
			}
		}
	}
	l.Info("preflight-env completed successfully")
	return nil
}
