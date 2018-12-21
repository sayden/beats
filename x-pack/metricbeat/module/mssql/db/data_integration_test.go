// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package db

import (
	mbtest "github.com/elastic/beats/metricbeat/mb/testing"
	mtest "github.com/elastic/beats/x-pack/metricbeat/module/mssql/testing"
	"testing"
)

func TestData(t *testing.T) {
	f := mbtest.NewReportingMetricSetV2(t, mtest.GetConfig("db"))

	err := mbtest.WriteEventsReporterV2(f, t, "")
	if err != nil {
		t.Fatal("write", err)
	}
}
