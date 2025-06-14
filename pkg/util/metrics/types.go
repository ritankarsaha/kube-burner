// Copyright 2020 The Kube-burner Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package metrics

import (
	"github.com/cloud-bulldozer/go-commons/v2/indexers"
	"github.com/kube-burner/kube-burner/pkg/alerting"
	"github.com/kube-burner/kube-burner/pkg/config"
	"github.com/kube-burner/kube-burner/pkg/prometheus"
	"github.com/kube-burner/kube-burner/pkg/util/fileutils"
)

// ScraperConfig holds data related to scraper and target indexer
type ScraperConfig struct {
	ConfigSpec      *config.Spec
	MetricsEndpoint string
	UserMetaData    string
	SummaryMetadata map[string]any
	MetricsMetadata map[string]any
	MetricsProfile  string
	AlertProfile    string
	EmbedCfg        *fileutils.EmbedConfiguration
}

// ScraperResponse holds parsed data related to scraper and target indexer
type Scraper struct {
	PrometheusClients []*prometheus.Prometheus
	AlertMs           []*alerting.AlertManager
	IndexerList       map[string]indexers.Indexer
	SummaryMetadata   map[string]any
	MetricsMetadata   map[string]any
}
