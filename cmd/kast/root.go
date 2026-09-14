package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kast",
	Short: "Kubernetes YAML templating tool",
	Long:  `Kast is a kubernetes yaml manifest templating tool.`,
}
