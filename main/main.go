/*
 * http://www.apache.org/licenses/LICENSE-2.0.txt
 *
 * Copyright 2017 OpsVision Solutions
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/opsvision/wordpress-client"
)

var site = "http://demo.wp-api.org"

/**
 * Initialize the command line arguments
 */
func init() {
	flag.StringVar(&site, "site", site, "the wordpress site address")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(-1)
	}
	flag.Parse()
}

/**
 * Main entry point for program
 */
func main() {
	// Make sure we have something to work with
	if len(os.Args) == 1 {
		flag.Usage()
	}

	client := wordpress.GetClient(site)
	pages, err := client.ListPages()
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(-1)
	}

	for _, page := range pages {
		fmt.Printf("%s\n", page.Slug)
	}
}
