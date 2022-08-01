package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/crypto/sha3"

	"github.com/tse-lao/ether-user/wallet"
	user "github.com/tse-lao/ether-user/wallet"
)

// App struct
type App struct {
	ctx context.Context
}

type Connections struct {
	IPFS     bool
	Ethereum bool
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Connect() string {

	//now we implement to the ethereum dial

	client, err := ethclient.Dial("wss://rpc-mainnet.matic.network")

	if err != nil {
		return fmt.Sprint(err)
	}
	_ = client

	//so now connected so we wqant tor ead somehting.
	account := common.HexToAddress("0xc94d737b36A32BbC4eaf545832C05420fa11B916")
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s", balance)

}

type Wallet struct {
	PrivateKey string
	PublicKey  string
	Address    string
}

//now we can generate new wallets for on the ethereum network
func (a *App) GenerateWallet() Wallet {

	//create new wallet here
	newWallet := Wallet{}
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		fmt.Println("Error generating key:", err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("Private Key Bytes: ", hexutil.Encode(privateKeyBytes)[2:]) // fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19
	newWallet.PrivateKey = hexutil.Encode(privateKeyBytes)[2:]

	publicKey := privateKey.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("Public Key bytes", hexutil.Encode(publicKeyBytes)[4:])

	newWallet.PublicKey = hexutil.Encode(publicKeyBytes)[4:]

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	fmt.Println("Address", address)

	newWallet.Address = address

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])

	return newWallet
	//now we want to create something that returns all the three different items as an array to the user.
}

//we would like to make sure that we can access and share the private keys together with new clients.

type Transact struct {
}

func (a *App) TransferFunds() {
	client, err := ethclient.Dial("http://127.0.0.1:7545")

	if err != nil {
		fmt.Println("this is the error", err)
		return
	}

	privateKey, err := crypto.HexToECDSA("b48555ad69c01f125c884ed5e70384920e859b53d0eaa0076d5f2f1ed9ea0c38")
	if err != nil {
		fmt.Println("this is the error", err)
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println("this is the error", err)
		return
	}

	value := big.NewInt(5000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)                // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("this is the error", err)
		return
	}

	toAddress := common.HexToAddress("0x23e2408c6d509aF73B77f5d6E44b6Ba79BC657E4")
	var data []byte

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println("this is the error", err)
		return
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		fmt.Println("this is the error", err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println("this is the error", err)
	}

	fmt.Println("tx sent: %s", signedTx.Hash().Hex())
}

func (a *App) GetBalance(address string) string {
	client, err := ethclient.Dial("http://127.0.0.1:7545")

	if err != nil {
		return fmt.Sprint(err)
	}
	_ = client

	//so now connected so we wqant tor ead somehting.
	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return fmt.Sprintf("The error is %s", err)
	}
	return fmt.Sprintf("The current balance is: %s", balance)
}

//on the startup we need to to getAccount.

func (a *App) GetAccount(password string) []user.Wallet {
	accountwallet := user.Wallet{}

	// now we want to read ls

	//just create the one that is latest settled.

	result := wallet.GetAccount(password)

	accountwallet = result

	wallets := []user.Wallet{accountwallet}
	return wallets
}

func (a *App) GenerateAccount(password string) user.Wallet {
	//Let's try and generate this connection via the files.
	result := user.Login("privatekey")

	//check if there is a folder with tempirorlay access.

	//return a notification of the user.Wallet.
	//which contains.
	//public and private key.

	//wallet.CreateNewAccount(password)

	//this creates a file that should be readable. Lets check where it is located
	//if not we need to create new pairing.

	return result
}

func (a *App) OpenMessage() {
	selection, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select File",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Images (*.png;*.jpg)",
				Pattern:     "*.png;*.jpg",
			}, {
				DisplayName: "Videos (*.mov;*.mp4)",
				Pattern:     "*.mov;*.mp4",
			},
		},
	})

	if err != nil {
		fmt.Println("Error")
	}

	fmt.Println(selection)

}

func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	dialog, err := runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type:    runtime.QuestionDialog,
		Title:   "Quit?",
		Message: "Are you sure you want to quit?",
	})

	if err != nil {
		return false
	}
	return dialog != "Yes"
}
