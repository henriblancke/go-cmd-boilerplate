package cmd

import (
	"fmt"
	"os"

	"github.com/henriblancke/go-cmd-boilerplate/pkg/adaptors"
	"github.com/henriblancke/go-cmd-boilerplate/pkg/log"
	"github.com/henriblancke/go-cmd-boilerplate/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	tracker = metrics.NewMetrics()
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "boil",
	Short: "This is a command line application",
	Run: func(cmd *cobra.Command, args []string) {
		tracker.CliCounter.With(prometheus.Labels{"command": "boil"}).Inc()

		message, _ := cmd.Flags().GetString("message")

		service := adaptors.HelloService{}
		user, err := service.Create(message)
		if err != nil {
			log.Fatal(err)
		}

		msg, err := service.SayMessage((user))
		if err != nil {
			log.Fatal(err)
		}

		log.Info(msg)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		push.New(viper.GetString("pushgateway.url"), viper.GetString("name")).Gatherer(tracker.Registry).Push()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()

	rootCmd.PersistentFlags().StringP("message", "m", "", "message for user")
	rootCmd.MarkPersistentFlagRequired("message")
}
