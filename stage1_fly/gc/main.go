// Copyright 2015 The rkt Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"io/ioutil"
	"log"
)

const (
	mountinfoPath = "/proc/self/mountinfo"
)

var (
	debug bool
)

func init() {
	flag.BoolVar(&debug, "debug", false, "Run in debug mode")
}

func main() {
	flag.Parse()

	if !debug {
		log.SetOutput(ioutil.Discard)
	}
	log.Printf("Not doing anything since stage0 is cleaning up the mounts")
	return
}
