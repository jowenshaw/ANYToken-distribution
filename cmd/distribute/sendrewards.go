package main

import (
	"fmt"

	"github.com/anyswap/ANYToken-distribution/cmd/utils"
	"github.com/anyswap/ANYToken-distribution/distributer"
	"github.com/anyswap/ANYToken-distribution/log"
	"github.com/urfave/cli/v2"
)

var (
	sendRewardsCommand = &cli.Command{
		Action:    sendRewards,
		Name:      "sendrewards",
		Usage:     "send rewards batchly",
		ArgsUsage: " ",
		Description: `
send rewards batchly according to verified input file with line format: <address> <rewards>
`,
		Flags: []cli.Flag{
			utils.GatewayFlag,
			utils.RewardTyepFlag,
			utils.ExchangeFlag,
			utils.RewardTokenFlag,
			utils.StartHeightFlag,
			utils.EndHeightFlag,
			utils.InputFileFlag,
			utils.SenderFlag,
			utils.KeyStoreFileFlag,
			utils.PasswordFileFlag,
			utils.GasLimitFlag,
			utils.GasPriceFlag,
			utils.AccountNonceFlag,
			utils.OutputFileFlag,
			utils.SaveDBFlag,
			utils.DryRunFlag,
		},
	}
)

func sendRewards(ctx *cli.Context) error {
	serverURL := ctx.String(utils.GatewayFlag.Name)
	if serverURL == "" {
		return fmt.Errorf("must specify gateway URL")
	}

	capi := utils.InitAppWithURL(ctx, true, serverURL)
	distributer.SetAPICaller(capi)

	opt, err := getOptionAndTxArgs(ctx)
	if err != nil {
		log.Fatalf("get option error: %v", err)
	}

	defer capi.CloseClient()
	return opt.SendRewardsFromFile()
}
