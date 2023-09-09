// SPDX-FileCopyrightText: The RamenDR authors
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"log"
	"os/exec"

	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to clusterset",
	Run:   runLogin,
}

func init() {
	rootCmd.AddCommand(loginCmd)
}

func runLogin(cmd *cobra.Command, args []string) {
	clusterset, err := loadClusterset(config)
	if err != nil {
		log.Fatal(err)
	}
	if verbose {
		log.Printf("Using kubeconfig %q\n", kubeconfig)
	}
	for _, cluster := range clusterset.Clusters {
		loginToCluster(cluster)
		renameContext(cluster)
		if clusterset.CurrentContext != "" {
			setCurrentContext(clusterset.CurrentContext)
		}
	}
}

func loginToCluster(cluster *Cluster) {
	if verbose {
		log.Printf("Logging in to cluster %q %q", cluster.Name, cluster.Url)
	}
	cmd := exec.Command(
		"oc", "login", cluster.Url,
		"--username", cluster.Username,
		"--password", cluster.Password,
		"--kubeconfig", kubeconfig,
	)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// renameContext rename the context refering to cluster to cluster name instead
// of the long and unusable context name set by oc login. It must be called
// after oc login since it assumes the curent context is the new logged in
// context.
func renameContext(cluster *Cluster) {
	config, err := clientcmd.LoadFromFile(kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	if verbose {
		log.Printf("Renaming context %q to %q\n", config.CurrentContext, cluster.Name)
	}

	// Remove previously updated context, since `oc login`.
	delete(config.Contexts, cluster.Name)

	// Rename current context.
	config.Contexts[cluster.Name] = config.Contexts[config.CurrentContext]
	delete(config.Contexts, config.CurrentContext)
	config.CurrentContext = cluster.Name

	if err := clientcmd.WriteToFile(*config, kubeconfig); err != nil {
		log.Fatal(err)
	}
}

func setCurrentContext(name string) {
	config, err := clientcmd.LoadFromFile(kubeconfig)
	if err != nil {
		log.Fatal(err)
	}

	if verbose {
		log.Printf("Setting current context to %q\n", name)
	}

	config.CurrentContext = name

	if err := clientcmd.WriteToFile(*config, kubeconfig); err != nil {
		log.Fatal(err)
	}
}
