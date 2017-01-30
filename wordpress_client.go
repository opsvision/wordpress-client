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

// WordPressClient is used to interact with a WordPress site using the REST
// API documented here: https://developer.wordpress.org/rest-api
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
 * ListPages retrieves an array of Page objects from a WordPress site
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
 * ListPages retrieves an array of Page objects from a WordPress site
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

// ListUsers retrieves an array of User objects
func (w *WordPressClient) ListUsers() (Users, error) {
	var users Users
	var buff bytes.Buffer

	// build url
	fmt.Fprintf(&buff, "%s/wp-json/wp/v2/users", site)

	resp, err := client.Get(buff.String())
	if err != nil {
		return users, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return users, fmt.Errorf("%s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		return users, err
	}

	return users, nil
}

// RetrieveUser retrieves a single User object based on the ID
func (w *WordPressClient) RetrieveUser(id int) (User, error) {
	var user User
	var buff bytes.Buffer

	// build url
	fmt.Fprintf(&buff, "%s/wp-json/wp/v2/users/%d", site, id)

	resp, err := client.Get(buff.String())
	if err != nil {
		return user, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return user, fmt.Errorf("%s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return user, err
	}

	return user, nil
}
