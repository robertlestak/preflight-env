package preflightenv

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

type PreflightEnv struct {
	EnvVars map[string]string `json:"envVars" yaml:"envVars"`
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
