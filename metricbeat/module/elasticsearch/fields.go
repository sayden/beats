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

package elasticsearch

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "elasticsearch", asset.ModuleFieldsPri, AssetElasticsearch); err != nil {
		panic(err)
	}
}

// AssetElasticsearch returns asset data.
// This is the base64 encoded gzipped contents of module/elasticsearch.
func AssetElasticsearch() string {
	return "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"
}
