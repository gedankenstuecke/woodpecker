// Copyright 2024 Woodpecker Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package registry

import (
	"context"
	"html/template"
	"os"

	"github.com/urfave/cli/v3"

	"go.woodpecker-ci.org/woodpecker/v2/cli/common"
	"go.woodpecker-ci.org/woodpecker/v2/cli/internal"
)

var registryShowCmd = &cli.Command{
	Name:   "show",
	Usage:  "show registry information",
	Action: registryShow,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "hostname",
			Usage: "registry hostname",
			Value: "docker.io",
		},
		common.FormatFlag(tmplRegistryList, true),
	},
}

func registryShow(ctx context.Context, c *cli.Command) error {
	var (
		hostname = c.String("hostname")
		format   = c.String("format") + "\n"
	)

	client, err := internal.NewClient(ctx, c)
	if err != nil {
		return err
	}

	registry, err := client.GlobalRegistry(hostname)
	if err != nil {
		return err
	}

	tmpl, err := template.New("_").Parse(format)
	if err != nil {
		return err
	}
	return tmpl.Execute(os.Stdout, registry)
}
