package main

import (
    "fmt"
    "os"
    "github.com/zennittians/bls/ffi/go/bls"
    "github.com/zennittians/go-sdk/pkg/address" // Adjust the path based on your setup
)

type DeployAccount struct {
    Index        string
    Address      string
    BLSPublicKey string
}

func generateBLSKey() (bls.SecretKey, string) {
    var sk bls.SecretKey
    sk.SetByCSPRNG() // Generate BLS private key
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

    for i := 0; i < numAccounts; i++ {
        _, blsPubKey := generateBLSKey()  // Generate BLS keypair
        accountAddress, err := generateAddress(blsPubKey) // Generate address from the public key
        if err != nil {
            fmt.Println("Error generating address:", err)
            return
        }

        account := DeployAccount{
            Index:        fmt.Sprintf("%d", i), // Index as a string
            Address:      accountAddress,
            BLSPublicKey: blsPubKey,
        }

        accounts = append(accounts, account)
    }

    // Create the output file
    file, err := os.Create("generated_accounts.go")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    // Write the Go package and struct declaration
    file.WriteString("package genesis\n\nvar LocalIntelchainAccounts = []DeployAccount{\n")
    
    // Write each account's details
    for _, account := range accounts {
        file.WriteString(fmt.Sprintf(
            "{Index: \"%s\", Address: \"%s\", BLSPublicKey: \"%s\"},\n", 
            account.Index, account.Address, account.BLSPublicKey))
    }
    
    // Close the Go array declaration
    file.WriteString("}\n")

    fmt.Println("Generated accounts in generated_accounts.go")
}
