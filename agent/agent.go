package agent

import (
	"context"
	"fmt"

	"github.com/openai/openai-go/v3"
)

type Agent struct {
	llm            *openai.Client
	getUserMessage func() (string, bool)
}

func New(llm *openai.Client, getUserMessage func() (string, bool)) *Agent {
	return &Agent{
		llm:            llm,
		getUserMessage: getUserMessage,
	}
}

func (a *Agent) Run(ctx context.Context) error {
	conversation := []openai.ChatCompletionMessageParamUnion{
		openai.SystemMessage("You are a fitness personal trainer and nutritionist called 'The Gains Father' you can also be refered to as Father Gains. You persona is similar to The God Father or that of The All Father (Odin)"),
	}

	fmt.Println("Chat with The Gains Father (use 'ctrl-c' to quit)")

	for {
		fmt.Print("\u001b[94mYou\u001b[0m: ")
		userInput, ok := a.getUserMessage()
		if !ok {
			break
		}
		userMessage := openai.UserMessage(userInput)
		conversation = append(conversation, userMessage)

		message, err := a.runInference(ctx, conversation)
		if err != nil {
			return err
		}
		conversation = append(conversation, message.Choices[0].Message.ToParam())
		for _, content := range message.Choices {
			fmt.Printf("\u001b[93mThe Gains Father\u001b[0m: %s\n", content.Message.Content)
		}
	}

	return nil
}

func (a *Agent) runInference(ctx context.Context, conversation []openai.ChatCompletionMessageParamUnion) (*openai.ChatCompletion, error) {

	message, err := a.llm.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Model:               openai.ChatModelGPT4o,
		MaxCompletionTokens: openai.Int(1024),
		Messages:            conversation,
	})
	return message, err
}
