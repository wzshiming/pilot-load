package cmd

import (
	"github.com/howardjohn/pilot-load/pkg/simulation"
	"github.com/howardjohn/pilot-load/pkg/simulation/model"
	"github.com/spf13/cobra"
)

var adscConfig = model.AdscConfig{}

func init() {
	adscCmd.PersistentFlags().IntVar(&adscConfig.Count, "count", adscConfig.Count, "number of adsc connections to make")
}

var adscCmd = &cobra.Command{
	Use:   "adsc",
	Short: "open simple ADS connection to Istiod",
	RunE: func(cmd *cobra.Command, _ []string) error {
		args, err := GetArgs()
		if err != nil {
			return err
		}
		args.AdsConfig = adscConfig
		logConfig(args.AdsConfig)
		return simulation.Adsc(args)
	},
}