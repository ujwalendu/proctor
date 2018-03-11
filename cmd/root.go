package cmd

import (
	"fmt"
	"os"

	"github.com/gojekfarm/proctor/cmd/jobs"
	"github.com/gojekfarm/proctor/cmd/version"
	"github.com/gojekfarm/proctor/config"
	"github.com/gojekfarm/proctor/engine"
	"github.com/gojekfarm/proctor/io"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:     "proctor",
		Long:    `A command-line tool to interact with proctor-engine, the heart of Proctor: An Automation Framework`,
		Version: "Ipsum Lorem",
	}
)

func Execute(printer io.Printer, proctorEngineClient engine.Client) {
	cobra.OnInitialize(config.InitConfig)

	versionCmd := version.NewCmd(printer)
	rootCmd.AddCommand(versionCmd)
	rootCmd.SetVersionTemplate("😊  Proctor: An Automation Framework v0.1.0\n")

	jobCmd := jobs.NewCmd(printer, proctorEngineClient)
	rootCmd.AddCommand(jobCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
