# jiractl
jiraをCLI経由で操作するコマンド

# Install
```
go get -u git@github.com:yuichi10/jiractl.git
```

# Usage
詳しくは`-h`オプションでみてください

## config
`config`コマンドはjiraに接続するための認証情報などを登録します。

**一連のログイン方法**
```sh
jiractl config set-credentials {cred name (anything)} -u {jira username}
jiractl config set-url {url name (anything)} --url {https://jira-api-url}
jiractl config set-context {context name (anything)} --url {url name} --user {cred name}
jiractl config use-context {context name}

# you can check config file at $Home/.jiractl.yaml
cat ~/.jiractl.yaml
```

## sprint
```sh
jiractl sprint issues -b "board name"
```


# 構成
clean architectureの勉強をかねて作ってみてた。

大まかな構成図は以下
```
├── cmd
│   ├── config
│   └── sprint
├── entity
├── infrastructure
│   └── view
├── interface
│   ├── api
│   ├── controller
│   ├── database
│   └── presenter
└── usecase
```

## 境界
|レイヤー|ディレクトリ名|
|:--|:--|
|Entities|entity|
|usecase|usecase|
|interface|interface|
|Frameworks & Drivers|infrastructure, cmd (cobra)|

## 処理の流れ
TBA
