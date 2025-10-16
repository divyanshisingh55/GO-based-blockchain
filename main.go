package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
)

// Block structure
type Block struct {
	Index     int    `json:"index"`
	Timestamp string `json:"timestamp"`
	FileName  string `json:"file_name"`
	FileHash  string `json:"file_hash"`
	PrevHash  string `json:"prev_hash"`
	Hash      string `json:"hash"`
}

// Blockchain storage
var Blockchain []Block

// Calculate hash of the block
func calculateHash(block Block) string {
	record := fmt.Sprintf("%d%s%s%s%s", block.Index, block.Timestamp, block.FileName, block.FileHash, block.PrevHash)
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

// Generate a hash of a file
func getFileHash(filePath string) (string, error) {
	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(fileData)
	return hex.EncodeToString(hash[:]), nil
}

// Create a new block
func createBlock(prevBlock Block, fileName string, fileHash string) Block {
	newBlock := Block{
		Index:     prevBlock.Index + 1,
		Timestamp: time.Now().String(),
		FileName:  fileName,
		FileHash:  fileHash,
		PrevHash:  prevBlock.Hash,
	}
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

// Create the Genesis block
func createGenesisBlock() Block {
	return Block{Index: 0, Timestamp: time.Now().String(), FileName: "Genesis Block", FileHash: "0", PrevHash: "0", Hash: ""}
}

func main() {
	// Initialize blockchain with the Genesis block
	Blockchain = append(Blockchain, createGenesisBlock())

	// Start the API server
	r := gin.Default()
	r.Static("/static", "./static")

	// Serve the UI
	r.GET("/", func(c *gin.Context) {
		c.File("index.html")
	})

	gin.SetMode(gin.ReleaseMode) // Reduces unnecessary debug logs

	// Custom logging middleware
	r.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		fmt.Printf("[%s] %s %s | Status: %d | Time: %v\n",
			time.Now().Format("2006-01-02 15:04:05"),
			c.Request.Method, c.Request.URL.Path,
			c.Writer.Status(), duration)
	})

	// Upload a file and add it to the blockchain
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{"error": "No file uploaded"})
			return
		}

		// Save file locally
		filePath := "uploads/" + file.Filename
		c.SaveUploadedFile(file, filePath)

		// Get file hash
		fileHash, err := getFileHash(filePath)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to hash file"})
			return
		}

		// Add to blockchain
		newBlock := createBlock(Blockchain[len(Blockchain)-1], file.Filename, fileHash)
		Blockchain = append(Blockchain, newBlock)

		c.JSON(200, gin.H{"message": "File stored on blockchain", "file_name": file.Filename, "hash": newBlock.Hash})
	})

	// Get blockchain data
	r.GET("/blocks", func(c *gin.Context) {
		c.JSON(200, Blockchain)
	})

	// Run the API server
	r.Run(":8080")
}