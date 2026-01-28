# 🏗️ アーキテクチャ設計書

## プロジェクト情報

| 項目 | 内容 |
|------|------|
| プロジェクト名 | go-review-cli |
| バージョン | v1.0.0 |
| 作成日 | 2026-01-27 |

---

## 目次
1. [アーキテクチャ概要](#アーキテクチャ概要)
2. [レイヤー設計](#レイヤー設計)
3. [ディレクトリ構造](#ディレクトリ構造)
4. [コンポーネント設計](#コンポーネント設計)
5. [データフロー](#データフロー)
6. [依存関係](#依存関係)
7. [設計パターン](#設計パターン)
8. [拡張性設計](#拡張性設計)

---

## アーキテクチャ概要

### アーキテクチャスタイル
**Clean Architecture + Hexagonal Architecture（ポート&アダプター）**

### アーキテクチャ原則
1. **依存性逆転の原則（DIP）**: 上位レイヤーは下位レイヤーに依存しない
2. **単一責任の原則（SRP）**: 各コンポーネントは単一の責任を持つ
3. **インターフェース分離の原則（ISP）**: クライアントに不要な依存を強制しない
4. **開放/閉鎖原則（OCP）**: 拡張に開いて、修正に閉じる

### レイヤー構成図
```
┌─────────────────────────────────────────────────────────┐
│                   Presentation Layer                     │
│                      (CLI Commands)                      │
│  ┌─────────────────────────────────────────────────┐   │
│  │  Cobra Commands  │  Survey Prompts  │  Output  │   │
│  └─────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────┘
                          │ (uses)
                          ▼
┌─────────────────────────────────────────────────────────┐
│                    Application Layer                     │
│                       (Use Cases)                        │
│  ┌─────────────────────────────────────────────────┐   │
│  │ CreateProject │ ManageTemplate │ ManageConfig  │   │
│  └─────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────┘
                          │ (uses)
                          ▼
┌─────────────────────────────────────────────────────────┐
│                     Domain Layer                         │
│                  (Business Logic)                        │
│  ┌─────────────────────────────────────────────────┐   │
│  │  Project  │  Template  │  Config  │  Generator │   │
│  └─────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────┘
                          │ (uses)
                          ▼
┌─────────────────────────────────────────────────────────┐
│                  Infrastructure Layer                    │
│               (External Dependencies)                    │
│  ┌─────────────────────────────────────────────────┐   │
│  │  Git  │  FileSystem  │  ConfigStore  │  Logger │   │
│  └─────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────┘
```

---

## レイヤー設計

### 1. Presentation Layer（プレゼンテーション層）
**責務**: ユーザーインターフェース、入出力処理

#### コンポーネント
- **Cobra Commands**: CLIコマンドの定義とルーティング
- **Survey Prompts**: インタラクティブプロンプト
- **Output Formatter**: 出力フォーマット（カラー、プログレスバー）

#### 依存関係
- Application Layer（Use Cases）に依存
- Domain Layerには直接依存しない

#### 実装場所
```
internal/cli/
├── commands/           # Cobraコマンド
├── prompts/            # インタラクティブプロンプト
└── output/             # 出力フォーマッター
```

---

### 2. Application Layer（アプリケーション層）
**責務**: ユースケースの実装、ビジネスフローの調整

#### コンポーネント
- **Use Cases**: ビジネスロジックの調整役
  - `CreateProjectUseCase`: プロジェクト生成
  - `ManageTemplateUseCase`: テンプレート管理
  - `ManageConfigUseCase`: 設定管理

#### 依存関係
- Domain Layerに依存
- Infrastructure Layerのインターフェースに依存（実装には依存しない）

#### 実装場所
```
internal/usecase/
├── create_project.go
├── manage_template.go
└── manage_config.go
```

---

### 3. Domain Layer（ドメイン層）
**責務**: ビジネスルール、ドメインロジック

#### コンポーネント
- **Entities（エンティティ）**: ビジネスの核となるオブジェクト
  - `Project`: プロジェクト情報
  - `Template`: テンプレート情報
  - `Config`: 設定情報

- **Value Objects（値オブジェクト）**: 不変のビジネス値
  - `ProjectName`: プロジェクト名
  - `ModuleName`: モジュール名
  - `ProjectType`: プロジェクトタイプ

- **Domain Services（ドメインサービス）**: エンティティに属さないビジネスロジック
  - `ProjectGenerator`: プロジェクト生成ロジック
  - `TemplateProcessor`: テンプレート処理ロジック

- **Repository Interfaces（リポジトリインターフェース）**: データ永続化の抽象化
  - `ProjectRepository`: プロジェクト操作
  - `TemplateRepository`: テンプレート操作
  - `ConfigRepository`: 設定操作

#### 依存関係
- 他のレイヤーに依存しない（完全に独立）
- 標準ライブラリのみ使用

#### 実装場所
```
internal/domain/
├── entity/             # エンティティ
│   ├── project.go
│   ├── template.go
│   └── config.go
├── value/              # 値オブジェクト
│   ├── project_name.go
│   ├── module_name.go
│   └── project_type.go
├── service/            # ドメインサービス
│   ├── project_generator.go
│   └── template_processor.go
└── repository/         # リポジトリインターフェース
    ├── project_repository.go
    ├── template_repository.go
    └── config_repository.go
```

---

### 4. Infrastructure Layer（インフラストラクチャ層）
**責務**: 外部システムとの連携、技術的詳細

#### コンポーネント
- **Repository Implementations**: リポジトリの具体実装
  - `FileSystemProjectRepository`
  - `FileSystemTemplateRepository`
  - `ViperConfigRepository`

- **External Services**: 外部サービスとの連携
  - `GitService`: Git操作（go-git使用）
  - `FileSystemService`: ファイルシステム操作（afero使用）

- **Logger**: ロギング実装

#### 依存関係
- Domain Layerのインターフェースを実装
- 外部ライブラリに依存

#### 実装場所
```
internal/infrastructure/
├── repository/         # リポジトリ実装
│   ├── filesystem_project_repository.go
│   ├── filesystem_template_repository.go
│   └── viper_config_repository.go
├── git/                # Git操作
│   └── git_service.go
├── filesystem/         # ファイルシステム
│   └── fs_service.go
└── logger/             # ロガー
    └── logger.go
```

---

## ディレクトリ構造

### 完全なディレクトリ構造
```
go-review-cli/
├── cmd/
│   └── go-review-cli/
│       └── main.go                         # エントリーポイント
│
├── internal/
│   ├── cli/                                # Presentation Layer
│   │   ├── commands/
│   │   │   ├── root.go                    # ルートコマンド
│   │   │   ├── create.go                  # createコマンド
│   │   │   ├── template.go                # templateコマンド
│   │   │   ├── config.go                  # configコマンド
│   │   │   └── version.go                 # versionコマンド
│   │   ├── prompts/
│   │   │   ├── project_prompt.go          # プロジェクト作成プロンプト
│   │   │   └── template_prompt.go         # テンプレート選択プロンプト
│   │   └── output/
│   │       ├── formatter.go               # 出力フォーマッター
│   │       ├── progress.go                # プログレスバー
│   │       └── color.go                   # カラー出力
│   │
│   ├── usecase/                            # Application Layer
│   │   ├── create_project.go              # プロジェクト生成ユースケース
│   │   ├── manage_template.go             # テンプレート管理ユースケース
│   │   └── manage_config.go               # 設定管理ユースケース
│   │
│   ├── domain/                             # Domain Layer
│   │   ├── entity/
│   │   │   ├── project.go                 # プロジェクトエンティティ
│   │   │   ├── template.go                # テンプレートエンティティ
│   │   │   └── config.go                  # 設定エンティティ
│   │   ├── value/
│   │   │   ├── project_name.go            # プロジェクト名値オブジェクト
│   │   │   ├── module_name.go             # モジュール名値オブジェクト
│   │   │   ├── project_type.go            # プロジェクトタイプ値オブジェクト
│   │   │   └── framework_type.go          # フレームワークタイプ値オブジェクト
│   │   ├── service/
│   │   │   ├── project_generator.go       # プロジェクト生成サービス
│   │   │   └── template_processor.go      # テンプレート処理サービス
│   │   └── repository/
│   │       ├── project_repository.go      # プロジェクトリポジトリIF
│   │       ├── template_repository.go     # テンプレートリポジトリIF
│   │       └── config_repository.go       # 設定リポジトリIF
│   │
│   └── infrastructure/                     # Infrastructure Layer
│       ├── repository/
│       │   ├── filesystem_project_repository.go
│       │   ├── filesystem_template_repository.go
│       │   └── viper_config_repository.go
│       ├── git/
│       │   ├── git_service.go             # Git操作サービス
│       │   └── git_service_interface.go   # Git操作IF
│       ├── filesystem/
│       │   ├── fs_service.go              # ファイルシステムサービス
│       │   └── fs_service_interface.go    # ファイルシステムIF
│       └── logger/
│           └── logger.go                  # ロガー実装
│
├── pkg/                                    # 公開パッケージ
│   └── version/
│       └── version.go                     # バージョン情報
│
├── templates/                              # 組み込みテンプレート
│   ├── go-api/
│   │   ├── template.yaml
│   │   └── files/
│   │       ├── cmd/
│   │       ├── internal/
│   │       └── ...
│   ├── go-cli/
│   │   ├── template.yaml
│   │   └── files/
│   │       └── ...
│   └── go-microservice/
│       ├── template.yaml
│       └── files/
│           └── ...
│
├── test/                                   # テストファイル
│   ├── integration/
│   │   ├── create_project_test.go
│   │   └── template_management_test.go
│   └── fixtures/
│       └── test_templates/
│
├── scripts/                                # スクリプト
│   ├── build.sh                           # ビルドスクリプト
│   └── install.sh                         # インストールスクリプト
│
├── docs/                                   # ドキュメント
│   ├── requirements.md                    # 要件定義書
│   ├── architecture.md                    # アーキテクチャ設計書
│   └── detailed_design.md                 # 詳細設計書
│
├── .gitignore
├── go.mod
├── go.sum
├── Makefile
├── README.md
└── LICENSE
```

---

## コンポーネント設計

### 主要コンポーネント詳細

#### 1. CreateProjectUseCase
**責務**: プロジェクト生成のオーケストレーション

```go
type CreateProjectUseCase struct {
    projectRepo    repository.ProjectRepository
    templateRepo   repository.TemplateRepository
    configRepo     repository.ConfigRepository
    gitService     git.GitServiceInterface
    fsService      filesystem.FileSystemServiceInterface
    generator      *service.ProjectGenerator
    logger         logger.Logger
}

func (uc *CreateProjectUseCase) Execute(input CreateProjectInput) (*CreateProjectOutput, error) {
    // 1. 入力のバリデーション
    // 2. テンプレートの取得
    // 3. プロジェクトエンティティの生成
    // 4. ディレクトリ構造の作成
    // 5. テンプレートファイルの展開
    // 6. Git初期化
    // 7. 結果の返却
}
```

**依存関係**
- Domain: `repository.ProjectRepository`, `service.ProjectGenerator`
- Infrastructure: `git.GitServiceInterface`, `filesystem.FileSystemServiceInterface`

---

#### 2. ProjectGenerator (Domain Service)
**責務**: プロジェクト生成のビジネスロジック

```go
type ProjectGenerator struct {
    templateProcessor *TemplateProcessor
}

func (g *ProjectGenerator) Generate(project *entity.Project, template *entity.Template) (*GeneratedProject, error) {
    // 1. テンプレート変数の準備
    // 2. テンプレートファイルの処理
    // 3. ファイル内容の生成
    // 4. 生成結果の返却
}
```

**ビジネスルール**
- プロジェクト名のバリデーション
- モジュール名の形式チェック
- テンプレート変数の置換ルール

---

#### 3. TemplateProcessor (Domain Service)
**責務**: テンプレート処理のビジネスロジック

```go
type TemplateProcessor struct {
    // テンプレートエンジンのラッパー
}

func (p *TemplateProcessor) Process(templateContent string, variables map[string]interface{}) (string, error) {
    // 1. テンプレート解析
    // 2. 変数の置換
    // 3. 結果の返却
}
```

**処理ルール**
- `{{.ProjectName}}` 形式の変数置換
- カスタム関数のサポート（ToLower, ToUpper等）
- エラーハンドリング

---

#### 4. GitService
**責務**: Git操作の実装

```go
type GitService struct {
    logger logger.Logger
}

func (s *GitService) Initialize(path string) error {
    // git init 実行
}

func (s *GitService) CreateBranch(path, branchName string) error {
    // ブランチ作成
}

func (s *GitService) Commit(path, message string) error {
    // コミット作成
}

func (s *GitService) SetupHooks(path string) error {
    // Git Hooks設定
}
```

**使用ライブラリ**: `github.com/go-git/go-git/v5`

---

#### 5. FileSystemService
**責務**: ファイルシステム操作の実装

```go
type FileSystemService struct {
    fs     afero.Fs
    logger logger.Logger
}

func (s *FileSystemService) CreateDirectory(path string) error {
    // ディレクトリ作成
}

func (s *FileSystemService) WriteFile(path, content string) error {
    // ファイル書き込み
}

func (s *FileSystemService) ReadFile(path string) (string, error) {
    // ファイル読み込み
}

func (s *FileSystemService) CopyDirectory(src, dst string) error {
    // ディレクトリコピー
}
```

**使用ライブラリ**: `github.com/spf13/afero`

---

## データフロー

### プロジェクト生成フロー

```
User Input (CLI)
    │
    ▼
┌─────────────────────────────────────────┐
│ 1. Cobra Command (create.go)            │
│    - フラグ解析                          │
│    - 対話型プロンプト（必要に応じて）    │
└─────────────────────────────────────────┘
    │
    ▼
┌─────────────────────────────────────────┐
│ 2. CreateProjectUseCase                 │
│    - 入力バリデーション                  │
│    - オーケストレーション                │
└─────────────────────────────────────────┘
    │
    ├─────────────────────────────────────┐
    │                                     │
    ▼                                     ▼
┌─────────────────────┐      ┌─────────────────────┐
│ 3a. TemplateRepo    │      │ 3b. ConfigRepo      │
│ - テンプレート取得   │      │ - 設定取得          │
└─────────────────────┘      └─────────────────────┘
    │                                     │
    └─────────────────┬───────────────────┘
                      ▼
    ┌─────────────────────────────────────────┐
    │ 4. ProjectGenerator (Domain Service)    │
    │    - Project Entity作成                 │
    │    - ビジネスルール適用                  │
    └─────────────────────────────────────────┘
                      │
                      ▼
    ┌─────────────────────────────────────────┐
    │ 5. TemplateProcessor (Domain Service)   │
    │    - テンプレート変数置換                │
    │    - ファイル内容生成                    │
    └─────────────────────────────────────────┘
                      │
                      ▼
    ┌─────────────────────────────────────────┐
    │ 6. FileSystemService                    │
    │    - ディレクトリ作成                    │
    │    - ファイル書き込み                    │
    └─────────────────────────────────────────┘
                      │
                      ▼
    ┌─────────────────────────────────────────┐
    │ 7. GitService                           │
    │    - Git初期化                          │
    │    - ブランチ作成                        │
    │    - 初回コミット                        │
    └─────────────────────────────────────────┘
                      │
                      ▼
                  Success
                      │
                      ▼
    ┌─────────────────────────────────────────┐
    │ 8. Output Formatter                     │
    │    - 成功メッセージ表示                  │
    │    - 次のステップガイド                  │
    └─────────────────────────────────────────┘
```

---

## 依存関係

### 依存関係図

```
┌──────────────────┐
│   cmd/main.go    │
└────────┬─────────┘
         │ creates
         ▼
┌──────────────────────────────────────────┐
│          Dependency Container            │
│  - Repositories                          │
│  - Services                              │
│  - Use Cases                             │
│  - Commands                              │
└────────┬─────────────────────────────────┘
         │ provides
         ▼
┌──────────────────┐       ┌──────────────────┐
│  CLI Commands    │◄──────│   Use Cases      │
│  (Presentation)  │       │  (Application)   │
└──────────────────┘       └────────┬─────────┘
                                    │ uses
                                    ▼
                           ┌──────────────────┐
                           │     Domain       │
                           │  - Entities      │
                           │  - Services      │
                           │  - Repositories  │
                           │    (Interfaces)  │
                           └────────┬─────────┘
                                    │ implements
                                    ▼
                           ┌──────────────────┐
                           │  Infrastructure  │
                           │  - Git Service   │
                           │  - FS Service    │
                           │  - Repositories  │
                           └──────────────────┘
```

### 依存性注入（DI）

**DIコンテナの役割**
- 全ての依存関係を一箇所で管理
- インターフェースと実装のバインディング
- テスト時のモック差し替えを容易に

**DIコンテナ実装例**
```go
type Container struct {
    // Infrastructure
    fsService  filesystem.FileSystemServiceInterface
    gitService git.GitServiceInterface
    logger     logger.Logger
    
    // Repositories
    projectRepo  repository.ProjectRepository
    templateRepo repository.TemplateRepository
    configRepo   repository.ConfigRepository
    
    // Domain Services
    projectGenerator   *service.ProjectGenerator
    templateProcessor  *service.TemplateProcessor
    
    // Use Cases
    createProjectUC   *usecase.CreateProjectUseCase
    manageTemplateUC  *usecase.ManageTemplateUseCase
    manageConfigUC    *usecase.ManageConfigUseCase
}

func NewContainer() *Container {
    // 依存関係の構築
}
```

---

## 設計パターン

### 1. Repository Pattern
**目的**: データアクセスの抽象化

```go
// Domain Layer (インターフェース定義)
type ProjectRepository interface {
    Create(project *entity.Project) error
    Exists(path string) (bool, error)
}

// Infrastructure Layer (実装)
type FileSystemProjectRepository struct {
    fsService filesystem.FileSystemServiceInterface
}

func (r *FileSystemProjectRepository) Create(project *entity.Project) error {
    // ファイルシステムへの書き込み実装
}
```

**メリット**
- データソースの切り替えが容易（ファイル → DB）
- テストでモックに差し替え可能

---

### 2. Dependency Injection Pattern
**目的**: 依存関係の疎結合化

```go
// Use Caseにインターフェースを注入
type CreateProjectUseCase struct {
    projectRepo  repository.ProjectRepository  // インターフェース
    gitService   git.GitServiceInterface       // インターフェース
}

func NewCreateProjectUseCase(
    projectRepo repository.ProjectRepository,
    gitService git.GitServiceInterface,
) *CreateProjectUseCase {
    return &CreateProjectUseCase{
        projectRepo: projectRepo,
        gitService:  gitService,
    }
}
```

**メリット**
- テスタビリティ向上
- 実装の差し替えが容易

---

### 3. Factory Pattern
**目的**: オブジェクト生成の抽象化

```go
// Template Factory
type TemplateFactory struct {
    templateRepo repository.TemplateRepository
}

func (f *TemplateFactory) CreateFromType(projectType value.ProjectType) (*entity.Template, error) {
    switch projectType {
    case value.ProjectTypeAPI:
        return f.createAPITemplate()
    case value.ProjectTypeCLI:
        return f.createCLITemplate()
    case value.ProjectTypeMicroservice:
        return f.createMicroserviceTemplate()
    default:
        return nil, errors.New("unknown project type")
    }
}
```

---

### 4. Strategy Pattern
**目的**: アルゴリズムの切り替え

```go
// Template Processing Strategy
type TemplateStrategy interface {
    Process(content string, variables map[string]interface{}) (string, error)
}

type GoTemplateStrategy struct{}
type CustomTemplateStrategy struct{}

// Use Caseで戦略を選択
func (uc *CreateProjectUseCase) selectStrategy(template *entity.Template) TemplateStrategy {
    if template.IsCustom {
        return &CustomTemplateStrategy{}
    }
    return &GoTemplateStrategy{}
}
```

---

### 5. Builder Pattern
**目的**: 複雑なオブジェクトの段階的構築

```go
// Project Builder
type ProjectBuilder struct {
    project *entity.Project
}

func NewProjectBuilder() *ProjectBuilder {
    return &ProjectBuilder{
        project: &entity.Project{},
    }
}

func (b *ProjectBuilder) WithName(name string) *ProjectBuilder {
    b.project.Name = value.NewProjectName(name)
    return b
}

func (b *ProjectBuilder) WithType(projectType value.ProjectType) *ProjectBuilder {
    b.project.Type = projectType
    return b
}

func (b *ProjectBuilder) Build() (*entity.Project, error) {
    // バリデーション
    if err := b.project.Validate(); err != nil {
        return nil, err
    }
    return b.project, nil
}

// 使用例
project, err := NewProjectBuilder().
    WithName("my-project").
    WithType(value.ProjectTypeAPI).
    WithModule("github.com/user/my-project").
    Build()
```

---

## 拡張性設計

### 1. 新言語対応の拡張ポイント

#### Language Enum追加
```go
// domain/value/language.go
type Language string

const (
    LanguageGo     Language = "go"
    LanguageRust   Language = "rust"     // 追加
    LanguagePython Language = "python"   // 追加
)
```

#### Template追加
```
templates/
├── rust-api/           # Rust API テンプレート
│   ├── template.yaml
│   └── files/
└── python-api/         # Python API テンプレート
    ├── template.yaml
    └── files/
```

#### Generator拡張
```go
// domain/service/project_generator.go
func (g *ProjectGenerator) Generate(project *entity.Project, template *entity.Template) (*GeneratedProject, error) {
    switch project.Language {
    case value.LanguageGo:
        return g.generateGo(project, template)
    case value.LanguageRust:
        return g.generateRust(project, template)  // 追加
    case value.LanguagePython:
        return g.generatePython(project, template)  // 追加
    }
}
```

---

### 2. プラグインシステム設計（v2.0.0）

#### Plugin Interface
```go
// pkg/plugin/plugin.go
type Plugin interface {
    Name() string
    Version() string
    Execute(ctx context.Context, input PluginInput) (PluginOutput, error)
}

type PluginInput struct {
    Project  *entity.Project
    Template *entity.Template
}

type PluginOutput struct {
    Files map[string]string
}
```

#### Plugin Registry
```go
type PluginRegistry struct {
    plugins map[string]Plugin
}

func (r *PluginRegistry) Register(plugin Plugin) error {
    r.plugins[plugin.Name()] = plugin
    return nil
}

func (r *PluginRegistry) Get(name string) (Plugin, error) {
    plugin, exists := r.plugins[name]
    if !exists {
        return nil, errors.New("plugin not found")
    }
    return plugin, nil
}
```

---

### 3. テンプレートリポジトリ対応（v2.0.0）

#### Remote Template Support
```go
// domain/value/template_source.go
type TemplateSource string

const (
    TemplateSourceBuiltin TemplateSource = "builtin"
    TemplateSourceLocal   TemplateSource = "local"
    TemplateSourceGit     TemplateSource = "git"      // 追加
    TemplateSourceHTTP    TemplateSource = "http"     // 追加
)

// infrastructure/repository/remote_template_repository.go
type RemoteTemplateRepository struct {
    httpClient *http.Client
    gitService git.GitServiceInterface
}

func (r *RemoteTemplateRepository) Fetch(url string) (*entity.Template, error) {
    // GitまたはHTTPでテンプレート取得
}
```

---

## セキュリティ設計

### 1. パストラバーサル防止
```go
func (s *FileSystemService) validatePath(path string) error {
    // 絶対パスに変換
    absPath, err := filepath.Abs(path)
    if err != nil {
        return err
    }
    
    // ベースディレクトリ外へのアクセスをチェック
    if !strings.HasPrefix(absPath, s.baseDir) {
        return errors.New("path traversal detected")
    }
    
    return nil
}
```

### 2. テンプレート変数のサニタイゼーション
```go
func (p *TemplateProcessor) sanitizeVariables(variables map[string]interface{}) map[string]interface{} {
    sanitized := make(map[string]interface{})
    for key, value := range variables {
        // HTMLエスケープ
        // シェルコマンド防止
        sanitized[key] = sanitizeValue(value)
    }
    return sanitized
}
```

---

## パフォーマンス設計

### 1. 並行処理
```go
// 複数ファイルの並行生成
func (g *ProjectGenerator) generateFiles(files []FileInfo) error {
    var wg sync.WaitGroup
    errChan := make(chan error, len(files))
    
    for _, file := range files {
        wg.Add(1)
        go func(f FileInfo) {
            defer wg.Done()
            if err := g.generateFile(f); err != nil {
                errChan <- err
            }
        }(file)
    }
    
    wg.Wait()
    close(errChan)
    
    // エラーチェック
    for err := range errChan {
        if err != nil {
            return err
        }
    }
    
    return nil
}
```

### 2. キャッシング
```go
// テンプレートキャッシュ
type TemplateCache struct {
    cache map[string]*template.Template
    mu    sync.RWMutex
}

func (c *TemplateCache) Get(name string) (*template.Template, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    tmpl, exists := c.cache[name]
    return tmpl, exists
}
```

---

## テスト戦略

### 1. 単体テスト
```go
// Use Caseのテスト
func TestCreateProjectUseCase_Execute(t *testing.T) {
    // モックの作成
    mockProjectRepo := &mocks.ProjectRepository{}
    mockGitService := &mocks.GitService{}
    
    // Use Case作成
    uc := usecase.NewCreateProjectUseCase(
        mockProjectRepo,
        mockGitService,
    )
    
    // テスト実行
    output, err := uc.Execute(input)
    
    // アサーション
    assert.NoError(t, err)
    assert.NotNil(t, output)
}
```

### 2. 統合テスト
```go
// エンドツーエンドテスト
func TestCreateProject_Integration(t *testing.T) {
    // 実際のファイルシステムを使用（テスト用ディレクトリ）
    tmpDir := t.TempDir()
    
    // 実際のコンポーネントを組み立て
    container := setupIntegrationContainer(tmpDir)
    
    // プロジェクト生成
    err := container.CreateProjectUseCase.Execute(input)
    
    // ファイルシステムを検証
    assert.NoError(t, err)
    assert.DirExists(t, filepath.Join(tmpDir, "my-project"))
}
```

---

## デプロイメント設計

### ビルド戦略
```makefile
# Makefile
.PHONY: build
build:
	go build -ldflags="-s -w" -o bin/go-review-cli cmd/go-review-cli/main.go

.PHONY: build-all
build-all:
	GOOS=linux GOARCH=amd64 go build -o bin/go-review-cli-linux-amd64
	GOOS=darwin GOARCH=amd64 go build -o bin/go-review-cli-darwin-amd64
	GOOS=windows GOARCH=amd64 go build -o bin/go-review-cli-windows-amd64.exe
```

### インストール方法
1. **バイナリ直接実行**: ダウンロードして実行
2. **go install**: `go install github.com/user/go-review-cli@latest`
3. **パッケージマネージャー**: Homebrew, apt等（将来対応）

---

## モニタリング・ロギング

### ロギング設計
```go
// Logger Interface
type Logger interface {
    Debug(msg string, fields ...Field)
    Info(msg string, fields ...Field)
    Warn(msg string, fields ...Field)
    Error(msg string, fields ...Field)
}

// 使用例
logger.Info("Project created",
    Field{"project_name", project.Name},
    Field{"project_type", project.Type},
)
```

---

## まとめ

### アーキテクチャの特徴
1. **Clean Architecture**: レイヤー分離による高い保守性
2. **Dependency Injection**: テスタビリティの向上
3. **インターフェース駆動**: 拡張性と柔軟性
4. **Go言語のベストプラクティス**: シンプルで効率的な設計

### 将来の拡張性
- 多言語対応
- プラグインシステム
- リモートテンプレート
- マイクロサービス統合

---

## 変更履歴

| バージョン | 日付 | 変更内容 |
|-----------|------|---------|
| 1.0.0 | 2026-01-27 | 初版作成 |