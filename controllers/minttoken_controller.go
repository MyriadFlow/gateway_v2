package controllers

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"os"

	contract "app.myriadflow.com/abicontract"
	"app.myriadflow.com/db"
	"app.myriadflow.com/models"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

func DelegateMintFanToken(c *gin.Context) {
	var req models.DelegateMintFanTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Connect to Ethereum client
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
	// Create contract instance
	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))
	instance, err := contract.NewAbicontract(contractAddress, client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create contract instance"})
		return
	}

	// Get the token ID using PhygitalcontractAddrToToken
	nftContractAddr := common.HexToAddress(req.NFTContractAddress)
	tokenID, err := instance.PhygitalcontractAddrToToken(&bind.CallOpts{}, nftContractAddr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get token ID"})
		return
	}

	req.TokenID = tokenID.String()
	// Convert parameters
	creatorWallet := common.HexToAddress(req.CreatorWallet)
	amount, ok := new(big.Int).SetString(req.Amount, 10)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Amount"})
		return
	}
	data := common.FromHex(req.Data)

	// Call DelegateMintFanToken function
	tx, err := instance.DelegateMintFanToken(auth, creatorWallet, tokenID, amount, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to call DelegateMintFanToken: %v", err)})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to save to database: %v", result.Error)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "DelegateMintFanToken transaction sent", "txHash": req.TxHash})
}


func CreateMintFanToken(c *gin.Context) {
	var fantoken models.DelegateMintFanTokenRequest
	if err := c.ShouldBindJSON(&fantoken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Create(&fantoken).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fantoken)
}


func GetAllMintFanToken(c *gin.Context) {
	var fantoken []models.DelegateMintFanTokenRequest
	if err := db.DB.Find(&fantoken).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, fantoken)
}

func GetMintFanTokenByWalletAddress(c *gin.Context) {
	walletAddress := c.Param("creator_wallet")
	var fantoken []models.DelegateMintFanTokenRequest
	if err := db.DB.Where("creator_wallet = ?", walletAddress).Find(&fantoken).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(fantoken) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No MintFanToken found for the specified manager_id"})
		return
	}
	c.JSON(http.StatusOK, fantoken)
}

func UpdateMintFanToken(c *gin.Context) {
	id := c.Param("id")
	var fantoken models.DelegateMintFanTokenRequest
	if err := db.DB.First(&fantoken, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "MintFanToken not found"})
		return
	}

	if err := c.ShouldBindJSON(&fantoken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Save(&fantoken).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, fantoken)
}

func DeleteMintFanToken(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&models.DelegateMintFanTokenRequest{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "MintFanToken deleted successfully"})
}
