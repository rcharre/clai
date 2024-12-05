package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/mistral"
)

func main() {
	key, exists := os.LookupEnv("MISTRAL_API_KEY")
	if !exists {
		panic("Should have a MISTRAL_API_KEY en var set")
	}

	model, exists := os.LookupEnv("MISTRAL_MODEL")
	if !exists {
		model = "codestral-2405"
	}

	args := os.Args
	sb := strings.Builder{}
	pipe := isPipe()
	if !pipe && len(args) < 2 {
		panic("should pipe or have an arg")
	}

	if len(args) == 2 {
		sb.WriteString(args[1])
	}

	if pipe {
		in, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		str := string(in)

		// show input before response
		fmt.Println(str)

		sb.WriteString("\n")
		sb.WriteString(str)
	}

	llm, err := mistral.New(mistral.WithAPIKey(key))
	if err != nil {
		panic(err)
	}

	prompt := sb.String()
	ctx := context.Background()
	res, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt,
		llms.WithTemperature(0.2),
		llms.WithModel(model),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func isPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}
