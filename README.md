# go-review-cli

> A powerful project scaffolding CLI tool for rapid development setup

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Version](https://img.shields.io/badge/version-1.0.0-green.svg)](https://github.com/Yamituki/go-review-cli/releases)

---

## 📋 概要

**go-review-cli** は、開発者がプロジェクトを素早く立ち上げるための対話型CLIツールです。統一されたプロジェクト構造、Git Flow管理、カスタムテンプレートサポートを提供し、将来的にマイクロサービスアーキテクチャの基盤となります。

### 主な特徴

- 🚀 **高速プロジェクト生成**: 数秒でプロジェクトを立ち上げ
- 🎯 **統一された構造**: Clean Architectureに基づいた一貫性のある設計
- 🔧 **Git Flow統合**: ブランチ戦略とConventional Commits対応
- 🎨 **カスタマイズ可能**: 独自のテンプレート追加が可能
- 🌐 **多言語対応準備**: Go以外の言語にも拡張可能な設計
- 💡 **対話型UI**: 使いやすいインタラクティブプロンプト

---

## 🎯 ユースケース

### プロジェクト立ち上げの時間短縮
従来30分かかっていたプロジェクトセットアップが5秒で完了

### チーム開発の統一化
全プロジェクトで同じ構造を採用し、新規メンバーのオンボーディングを加速

### 将来のマイクロサービス基盤
複数の言語・フレームワークで統一されたサービスを生成可能

---

## 📦 インストール

### バイナリから（推奨）

```bash
# 最新版をダウンロード
curl -L https://github.com/Yamituki/go-review-cli/releases/latest/download/go-review-cli-$(uname -s)-$(uname -m) -o go-review-cli

# 実行権限を付与
chmod +x go-review-cli

# パスに追加
sudo mv go-review-cli /usr/local/bin/
```

### go installから

```bash
go install github.com/Yamituki/go-review-cli@latest
```

### ソースからビルド

```bash
git clone https://github.com/Yamituki/go-review-cli.git
cd go-review-cli
make build
```

---

## 🚀 クイックスタート

### 対話型モード

```bash
go-review-cli create
```

プロンプトに従って、プロジェクト名、タイプ、フレームワークなどを選択するだけ！

### 非対話型モード

```bash
go-review-cli create my-api \
  --type api \
  --framework gin \
  --module github.com/user/my-api \
  --description "My awesome API"
```

### 生成されるプロジェクト構造

```
my-api/
├── cmd/
│   └── my-api/
│       └── main.go
├── internal/
│   ├── domain/
│   ├── usecase/
│   ├── handler/
│   └── infrastructure/
├── pkg/
├── test/
├── configs/
├── docs/
├── .gitignore
├── go.mod
├── Makefile
├── README.md
└── .env.example
```

---

## 📖 使い方

### プロジェクト生成

#### Go API（RESTful）

```bash
# Ginフレームワーク
go-review-cli create my-api --type api --framework gin --module github.com/user/my-api

# Echoフレームワーク
go-review-cli create my-api --type api --framework echo --module github.com/user/my-api
```

#### Go CLI Tool

```bash
go-review-cli create my-cli --type cli --module github.com/user/my-cli
```

#### Go Microservice（gRPC）

```bash
go-review-cli create my-service --type microservice --module github.com/user/my-service
```

---

### テンプレート管理

#### テンプレート一覧

```bash
go-review-cli template list
```

**出力例**:
```
Available Templates:
  [built-in] go-api          - Go RESTful API
  [built-in] go-cli          - Go CLI Tool
  [built-in] go-microservice - Go gRPC Microservice
  [custom]   my-template     - My custom template
```

#### カスタムテンプレート追加

```bash
go-review-cli template add my-template /path/to/template
```

**テンプレート構造**:
```
my-template/
├── template.yaml          # メタデータ
├── files/                 # テンプレートファイル
│   ├── cmd/
│   ├── internal/
│   └── ...
└── README.md
```

**template.yaml 例**:
```yaml
name: my-template
version: 1.0.0
description: My custom template
language: go
type: api
variables:
  - name: ProjectName
    description: Project name
    required: true
  - name: ModuleName
    description: Go module name
    required: true
```

#### テンプレート削除

```bash
go-review-cli template remove my-template
```

---

### 設定管理

#### 設定一覧

```bash
go-review-cli config list
```

#### 設定取得

```bash
go-review-cli config get project.default_path
```

#### 設定更新

```bash
go-review-cli config set project.default_path ~/projects
go-review-cli config set project.default_framework gin
go-review-cli config set git.user_name "Your Name"
go-review-cli config set git.user_email "your@email.com"
```

#### 設定リセット

```bash
go-review-cli config reset
```

---

### Git統合

プロジェクト生成時に自動的に：
- Git リポジトリを初期化
- `main` と `develop` ブランチを作成
- Conventional Commits対応のGit Hooksを設定
- 初回コミットを作成（`🎉 chore: Initial commit`）

#### Git Hooks

コミット時に対話型でプレフィックスを選択：

```
🎯 コミットタイプを選択:
  1) ✨ feat: 新機能
  2) 🐛 fix: バグ修正
  3) 📝 docs: ドキュメント
  4) 🎨 style: フォーマット
  5) ♻️ refactor: リファクタリング
  6) ⚡️ perf: パフォーマンス
  7) ✅ test: テスト
  8) 🔧 chore: その他
```

#### Git初期化をスキップ

```bash
go-review-cli create my-project --no-git
```

---

## ⚙️ 設定ファイル

設定ファイルは `~/.go-review-cli/config.yaml` に保存されます。

```yaml
project:
  default_path: ~/projects
  default_framework: gin

git:
  user_name: Your Name
  user_email: your@email.com
  enable_hooks: true

template:
  directory: ~/.go-review-cli/templates

ui:
  color_enabled: true
```

---

## 🔧 開発

### 必要な環境

- Go 1.21以上
- Git 2.0以上
- Make

### セットアップ

```bash
# リポジトリをクローン
git clone https://github.com/Yamituki/go-review-cli.git
cd go-review-cli

# 依存関係をインストール
go mod download

# ビルド
make build

# テスト実行
make test

# ローカルで実行
go run cmd/go-review-cli/main.go
```

### Makefileコマンド

```bash
make build        # ビルド
make test         # テスト実行
make test-cover   # カバレッジ付きテスト
make lint         # リント
make clean        # クリーンアップ
make install      # インストール
make build-all    # 全OS向けビルド
```

---

## 🏗️ アーキテクチャ

Clean Architectureに基づいた4層構造：

```
┌─────────────────────────────────┐
│   Presentation Layer (CLI)      │
│   - Cobra Commands               │
│   - Interactive Prompts          │
└─────────────┬───────────────────┘
              │
┌─────────────▼───────────────────┐
│   Application Layer (Use Cases) │
│   - CreateProject                │
│   - ManageTemplate               │
│   - ManageConfig                 │
└─────────────┬───────────────────┘
              │
┌─────────────▼───────────────────┐
│   Domain Layer (Business Logic) │
│   - Entities                     │
│   - Value Objects                │
│   - Domain Services              │
└─────────────┬───────────────────┘
              │
┌─────────────▼───────────────────┐
│   Infrastructure Layer           │
│   - Git Operations               │
│   - File System                  │
│   - Config Storage               │
└─────────────────────────────────┘
```

詳細は[アーキテクチャ設計書](docs/architecture.md)を参照してください。

---

## 📚 ドキュメント

- [要件定義書](docs/requirements.md)
- [アーキテクチャ設計書](docs/architecture.md)
- [詳細設計書](docs/detailed_design.md)

---

## 🛠️ 技術スタック

| カテゴリ | ライブラリ | 用途 |
|---------|-----------|------|
| CLI Framework | [Cobra](https://github.com/spf13/cobra) | コマンド構造 |
| Configuration | [Viper](https://github.com/spf13/viper) | 設定管理 |
| Interactive UI | [Survey](https://github.com/AlecAivazis/survey) | プロンプト |
| Git Operations | [go-git](https://github.com/go-git/go-git) | Git操作 |
| File System | [Afero](https://github.com/spf13/afero) | FS抽象化 |
| Color Output | [Color](https://github.com/fatih/color) | カラー表示 |
| Progress Bar | [progressbar](https://github.com/schollz/progressbar) | 進捗表示 |
| Testing | [Testify](https://github.com/stretchr/testify) | テスト |

---

## 🗺️ ロードマップ

### v1.0.0 ✅
- [x] Go言語プロジェクトテンプレート
- [x] Git Flow統合
- [x] カスタムテンプレート管理
- [x] 対話型UI
- [x] 設定管理

### v2.0.0 (計画中)
- [ ] Rust言語対応
- [ ] Python言語対応
- [ ] プラグインシステム
- [ ] CI/CD設定生成
- [ ] リモートテンプレートリポジトリ

### v3.0.0 (計画中)
- [ ] マイクロサービスオーケストレーション
- [ ] Kubernetes マニフェスト生成
- [ ] Docker Compose自動生成
- [ ] サービスメッシュ統合

---

## 🤝 コントリビューション

コントリビューションを歓迎します！

### コントリビューション方法

1. このリポジトリをFork
2. Feature ブランチを作成 (`git checkout -b feature/amazing-feature`)
3. 変更をCommit (`git commit -m '✨ feat: Add amazing feature'`)
4. ブランチをPush (`git push origin feature/amazing-feature`)
5. Pull Requestを作成

### 開発ガイドライン

- Conventional Commitsに従う
- テストを書く（カバレッジ80%以上）
- ドキュメントを更新する
- コードレビューを受ける

---

## 📝 Git Flow

このプロジェクトはGit Flowを採用しています：

- `main`: 本番リリースブランチ
- `develop`: 開発ブランチ
- `feature/*`: 機能開発ブランチ
- `release/*`: リリース準備ブランチ
- `hotfix/*`: 緊急修正ブランチ

---

## 🐛 バグ報告

バグを見つけた場合は、[Issues](https://github.com/Yamituki/go-review-cli/issues)で報告してください。

**バグ報告に含めるもの**:
- 再現手順
- 期待される動作
- 実際の動作
- 環境情報（OS、Goバージョン等）

---

## 💡 機能リクエスト

新機能のアイデアがある場合は、[Issues](https://github.com/Yamituki/go-review-cli/issues)で提案してください。

---

## 📄 ライセンス

このプロジェクトは[MITライセンス](LICENSE)の下で公開されています。

---

## 👤 作成者

**闇月**

- GitHub: [@Yamituki](https://github.com/Yamituki)

---

## 🙏 謝辞

このプロジェクトは以下のオープンソースプロジェクトに影響を受けています：

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Cookiecutter](https://github.com/cookiecutter/cookiecutter) - Project templates
- [Yeoman](https://yeoman.io/) - Scaffolding tool

---

## 📊 統計

![Lines of Code](https://img.shields.io/tokei/lines/github/Yamituki/go-review-cli)
![Code Size](https://img.shields.io/github/languages/code-size/Yamituki/go-review-cli)
![Last Commit](https://img.shields.io/github/last-commit/Yamituki/go-review-cli)

---

## 🔗 関連プロジェクト

- [go-review-logagg](https://github.com/Yamituki/go-review-logagg) - ログ集約ツール
- 今後、マイクロサービスプロジェクトが追加予定

---

**Happy Coding! 🚀**