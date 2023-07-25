package cmd

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shamimice03/golang-gin-mysql-gorm/pkg/config"
	"github.com/shamimice03/golang-gin-mysql-gorm/pkg/routing"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Server app on dev server",
	Long:  `Application will be served on host and port defined in config.yaml file`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {

	config.Set()
	routing.Init()
	router := routing.GetRouter()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":  "pong",
			"app name": viper.Get("App.Name"),
		})
	})

	routing.Serve()
}
