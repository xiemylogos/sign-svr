package cmd

import (
	"encoding/hex"
	"fmt"
	"github.com/howeyc/gopass"
	"github.com/ontio/ontology-crypto/keypair"
	"github.com/ontio/ontology-crypto/signature"
	sdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology-go-sdk/utils"
	"github.com/urfave/cli"
)

var SignCommand = cli.Command{
	Name:        "signtx",
	Usage:       "ont tx rawsign",
	Description: "sign rawTx use wallet",
	Action:      SignTx,
	Flags: []cli.Flag{
		AddrFlag,
		PayerFlag,
		TxRawFlag,
	},
}

var SignHashCommand = cli.Command{
	Name:        "signtxhash",
	Usage:       "ont tx hash sign",
	Description: "sign tx hash use wallet",
	Action:      SignTxHash,
	Flags: []cli.Flag{
		AddrFlag,
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
	acc, err := OpenAccount(optionFile, ontSdk, ctx.String(GetFlagName(AddrFlag)))
	if err != nil {
		fmt.Errorf("open account err:%s", err)
		return fmt.Errorf("open account err:%s", err)
	}
	err = ontSdk.SignToTransaction(mutableTx, acc)
	if err != nil {
		return fmt.Errorf("sign tx err:%s", err)
	}
	if ctx.String(GetFlagName(PayerFlag)) != "" {
		payer, err := OpenAccount(optionFile, ontSdk, ctx.String(GetFlagName(PayerFlag)))
		if err != nil {
			fmt.Errorf("open account err:%s", err)
			return fmt.Errorf("open account err:%s", err)
		}
		err = ontSdk.SignToTransaction(mutableTx, payer)
		if err != nil {
			return fmt.Errorf("sign tx err:%s", err)
		}
	}
	for _, sig := range mutableTx.Sigs {
		sigdata := sig.SigData[0]
		fmt.Printf("publickey:= %s\n", hex.EncodeToString(keypair.SerializePublicKey(sig.PubKeys[0])))

		fmt.Printf("sigdata:= %s\n", hex.EncodeToString(sigdata))
	}
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
	optionFile := checkFileName(ctx)
	txHash, err := hex.DecodeString(ctx.String(GetFlagName(TxHashFlag)))
	if err != nil {
		fmt.Errorf("DecodeString txHash err:%s\n", err)
		return fmt.Errorf("decode string err:%sn", err)
	}
	acc, err := OpenAccount(optionFile, ontSdk, ctx.String(GetFlagName(AddrFlag)))
	if err != nil {
		fmt.Errorf("open account err:%s", err)
		return fmt.Errorf("open account err:%s", err)
	}
	sigData, err := signature.Sign(acc.SigScheme, acc.PrivateKey, txHash, nil)
	if err != nil {
		return fmt.Errorf("sign error:%s", err)
	}
	if !signature.Verify(acc.PublicKey, txHash, sigData) {
		fmt.Println("verify sign err")
		return fmt.Errorf("verify sign err")
	}
	buf, err := signature.Serialize(sigData)
	if err != nil {
		fmt.Println("sig serialize err")
		return fmt.Errorf("sig serialize err")
	}
	fmt.Printf("signed txHash:%s\n", hex.EncodeToString(buf))
	return nil
}

func checkFileName(ctx *cli.Context) string {
	if ctx.IsSet(GetFlagName(WalletFileFlag)) {
		return ctx.String(GetFlagName(WalletFileFlag))
	} else {
		return DEFAULT_WALLET_FILE_NAME
	}
}
func OpenAccount(path string, ontSdk *sdk.OntologySdk, addr string) (*sdk.Account, error) {
	wallet, err := ontSdk.OpenWallet(path)
	if err != nil {
		return nil, err
	}
	pwd, err := GetPassword()
	if err != nil {
		return nil, err
	}
	defer ClearPasswd(pwd)
	account, err := wallet.GetAccountByAddress(addr, pwd)
	//account, err := wallet.GetDefaultAccount(pwd)
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
