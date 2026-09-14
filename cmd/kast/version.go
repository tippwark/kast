package main

import (
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",

	PreRunE: func(cmd *cobra.Command, args []string) error {
		format, _ := cmd.Flags().GetString("format")
		if format != "text" && format != "json" && format != "yaml" {
			return fmt.Errorf("invalid format: %s. Valid formats are: text, json, yaml", format)
		}

		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		versionInfo := map[string]any{
			"version":     kastVersion,
			"compileDate": compileDate,
			"goVersion":   runtime.Version(),
		}

		var infoString string
		if format, _ := cmd.Flags().GetString("format"); format == "json" {
			info, _ := json.Marshal(versionInfo)
			infoString = string(info)
		} else if format == "yaml" {
			info, _ := yaml.Marshal(versionInfo)
			infoString = string(info)
		} else {
			infoString = fmt.Sprintf("Kast Version: %s\nCompile Date: %s\nGo Version: %s", kastVersion, compileDate, runtime.Version())
		}

		fmt.Println(infoString)
	},
}

func init() {
	versionCmd.Flags().StringP("format", "f", "text", "output format (text, json, yaml)")
	rootCmd.AddCommand(versionCmd)
}
