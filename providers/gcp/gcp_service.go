// Copyright 2018 The Terraformer Authors.
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

package gcp

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
)

type GCPService struct { //nolint
	terraformutils.Service
}

func (s *GCPService) applyCustomProviderType(resources []terraformutils.Resource, providerName string) []terraformutils.Resource {
	editedResources := []terraformutils.Resource{}
	for _, r := range resources {
		r.Item["provider"] = providerName
		editedResources = append(editedResources, r)
	}
	return editedResources
}
