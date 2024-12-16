package main

import (
	"fmt"
	"os"

	"github.com/intelchain-itc/bls/ffi/go/bls"
	"github.com/intelchain-itc/go-sdk/pkg/address" // Adjust the path based on your setup
)

type DeployAccount struct {
	Index        string
	Address      string
	BLSPublicKey string
	PrivateKey   string
}

func generateBLSKey() (bls.SecretKey, string) {
	var sk bls.SecretKey
	sk.SetByCSPRNG()                                // Generate BLS private key
	pubKey := sk.GetPublicKey().SerializeToHexStr() // Get the public key as a hex string
	return sk, pubKey
}

func generateAddress(pubKey string) (string, error) {
	// Convert the public key to bytes and generate the Bech32 address using the provided address package
	return address.ToBech32(address.Parse(pubKey)), nil
}

func main() {
	accounts := make([]DeployAccount, 0)
	numAccounts := 10 // Adjust the number of accounts you want to generate

	// Open the file to write public keys
	pubKeyFile, err := os.Create("public_keys.txt")
	if err != nil {
		fmt.Println("Error creating public keys file:", err)
		return
	}
	defer pubKeyFile.Close()

	// Open the file to write private keys
	privKeyDir := "credentials/bls/"
	err = os.MkdirAll(privKeyDir, 0755)
	if err != nil {
		fmt.Println("Error creating directory for private keys:", err)
		return
	}

	for i := 0; i < numAccounts; i++ {
		// Generate BLS keypair
		sk, blsPubKey := generateBLSKey()

		// Save private key to a file
		privKeyFilePath := fmt.Sprintf("%s%d.key", privKeyDir, i)
		privKeyFile, err := os.Create(privKeyFilePath)
		if err != nil {
			fmt.Println("Error creating private key file:", err)
			return
		}
		defer privKeyFile.Close()

		// Write the private key to the file
		privKeyFile.WriteString(sk.SerializeToHexStr())

		// Generate address from the public key
		accountAddress, err := generateAddress(blsPubKey)
		if err != nil {
			fmt.Println("Error generating address:", err)
			return
		}

		// Create a DeployAccount struct
		account := DeployAccount{
			Index:        fmt.Sprintf("%d", i), // Index as a string
			Address:      accountAddress,
			BLSPublicKey: blsPubKey,
			PrivateKey:   privKeyFilePath,
		}

		// Append the account to the list
		accounts = append(accounts, account)

		// Write the public key to the file
		pubKeyFile.WriteString(fmt.Sprintf("PublicKey for Account %d: %s\n", i, blsPubKey))
	}

	// Create the output Go file for the accounts
	file, err := os.Create("generated_accounts.go")
	if err != nil {
		fmt.Println("Error creating Go file:", err)
		return
	}
	defer file.Close()

	// Write the Go package and struct declaration
	file.WriteString("package genesis\n\nvar LocalIntelchainAccounts = []DeployAccount{\n")

	// Write each account's details
	for _, account := range accounts {
		file.WriteString(fmt.Sprintf(
			"{Index: \"%s\", Address: \"%s\", BLSPublicKey: \"%s\", PrivateKey: \"%s\"},\n",
			account.Index, account.Address, account.BLSPublicKey, account.PrivateKey))
	}

	// Close the Go array declaration
	file.WriteString("}\n")

	fmt.Println("Generated accounts and keys successfully")
}
