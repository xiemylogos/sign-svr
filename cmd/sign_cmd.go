package cmd

import (
	"encoding/hex"
	"fmt"
	"github.com/howeyc/gopass"
	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/urfave/cli"
)

var SignCommand = cli.Command{
	Name:        "signTx",
	Usage:       "ont tx rawsign",
	Description: "sign rawTx use wallet",
	Action:      SignTx,
	Flags: []cli.Flag{
		TxRawFlag,
	},
}

var SignHashCommand = cli.Command{
	Name:        "signTxHash",
	Usage:       "ont tx hash sign",
	Description: "sign tx hash use wallet",
	Action:      SignTxHash,
	Flags: []cli.Flag{
		TxHashFlag,
	},
}

func SignTx(ctx *cli.Context) error {
	ontSdk := sdk.NewOntologySdk()
	rawTx := ctx.String(GetFlagName(TxRawFlag))
	mutableTx, err := ontSdk.GetMutableTx(rawTx)
	if err != nil {
		fmt.Printf("SignSvr GetMutableTx failed:%s", err)
		return fmt.Errorf("SignSvr GetMutableTx failed:%s", err)
	}
	optionFile := checkFileName(ctx)
	acc, err := OpenAccount(optionFile, ontSdk)
	if err != nil {
		fmt.Errorf("open account err:%s", err)
		return fmt.Errorf("open account err:%s", err)
	}
	err = ontSdk.SignToTransaction(mutableTx, acc)
	txData, err := ontSdk.GetTxData(mutableTx)
	if err != nil {
		fmt.Errorf("SignSvr GetTxData failed:%s", err)
		return fmt.Errorf("SignSvr GetTxData failed:%s", err)
	}
	fmt.Printf("signed tx: %s\n", txData)
	txn, err := utils.TransactionFromHexString(txData)
	if err != nil {
		fmt.Errorf("transaction from raw err:%s", err)
		return fmt.Errorf("transaction from raw err:%s", err)
	}
	var hash = txn.Hash()
	fmt.Printf("signed txHash: %s\n", hash.ToHexString())
	return nil
}

func SignTxHash(ctx *cli.Context) error {
	ontSdk := sdk.NewOntologySdk()
	txHash := ctx.String(GetFlagName(TxHashFlag))
	optionFile := checkFileName(ctx)
	acc, err := OpenAccount(optionFile, ontSdk)
	if err != nil {
		fmt.Errorf("open account err:%s", err)
		return fmt.Errorf("open account err:%s", err)
	}
	sigData, err := acc.Sign([]byte(txHash))
	if err != nil {
		return fmt.Errorf("sign error:%s", err)
	}
	fmt.Printf("signed txHash:%s\n",hex.EncodeToString(sigData))
	return nil
}

func checkFileName(ctx *cli.Context) string {
	if ctx.IsSet(GetFlagName(WalletFileFlag)) {
		return ctx.String(GetFlagName(WalletFileFlag))
	} else {
		return DEFAULT_WALLET_FILE_NAME
	}
}
func OpenAccount(path string, ontSdk *sdk.OntologySdk) (*sdk.Account, error) {
	wallet, err := ontSdk.OpenWallet(path)
	if err != nil {
		return nil, err
	}
	pwd, err := GetPassword()
	if err != nil {
		return nil, err
	}
	defer ClearPasswd(pwd)
	account, err := wallet.GetDefaultAccount(pwd)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func GetPassword() ([]byte, error) {
	fmt.Printf("Password:")
	passwd, err := gopass.GetPasswd()
	if err != nil {
		return nil, err
	}
	return passwd, nil
}

func ClearPasswd(passwd []byte) {
	size := len(passwd)
	for i := 0; i < size; i++ {
		passwd[i] = 0
	}
}
