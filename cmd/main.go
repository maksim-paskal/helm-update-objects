package main

import (
	"flag"

	"github.com/maksim-paskal/helm-update-objects/internal"
	"github.com/maksim-paskal/helm-update-objects/pkg/client"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

var logLevel = flag.String("log-level", "INFO", "log level")

func main() {
	flag.Parse()

	logLevel, err := log.ParseLevel(*logLevel)
	if err != nil {
		log.WithError(err).Fatal()
	}

	log.SetLevel(logLevel)

	if err := client.Init(); err != nil {
		log.WithError(err).Fatal()
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := internal.Run(ctx); err != nil {
		log.WithError(err).Fatal()
	}
}
