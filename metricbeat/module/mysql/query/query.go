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

/*
Package status fetches MySQL server status metrics.

For more information on the query it uses, see:
http://dev.mysql.com/doc/refman/5.7/en/show-status.html
*/
package query

import (
	"context"

	"github.com/pkg/errors"

	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/common/cfgwarn"
	"github.com/elastic/beats/v7/metricbeat/helper/sql"
	"github.com/elastic/beats/v7/metricbeat/mb"
)

func init() {
	mb.Registry.MustAddMetricSet("mysql", "query", New,
		mb.DefaultMetricSet(),
	)
}

type query struct {
	Namespace      string `config:"query_namespace" validate:"nonzero,required"`
	Query          string `config:"query" validate:"nonzero,required"`
	ResponseFormat string `config:"response_format" validate:"nonzero,required"`
	DeDotEnabled   bool   `config:"dedot.enabled"`
}

// MetricSet for fetching MySQL server status.
type MetricSet struct {
	mb.BaseMetricSet
	db     *sql.DbClient
	Config struct {
		Queries   []query `config:"queries" validate:"nonzero,required"`
		Namespace string  `config:"namespace" validate:"nonzero,required"`
	}
}

// New creates and returns a new MetricSet instance.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	cfgwarn.Beta("The mysql 'query' metricset is beta.")

	b := &MetricSet{BaseMetricSet: base}

	if err := base.Module().UnpackConfig(&b.Config); err != nil {
		return nil, err
	}

	return b, nil
}

// Fetch fetches status messages from a mysql host.
func (m *MetricSet) Fetch(ctx context.Context, reporter mb.ReporterV2) error {
	if m.db == nil {
		var err error
		m.db, err = sql.DB("mysql", m.HostData().URI, m.Logger())
		if err != nil {
			return errors.Wrap(err, "mysql-status fetch failed")
		}
	}

	for _, q := range m.Config.Queries {
		err := m.fetchQuery(ctx, q, reporter)
		if err != nil {
			m.Logger().Error("error doing query", err)
		}
	}

	return nil
}

func (m *MetricSet) fetchQuery(ctx context.Context, query query, reporter mb.ReporterV2) error {
	if query.ResponseFormat == "table" {
		mss, err := m.db.FetchTableMode(ctx, query.Query)
		if err != nil {
			return err
		}

		for _, ms := range mss {
			event := m.transformMapStrToEvent(query, ms)
			reporter.Event(event)
		}
	} else {
		ms, err := m.db.FetchVariableMode(ctx, query.Query)
		if err != nil {
			return err
		}

		event := m.transformMapStrToEvent(query, ms)
		reporter.Event(event)
	}

	return nil
}

func (m *MetricSet) transformMapStrToEvent(query query, ms common.MapStr) mb.Event {
	event := mb.Event{ModuleFields: common.MapStr{m.Config.Namespace: common.MapStr{}}}

	data := ms
	if query.DeDotEnabled {
		data = sql.ToDotKeys(ms)
	}

	if query.Namespace != "" {
		event.ModuleFields[m.Config.Namespace] = common.MapStr{query.Namespace: data}
	} else {
		event.ModuleFields[m.Config.Namespace] = data
	}

	return event
}

// Close closes the database connection and prevents future queries.
func (m *MetricSet) Close() error {
	if m.db == nil {
		return nil
	}
	return errors.Wrap(m.db.Close(), "failed to close mysql database client")
}
