package server

import (
	"context"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "start server",
	Run: func(cmd *cobra.Command, args []string) {
		router, err := initializerRoute()
		if err != nil {
			panic("Failed to create router: " + err.Error())
		}

		err = router.Start(context.Background())
		if err != nil {
			panic("Failed to start router: " + err.Error())
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
