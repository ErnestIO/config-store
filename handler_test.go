/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestExtractService(t *testing.T) {
	Convey("Given a service full name of a service", t, func() {
		serviceFullName := "config.redis"
		Convey("gets its name", func() {
			service := extractService(serviceFullName)
			So(service, ShouldEqual, "redis")
		})
	})

	Convey("Given an invalid full name of a service", t, func() {
		serviceFullName := "foo"
		Convey("can't get the name, get the entire string", func() {
			service := extractService(serviceFullName)
			So(service, ShouldEqual, "foo")
		})
	})
}

func TestLoadServiceConfig(t *testing.T) {
	Convey("Given a valid configuration fixture", t, func() {
		configPath := "./fixtures/config.json"
		Convey("Load a service configuration correctly", func() {
			service := loadServiceConfig("salt", configPath)
			So(string(service), ShouldEqual, "{\"password\":\"bar\",\"user\":\"foo\"}")
		})
	})
}
