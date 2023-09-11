// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"
)

// Global flags
var config string
var kubeconfig string
var verbose bool

var example = `  # Log in to all clusterss in config.yaml:
  oc clusterset login --config config.yaml

  # Log out from all clustes in config.yaml:
  oc clusterset logout --config config.yaml

  # Example config.yaml:
  clusters:
  - name: hub
    url: cluster1.example.com:8443
    username: kubeadmin
    password: PeSkM-R6YcH-LyPZa-oTOO1
  - name: c1
    url: cluster2.example.com:8443
    username: kubeadmin
    password: ZjIZn-SFUyR-aE4gI-fJcfL
  - name: c2
    url: cluster3.example.com:8443
    username: kubeadmin
    password: 7C700-oVS3Q-25rtx-YMew5
  current-context: hub
`

var rootCmd = &cobra.Command{
	Use:     "oc-clusterset",
	Short:   "manage clusterset logins",
	Example: example,
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&config, "config", "c", "config.yaml", "clusterset configuration file")
	rootCmd.PersistentFlags().StringVarP(&kubeconfig, "kubeconfig", "", defaultKubeconfig(), "kubeconfig file")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "be more verbose")
}

func defaultKubeconfig() string {
	env := os.Getenv("KUBECONFIG")
	if env != "" {
		return env
	}
	return clientcmd.RecommendedHomeFile
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		errlog.Fatal(err)
	}
}
