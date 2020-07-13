// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package mysql

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "mysql", asset.ModuleFieldsPri, AssetMysql); err != nil {
		panic(err)
	}
}

// AssetMysql returns asset data.
// This is the base64 encoded gzipped contents of module/mysql.
func AssetMysql() string {
	return "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"
}
