# Calculator API

シンプルな計算機能を提供するGolang製のRESTful APIサンプルプロジェクト

## 特徴

- 基本的な四則演算（加算、減算、乗算、除算）
- 統計計算（平均、中央値、標準偏差）
- クリーンアーキテクチャに基づいたディレクトリ構成
- JSONレスポンス
- エラーハンドリング

## ディレクトリ構成

```
.
├── main.go                 # エントリーポイント
├── go.mod                  # Go modules定義
├── README.md              
├── cmd/
│   └── server/
│       └── main.go        # サーバー起動処理
├── internal/
│   ├── handler/           # HTTPハンドラー
│   │   └── calculator.go
│   ├── service/           # ビジネスロジック
│   │   └── calculator.go
│   └── model/             # データモデル
│       └── request.go
└── pkg/
    └── calculator/        # 計算ロジック（再利用可能）
        ├── basic.go
        └── statistics.go
```

## セットアップ

### 必要要件

- Go 1.21以上

### インストール

```bash
# リポジトリのクローン
git clone <your-repo-url>
cd calculator-api

# 依存関係のインストール
go mod download

# ビルド
go build -o calculator-api ./cmd/server
```

## 実行

```bash
# 開発モードで実行
go run cmd/server/main.go

# またはビルド済みバイナリを実行
./calculator-api
```

サーバーは `http://localhost:8080` で起動します。

## API エンドポイント

### 1. 基本計算

#### 加算
```bash
POST /api/v1/calculate/add
Content-Type: application/json

{
  "numbers": [10, 20, 30]
}

# レスポンス
{
  "result": 60,
  "operation": "addition"
}
```

#### 減算
```bash
POST /api/v1/calculate/subtract
Content-Type: application/json

{
  "numbers": [100, 30, 20]
}

# レスポンス
{
  "result": 50,
  "operation": "subtraction"
}
```

#### 乗算
```bash
POST /api/v1/calculate/multiply
Content-Type: application/json

{
  "numbers": [5, 4, 3]
}

# レスポンス
{
  "result": 60,
  "operation": "multiplication"
}
```

#### 除算
```bash
POST /api/v1/calculate/divide
Content-Type: application/json

{
  "numbers": [100, 2, 5]
}

# レスポンス
{
  "result": 10,
  "operation": "division"
}
```

### 2. 統計計算

#### 平均
```bash
POST /api/v1/statistics/average
Content-Type: application/json

{
  "numbers": [10, 20, 30, 40, 50]
}

# レスポンス
{
  "result": 30,
  "operation": "average",
  "count": 5
}
```

#### 中央値
```bash
POST /api/v1/statistics/median
Content-Type: application/json

{
  "numbers": [1, 3, 5, 7, 9]
}

# レスポンス
{
  "result": 5,
  "operation": "median",
  "count": 5
}
```

#### 標準偏差
```bash
POST /api/v1/statistics/stddev
Content-Type: application/json

{
  "numbers": [2, 4, 6, 8, 10]
}

# レスポンス
{
  "result": 2.8284271247461903,
  "operation": "standard_deviation",
  "count": 5
}
```

### 3. ヘルスチェック

```bash
GET /health

# レスポンス
{
  "status": "ok"
}
```

## テスト

```bash
# 全テストの実行
go test ./...

# カバレッジ付き
go test -cover ./...

# 詳細出力
go test -v ./...
```

## 開発

### 新しい計算機能の追加

1. `pkg/calculator/` に新しい計算ロジックを追加
2. `internal/service/calculator.go` にサービスメソッドを追加
3. `internal/handler/calculator.go` にハンドラーを追加
4. `cmd/server/main.go` でルーティングを設定

## ライセンス

MIT License
