// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"sigs.k8s.io/yaml"
)

type Cluster struct {
	Name     string `json:"name"`
	Url      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Clusterset struct {
	Clusters       []*Cluster `json:"clusters"`
	CurrentContext string     `json:"current-context,omitempty"`
}

func loadClusterset(filename string) (*Clusterset, error) {
	if verbose {
		log.Printf("Loading config from %q\n", filename)
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	clusterset := Clusterset{}
	if err := yaml.UnmarshalStrict(data, &clusterset); err != nil {
		return nil, err
	}

	if err := clusterset.Validate(); err != nil {
		return nil, err
	}

	return &clusterset, nil
}

func (c *Clusterset) Validate() error {
	if len(c.Clusters) == 0 {
		return fmt.Errorf("invalid config: no clusters")
	}

	if c.CurrentContext != "" {
		for _, cluster := range c.Clusters {
			if cluster.Name == c.CurrentContext {
				return nil
			}
		}

		return fmt.Errorf(
			"invalid config: invalid %q: no cluster named %q",
			"current-context", c.CurrentContext,
		)
	}

	for _, cluster := range c.Clusters {
		if err := cluster.Validate(); err != nil {
			return fmt.Errorf("invalid config: %s", err)
		}
	}

	return nil
}

func (c *Cluster) Validate() error {
	var missing []string

	if c.Name == "" {
		missing = append(missing, "name")
	}
	if c.Url == "" {
		missing = append(missing, "url")
	}
	if c.Username == "" {
		missing = append(missing, "username")
	}
	if c.Password == "" {
		missing = append(missing, "password")
	}

	if len(missing) > 0 {
		return fmt.Errorf("missing required keys: %s", strings.Join(missing, ", "))
	}

	return nil
}
