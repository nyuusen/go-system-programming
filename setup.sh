#!/bin/bash

# 1. 引数チェック（引数がない場合は使い方を表示して終了）
if [ -z "$1" ]; then
    echo "Usage: $0 <suffix> (e.g., $0 2-1)"
    exit 1
fi

# 2. ディレクトリ名の組み立て
DIR_NAME="ch$1"

# 3. ディレクトリの作成
# -p オプションをつけることで、既に存在してもエラーにならず、必要なら親ディレクトリも作ります
mkdir -p "$DIR_NAME"

# 4. 空の main.go を作成
# touch はファイルがなければ作成し、あればタイムスタンプを更新します
touch "$DIR_NAME/main.go"

echo "package main\n\nfunc main() {\n\t\n}" > "$DIR_NAME/main.go"

echo "Created: $DIR_NAME/main.go"