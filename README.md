# gojudge CLI Tool
**gojudge**は、RareTECH Go Communityで扱う提出コードを問題のテストケースでAC/WA判定してくれるCLIツールです。

- 正答の場合
    - ✅ AC - <問題のTitle>

- 誤答の場合
    - ❌ WA - <問題のTitle>

## Usage
> [!IMPORTANT]
> Go 1.16以上の環境を用意してください。

### 1. gojudgeコマンドのセットアップ
```sh
go get github.com/anton-fuji/gojudge/cmd
```
or 

#### cloneする場合は以下の手順で進めてください。
```
git clone https://github.com/anton-fuji/gojudge.git
cd gojudge
```

#### 次に、依存関係のインストール
```sh
go mod tidy
```

#### ビルド
```sh
go build -o gojudge
```

### 2. 基本コマンド
```sh
./gojudge check <file.go> -p <problem_id>
```

**例**
```sh
// test.goファイルに問題を書き込んでいることを想定
./gojudge check test.go -p 1 
```

