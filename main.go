/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"flag"
	"log"
	"os"
	"runtime"

	ecc "github.com/ernestio/ernest-config-client"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "config.json", "The path to the shared config file")
	flag.Parse()

	c := ecc.NewConfig(os.Getenv("NATS_URI"))
	n := c.Nats()

	h := Handler{Nats: n, ConfigPath: configPath}

	log.Println("Started")

	_, err := n.Subscribe("config.get.*", h.ConfigGet)
	if err != nil {
		log.Fatal(err)
	}

	_, err = n.Subscribe("config.set.*", h.ConfigSet)
	if err != nil {
		log.Fatal(err)
	}

	runtime.Goexit()
	log.Println("Stopped")
}
