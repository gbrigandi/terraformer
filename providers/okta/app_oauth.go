// Copyright 2021 The Terraformer Authors.
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

package okta

import (
	"context"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/okta/okta-sdk-golang/v2/okta"
)

type AppOAuthGenerator struct {
	OktaService
}

func (g AppOAuthGenerator) createResources(appList []*okta.Application) []terraformutils.Resource {
	var resources []terraformutils.Resource
	for _, app := range appList {
		resource := terraformutils.NewSimpleResource(
			app.Id,
			normalizeResourceName(app.Id+"_"+app.Name),
			"okta_app_oauth",
			"okta",
			[]string{})

		resource.IgnoreKeys = append(resource.IgnoreKeys,
			"^implicit_assignment$",
		)

		resources = append(resources, resource)
	}

	return resources
}

// Generate Terraform Resources from Okta API,
func (g *AppOAuthGenerator) InitResources() error {
	ctx, client, e := g.Client()
	if e != nil {
		return e
	}

	apps, err := getOAuthApplications(ctx, client)
	if err != nil {
		return err
	}

	g.Resources = g.createResources(apps)
	return nil
}

func getOAuthApplications(ctx context.Context, client *okta.Client) ([]*okta.Application, error) {
	signOnMode := "OPENID_CONNECT"
	apps, err := getApplications(ctx, client, signOnMode)
	if err != nil {
		return nil, err
	}
	return apps, nil
}
