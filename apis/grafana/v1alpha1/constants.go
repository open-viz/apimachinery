/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

// **************** Common Constants *******************

type GrafanaPhase string

const (
	GrafanaPhaseProcessing  GrafanaPhase = "Processing"
	GrafanaPhaseTerminating GrafanaPhase = "Terminating"
	GrafanaPhaseSuccess     GrafanaPhase = "Success"
	GrafanaPhaseFailed      GrafanaPhase = "Failed"
)

// ************** Datasource Constants ****************

type DatasourceType string

const (
	DatasourceTypePrometheus DatasourceType = "prometheus"
)

type DatasourceAccessType string

const (
	DatasourceAccessTypeProxy  DatasourceAccessType = "proxy"
	DatasourceAccessTypeDirect DatasourceAccessType = "direct"
)
