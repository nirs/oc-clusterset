// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"log"
	"os"

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

	return &clusterset, nil
}
