/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/nats-io/go-nats"
	"github.com/r3labs/akira"
)

// Handler describes a NATS handler and its dependencies.
type Handler struct {
	Nats       akira.Connector
	ConfigPath string
}

// ConfigGet handles requests to config.get.*
func (h *Handler) ConfigGet(msg *nats.Msg) {
	service := extractService(msg.Subject)
	data := getServiceConfig(service, h.ConfigPath)
	err := h.Nats.Publish(msg.Reply, data)
	if err != nil {
		log.Println(err)
	}

}

// ConfigSet handles requests to config.set.*
func (h *Handler) ConfigSet(msg *nats.Msg) {
	service := extractService(msg.Subject)
	data := setServiceConfig(service, h.ConfigPath, msg.Data)
	err := h.Nats.Publish(msg.Reply, data)
	if err != nil {
		log.Println(err)
	}

}

func extractService(subject string) string {
	s := strings.Split(subject, ".")
	return s[len(s)-1]
}

func getServiceConfig(service, configPath string) []byte {
	c := make(map[string]interface{})

	file, err := os.Open(configPath)
	if err != nil {
		log.Println(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(data, &c)
	if err != nil {
		log.Println(err)
	}

	data, err = json.Marshal(c[service])
	if err != nil {
		log.Println(err)
	}

	return data
}

func setServiceConfig(service, configPath string, msg []byte) []byte {
	c := make(map[string]interface{})

	file, err := os.Open(configPath)
	if err != nil {
		log.Println(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(data, &c)
	if err != nil {
		log.Println(err)
	}

	msg = []byte(`{"` + service + `": ` + string(msg) + `}`)
	err = json.Unmarshal(msg, &c)
	if err != nil {
		log.Println(err)
	}

	data, err = json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Println(err)
	}

	err = ioutil.WriteFile(configPath, data, 0644)
	if err != nil {
		log.Println(err)
	}

	return data
}
