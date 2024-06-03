# ミーティング日時調整webサービスケーション Let’s Schedule 要件定義書


# アジェンダ


## 目的

本ドキュメントでは筆者が開発するミーティング日程調整サービスに必要な要件を定義する。

## サービス概要

 本サービスはミーティングの日程調整時に、各参加者の予定を元に全員が参加可能な日時を提案するwebサービスである。

具体的にはミーティングの候補日時の1つずつにリアクション（例: ◯/△/×）するのではなく、主催者・参加者共に範囲で入力できる

確定したミーティングの日時はicsファイル、Googleカレンダー、Outlookカレンダーなど、様々なフォーマットで外部に提供できる。

## 既存サービスに対する優位性

これらの点について以下の既存サービスと比較することで、本サービスの優位性を明確にする。

| | 主催者日時入力 | 参加者通知 | 参加者日時入力 | 日時決定 | カレンダーへの登録 | サービスへの使用条件 |
| ---- | ---- | ---- | ---- | ---- | ---- |---- |
| Let's Schedule | 候補入力 | サービス内で完結 | 参加できる日時を範囲入力 | 主催者が決定 | icsファイル Google Calendar・Outlook Calendarのリンク発行 | メアド制限なし |
| 調整さん | 候補日時1つずつ入力| 公開URL | 候補日時ごとに◯/X/△ | 全員が参加できる日時を色付け | なし | メアド
制限なし |
| Google Calendar | 候補日時範囲入力/自動抽出 | メール | 候補日時ごとにYes/No/Maybe | 主催者が決定 | Google
Calendarに登録| メアド制限なし (ただしGmail以外は別途連携必要) |
| Microsoft Outlook | 候補日時範囲入力/自動抽出 | メール | 候補日時ごとにYes/No | 主催者が決定 | Outlook
Calendarに登録 | Outlook Emailユーザーのみ |


## 技術要件

本webサービスに導入する技術は以下の通りである。

### 前提条件 

webブラウザからアクセスできる
Htmlファイルが使えるブラウザを使用する
端末のディスプレイに応じた画面サイズに応じた表示をする

### 制約条件 

webサービスの管理者は開発者(中井)のみ
稼働率は99.5%である (24時間365日の稼働時間は保証しない)
同時接続数は1100である

### サービスケーションのアーキテクチャ

選定技術: SPA (single page architecture)
選定理由: 
必要最低限のデータを読み込むのでページの遷移が高速になる
フロントエンド側・バックエンド側の実装が分離できるので、それぞれの実装に専念できる
 
選定技術: MVCアーキテクチャ
選定理由: 
Model (DBとの通信やデータの制約の定義)
View(ページを表示するUI)
Controller (REST APIのエンドポイントでの処理)の実装が分離できる

フロントエンド 
選定技術: Vue.js
選定理由: 
HTMLとJavaScript部分を分離できるので可読性が高い 
 (参考) ReactのJSXはJavaScriptの中にhtmlを書くので読みづらくなる恐れがある
    Component-basedな実装方法の新規学習コストが高い

バックエンド
選定技術: Golang (Echo)
Golangの選定理由:
記法がシンプルで可読性が高い
シングルバイナリ (ビルド・デプロイが短時間でできる)
処理が早い 
Echoの選定理由: 
小規模フレームワーク


### インフラのアーキテクチャ

選定技術: 3層アーキテクチャ
選定理由: 
MVCアーキテクチャを物理的に反映できる

Webサーバー: Viewをデプロイサーバー
APサーバー:  Model・Controllerを提供するサーバー
DBサーバー: Modelで定義されたデータを保存するサーバー

インフラ
選定技術: AWS (EC2)
選定理由: 
業務で使用経験があるので馴染みがある

DB 
選定技術: MySQL
選定理由: 
PostgreSQLが得意な複雑なクエリの実装・大規模データ処理が必要ない

## 機能要件

サービスに要求される機能

### セキュリティ

ユーザーは前提としてステータスなし
練習としてCookie実装 

認証にはJWTを使用する


### 証明書

Webサーバーの証明書はLet's Encryptで取得した証明書を、APIサーバーには自己署名した証明書を使用する。


### ユーザー共通

- ユーザー新規作成
- ログイン
- ログアウト
- ミーティング情報の取得・外部登録 (ics, Google Calendar, Outlook Calendar)

### ミーティング主催者

ミーティング新規作成
ミーティング情報編集
参加できる時間を入力
参加できる時間を送信
最終的な日時を選択
選択した日時を参加者に通知

### ミーティング参加者  (主催者以外)

参加できる時間を入力
参加できる時間を送信


## DB設計

### ユーザーテーブル (usersテーブル)

| 項目名 (日本語) | 項目名 (変数) | 型 | 備考 | 
| ---- | ---- | ---- | ---- |
| ID | id | int | NOT NULL AUTO INCREMENT |
| ユーザー名 | user_name | string | NOT NULL |
| メールアドレス | email_address | string | NOT NULL |
| パスワード | password | string | NOT NULL |
| 管理者かどうか | is_admin | bool | NOT NULL |
| ログイン可能か | can_login | bool | NOT NULL |
| 登録日時 | created_at | datetime | NOT NULL DEFAULT CURRENT_DATETIME |
| 更新日時 | updated_at | datetime | NOT NULL DEFAULT CURRENT_DATETIME ON UPDATE CURRENT_DATETIME |


### ミーティングテーブル (meetingsテーブル)

| 項目名 (日本語) | 項目名 (変数) | 型 | 備考 | 
| ---- | ---- | ---- | ---- |
| ID | id | int | NOT NULL  AUTO INCREMENT |
| ミーティング名 | title | string | NOT NULL |
| 概要 | description | text | NOT NULL |
| 形式 | type | string | 現地 or オンライン |
| 集合場所 | place | string | |
| ミーティングURL | url | string | |
| 全員が回答したか | all_participants_responded | bool | NOT NULL |
| 確定したか | is_confirmed | bool | NOT NULL |
| 開始日時 | start_time | datetime | NOT NULL |
| 終了日時 | end_time | datetime | NOT NULL |
| 登録日時 | created_at | datetime | NOT NULL DEFAULT CURRENT_DATETIME |
| 更新日時 | updated_at | datetime | NOT NULL DEFAULT CURRENT_DATETIME ON UPDATE CURRENT_DATETIME |


### 参加者テーブル (participantsテーブル)

| 項目名 (日本語) | 項目名 (変数) | 型 | 備考 | 
| ---- | ---- | ---- | ---- |
| ID | id  | int | NOT NULL  AUTO INCREMENT |
| ユーザーID | user_id | int | NOT NULL |
| ミーティングID | meeting_id | int | NOT NULL |
| 主催者かどうか | is_host | bool | NOT NULL |
| 回答の有無 | has_responded | bool | NOT NULL |
| 登録日時 | created_at | datetime | NOT NULL DEFAULT CURRENT_DATETIME |
| 更新日時 | updated_at | datetime | NOT NULL DEFAULT CURRENT_DATETIME ON UPDATE CURRENT_DATETIME |



### 候補日時テーブル (candidate_timeテーブル)

| 項目名 (日本語) | 項目名 (変数) | 型 | 備考 | 
| ---- | ---- | ---- | ---- |
| ID | id  | int | NOT NULL  AUTO INCREMENT |
| ユーザーID | user_id | int | NOT NULL |
| ミーティングID | meeting_id | int | NOT NULL |
| 開始時間 | start_time | datetime | NOT NULL |
| 終了時間 | end_time | datetime | NOT NULL |
| 登録日時 | created_at | datetime | NOT NULL DEFAULT CURRENT_DATETIME |
| 更新日時 | updated_at | datetime | NOT NULL DEFAULT CURRENT_DATETIME ON UPDATE CURRENT_DATETIME |


## API仕様

### 共通事項

各APIは次のフォーマットのURLとする

https://lets-schedule.net/<エンドポイントのパス>

通信プロトコル: HTTPS
APIの種類: REST API
インターフェース: JSON
文字コード: UTF-8

ユーザー関連API 


#### 新規ユーザー作成

エンドポイント POST /YXBpL3NpZ251cA==

リクエストパラメータ

| キー名 | 型 (変数) | 概要 | 備考 |
| -- | -- | -- | -- | 
| id | int | ユーザーID | -- | 
| username | string | ユーザー名 | |
| email_address | string | メールアドレス | |
| password | string | パスワード | |
| is_admin | bool | 管理者かどうか | デフォルトはfalse |
| can_login | bool | ログイン可能かどうか | デフォルトはtrue |


ステータスコード

| コード | 意味 |
| -- | -- | 
| 200 | ユーザー登録成功 |
| 400 | ユーザー登録失敗 |

レスポンスパラメータ

| シチュエーション | キー名 | 型 | 概要 |
| -- | -- | -- | -- | 
| 成功時 | id | int | ユーザーID |
| | username | string | ユーザー名 |
| | email_address | string | メールアドレス |
| | password | string | パスワード |
| | is_admin | bool | 管理者かどうか |
| | can_login | bool | ログイン可能かどうか | 
| 失敗時 | errorMessageList | array\[string\] | エラーメッセージの配列 |



### ミーティング関連API エンドポイント 

#### ミーティング新規作成

エンドポイント POST /YXBpL3Jlc3RyaWN0ZWQvbWVldGluZ3MvbmV3

パラメータ

| キー名 | 型 (変数) | 概要 | 
| -- | -- | -- |
| title | string | ミーティング名 |
| description | text | 概要 |
| type | string | 形式 |
| place | string | 集合場所 |
| url | string | ミーティングURL |
| is_confirmed | bool | 日時が決定したかどうか |


ステータスコード

| コード | 意味 |
| -- | -- |
| 200 | ミーティング登録成功 |
| 400 | ミーティング登録失敗 |


レスポンスパラメータ

| シチュエーション | キー名 | 型 | 概要 |
| -- | -- | -- | -- | 
| 成功時 | id | int | ユーザーID |
| | title | string | ミーティング名 |
| |  description | text | 概要 |
| |  type | string | 形式 |
| |  place | string | 集合場所 |
| |  url | string | ミーティングURL |
| |  is_confirmed | bool | 日時が決定したかどうか |
| 失敗時 | errorMessageList | array\[string\] | エラーメッセージの配列 |


#### ミーティング情報取得

エンドポイント GET /YXBpL3Jlc3RyaWN0ZWQvbWVldGluZ3MvdXNlcg==/:user_id

パラメータ

| キー名 | 型 (変数) | 概要 | 
| -- | -- | -- |
| user_id | string | ユーザーID |


ステータスコード

| コード | 意味 |
| -- | -- |
| 200 | ミーティング情報取得成功 |
| 400 | ミーティング情報取得失敗 |


レスポンスパラメータ

| シチュエーション | キー名 | 型 (変数) | 概要 | 
| -- | -- | -- | -- |
| 成功時 | id| string | ミーティングID |
| | title | string | ミーティング名 |
| | description | text | 概要 |
| | type | string | 形式 |
| | meeting_place | string | 集合場所 |
| | meeting_url | string | ミーティングURL |
| | all_participants_responded | bool | 全員から返信があったか |
| | is_confirmed | bool | 日時が決まったか |
| | start_time | datetime | ミーティングの開始日時 |
| | end_time | datetime | ミーティングの終了日時 |
| | hour | float | ミーティングの時間 (h) |
| 失敗時 | error | string | エラー文 |

#### 日時が決定した主催ミーティング情報取得

エンドポイント GET /YXBpL3Jlc3RyaWN0ZWQvbWVldGluZ3MvaG9zdC9jb25maXJtZWQ=/:user_id

パラメータ

| キー名 | 型 (変数) | 概要 | 
| -- | -- | -- |
| user_id | string | ユーザーID |


ステータスコード

| コード | 意味 |
| -- | -- |
| 200 | ミーティング情報取得成功 |
| 400 | ミーティング情報取得失敗 |


レスポンスパラメータ

| シチュエーション | キー名 | 型 (変数) | 概要 | 
| -- | -- | -- | -- |
| 成功時 | meetings | Meeting[] | ミーティング情報の配列 |
| 失敗時 | error | string | エラー文 |


#### 日時が決定していない主催ミーティング情報取得

エンドポイント GET /aYXBpL3Jlc3RyaWN0ZWQvbWVldGluZ3MvaG9zdC9ub3QtY29uZmlybWVk/:user_id

パラメータ

| キー名 | 型 (変数) | 概要 | 
| -- | -- | -- |
| user_id | string | ユーザーID |


ステータスコード

| コード | 意味 |
| -- | -- |
| 200 | ミーティング情報取得成功 |
| 400 | ミーティング情報取得失敗 |


レスポンスパラメータ

| シチュエーション | キー名 | 型 (変数) | 概要 | 
| -- | -- | -- | -- |
| 成功時 | meetings | Meeting[] | ミーティング情報の配列 |
| 失敗時 | error | string | エラー文 |


#### 返信していない主催ミーティング情報取得

エンドポイント GET /YXBpL3Jlc3RyaWN0ZWQvbWVldGluZ3MvaG9zdC9ub3QtcmVzcG9uZGVk/:user_id

パラメータ

| キー名 | 型 (変数) | 概要 | 
| -- | -- | -- |
| user_id | string | ユーザーID |


ステータスコード

| コード | 意味 |
| -- | -- |
| 200 | ミーティング情報取得成功 |
| 400 | ミーティング情報取得失敗 |


レスポンスパラメータ

| シチュエーション | キー名 | 型 (変数) | 概要 | 
| -- | -- | -- | -- |
| 成功時 | meetings | Meeting[] | ミーティング情報の配列 |
| 失敗時 | error | string | エラー文 |


#### 日時が決定している参加ミーティング情報取得

エンドポイント GET /YXBpL3Jlc3RyaWN0ZWQvbWVldGluZ3MvZ3Vlc3QvY29uZmlybWVk/:user_id

パラメータ

| キー名 | 型 (変数) | 概要 | 
| -- | -- | -- |
| user_id | string | ユーザーID |


ステータスコード

| コード | 意味 |
| -- | -- |
| 200 | ミーティング情報取得成功 |
| 400 | ミーティング情報取得失敗 |


レスポンスパラメータ

| シチュエーション | キー名 | 型 (変数) | 概要 | 
| -- | -- | -- | -- |
| 成功時 | meetings | Meeting[] | ミーティング情報の配列 |
| 失敗時 | error | string | エラー文 |


#### 日時が決定していない参加ミーティング情報取得

エンドポイント GET /YXBpL3Jlc3RyaWN0ZWQvbWVldGluZ3MvZ3Vlc3Qvbm90LWNvbmZpcm1lZA==/:user_id

パラメータ

| キー名 | 型 (変数) | 概要 | 
| -- | -- | -- |
| user_id | string | ユーザーID |


ステータスコード

| コード | 意味 |
| -- | -- |
| 200 | ミーティング情報取得成功 |
| 400 | ミーティング情報取得失敗 |


レスポンスパラメータ

| シチュエーション | キー名 | 型 (変数) | 概要 | 
| -- | -- | -- | -- |
| 成功時 | meetings | Meeting[] | ミーティング情報の配列 |
| 失敗時 | error | string | エラー文 |


#### 返信していない参加ミーティング情報取得

エンドポイント GET /YXBpL3Jlc3RyaWN0ZWQvbWVldGluZ3MvZ3Vlc3Qvbm90LXJlc3BvbmRlZA===/:user_id

パラメータ

| キー名 | 型 (変数) | 概要 | 
| -- | -- | -- |
| user_id | string | ユーザーID |


ステータスコード

| コード | 意味 |
| -- | -- |
| 200 | ミーティング情報取得成功 |
| 400 | ミーティング情報取得失敗 |


レスポンスパラメータ

| シチュエーション | キー名 | 型 (変数) | 概要 | 
| -- | -- | -- | -- |
| 成功時 | meetings | Meeting[] | ミーティング情報の配列 |
| 失敗時 | error | string | エラー文 |



### 候補日時関連API 

#### 候補日時の新規登録

エンドポイントPOST /YXBpL3Jlc3RyaWN0ZWQvY2FuZGlkYXRlX3RpbWVzL25ldw==

リクエストパラメータ

| キー名 | 型 (変数) | 概要 |
| -- | -- | -- |
| candidate_time_list | CandidateTime[] | 候補日時の配列 |


レスポンスパラメータ

| シチュエーション | キー名 | 型 (変数) | 概要 | 
| -- | -- | -- | -- |
| 成功時 | candidate_time | CandidateTime[] | 候補日時の配列 |
| 失敗時 | error | string | エラー文 |


#### 候補日時の取得

エンドポイント GET /YXBpL3Jlc3RyaWN0ZWQvY2FuZGlkYXRlX3RpbWVzL3VzZXI=/:user_id/bWVldGluZw==/:meeting_id

リクエストパラメータ

| キー名 | 型 (変数) | 概要 |
| -- | -- | -- |
| user_id | int | ユーザーID |
| meeting_id | int | ミーティングID |


レスポンスパラメータ

| シチュエーション | キー名 | 型 (変数) | 概要 | 
| -- | -- | -- | -- |
| 成功時 | candidate_time | CandidateTime | 候補日時の配列 |
| 失敗時 | error | string | エラー文 |




#### 候補日時の編集

エンドポイント PUT /YXBpL3Jlc3RyaWN0ZWQvY2FuZGlkYXRlX3RpbWVzL3VzZXI=/:user_id/bWVldGluZw==/:meeting_id

リクエストパラメータ

| キー名 | 型 (変数) | 概要 |
| -- | -- | -- |
| user_id | int | ユーザーID |
| meeting_id | int | ミーティングID |


レスポンスパラメータ

| シチュエーション | キー名 | 型 (変数) | 概要 | 
| -- | -- | -- | -- |
| 成功時 | candidate_time | CandidateTime[] | 候補日時の配列 |
| 失敗時 | error | string | エラー文 |



### 参加者関連API 

#### 参加者の新規登録

エンドポイントPOST /YXBpL3Jlc3RyaWN0ZWQvcGFydGljaXBhbnRzL25ldw==

リクエストパラメータ

| キー名 | 型 (変数) | 概要 |
| -- | -- | -- |
| participants | Participant[] | 候補日時の配列 |


レスポンスパラメータ

| シチュエーション | キー名 | 型 (変数) | 概要 | 
| -- | -- | -- | -- |
| 成功時 | participants | Participant[] | 候補日時の配列 |
| 失敗時 | error | string | エラー文 |


#### 参加者の取得

エンドポイント GET /YXBpL3Jlc3RyaWN0ZWQvcGFydGljaXBhbnRz/:meeting_id

リクエストパラメータ

| キー名 | 型 (変数) | 概要 |
| -- | -- | -- |
| meeting_id | int | ミーティングID |


レスポンスパラメータ

| シチュエーション | キー名 | 型 (変数) | 概要 | 
| -- | -- | -- | -- |
| 成功時 | participants | Participant[] | 候補日時の配列 |
| 失敗時 | error | string | エラー文 |




#### 参加者の編集

エンドポイント PUT /YXBpL3Jlc3RyaWN0ZWQvcGFydGljaXBhbnRzL21lZXRpbmc=/:meeting_id

リクエストパラメータ

| キー名 | 型 (変数) | 概要 |
| -- | -- | -- |
| meeting_id | int | ミーティングID |


レスポンスパラメータ

| シチュエーション | キー名 | 型 (変数) | 概要 | 
| -- | -- | -- | -- |
| 成功時 | participants | Participant[] | 参加者の配列 |
| 失敗時 | error | string | エラー文 |




## 画面

### フロントページ

#### ログイン画面
- ユーザー名
- パスワード
- パスワードを忘れた場合

#### ユーザー登録画面
- メアド
- ユーザー名
- パスワード

#### パスワード再設定画面
- メアド
- パスワード
- パスワード (確認用)

#### ダッシュボード
- 主催ミーティング一覧
- 参加ミーティング一覧
- ミーティング作成ボタン

#### ミーティング新規作成画面
- ミーティング名
- 概要
- 予定日時
- 形式
- 参加者
- 作成ボタン
- 止めるボタン


#### 主催しているミーティング詳細画面
- ミーティング名
- 概要
- 予定日時
- 形式
- 参加者
- 誰が回答してるか / してないか
- 最適な時間帯
- 最終決定ボタン
- 編集ボタン
- 削除ボタン


#### 参加しているミーティング詳細画面
- ミーティング名
- 概要
- 予定日時
- 形式
- 参加者
- 回答欄
- 編集ボタン
- 削除ボタン
- 設定画面
- メアド変更
- ユーザー名変更


## 非機能要件

### 同時接続数

研究室のメンバー (教員・学生含め) 110人 
合唱団のメンバー 38人
日常的に連絡を取っている友人 10人 
友人が参加する集まりの参加者数 10人程度

合計 (110 + 38 + 10 * 10 =) 248 接続数

レスポンスタイム (クライアントからの通信 ~ クライアントへの受信) 
100ms


### インフラ構成

<img width="670" alt="Screenshot 2024-06-01 at 21 51 32" src="https://github.com/makotonakai/lets-schedule/assets/45162150/26310471-c6a6-4663-820f-77c3e2ca5e50">


可用性 
約 99.6%
単一インスタンスの可用性最低99.5%
(https://aws.amazon.com/compute/sla/?nc1=h_ls)

 1 - (1 - 0.995) ** 2 ≒ 0.9996059576745632…


