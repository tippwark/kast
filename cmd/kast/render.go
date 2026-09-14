package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tippwark/kast/internal/kast"
)

var renderCmd = &cobra.Command{
	Use:   "render [template|dir]",
	Short: "Render templates",
	Args:  cobra.ExactArgs(1),

	PreRunE: func(cmd *cobra.Command, args []string) error {
		inputTemplate := args[0]

		if _, err := os.Stat(inputTemplate); os.IsNotExist(err) {
			return fmt.Errorf("input template does not exist: %s", inputTemplate)
		}

		return nil
	},

	RunE: kast.Render(),
}

func init() {
	rootCmd.AddCommand(renderCmd)

	renderCmd.Flags().StringP("values", "i", "", "values file (YAML)")
	renderCmd.Flags().BoolP("env", "e", false, "add env vars")
	renderCmd.Flags().StringP("output", "o", "-", "output file")
	renderCmd.Flags().BoolP("recursive", "r", false, "recurse into subdirectories")
}
