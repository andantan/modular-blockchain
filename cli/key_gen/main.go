package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/andantan/modular-blockchain/crypto"
	"os"
)

func main() {
	outputFile := flag.String("o", "", "Output file path for the private key (e.g., mykey.hex)")
	flag.Parse()

	privKey, err := crypto.GeneratePrivateKey()
	if err != nil {
		panic(fmt.Sprintf("❌  Failed to generate private key: %v", err))
	}

	privKeyHex := hex.EncodeToString(privKey.Bytes())

	if *outputFile != "" {
		if err = os.WriteFile(*outputFile, []byte(privKeyHex), 0600); err != nil {
			panic(fmt.Sprintf("❌  Failed to write private key to file: %v", err))
		}
		fmt.Printf("✅  Private key saved successfully to '%s'.\n", *outputFile)
	} else {
		pubKeyHex := hex.EncodeToString(privKey.PublicKey().Bytes())
		addressHex := hex.EncodeToString(privKey.PublicKey().Address().Bytes())
		fmt.Println("🔑  New key pair generated successfully!")
		fmt.Println("==================================================================")
		fmt.Printf("Private Key: %s\n", privKeyHex)
		fmt.Printf("Public Key : %s\n", pubKeyHex)
		fmt.Printf("Address: %s\n", addressHex)
		fmt.Println("==================================================================")
		fmt.Println("⚠️  IMPORTANT: Store your private key in a secure location and never share it.")
	}
}
