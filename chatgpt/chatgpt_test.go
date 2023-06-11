package chatgpt

import (
	"context"
	"os"
	"testing"

	"github.com/xpzouying/gollm"
)

func TestChatGPT(t *testing.T) {

	var (
		token = os.Getenv("OPENAI_TOKEN")
	)

	c := NewChatGPT(token)

	answer, err := c.Chat(context.Background(), []gollm.LlmMessage{
		{
			Role:    gollm.RoleUser,
			Content: "Tell me Kobe in 10 words.",
		},
	})
	if err != nil {
		t.Errorf("Chat with chatgpt failed: %v", err)
	}

	a := answer.Content
	if a == "" {
		t.Error("Chat with chatgpt got empty answer")
	}

	t.Log(a)
}
