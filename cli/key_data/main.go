package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/andantan/modular-blockchain/crypto"
	"log"
	"os"
)

func main() {
	privKeyFile := flag.String("k", "", "Path to the private key file (e.g., mykey.hex)")
	flag.Parse()

	if *privKeyFile == "" {
		flag.Usage()
		panic("❌  Private key file path is required. Please use the -k flag.")
	}

	keyHexBytes, err := os.ReadFile(*privKeyFile)
	if err != nil {
		log.Fatalf("❌  Failed to read private key file '%s': %v", *privKeyFile, err)
	}

	privKeyBytes, err := hex.DecodeString(string(keyHexBytes))
	if err != nil {
		log.Fatalf("❌  Failed to decode hex string from file: %v", err)
	}

	privKey := crypto.PrivateKeyFromBytes(privKeyBytes)
	privKeyHex := hex.EncodeToString(privKey.Bytes())
	pubKeyHex := hex.EncodeToString(privKey.PublicKey().Bytes())
	addressHex := hex.EncodeToString(privKey.PublicKey().Address().Bytes())
	fmt.Println("🔑  Key pair data")
	fmt.Println("==================================================================")
	fmt.Printf("Private Key: 0x%s\n", privKeyHex)
	fmt.Printf("Public Key : 0x%s\n", pubKeyHex)
	fmt.Printf("Address: 0x%s\n", addressHex)
	fmt.Println("==================================================================")
	fmt.Println("⚠️  IMPORTANT: Store your private key in a secure location and never share it.")
}
