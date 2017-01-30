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

type Page struct {
	Date          string  `json:"date"`
	DateGmt       string  `json:"date_gmt"`
	GUID          GUID    `json:"guid"`
	ID            int     `json:"id"`
	Link          string  `json:"link"`
	Modified      string  `json:"modified"`
	ModifiedGmt   string  `json:"modified_gmt"`
	Slug          string  `json:"slug"`
	Status        string  `json:"status"`
	Type          string  `json:"type"`
	Parent        int     `json:"parent"`
	Title         Title   `json:"title"`
	Content       Content `json:"content"`
	Author        int     `json:"author"`
	Excerpt       Excerpt `json:"excerpt"`
	FeaturedMedia int     `json:"featured_media"`
	CommentStatus string  `json:"comment_status"`
	PingStatus    string  `json:"ping_status"`
	MenuOrder     int     `json:"menu_order"`
	Meta          Metas   `json:"meta"`
	Template      string  `json:"template"`
}

type Pages []Page
