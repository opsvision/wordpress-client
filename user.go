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

type User struct {
	ID                int           `json:"id"`
	Username          string        `json:"username"`
	Name              string        `json:"name"`
	FirstName         string        `json:"first_name"`
	LastName          string        `json:"last_name"`
	Email             string        `json:"email"`
	URL               string        `json:"url"`
	Description       string        `json:"description"`
	Link              string        `json:"link"`
	Locale            string        `json:"locale"`
	Nickname          string        `json:"nickname"`
	Slug              string        `json:"slug"`
	RegisteredDate    string        `json:"registered_date"`
	Roles             []interface{} `json:"roles"`
	Password          string        `json:"password"`
	Capabilities      interface{}   `json:"capabilities"`
	ExtraCapabilities interface{}   `json:"extra_capabilities"`
	AvatarURLs        interface{}   `json:"avatar_urls"`
	Meta              Metas         `json:"meta"`
}

type Users []User

func (u *User) CreateUser() {}

func (u *User) UpdateUser() {}

func (u *User) DeleteUser() {}
