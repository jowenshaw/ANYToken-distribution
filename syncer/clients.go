package syncer

import (
	"context"
	"math/big"

	"github.com/anyswap/ANYToken-distribution/log"
	"github.com/fsn-dev/fsn-go-sdk/efsn/common"
	"github.com/fsn-dev/fsn-go-sdk/efsn/core/types"
	"github.com/fsn-dev/fsn-go-sdk/efsn/ethclient"
)

var (
	clients    []*ethclient.Client
	cliContext = context.Background()
)

func dialServer() (err error) {
	var client *ethclient.Client
	for _, url := range serverURL {
		client, err = ethclient.Dial(url)
		if err != nil {
			log.Error("[syncer] client connection error", "server", url, "err", err)
			return err
		}
		log.Info("[syncer] client connection succeed", "server", url)
		clients = append(clients, client)
	}
	return nil
}

func closeClient() {
	for _, client := range clients {
		if client != nil {
			client.Close()
		}
	}
}

func getHeaderByNumber(number *big.Int) (header *types.Header, err error) {
	for _, client := range clients {
		header, err = client.HeaderByNumber(cliContext, number)
		if err == nil {
			return
		}
	}
	return
}

func getBlockByNumber(number *big.Int) (block *types.Block, err error) {
	for _, client := range clients {
		block, err = client.BlockByNumber(cliContext, number)
		if err == nil {
			return
		}
	}
	return
}

func getTransactionReceipt(txHash common.Hash) (receipt *types.Receipt, err error) {
	for _, client := range clients {
		receipt, err = client.TransactionReceipt(cliContext, txHash)
		if err == nil {
			return
		}
	}
	return
}
