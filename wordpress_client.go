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
package wordpress

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

var (
	client   *http.Client
	instance *WordPressClient
	once     sync.Once
	site     string
)

type WordPressClient struct{}

func GetClient(siteName string) *WordPressClient {
	once.Do(func() {
		instance = &WordPressClient{}
		site = siteName

		tr := &http.Transport{
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
			DisableCompression: true,
		}

		client = &http.Client{Transport: tr}
	})

	return instance
}

/**
 * ListPages fetches an array of Page objects from a WordPress site
 */
func (w *WordPressClient) ListPages() (Pages, error) {
	var pages Pages
	var buff bytes.Buffer

	// build url
	fmt.Fprintf(&buff, "%s/wp-json/wp/v2/pages", site)

	resp, err := client.Get(buff.String())
	if err != nil {
		return pages, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return pages, fmt.Errorf("%s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&pages)
	if err != nil {
		return pages, err
	}

	return pages, nil
}

/**
 * ListPages fetches an array of Page objects from a WordPress site
 */
func (w *WordPressClient) RetrievePage(id int) (Page, error) {
	var page Page
	var buff bytes.Buffer

	// build url
	fmt.Fprintf(&buff, "%s/wp-json/wp/v2/pages/%d", site, id)

	resp, err := client.Get(buff.String())
	if err != nil {
		return page, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return page, fmt.Errorf("%s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&page)
	if err != nil {
		return page, err
	}

	return page, nil
}
