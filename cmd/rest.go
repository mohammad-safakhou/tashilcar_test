package cmd

import (
	"fmt"
	"log"
	"github.com/spf13/cobra"
	"github.com/labstack/echo/v4"
)

func init() {
	rootCmd.AddCommand(restCmd)
}

var restCmd = &cobra.Command{
	Use:   "rest-server",
	Short: "Starting rest server",
	Run: func(cmd *cobra.Command, args []string) {
		e := echo.New()


		log.Fatal(e.Start(":" + vConfig.GetString("server.port")))
	},
}

