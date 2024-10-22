package cli

import (
	"context"
	"fmt"
	"sync"

	"github.com/manifoldco/promptui"
	"github.com/uptrace/bun"
	"github.com/vinibgoulart/gitbook-llm/packages/page"
)

func Exec(ctx *context.Context, db *bun.DB, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	promptAiPrompt := promptui.Prompt{
		Label: "AI Prompt",
	}

	aiPrompt, aiPromptErr := promptAiPrompt.Run()
	if aiPromptErr != nil {
		panic(aiPromptErr)
	}

	page, err := page.GetEmbedded(ctx, db)(&aiPrompt)
	if err != nil {
		panic(err)
	}

	fmt.Println(page.Text)
}
