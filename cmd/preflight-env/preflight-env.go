package main

import (
	"flag"
	"os"
	"strings"

	"github.com/robertlestak/preflight-env/pkg/preflightenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	ll, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		ll = log.InfoLevel
	}
	log.SetLevel(ll)
}

type envVarList []string

func (e *envVarList) String() string {
	return strings.Join(*e, ",")
}

func (e *envVarList) Set(value string) error {
	*e = append(*e, value)
	return nil
}

func main() {
	l := log.WithFields(log.Fields{
		"app": "preflight-env",
	})
	l.Debug("starting preflight-env")
	preflightFlags := flag.NewFlagSet("preflight-env", flag.ExitOnError)
	logLevel := preflightFlags.String("log-level", log.GetLevel().String(), "log level")
	var envList envVarList
	preflightFlags.Var(&envList, "e", "enviornment variable to check in the form of KEY=VALUE. if VALUE is omitted, only checks if KEY is set.")
	preflightFlags.Parse(os.Args[1:])
	ll, err := log.ParseLevel(*logLevel)
	if err != nil {
		ll = log.InfoLevel
	}
	log.SetLevel(ll)
	envVars := make(map[string]string)
	for _, e := range envList {
		// split on "=" to get key and value
		s := strings.Split(e, "=")
		if len(s) != 2 {
			envVars[s[0]] = ""
		} else {
			envVars[s[0]] = s[1]
		}
	}
	pf := &preflightenv.PreflightEnv{
		EnvVars: envVars,
	}
	if err := pf.Run(); err != nil {
		l.WithError(err).Error("error running preflight-env")
		os.Exit(1)
	}
}
