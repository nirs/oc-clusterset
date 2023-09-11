// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out from clusterset",
	Run:   runLogout,
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}

func runLogout(cmd *cobra.Command, args []string) {
	clusterset, err := loadClusterset(config)
	if err != nil {
		log.Fatal(err)
	}
	if verbose {
		log.Printf("Using kubeconfig %q\n", kubeconfig)
	}
	for _, cluster := range clusterset.Clusters {
		logoutFromCluster(cluster)
	}
}

func logoutFromCluster(cluster *Cluster) {
	if verbose {
		log.Printf("Logging out from cluster %q\n", cluster.Name)
	}
	cmd := exec.Command(
		"oc", "logout",
		"--kubeconfig", kubeconfig,
		"--context", cluster.Name,
	)
	// oc may write useful errors to stdout.
	out, err := cmd.CombinedOutput()
	if err != nil {
		// oc writes error message to stdout.
		log.Fatalf("Logout failed: [%s] %s", err, out)
	}
}
