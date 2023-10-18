package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

type proof struct {
	Input  string
	Result string
}

type chatGPTClient struct{}

func newChatGPTClient() *chatGPTClient {
	return &chatGPTClient{}
}

// System Message
const sys = `## 命令
あなたは校正のプロです。
「入力テキスト」に対して、記者ハンドブックをベースに、日本語表記のルールについてレビューを行ってください。
なお、レビュー結果は次の「出力」に記載のJSONフォーマットに従って出力してください。
レビュー結果に指摘する点がない場合、レビュー結果は「なし」としてください。

## 出力
{"input": 入力テキストの内容, "result": レビュー結果}`

func (c *chatGPTClient) requestChatGPT(text string, model string) (*proof, error) {
	apiKey, ok := os.LookupEnv("OPENAI_API_KEY")
	if !ok {
		return &proof{"", ""}, fmt.Errorf("missing api key")
	}
	client := openai.NewClient(apiKey)
	input := "## 入力テキスト\n" + text

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: sys,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: input,
				},
			},
		},
	)

	if err != nil {
		return &proof{"", ""}, fmt.Errorf("ChatGPT request error")
	}

	res := resp.Choices[0].Message.Content
	res = strings.Replace(res, "## 出力\n", "", -1)

	var p proof

	if err := json.Unmarshal([]byte(res), &p); err != nil {
		return &proof{"", ""}, fmt.Errorf("json parse error")
	}

	return &p, nil
}
