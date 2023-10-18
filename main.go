package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

type Proof struct {
	Input  string
	Result string
}

func main() {
	var filePath = flag.String("f", "", "Target file path")
	var printSuccess = flag.Bool("s", false, "Print success line")
	var model = flag.String("model", "gpt-4", "Model name")
	flag.Parse()

	var m string
	switch *model {
	case "gpt-4", "gpt4":
		m = openai.GPT4
	case "gpt-3.5", "gpt3.5":
		m = openai.GPT3Dot5Turbo
	default:
		m = ""
	}

	fp, err := os.Open(*filePath)
	if err != nil {
		fmt.Println("file cannot open")
		panic(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	line := 0

	for scanner.Scan() {
		line += 1
		if strings.TrimSpace(scanner.Text()) == "" {
			continue
		}
		if len(scanner.Text()) == len([]rune(scanner.Text())) {
			continue
		}

		c := newChatGPTClient()
		p, err := c.requestChatGPT(scanner.Text(), m)
		if err != nil {
			fmt.Println(err.Error())
			panic(err)
		}

		if p.Result == "なし" || p.Result == "" {
			if *printSuccess == true {
				fmt.Printf("\x1b[34m%s行目: %s\n指摘内容: 指摘事項はありませんでした\n\x1b[0m", strconv.Itoa(line), p.Input)
			}
			continue
		}
		fmt.Printf("\x1b[31m%s行目: %s\n指摘内容: %s\n\x1b[0m", strconv.Itoa(line), p.Input, p.Result)
	}
}
