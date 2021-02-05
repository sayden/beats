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
