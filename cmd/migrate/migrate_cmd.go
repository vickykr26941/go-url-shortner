package migrate

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	config2 "github.com/vickykumar/url_shortner/internal/config"
	"github.com/vickykumar/url_shortner/internal/database"
)

// migrateUpCmd represents the "migrate up" command
var migrateUpCmd = &cobra.Command{
	Use:   "migrate:up",
	Short: "Run all pending SQL migrations",
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := config2.Load("CONFIG_JSON")
		if err != nil {
			panic("Failed to load configuration: " + err.Error())
		}
		mysql, err := database.NewMySqlConnection(&config.Database)
		if err != nil {
			panic("Failed to connect to DB: " + err.Error())
		}
		return runMigrationsUp(mysql.GetDB())
	},
}

func init() {
	migrateUpCmd.AddCommand(migrateUpCmd)
}
