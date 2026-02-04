# 📋 要件定義書

## プロジェクト情報

| 項目 | 内容 |
|------|------|
| プロジェクト名 | go-review-cli |
| 英語名 | Project Scaffolding Tool |
| バージョン | v1.0.0 |
| リポジトリ | go-review-cli |
| プロジェクトパス | `\Project\GoLang\review\go-review-cli` |
| 作成日 | 2026-01-27 |

---

## プロジェクト概要

開発者がプロジェクトを素早く立ち上げるための対話型CLIツール。統一されたプロジェクト構造、Git Flow管理、カスタムテンプレートサポートを提供し、将来的にマイクロサービスアーキテクチャの基盤となる。

### 背景・目的
- プロジェクト立ち上げの時間短縮
- プロジェクト構造の統一化（バグの温床削減）
- Git Flow管理の簡素化
- 将来的な多言語マイクロサービスエコシステムの基盤構築

### スコープ
**v1.0.0の範囲**
- Go言語プロジェクトのテンプレート生成
- 基本的なGit管理機能
- カスタムテンプレート管理
- 設定ファイル管理

**将来のスコープ（v2.0.0以降）**
- Rust、Python、TypeScript等の多言語対応
- プラグインシステム
- CI/CD設定生成
- マイクロサービスオーケストレーション

---

## 機能要件

### 1. プロジェクト生成機能

#### 1.1 対話型プロジェクト作成
**説明**: ユーザーと対話しながらプロジェクトを作成

**入力項目**
- プロジェクト名（必須、英数字とハイフンのみ）
- プロジェクトタイプ（必須、選択式）
  - Go API (RESTful API)
  - Go CLI (Command Line Tool)
  - Go Microservice (gRPC Service)
- Webフレームワーク（APIタイプの場合）
  - Gin
  - Echo
- モジュール名（必須、Go module形式）
- 説明文（任意）
- 作成パス（デフォルトあり、変更可能）

**処理フロー**
1. プロジェクト名のバリデーション
2. プロジェクトタイプの選択
3. 追加設定の入力
4. 確認プロンプト表示
5. プロジェクト生成実行

**出力**
- 生成されたディレクトリパス
- 生成されたファイル一覧
- 次のステップのガイド

#### 1.2 非対話型プロジェクト作成
**説明**: コマンドラインフラグで全設定を指定

**コマンド例**
```bash
go-review-cli create myproject \
  --type api \
  --framework gin \
  --module github.com/user/myproject \
  --description "My API project" \
  --path /path/to/projects
```

**フラグ**
- `--type, -t`: プロジェクトタイプ
- `--framework, -f`: Webフレームワーク
- `--module, -m`: モジュール名
- `--description, -d`: 説明文
- `--path, -p`: 作成パス
- `--no-git`: Git初期化をスキップ

#### 1.3 テンプレートベース生成

**Go API テンプレート**
- Gin または Echo フレームワーク
- Clean Architecture構造
- ルーティング設定
- ミドルウェア（CORS、Logger）
- 環境変数管理（.env）
- Makefile
- Docker対応（Dockerfile、docker-compose.yml）

**Go CLI テンプレート**
- Cobra フレームワーク
- サブコマンド構造
- 設定ファイル管理（Viper）
- フラグ管理
- Makefile

**Go Microservice テンプレート**
- gRPC サーバー/クライアント
- Protocol Buffers定義
- Clean Architecture構造
- ヘルスチェックエンドポイント
- Docker対応
- Makefile

#### 1.4 ディレクトリ構造生成

**共通構造**
```
project-name/
├── cmd/                    # エントリーポイント
│   └── [app-name]/
│       └── main.go
├── internal/               # プライベートパッケージ
│   ├── domain/            # ドメインロジック
│   ├── usecase/           # ユースケース
│   ├── handler/           # ハンドラー（API/CLI）
│   └── infrastructure/    # インフラ実装
├── pkg/                   # 公開パッケージ
├── test/                  # テストファイル
├── configs/               # 設定ファイル
├── scripts/               # スクリプト
├── docs/                  # ドキュメント
├── .gitignore
├── go.mod
├── go.sum
├── Makefile
├── README.md
└── .env.example
```

---

### 2. Git管理機能

#### 2.1 Git初期化
**説明**: プロジェクト生成時にGitリポジトリを初期化

**処理内容**
1. `git init` 実行
2. `.gitignore` ファイル生成（Go用）
3. Git Flow用ブランチ作成
   - `main` ブランチ（本番リリース用）
   - `develop` ブランチ（開発用）
4. 初回コミット作成
   - コミットメッセージ: `🎉 chore: Initial commit`

**オプション**
- `--no-git`: Git初期化をスキップ
- `--remote [url]`: リモートリポジトリを設定

#### 2.2 Git Hooks設定
**説明**: Conventional Commits対応のprepare-commit-msgフックを配置

**Hooksの内容**
- 対話型コミットメッセージ選択
- プレフィックス付与
  - ✨ feat: 新機能
  - 🐛 fix: バグ修正
  - 📝 docs: ドキュメント
  - 🎨 style: フォーマット
  - ♻️ refactor: リファクタリング
  - ⚡️ perf: パフォーマンス改善
  - ✅ test: テスト追加
  - 🔧 chore: その他

**配置パス**
```
.git/hooks/prepare-commit-msg
```

---

### 3. テンプレート管理機能

#### 3.1 組み込みテンプレート
**説明**: CLIに同梱される3種類の基本テンプレート

**テンプレート一覧**
1. `go-api`: Go RESTful API（Gin/Echo）
2. `go-cli`: Go CLI Tool（Cobra）
3. `go-microservice`: Go gRPC Microservice

**格納場所**
- バイナリに埋め込み（`embed`ディレクティブ使用）
- または `~/.go-review-cli/templates/` にコピー

#### 3.2 カスタムテンプレート追加
**説明**: ユーザー定義のテンプレートを追加

**コマンド**
```bash
go-review-cli template add my-template /path/to/template
```

**テンプレート構造**
```
my-template/
├── template.yaml          # テンプレートメタデータ
├── files/                 # テンプレートファイル
│   ├── cmd/
│   ├── internal/
│   └── ...
└── README.md
```

**template.yaml 例**
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

**変数置換**
- `{{.ProjectName}}`: プロジェクト名
- `{{.ModuleName}}`: モジュール名
- `{{.Description}}`: 説明文
- `{{.Author}}`: 作成者
- `{{.Year}}`: 現在の年

#### 3.3 テンプレート一覧表示
**コマンド**
```bash
go-review-cli template list
```

**出力形式**
```
Available Templates:
  [built-in] go-api          - Go RESTful API
  [built-in] go-cli          - Go CLI Tool
  [built-in] go-microservice - Go gRPC Microservice
  [custom]   my-template     - My custom template
```

#### 3.4 テンプレート削除
**コマンド**
```bash
go-review-cli template remove my-template
```

**制約**
- 組み込みテンプレートは削除不可
- 確認プロンプト表示

---

### 4. 設定管理機能

#### 4.1 グローバル設定
**説明**: ユーザーごとのデフォルト設定を管理

**設定項目**
- `project.default_path`: デフォルトのプロジェクト作成パス
- `project.default_framework`: デフォルトのWebフレームワーク
- `git.user_name`: Gitユーザー名
- `git.user_email`: Gitメールアドレス
- `git.enable_hooks`: Git Hooksを有効化（デフォルト: true）
- `template.directory`: カスタムテンプレートディレクトリ
- `ui.color_enabled`: カラー出力を有効化（デフォルト: true）

#### 4.2 設定ファイル形式
**ファイルパス**
```
~/.go-review-cli/config.yaml
```

**設定ファイル例**
```yaml
project:
  default_path: ~/projects
  default_framework: gin

git:
  user_name: Your Name
  user_email: your.email@example.com
  enable_hooks: true

template:
  directory: ~/.go-review-cli/templates

ui:
  color_enabled: true
```

#### 4.3 設定コマンド

**設定値取得**
```bash
go-review-cli config get project.default_path
# Output: ~/projects
```

**設定値設定**
```bash
go-review-cli config set project.default_path /path/to/projects
# Output: Configuration updated: project.default_path = /path/to/projects
```

**全設定表示**
```bash
go-review-cli config list
# Output: 全設定をYAML形式で表示
```

---

### 5. CLIコマンド設計

#### 5.1 コマンド一覧

```bash
# プロジェクト生成
go-review-cli create [project-name]                      # 対話型
go-review-cli create [project-name] [flags]              # 非対話型

# テンプレート管理
go-review-cli template list                              # テンプレート一覧
go-review-cli template add [name] [path]                 # テンプレート追加
go-review-cli template remove [name]                     # テンプレート削除
go-review-cli template show [name]                       # テンプレート詳細表示

# 設定管理
go-review-cli config get [key]                           # 設定取得
go-review-cli config set [key] [value]                   # 設定更新
go-review-cli config list                                # 全設定表示
go-review-cli config reset                               # 設定リセット

# その他
go-review-cli version                                    # バージョン表示
go-review-cli help                                       # ヘルプ表示
go-review-cli help [command]                             # コマンドヘルプ
```

#### 5.2 グローバルフラグ
- `--verbose, -v`: 詳細ログ出力
- `--quiet, -q`: 最小限の出力
- `--no-color`: カラー出力を無効化
- `--config [path]`: 設定ファイルパス指定

---

## 非機能要件

### パフォーマンス
| 項目 | 要件 |
|------|------|
| プロジェクト生成時間 | 5秒以内（通常サイズのテンプレート） |
| 対話型UIの応答性 | 100ms以内 |
| 起動時間 | 500ms以内 |
| メモリ使用量 | 50MB以内 |

### 拡張性
- 新言語テンプレート追加可能な設計
- プラグインシステムの考慮（インターフェース設計）
- カスタムコマンド追加の仕組み（v2.0.0以降）

### ユーザビリティ
- カラフルで見やすい出力
- プログレスバー表示（長時間処理時）
- エラーメッセージは明確で修正方法を提示
- Tab補完対応（bash/zsh）
- 一貫性のあるコマンド構造

### 互換性
- **OS**: Windows, macOS, Linux
- **Go**: 1.21以上
- **Git**: 2.0以上
- **シェル**: bash, zsh, PowerShell

### セキュリティ
- ファイルパスのバリデーション（パストラバーサル防止）
- 既存ファイルの上書き前に確認
- Git認証情報を平文で保存しない

### 保守性
- 単体テストカバレッジ 80%以上
- 統合テストの実装
- ドキュメントの整備
- CI/CDパイプライン構築

---

## 技術スタック

### コアライブラリ
| ライブラリ | バージョン | 用途 |
|-----------|----------|------|
| github.com/spf13/cobra | v1.8+ | CLIフレームワーク |
| github.com/spf13/viper | v1.18+ | 設定管理 |
| github.com/AlecAivazis/survey/v2 | v2.3+ | インタラクティブプロンプト |
| text/template | 標準ライブラリ | テンプレートエンジン |

### サポートライブラリ
| ライブラリ | バージョン | 用途 |
|-----------|----------|------|
| github.com/go-git/go-git/v5 | v5.11+ | Git操作 |
| github.com/fatih/color | v1.16+ | カラー出力 |
| github.com/spf13/afero | v1.11+ | ファイルシステム抽象化 |
| github.com/stretchr/testify | v1.8+ | テストフレームワーク |
| github.com/schollz/progressbar/v3 | v3.14+ | プログレスバー |

---

## データ設計

### 設定ファイル構造
**ファイルパス**: `~/.go-review-cli/config.yaml`

```yaml
project:
  default_path: string       # デフォルトのプロジェクト作成パス
  default_framework: string  # デフォルトのWebフレームワーク

git:
  user_name: string          # Gitユーザー名
  user_email: string         # Gitメールアドレス
  enable_hooks: boolean      # Git Hooksを有効化

template:
  directory: string          # カスタムテンプレートディレクトリ

ui:
  color_enabled: boolean     # カラー出力を有効化
```

### テンプレートメタデータ
**ファイル名**: `template.yaml`

```yaml
name: string               # テンプレート名
version: string            # テンプレートバージョン
description: string        # 説明
language: string           # 対応言語（go, rust, python等）
type: string               # タイプ（api, cli, microservice等）
variables:                 # 変数定義
  - name: string           # 変数名
    description: string    # 説明
    required: boolean      # 必須フラグ
    default: string        # デフォルト値（任意）
```

---

## エラーハンドリング

### エラー分類
1. **ユーザー入力エラー**
   - 不正なプロジェクト名
   - 存在しないテンプレート指定
   - 不正な設定値

2. **システムエラー**
   - ファイルシステムエラー
   - Git操作エラー
   - テンプレート読み込みエラー

3. **ネットワークエラー**
   - リモートリポジトリ接続エラー（将来対応）

### エラーメッセージ形式
```
❌ Error: [エラー内容]

Reason: [詳細な理由]

Solution: [解決方法]
```

**例**
```
❌ Error: Invalid project name 'my project'

Reason: Project name contains whitespace

Solution: Use hyphens instead of spaces (e.g., 'my-project')
```

---

## テスト要件

### 単体テスト
- 各レイヤーの個別テスト
- モックを使用したテスト
- テーブル駆動テスト
- カバレッジ: 80%以上

### 統合テスト
- エンドツーエンドのプロジェクト生成テスト
- Git操作の統合テスト
- テンプレート管理の統合テスト

### テストケース例
1. 正常系: プロジェクトが正しく生成される
2. 異常系: 既存ディレクトリへの生成
3. 境界値: 長いプロジェクト名
4. エッジケース: 特殊文字を含むパス

---

## 開発計画（Git Flow）

### ブランチ戦略
- `main`: 本番リリースブランチ
- `develop`: 開発ブランチ
- `feature/*`: 機能開発ブランチ
- `release/*`: リリース準備ブランチ
- `hotfix/*`: 緊急修正ブランチ

### マイルストーン

#### v0.1.0 - 基盤実装（2週目）
- [ ] プロジェクト構造セットアップ
- [ ] Cobra統合
- [ ] Viper統合
- [ ] 基本的なコマンド構造
- [ ] CI/CD設定

#### v0.2.0 - プロジェクト生成機能（3-4週目）
- [ ] テンプレートエンジン実装
- [ ] Go APIテンプレート
- [ ] ファイル生成機能
- [ ] ディレクトリ構造生成

#### v0.3.0 - Git統合（5週目）
- [ ] Git初期化機能
- [ ] Git Flow対応
- [ ] Hooks設定
- [ ] .gitignore生成

#### v0.4.0 - テンプレート管理（6週目）
- [ ] カスタムテンプレート追加
- [ ] テンプレート一覧・削除
- [ ] テンプレートメタデータ処理

#### v0.5.0 - 設定管理（7週目）
- [ ] グローバル設定機能
- [ ] 設定の永続化
- [ ] 設定コマンド実装

#### v0.6.0 - インタラクティブUI（8週目）
- [ ] Survey統合
- [ ] プログレスバー
- [ ] カラー出力
- [ ] エラーメッセージ改善

#### v1.0.0 - 最終調整（9-10週目）
- [ ] 全機能統合テスト
- [ ] ドキュメント整備
- [ ] バイナリビルド
- [ ] リリースノート作成

---

## 制約事項

### v1.0.0の制約
- **対応言語**: Goのみ（Rust対応はv2.0.0以降）
- **テンプレート**: 3種類の固定テンプレート + カスタム
- **Git操作**: 基本的な初期化とブランチ作成のみ
- **プラグイン**: 未対応

### 技術的制約
- Go 1.21以上必須
- Git 2.0以上必須
- ファイルシステムへの書き込み権限必須

### 将来拡張予定（v2.0.0以降）
- 多言語対応（Rust, Python, TypeScript, Java等）
- プラグインシステム
- CI/CD設定生成（GitHub Actions, GitLab CI等）
- Dockerファイル自動生成
- Kubernetes マニフェスト生成
- マイクロサービスオーケストレーション機能
- リモートテンプレートリポジトリ対応
- テンプレート共有機能

---

## 成功基準

### v1.0.0リリース条件
- [ ] 全機能が要件通り実装されている
- [ ] 単体テストカバレッジ 80%以上
- [ ] 統合テストが全てパスする
- [ ] ドキュメントが完備されている
- [ ] 3つのOSでビルド・動作確認完了
- [ ] パフォーマンス要件を満たす

### ユーザー体験目標
- プロジェクト立ち上げ時間を80%削減
- Git Flow管理の学習コストを50%削減
- プロジェクト構造の統一により、新規メンバーのオンボーディング時間を30%短縮

---

## 参考資料

### 参照ドキュメント
- Cobra: https://github.com/spf13/cobra
- Viper: https://github.com/spf13/viper
- Survey: https://github.com/AlecAivazis/survey
- Go Template: https://pkg.go.dev/text/template

### 競合製品分析
- `cookiecutter`: Python製、多言語対応
- `yeoman`: Node.js製、Web開発向け
- `cargo generate`: Rust製、Rust特化

---

## 変更履歴

| バージョン | 日付 | 変更内容 |
|-----------|------|---------|
| 1.0.0 | 2026-01-27 | 初版作成 |