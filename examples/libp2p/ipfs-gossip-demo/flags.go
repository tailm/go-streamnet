package main

import (
	"flag"
)

// Config contains ...
type Config struct {
	Seed string
	Port string
}

// ParseFlags parsing arguments for app
func ParseFlags() (Config, error) {
	config := Config{}
	flag.StringVar(&config.Seed, "seed", "/ip4/127.0.0.1/tcp/45759/ipfs/QmWjz6xb8v9K4KnYEwP5Yk75k5mMBCehzWFLCvvQpYxF3d", "while starting you peer, a seed should be specify")
	flag.StringVar(&config.Port, "port", "45759", "listening port")
	flag.Parse()

	return config, nil
}
