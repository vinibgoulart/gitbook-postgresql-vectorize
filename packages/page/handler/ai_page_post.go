package page

import (
	"context"
	"net/http"

	"github.com/uptrace/bun"
	"github.com/vinibgoulart/gitbook-llm/packages/page"
	"github.com/vinibgoulart/gitbook-llm/packages/utils"
	"github.com/vinibgoulart/zius"
)

type AiPrompt struct {
	Prompt string `json:"prompt"`
}

func AiPromptPost(ctx *context.Context, db *bun.DB) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		var aiPrompt AiPrompt

		errJsonDecode := utils.JsonDecode(res, req, &aiPrompt)
		if errJsonDecode != nil {
			http.Error(res, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		_, errValidate := zius.Validate(aiPrompt)
		if errValidate != nil {
			http.Error(res, errValidate.Error(), http.StatusBadRequest)
			return
		}

		page, err := page.GetEmbedded(ctx, db)(&aiPrompt.Prompt)
		if err != nil {

			return
		}

		res.Write([]byte(page.Text))
	}
}