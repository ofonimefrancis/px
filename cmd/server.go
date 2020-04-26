package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/ofonimefrancis/pixels/pkg/datastore/driver"
	"github.com/ofonimefrancis/pixels/pkg/graphql"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	IsProduction    bool   `mapstructure:"is_production"`
	AppName         string `mapstructure:"app_name"`
	Port            string `mapstructure:"port"`
	DBDriver        string `mapstructure:"database_driver"`
	DatabaseName    string `mapstructure:"database_name"`
	DatabaseURI     string `mapstructure:"database_uri"`
	DatabaseTimeout int    `mapstructure:"database_timeout"`
}

var config Config

var rootCmd = &cobra.Command{
	Use:   "pixels",
	Short: "This is the API",
	Long:  `Long Description will go here`,
	Run: func(cmd *cobra.Command, args []string) {
		driver := driver.NewConnectionDriver()
		//pass the driver to setup services
		router := chi.NewMux()
		graphQLFacade := graphql.NewGraphQLFacade(driver)
		graphQLFacade.RegisterRoutes(router)

		log.Println("Application running on http://localhost:8000")
		http.ListenAndServe(":8000", router)
	},
}

func Execute() {
	var cfgFile string
	cobra.OnInitialize(initConfig(cfgFile))
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pixels.yaml)")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig(cfgFile string) func() {
	return func() {
		if cfgFile != "" {
			viper.SetConfigFile(cfgFile)
		} else {
			home, err := homedir.Dir()
			if err != nil {
				log.Fatal(err)
			}

			viper.AddConfigPath(home)
			viper.AddConfigPath(".")
			viper.SetConfigName(".pixels")
		}

		//viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}

		err := viper.Unmarshal(&config)
		if err != nil {
			log.Panicf("unable to decode into struct, %v", err)
		}

		fmt.Println(config)
	}

}
