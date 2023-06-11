# Go LLM

Package for each llm, like chatgpt or others.


## Run example


```go

func main() {
    token := "sk-xxx"

    c := NewChatGPT(token)

	answer, err := c.Chat(context.Background(), []gollm.LlmMessage{
        {
            Role:    gollm.RoleUser,
			Content: "Tell me Kobe in 10 words.",
		},
	})


    log.Println(answer.Content)
}
```

```bash
go run .

# Legend, champion, scoring machine, skilled, charming, determined, Mamba mentality.
```