package cmd

import (
	"github.com/urfave/cli"
	"strings"
)

const (
	DEFAULT_WALLET_FILE_NAME = "./wallet.dat"
	DEFAULT_TX_RAW           = ""
	DEFAULT_TX_HASH
)

var (
	TxRawFlag = cli.StringFlag{
		Name:  "rawTx",
		Usage: "raw tx",
		Value: DEFAULT_TX_RAW,
	}
	TxHashFlag = cli.StringFlag{
		Name:  "txHash",
		Usage: "tx hash",
		Value: DEFAULT_TX_HASH,
	}
	WalletFileFlag = cli.StringFlag{
		Name:  "wallet,w",
		Value: DEFAULT_WALLET_FILE_NAME,
		Usage: "Wallet `<file>`",
	}
)

func GetFlagName(flag cli.Flag) string {
	name := flag.GetName()
	if name == "" {
		return ""
	}
	return strings.TrimSpace(strings.Split(name, ",")[0])
}
