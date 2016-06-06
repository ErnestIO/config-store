/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/nats-io/nats"
)

// Config stores all configuration for redis
type Config map[string]interface{}

func errCheck(err error) {
	if err != nil {
		log.Panic("Error: ", err)
	}
}

func extractService(subject string) string {
	s := strings.Split(subject, ".")
	return s[len(s)-1]
}

func loadServiceConfig(service string, configPath string) []byte {
	c := Config{}

	file, err := os.Open(configPath)
	errCheck(err)

	data, err := ioutil.ReadAll(file)
	errCheck(err)

	err = json.Unmarshal(data, &c)
	errCheck(err)

	data, err = json.Marshal(c[service])
	errCheck(err)

	return data
}

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "config.json", "The path to the shared config file")
	flag.Parse()

	natsURI := os.Getenv("NATS_URI")
	if natsURI == "" {
		log.Panic("No NATS_URI specified")
	}

	log.Println("starting")

	n, err := nats.Connect(natsURI)
	if err != nil {
		log.Panic(err)
	}

	n.Subscribe("config.get.*", func(msg *nats.Msg) {
		service := extractService(msg.Subject)
		data := loadServiceConfig(service, configPath)
		n.Publish(msg.Reply, data)
	})

	runtime.Goexit()
	log.Println("exiting")
}
