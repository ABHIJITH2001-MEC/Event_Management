package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/openai/openai-go/v1"
)

func main() {
	router := gin.Default()

	// GET: /api/openai/usechatgpt
	router.GET("/api/openai/usechatgpt", UseChatGPT)

	router.Run(":8080")
}

func UseChatGPT(c *gin.Context) {
	doubt := c.Query("doubt")

	openaiClient := openai.NewClient("your_key_here")

	completionRequest := &openai.CompletionRequest{
		Prompt:    doubt,
		Model:     "davinci-codex", // Note: Use the appropriate model name for your use case
		MaxTokens: 1024,
	}

	completions, _, err := openaiClient.Completions.CreateCompletion(completionRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get completion from OpenAI"})
		return
	}

	var outputResult string
	for _, completion := range completions.Choices {
		outputResult += completion.Text
	}

	c.JSON(http.StatusOK, gin.H{"result": outputResult})
}
