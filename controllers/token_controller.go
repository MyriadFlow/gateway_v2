package controllers

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"net/http"
	"os"

	contract "app.myriadflow.com/abicontract" // ABI encoded code
	"app.myriadflow.com/db"
	"app.myriadflow.com/models"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

func CreateMainnetFanTokenRequest(c *gin.Context) {
	var req models.MainnetFanToken
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client, err := ethclient.Dial(os.Getenv("BASE_RPC_URL"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to Ethereum client"})
		return
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load private key"})
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cast public key to ECDSA"})
		return
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get nonce"})
		return
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get gas price"})
		return
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get chain ID"})
		return
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transactor"})
		return
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // Adjust as needed
	auth.GasPrice = gasPrice

	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	instance, err := contract.NewAbicontract(contractAddress, client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create contract instance"})
		return
	}

	// Convert nftContractAddress to common.Address
	nftContractAddr := common.HexToAddress(req.NFTContractAddress)

	data := common.FromHex(req.Data)

	tx, err := instance.CreateFanToken(auth, nftContractAddr, data, req.URI)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to call CreateFanToken: %v", err)})
		return
	}
	// Convert tx to JSON
	txJSON, err := json.MarshalIndent(tx, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling transaction to JSON: %v\n", err)
	} else {
		fmt.Printf("Transaction details:\n%s\n", string(txJSON))
	}
	
	// Save the response to the database
	req.TxHash = tx.Hash().Hex()
	result := db.DB.Create(&req)
	if result.Error != nil {
		fmt.Printf("Database error: %v\n", result.Error)
		// Check if it's a validation error
		if result.Error.Error() == "Error 1062: Duplicate entry" {
			c.JSON(http.StatusConflict, gin.H{"error": "Duplicate entry in database"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to save to database: %v", result.Error)})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "CreateFanToken transaction sent", "txHash": req.TxHash})
}
