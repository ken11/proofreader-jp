# proofreader-jp
<img width="760" alt="proofreader-jp" src="https://github.com/ken11/proofreader-jp/assets/2043460/adc898bd-2a73-4984-8a9b-f9b1cacb3665">

## これなに
日本語の校正をChatGPTにお願いしちゃおうツール  
記者ハン基準でよしなにがんばってくれ！ってしてるだけなので、どの程度かはわからないが、誤字脱字とかは結構指摘してくれる印象(個人の感想です)  

## 使い方
```sh
$ go install github.com/ken11/proofreader-jp@latest
$ OPENAI_API_KEY=hogehoge proofreader-jp -f fuga.md
```

## 各種オプション
```sh
-f 解析したいファイルパスを指定
-s 指摘事項がない行も出力
-model 実行したいモデルを指定(デフォルトGPT-4)
```


