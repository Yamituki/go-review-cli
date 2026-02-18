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
- 💡 **対話型UI**: 使いやすいインタラクティブプロンプト

---

## 📦 インストール

### ソースからビルド

```bash
git clone https://github.com/Yamituki/go-review-cli.git
cd go-review-cli
go build -o bin/go-review-cli cmd/go-review-cli/main.go
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
└── README.md
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
利用可能なテンプレート:
  [built-in] go-api               - Go RESTful API template
  [built-in] go-cli               - Go CLI tool template
  [built-in] go-microservice      - Go microservice template
  [custom]   my-template          - User-defined template
```

#### テンプレート詳細表示

```bash
go-review-cli template show go-api
```

#### カスタムテンプレート追加

任意のディレクトリをカスタムテンプレートとして追加できます。  
追加したテンプレートは `~/.go-review-cli/templates/[name]` にコピーされます。

```bash
go-review-cli template add my-template /path/to/template-dir
```

**制約**:
- 組み込みテンプレート（go-api, go-cli, go-microservice）と同名は追加不可
- 同名のカスタムテンプレートが既に存在する場合は追加不可

#### テンプレート削除

```bash
go-review-cli template remove my-template
# エイリアス: rm, delete
```

**制約**:
- 組み込みテンプレートは削除不可
- 削除前に確認プロンプトが表示される

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
- Conventional Commits対応のGit Hooksを設定（prepare-commit-msg）
- 初回コミットを作成（`🎉 chore: Initial commit`）

#### Git Hooks

コミット時に対話型でプレフィックスを選択：

```
╔════════════════════════════════════════════════╗
║      コミットタイプを選択してください          ║
╚════════════════════════════════════════════════╝

  1️⃣  ✨ feat      新機能
  2️⃣  🐛 fix       バグ修正
  3️⃣  📝 docs      ドキュメント
  4️⃣  🎨 style     スタイル
  5️⃣  ♻️  refactor  リファクタリング
  6️⃣  ⚡️ perf      パフォーマンス
  7️⃣  ✅ test      テスト
  8️⃣  🔧 chore     その他
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

### セットアップ

```bash
# リポジトリをクローン
git clone https://github.com/Yamituki/go-review-cli.git
cd go-review-cli

# 依存関係をインストール
go mod download

# ビルド
go build -o bin/go-review-cli cmd/go-review-cli/main.go

# テスト実行
go test ./...

# リリース用ビルド（バージョン情報埋め込み）
go build -ldflags "-X github.com/Yamituki/go-review-cli/pkg/version.GitCommit=$(git rev-parse HEAD) -X github.com/Yamituki/go-review-cli/pkg/version.BuildDate=$(date -u +%Y-%m-%dT%H:%M:%SZ)" -o bin/go-review-cli cmd/go-review-cli/main.go
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

---

## 🛠️ 技術スタック

| カテゴリ | ライブラリ | 用途 |
|---------|-----------|------|
| CLI Framework | [Cobra](https://github.com/spf13/cobra) | コマンド構造 |
| Configuration | [Viper](https://github.com/spf13/viper) | 設定管理 |
| Interactive UI | [Survey](https://github.com/AlecAivazis/survey) | プロンプト |
| Git Operations | [go-git](https://github.com/go-git/go-git) | Git操作 |
| File System | [Afero](https://github.com/spf13/afero) | FS抽象化 |

---

## 🗺️ ロードマップ

### v1.0.0 ✅
- [x] Go言語プロジェクトテンプレート（API/CLI/Microservice）
- [x] Git Flow統合（初期化、ブランチ作成、Hooks）
- [x] カスタムテンプレート管理（list/show/add/remove）
- [x] 対話型UI（Survey）
- [x] 設定管理（config get/set/list/reset）
- [x] versionコマンド

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

---

## 🤝 コントリビューション

コントリビューションを歓迎します！

### コントリビューション方法

1. このリポジトリをFork
2. Feature ブランチを作成 (`git checkout -b feature/amazing-feature`)
3. 変更をCommit (`git commit -m '✨ feat: Add amazing feature'`)
4. ブランチをPush (`git push origin feature/amazing-feature`)
5. Pull Requestを作成

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

## 📄 ライセンス

このプロジェクトは[MITライセンス](LICENSE)の下で公開されています。

---

## 👤 作成者

**闇月**

- GitHub: [@Yamituki](https://github.com/Yamituki)

---

**Happy Coding! 🚀**