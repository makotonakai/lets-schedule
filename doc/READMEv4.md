# 日程調整webアプリケーションLet's schedule要件定義書

1. [概要](#概要)
   1. [前提条件](#前提条件)
   2. [制約条件](#制約条件)
   3. [機能要件](#機能要件)
   4. [基本機能](#基本機能)
      1. [ユーザー管理](#ユーザー管理)
      2. [日程調整](#日程調整)
      3. [カレンダー連携](#カレンダー連携)
   5. [オブジェクト定義](#オブジェクト定義)
      1. [Userオブジェクト](#userオブジェクト)
      2. [Meetingオブジェクト](#meetingオブジェクト)
      3. [CandidateTimeオブジェクト](#candidatetimeオブジェクト)
      4. [Participantオブジェクト](#participantオブジェクト)
   6. [フロントエンド](#フロントエンド)
      1. [UI](#ui)
      2. [ページ遷移図](#ページ遷移図)
   7. [バックエンド](#バックエンド)
      1. [API定義](#api定義)
   8. [DB](#db)
      1. [前提](#前提)
      2. [Usersテーブル](#usersテーブル)
      3. [Meetingsテーブル](#meetingsテーブル)
      4. [Participantテーブル](#participantテーブル)
      5. [CandidateTimeテーブル](#candidatetimeテーブル)
      6. [キー一覧](#キー一覧)
2. [非機能要件](#非機能要件)
   1. [可用性](#可用性)
      1. [SLA](#sla)
   2. [パフォーマンス](#パフォーマンス)
      1. [同時接続数](#同時接続数)
   3. [セキュリティ](#セキュリティ)
      1. [証明書](#証明書)
      2. [アクセス制御](#アクセス制御)
   4. [オブザーバビリティ](#オブザーバビリティ)
      1. [ホストVM](#ホストvm)
      2. [Webサーバー](#webサーバー)
      3. [APIサーバー](#apiサーバー)
      4. [DBサーバー](#dbサーバー)
   5. [保守性](#保守性)
      1. [インフラ構成図](#インフラ構成図)
      2. [ネットワーク構成](#ネットワーク構成)
      3. [セキュリティグループ](#セキュリティグループ)
3. [使用技術](#使用技術)
   1. [フロントエンド](#フロントエンド-1)
   2. [バックエンド](#バックエンド-1)
   3. [DB](#db-1)
   4. [インフラ](#インフラ)
4. [参考文献](#参考文献)


## 概要

日程調整はビジネスにおいて重要なプロセスであり、効率的に行うためには適切なツールが必要である。本プロジェクトでは、ユーザーが簡単かつ迅速にミーティングのスケジュールを調整できるアプリケーションを開発することを目的としている。

### 前提条件

- インターネット接続が必須
- ユーザーはメールアドレスを持っていること
- 対象デバイスはPC、タブレット、スマートフォンとする

### 制約条件

- プライバシー保護のため、データは暗号化して保存する
- サービスの稼働率は99.9%以上を目指す

### 機能要件

### 基本機能

#### ユーザー管理

- ユーザー登録
- ログイン / ログアウト

#### 日程調整

- 新規ミーティング作成
- 候補日時の編集

#### カレンダー連携

- Googleカレンダー連携
- Outlookカレンダー連携
- カレンダーのエクスポート

### オブジェクト定義

#### Userオブジェクト

##### プロパティ

| 変数名 | 型 | 説明 | 
| ---- | ---- | ---- |
| id | int | ユーザーID |
| user_name | string |  ユーザー名 | 
| email_address | string | メールアドレス | 
| password | string | パスワード | 
| is_admin | bool | 管理者かどうか | 
| can_login | bool | ログイン可能か | 
| created_at | datetime | 登録日時 | 
| updated_at | datetime | 更新日時 | 


##### IsEmailAddressValidメソッド

与えられたメールアドレスが有効か判定する関数

メールアドレスの正規表現は[\b[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}\b](https://www.regular-expressions.info/index.html)を採用した。

引数

| 変数名 | 型 | 
| ---- | ---- | 
| e | string |

返り値

| 型 | 説明 |  
| ---- | ---- | 
| bool | 与えられたメールアドレスが有効かどうか |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | true / false  | nil |
| 異常レスポンス (与えられたメールアドレスが空の場合) | false | errors.New("The given email address is empty") |


##### IsEmailAddressEmptyOrNullメソッド

与えられたメールアドレスが空白か判定する関数 

引数

| 変数名 | 型 | 
| ---- | ---- | 
| u | User |

返り値

| 型 | 説明 |  
| ---- | ---- | 
| bool | 与えられたメールアドレスが有効かどうか |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | true / false  | nil |
| 異常レスポンス (uのポイントがnilの場合) | false | errors.New("The given User object is nil") |


##### IsUserNameEmptyOrNullメソッド

与えられたユーザー名が空白か判定する関数

返り値

| 型 | 説明 |  
| ---- | ---- | 
| bool | 与えられたメールアドレスが有効かどうか |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | true / false  | nil |
| 異常レスポンス (uのポイントがnilの場合) | false | errors.New("The given User object is nil") |

##### IsPasswordEmptyOrNullメソッド

与えられたパスワードが空白か判定する関数

返り値

| 型 | 説明 |  
| ---- | ---- | 
| bool | 与えられたメールアドレスが有効かどうか |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | true / false  | nil |
| 異常レスポンス (uのポイントがnilの場合) | false | errors.New("The given User object is nil") |

##### ErrorsExistメソッド

与えられたエラーが存在するか判定する関数

引数

| 変数名 | 型 | 
| ---- | ---- | 
| errorMessageList | []string |

返り値

| 型 | 説明 |  
| ---- | ---- | 
| bool | 与えられたメールアドレスが有効かどうか |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | true / false  | nil |
| 異常レスポンス (errorMessageListのポイントがnilの場合) | false | errors.New("The list of error messages doesn't exist") |

##### AlreadyExistsメソッド

与えられたユーザーが既に登録されているか登録する関数

引数

| 変数名 | 型 | 
| ---- | ---- | 
| db | ([*gorm.DB](https://pkg.go.dev/gorm.io/gorm#DB)型の構造体ポインタ) |
| u | User |

返り値

| 型 | 説明 |  
| ---- | ---- | 
| bool | 与えられたメールアドレスが有効かどうか |
| error | 発生したエラー |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| ユーザー名とメールアドレスが存在する時 | true | nil | nil |
| メールアドレスが見つからない時 | true | errors.New("Email address not found") | nil |
| ユーザー名が見つからない時 | true | nil | errors.New("Username not found") |
| ユーザー名もメールアドレスが見つからない時 | false | errors.New("Email address not found") | errors.New("Username not found") |


#### Meetingオブジェクト

##### プロパティ

| 変数名 | 型 | 説明 | 
| ---- | ---- | ---- |
| id | int | ミーティングID | 
| title | string | ミーティング名 | 
| description | text | 概要 | 
| is_onsite | bool | オンサイト開催か (falseの場合はオンライン開催) | 
| place | string | 集合場所 | 
| url | string | ミーティングURL | 
| all_participants_responded | bool | 全員が回答したか | 
| is_confirmed | bool | 確定したか | 
| start_time | datetime | 開始日時 | 
| end_time | datetime | 終了日時 | 
| created_at | datetime | 登録日時 | 
| updated_at | datetime | 更新日時 | 

##### メソッド

| メソッド名 | 説明 | 
| ---- | ---- | 
| IsTitleEmpty | 与えられたミーティングのタイトルが空白かどうか判断する関数 |
| IsHourEmpty | 与えられたミーティングの時間が空白かどうか判断する関数 |
| IsOnsiteButNoPlaceSpecified | 与えられた物理開催のミーティングで開催場所が指定されているか判定する関数 |
| IsOnlineButNoURLSpecified | 与えられたオンライン開催のミーティングで開催URLが指定されているか判定する関数 |
| IsHybridButNeitherPlaceOrURLSpecified | 与えられたハイブリッド開催のミーティングで開催場所と開催URLが両方指定されているか判定する関数 |
| GetMeetingById | 与えられたミーティングIDからミーティング情報を取得する関数 |
| GetMeetingsByUserId | 与えられたユーザーIDから参加するミーティング情報を取得する関数 |
| GetConfirmedMeetingsForHostByUserId | 与えられたユーザーIDからホスト側の確定済みミーティング情報を取得する関数 |
| GetNotConfirmedMeetingsForHostByUserId | 与えられたユーザーIDからホスト側の未確定のミーティング情報を取得する関数 |
| GetNotRespondedMeetingsForHostByUserId | 与えられたユーザーIDからホスト側の未返信のミーティング情報を取得する関数 |
| GetConfirmedMeetingsForGuestByUserId | 与えられたユーザーIDからゲスト側の確定済みミーティング情報を取得する関数 |
| GetNotConfirmedMeetingsForGuestByUserId | 与えられたユーザーIDからゲスト側の未確定のミーティング情報を取得する関数 |
| GetNotRespondedMeetingsForGuestByUserId | 与えられたユーザーIDからゲスト側の未返信のミーティング情報を取得する関数 |

##### IsTitleEmptyメソッド

与えられたミーティングのタイトルが空白かどうか判断する関数 

引数

| 変数名 | 型 | 
| ---- | ---- | 
| m | Meeting |

返り値

| 型 | 説明 |  
| ---- | ---- | 
| bool | 与えられたミーティングのタイトルが空白かどうか |
| error | 発生しうるエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | true / false  | nil |
| 異常レスポンス (mのポインタがnilの時) | false | errors.New("The given Meeting object doesn't exist") |

##### IsHourEmptyメソッド

与えられたミーティングの時間が空白かどうか判断する関数

引数

| 変数名 | 型 | 
| ---- | ---- | 
| m | Meeting |

返り値

| 型 | 説明 |  
| ---- | ---- | 
| bool | 与えられたミーティングの時間が空白かどうか |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | true / false  | nil |
| 異常レスポンス (mのポインタがnilの時) | false | errors.New("The given Meeting object doesn't exist") |

##### IsOnsiteButNoPlaceSpecifiedメソッド

与えられた物理開催のミーティングで開催場所が指定されているか判定する関数

引数

| 変数名 | 型 | 
| ---- | ---- | 
| m | Meeting |

返り値

| 型 | 説明 |  
| ---- | ---- | 
| bool | 与えられた物理開催のミーティングで開催場所が指定されているか |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | true / false  | nil |
| 異常レスポンス (mのポインタがnilの時) | false | errors.New("The given Meeting object doesn't exist") |

##### IsOnlineButNoURLSpecifiedメソッド

与えられたオンライン開催のミーティングで開催URLが指定されているか判定する関数

引数

| 変数名 | 型 | 
| ---- | ---- | 
| m | Meeting |

返り値

| 型 | 説明 |  
| ---- | ---- | 
| bool | 与えられた物理開催のミーティングで開催場所が指定されているか |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | true / false  | nil |
| 異常レスポンス (mのポインタがnilの時) | false | errors.New("The given Meeting object doesn't exist") |

##### IsHybridButNeitherPlaceOrURLSpecifiedメソッド

与えられたハイブリッド開催のミーティングで開催場所と開催URLが両方指定されているか判定する関数

引数

| 変数名 | 型 | 
| ---- | ---- | 
| m | Meeting |

返り値

| 型 | 説明 |  
| ---- | ---- | 
| bool | 与えられたハイブリッド開催のミーティングで開催場所と開催URLが両方指定されているか |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | true / false  | nil |
| 異常レスポンス (mのポインタがnilの時) | false | errors.New("The given Meeting object doesn't exist") |

##### GetMeetingByIdメソッド

与えられたミーティングIDからミーティング情報を取得する関数

引数

| 変数名 | 型 | 
| ---- | ---- | 
| db | ([*gorm.DB](https://pkg.go.dev/gorm.io/gorm#DB)型の構造体ポインタ) |
| Id | int |

返り値

| 型 | 説明 |  
| ---- | ---- | 
| Meeting | 与えられたIdを持つMeetingオブジェクト |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | Meeting  | nil |
| 異常レスポンス (SQLの実行に失敗した時) | Meeting | err |

##### GetMeetingsByUserIdメソッド

与えられたユーザーIDから参加するミーティング情報を取得する関数

引数

| 変数名 | 型 | 
| ---- | ---- | 
| db | ([*gorm.DB](https://pkg.go.dev/gorm.io/gorm#DB)型の構造体ポインタ) |
| UserId | int |

返り値

| 型 | 説明 |  
| ---- | ---- | 
| []Meeting | 与えられたIdを持つMeetingオブジェクトの配列 |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | []Meeting | nil |
| 異常レスポンス (SQLの実行に失敗した時) | []Meeting | err |

##### GetConfirmedMeetingsForHostByUserIdメソッド

与えられたユーザーIDからホスト側の確定済みミーティング情報を取得する関数

引数

| 変数名 | 型 | 
| ---- | ---- | 
| db | ([*gorm.DB](https://pkg.go.dev/gorm.io/gorm#DB)型の構造体ポインタ) |
| UserId | int |

返り値

| 型 | 説明 |  
| ---- | ---- | 
| []Meeting | 与えられたIdを持つMeetingオブジェクトの配列 |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | []Meeting | nil |
| 異常レスポンス (SQLの実行に失敗した時) | []Meeting | err |

##### GetNotConfirmedMeetingsForHostByUserIdメソッド

与えられたユーザーIDからホスト側の未確定のミーティング情報を取得する関数 

引数

| 変数名 | 型 | 
| ---- | ---- | 
| db | ([*gorm.DB](https://pkg.go.dev/gorm.io/gorm#DB)型の構造体ポインタ) |
| UserId | int |

返り値

| 型 | 説明 |  
| ---- | ---- | 
| []Meeting | 与えられたIdを持つMeetingオブジェクトの配列 |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | []Meeting | nil |
| 異常レスポンス (SQLの実行に失敗した時) | []Meeting | err |

##### GetNotRespondedMeetingsForHostByUserIdメソッド

与えられたユーザーIDからホスト側の未返信のミーティング情報を取得する関数

引数

| 変数名 | 型 | 
| ---- | ---- | 
| db | ([*gorm.DB](https://pkg.go.dev/gorm.io/gorm#DB)型の構造体ポインタ) |
| UserId | int |

返り値

| 型 | 説明 |  
| ---- | ---- | 
| []Meeting | 与えられたIdを持つMeetingオブジェクトの配列 |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | []Meeting | nil |
| 異常レスポンス (SQLの実行に失敗した時) | []Meeting | err |

##### GetConfirmedMeetingsForGuestByUserIdメソッド

与えられたユーザーIDからゲスト側の確定済みミーティング情報を取得する関数

引数

| 変数名 | 型 | 
| ---- | ---- | 
| db | ([*gorm.DB](https://pkg.go.dev/gorm.io/gorm#DB)型の構造体ポインタ) |
| UserId | int |

返り値

| 型 | 説明 |  
| ---- | ---- | 
| []Meeting | 与えられたIdを持つMeetingオブジェクトの配列 |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | []Meeting | nil |
| 異常レスポンス (SQLの実行に失敗した時) | []Meeting | err |

##### GetNotConfirmedMeetingsForGuestByUserIdメソッド

引数

| 変数名 | 型 | 
| ---- | ---- | 
| db | ([*gorm.DB](https://pkg.go.dev/gorm.io/gorm#DB)型の構造体ポインタ) |
| UserId | int |

返り値

| 型 | 説明 |  
| ---- | ---- | 
| []Meeting | 与えられたIdを持つMeetingオブジェクトの配列 |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | []Meeting | nil |
| 異常レスポンス (SQLの実行に失敗した時) | []Meeting | err |

##### GetNotRespondedMeetingsForGuestByUserIdメソッド

引数

| 変数名 | 型 | 
| ---- | ---- | 
| db | ([*gorm.DB](https://pkg.go.dev/gorm.io/gorm#DB)型の構造体ポインタ) |
| UserId | int |

返り値

| 型 | 説明 |  
| ---- | ---- | 
| []Meeting | 与えられたIdを持つMeetingオブジェクトの配列 |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | []Meeting | nil |
| 異常レスポンス (SQLの実行に失敗した時) | []Meeting | err |

#### CandidateTimeオブジェクト

##### プロパティ

| 変数名 | 型 | 説明 | 
| ---- | ---- | ---- |
| id  | int | ID | 
| user_id | int | ユーザーID | 
| meeting_id | int | ミーティングID |
| start_time | datetime | 開始時間 | 
| end_time | datetime | 終了時間 | 
| created_at | datetime | 登録日時 | 
| updated_at | datetime | 更新日時 | 

##### メソッド

| メソッド名 | 説明 | 
| ---- | ---- | 
| GetCandidateTimeByMeetingId | 与えられたミーティングIDからミーティング候補時間を取得する関数 |
| GetCandidateTimeByMeetingIdAndUserId | 与えられたミーティングIDとユーザーIDからミーティング候補時間を取得する関数 |
| GetAvailableTimeByMeetingId | 与えられたミーティングIDからミーティング可能な時間帯を取得する関数 |
| Include | 与えられた数が与えられた配列に含まれているか判定する関数 |
| GetLatestStartTime | 与えられた1番遅いミーティング開始時間を取得する関数 |
| GetEarliestEndTime | 与えられた1番早いミーティング終了時間を取得する関数 |
| CreateUserIdList | 与えられたユーザIDリストを作成する関数 |
| IsSameSlice | 与えられた2つの配列が一致しているか判定する関数 |
| SortByStartTime | 与えられたミーティング候補時間の配列を開始時間で並べ替える関数 |
| AvailableTimeIsNotFound | 与えられたミーティング可能時間が存在しないか判定する関数 |

##### GetCandidateTimeByMeetingIdメソッド

与えられたミーティングIDからミーティング候補時間を取得する関数

引数

| 変数名 | 型 | 
| ---- | ---- | 
| db | ([*gorm.DB](https://pkg.go.dev/gorm.io/gorm#DB)型の構造体ポインタ) |
| UserId | int |


返り値

| 型 | 説明 |  
| ---- | ---- | 
| []CandidateTime | 与えられたMeetingIdを持つMeetingオブジェクトの配列 |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | []CandidateTime | nil |
| 異常レスポンス (SQLの実行に失敗した時) | []Candidate | err |

##### GetCandidateTimeByMeetingIdAndUserIdメソッド

与えられたミーティングIDとユーザーIDからミーティング候補時間を取得する関数

引数

| 変数名 | 型 | 
| ---- | ---- | 
| db | ([*gorm.DB](https://pkg.go.dev/gorm.io/gorm#DB)型の構造体ポインタ) |
| MeetingId | int |
| UserId | int |

返り値

| 型 | 説明 |  
| ---- | ---- | 
| []CandidateTime | 与えられたMeetingIdを持つMeetingオブジェクトの配列 |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | []CandidateTime | nil |
| 異常レスポンス (SQLの実行に失敗した時) | []Candidate | err |


##### GetAvailableTimeByMeetingIdメソッド

与えられたミーティングIDからミーティング可能な時間帯を取得する関数

引数

| 変数名 | 型 | 
| ---- | ---- | 
| db | ([*gorm.DB](https://pkg.go.dev/gorm.io/gorm#DB)型の構造体ポインタ) |
| MeetingId | int |
| UserId | int |

返り値

| 型 | 説明 |  
| ---- | ---- | 
| []CandidateTime | 与えられたMeetingIdを持つMeetingオブジェクトの配列 |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | []CandidateTime | nil |
| 異常レスポンス (SQLの実行に失敗した時) | []Candidate | err |

##### Includeメソッド

与えられた数が与えられた配列に含まれているか判定する関数

引数

| 変数名 | 型 | 
| ---- | ---- | 
| numList | []int |
| num | int |

返り値

| 型 | 説明 |  
| ---- | ---- | 
| bool | 与えられたMeetingIdを持つMeetingオブジェクトの配列 |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | true/false | nil |
| 異常レスポンス (numlistのポインタがnilの場合) | false | errors.New("The given int array is nil") |

##### GetLatestStartTimeメソッド

与えられた1番遅いミーティング開始時間を取得する関数

引数

| 変数名 | 型 | 
| ---- | ---- | 
| candidateTimeList | []CandidateTime |

返り値

| 型 | 説明 |  
| ---- | ---- | 
| time.Time | 与えられた1番遅いミーティング開始時間 |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | time.Time | nil |
| 異常レスポンス (candidateTimeListのポインタがnilの場合) | time.Time | errors.New("The given list of candidateTime is nil") |

##### GetEarliestEndTimeメソッド

与えられた1番早いミーティング終了時間を取得する関数

引数

| 変数名 | 型 | 
| ---- | ---- | 
| candidateTimeList | []CandidateTime |

返り値

| 型 | 説明 |  
| ---- | ---- | 
| time.Time | 与えられた1番早いミーティング終了時間 |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | time.Time | nil |
| 異常レスポンス (candidateTimeListのポインタがnilの場合) | time.Time | errors.New("The given list of candidateTime is nil") |

##### CreateUserIdListメソッド

与えられたユーザIDリストを作成する関数 


引数

| 変数名 | 型 | 
| ---- | ---- | 
| candidateTimeList | []CandidateTime |

返り値

| 型 | 説明 |  
| ---- | ---- | 
| []int | ユーザIDの配列 |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | []int | nil |
| 異常レスポンス (candidateTimeListのポインタがnilの場合) | []int | errors.New("The given list of candidateTime is nil") |

##### IsSameSliceメソッド

与えられた2つの配列が一致しているか判定する関数

引数

| 変数名 | 型 | 
| ---- | ---- | 
| slice1 | []int |
| slice2 | []int |

返り値

| 型 | 説明 |  
| ---- | ---- | 
| bool | 与えられた2つの配列が一致しているか |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | true/false | nil |
| 異常レスポンス (slice1とslice2の少なくともいずれかが空だった時) | false | errors.New("The given int array is empty") | 

##### SortByStartTimeメソッド

与えられたミーティング候補時間の配列を開始時間で並べ替える関数

返り値

| 型 | 説明 |  
| ---- | ---- | 
| []CandidateTime | 与えられたMeetingIdを持つMeetingオブジェクトの配列 |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | []CandidateTime | nil |
| 異常レスポンス (SQLの実行に失敗した時) | []Candidate | err |

##### AvailableTimeIsNotFoundメソッド

与えられたミーティング可能時間が存在しないか判定する関数

引数

| 変数名 | 型 | 
| ---- | ---- | 
| candidateTimeList | []CandidateTime |

返り値

| 型 | 説明 |  
| ---- | ---- | 
| bool | 与えられたミーティング可能時間が存在しないか |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | true/false | nil |
| 異常レスポンス (SQLの実行に失敗した時) | []Candidate | errors.New("The given list of candidateTime is nil") |

#### Participantオブジェクト

##### プロパティ

| 変数名 | 型 | 説明 | 
| ---- | ---- | ---- |
| id  | int | ID | 
| user_id | int | ユーザーID | 
| meeting_id | int | ミーティングID |
| is_host | bool | 主催者かどうか | 
| has_responded | bool | 回答の有無 | 
| created_at | datetime | 登録日時 | 
| updated_at | datetime | 更新日時 | 

##### メソッド

| メソッド名 | 説明 | 
| ---- | ---- | 
| GetParticipantListByMeetingId | 与えられたミーティングIDから参加者リストを取得する関数 |
| GetParticipantByUserIdAndMeetingId | 与えられたミーティングIDとユーザーIDから参加者の情報を取得する関数 |
| ConvertToParticipant | 与えられたParticipantWithUserNameオブジェクトをParticipantオブジェクトに変換する関数 |
| ConvertToParticipantWithUserName | 与えられたParticipantWithUserNameオブジェクトをParticipantオブジェクトに変換する関数 |
| ConvertToParticipantWithUserNameList | 与えられたParticipantオブジェクトの配列をParticipantWithUserNameオブジェクトの配列に変換する関数 |
| ConvertToParticipantList | 与えられたParticipantWithUserNameオブジェクトの配列をParticipantオブジェクトの配列に変換する関数 |
| Min | 与えられた2つの整数の小さい方を取得する関数 |
| HostIsInParticipant | 与えられたParticipantオブジェクトの配列の中にHostが含まれているか判断する関数 |

##### GetParticipantListByMeetingIdメソッド

与えられたミーティングIDから参加者リストを取得する関数

引数

| 変数名 | 型 | 
| ---- | ---- | 
| db | ([*gorm.DB](https://pkg.go.dev/gorm.io/gorm#DB)型の構造体ポインタ) |
| Id | int |

返り値

| 型 | 説明 |  
| ---- | ---- | 
| []Participant | 与えられたミーティングIDのミーティングの参加者リスト |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | []Participant | nil |
| 異常レスポンス (SQLの実行に失敗した時) | []Participante | err |

##### GetParticipantByUserIdAndMeetingIdメソッド

与えられたミーティングIDとユーザーIDから参加者の情報を取得する関数

引数

| 変数名 | 型 | 
| ---- | ---- | 
| db | ([*gorm.DB](https://pkg.go.dev/gorm.io/gorm#DB)型の構造体ポインタ) |
| UserId | int |
| MeetingId | int |

返り値

| 型 | 説明 |  
| ---- | ---- | 
| []Participant | 与えられたミーティングIDのミーティングの参加者リスト |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | []Participant | nil |
| 異常レスポンス (SQLの実行に失敗した時) | []Participante | err |

##### ConvertToParticipantメソッド

与えられたParticipantWithUserNameオブジェクトをParticipantオブジェクトに変換する関数

引数

| 変数名 | 型 | 
| ---- | ---- | 
| db | ([*gorm.DB](https://pkg.go.dev/gorm.io/gorm#DB)型の構造体ポインタ) |
| pw | ParticipantWithUserName |

返り値

| 型 | 説明 |  
| ---- | ---- | 
| Participant | Participantオブジェクト |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | Participant | nil |
| 異常レスポンス (pwのポインタがnilだった時) | Participant | errors.New("The given ParticipantWithUserName object is nil") |

##### ConvertToParticipantWithUserNameメソッド

与えられたParticipantオブジェクトをParticipantWithUserNameオブジェクトに変換する関数

引数

| 変数名 | 型 | 
| ---- | ---- | 
| db | ([*gorm.DB](https://pkg.go.dev/gorm.io/gorm#DB)型の構造体ポインタ) |
| p | Participant |

返り値

| 型 | 説明 |  
| ---- | ---- | 
| ParticipantWithUserName | 与えられたミーティングIDのミーティングの参加者リスト |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | ParticipantWithUserName | nil |
| 異常レスポンス (pのポインタがnilだった時) | ParticipantWithUserName | errors.New("The given Participant object is nil") |

##### ConvertToParticipantWithUserNameListメソッド

与えられたParticipantオブジェクトの配列をParticipantWithUserNameオブジェクトの配列に変換する関数

引数

| 変数名 | 型 | 
| ---- | ---- | 
| db | ([*gorm.DB](https://pkg.go.dev/gorm.io/gorm#DB)型の構造体ポインタ) |
| plist | []Participant |

返り値

| 型 | 説明 |  
| ---- | ---- | 
| []ParticipantWithUserName | ParticipantWithUserNameの配列 |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | ParticipantWithUserName | nil |
| 異常レスポンス (与えられた配列が空の時) | ParticipantWithUserName | errors.New("The given ParticipantWithUserName list is empty") |

##### ConvertToParticipantListメソッド

与えられたParticipantWithUserNameオブジェクトの配列をParticipantオブジェクトの配列に変換する関数


引数

| 変数名 | 型 | 
| ---- | ---- | 
| db | ([*gorm.DB](https://pkg.go.dev/gorm.io/gorm#DB)型の構造体ポインタ) |
| pwlist | []ParticipantWithUserName |

返り値

| 型 | 説明 |  
| ---- | ---- | 
| []Participant | 与えられたミーティングIDのミーティングの参加者リスト |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | []Participant | nil |
| 異常レスポンス (与えられた配列が空の時) | []Participant | errors.New("The given Participant list is empty") |

##### Minメソッド

与えられた2つの整数の小さい方を取得する関数

引数

| 変数名 | 型 | 
| ---- | ---- | 
| a | int |
| b | int |

返り値

| 型 | 説明 |  
| ---- | ---- | 
| int | 与えられた2つの整数の小さい方 |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス (aがb以下の時) | a | nil |
| 正常レスポンス (bがa未満の時) | b | nil |
| 異常レスポンス (aかbのポインタがnilだった場合) | -1 | errors.New("The given integer is nil") |

##### HostIsInParticipantメソッド

与えられたParticipantオブジェクトの配列の中にHostが含まれているか判断する関数

引数

| 変数名 | 型 | 
| ---- | ---- | 
| plist | []Participant |

返り値

| 型 | 説明 |  
| ---- | ---- | 
| bool | 与えられたParticipantオブジェクトの配列の中にHostが含まれているか |
| error | 発生したエラー |

| シチュエーション | bool | error |  
| ---- | ---- | ---- | 
| 正常レスポンス | true/false | nil |
| 異常レスポンス (与えられた配列が空の時) | false | errors.New("The given Participant list is empty") |


### フロントエンド

#### UI

##### トップページ

![トップページ](./img/toppage.png)

##### 新規作成ページ

![新規作成ページ](./img/signup.png)

##### ログインページ

![ログインページ](./img/login.png)

##### ミーティング新規作成ページ

![ミーティング新規作成ページ](./img/newmeeting.png)

##### 主催・参加ミーティング一覧ページ (日程確定)

![主催・参加ミーティング一覧ページ (日程確定)](./img/confirmed.png)

##### 主催・参加ミーティング一覧ページ (日程未確定)

![主催・参加ミーティング一覧ページ (日程未確定)](./img/non-confirmed.png)

##### 参加ミーティング一覧ページ (未返信)

![参加ミーティング一覧ページ (未返信)](./img/non-responded.png)

##### ミーティング時間決定ページ 

![ミーティング時間決定ページ](./img/choose-time.png)

##### ミーティング候補時間決定ページ

![ミーティング候補時間決定ページ](./img/choose-candidate-time.png)

#### ページ遷移図

![general](./img/diagram.png)
![dashboard](./img/diagram2.png)

### バックエンド



#### API定義

##### 共通設定

各APIは次のフォーマットのURLとする

```
https://lets-schedule.net/<エンドポイントのパス>
```

通信プロトコル: HTTPS
APIの種類: REST API
インターフェース: JSON
文字コード: UTF-8

##### ユーザー関連API

###### 新規ユーザー作成

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

##### ミーティング関連API

###### ミーティング新規作成

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


###### ミーティング情報取得

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
| 失敗時 | error | string | デフォルトは「エラーが発生しました」 |

###### 日時が決定した主催ミーティング情報取得

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
| 失敗時 | error | string | デフォルトは「エラーが発生しました」 |


###### 日時が決定していない主催ミーティング情報取得

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
| 失敗時 | error | string | デフォルトは"エラーが発生しました |


###### 返信していない主催ミーティング情報取得

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
| 失敗時 | error | string | デフォルトは「エラーが発生しました」 |


###### 日時が決定している参加ミーティング情報取得

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
| 失敗時 | error | string | デフォルトは「エラーが発生しました」 |


###### 日時が決定していない参加ミーティング情報取得

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
| 失敗時 | error | string | デフォルトは |


###### 返信していない参加ミーティング情報取得

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
| 失敗時 | error | string | デフォルトは「エラーが発生しました」 |

##### 候補日時関連API

###### 候補日時の新規登録

エンドポイントPOST /YXBpL3Jlc3RyaWN0ZWQvY2FuZGlkYXRlX3RpbWVzL25ldw==

リクエストパラメータ

| キー名 | 型 (変数) | 概要 |
| -- | -- | -- |
| candidate_time_list | CandidateTime[] | 候補日時の配列 |


レスポンスパラメータ

| シチュエーション | キー名 | 型 (変数) | 概要 | 
| -- | -- | -- | -- |
| 成功時 | candidate_time | CandidateTime[] | 候補日時の配列 |
| 失敗時 | error | string | デフォルトは「エラーが発生しました」 |


###### 候補日時の取得

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
| 失敗時 | error | string | デフォルトは「エラーが発生しました」 |




###### 候補日時の編集

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
| 失敗時 | error | string | デフォルトは「エラーが発生しました」 |


##### 参加者関連API 

###### 参加者の新規登録

エンドポイントPOST /YXBpL3Jlc3RyaWN0ZWQvcGFydGljaXBhbnRzL25ldw==

リクエストパラメータ

| キー名 | 型 (変数) | 概要 |
| -- | -- | -- |
| participants | Participant[] | 候補日時の配列 |


レスポンスパラメータ

| シチュエーション | キー名 | 型 (変数) | 概要 | 
| -- | -- | -- | -- |
| 成功時 | participants | Participant[] | 候補日時の配列 |
| 失敗時 | error | string | デフォルトは「エラーが発生しました」 |


###### 参加者の取得

エンドポイント GET /YXBpL3Jlc3RyaWN0ZWQvcGFydGljaXBhbnRz/:meeting_id

リクエストパラメータ

| キー名 | 型 (変数) | 概要 |
| -- | -- | -- |
| meeting_id | int | ミーティングID |


レスポンスパラメータ

| シチュエーション | キー名 | 型 (変数) | 概要 | 
| -- | -- | -- | -- |
| 成功時 | participants | Participant[] | 候補日時の配列 |
| 失敗時 | error | string | デフォルトは「エラーが発生しました」 |

###### 参加者の編集

エンドポイント PUT /YXBpL3Jlc3RyaWN0ZWQvcGFydGljaXBhbnRzL21lZXRpbmc=/:meeting_id

リクエストパラメータ

| キー名 | 型 (変数) | 概要 |
| -- | -- | -- |
| meeting_id | int | ミーティングID |


レスポンスパラメータ

| シチュエーション | キー名 | 型 (変数) | 概要 | 
| -- | -- | -- | -- |
| 成功時 | participants | Participant[] | 参加者の配列 |
| 失敗時 | error | string | デフォルトは「エラーが発生しました」 |


### DB

#### 前提

- 文字コードはUTF-8を用いる
- タイムゾーンはAsia/Tokyoを用いる
- sys.columnsは使用しない

#### Usersテーブル

| 項目名 (日本語) | 項目名 (変数) | 型 | NULL許容か | 備考 | 
| ---- | ---- | ---- | ---- | ---- |
| ID | id | bigint unsigned(11) | False | primary key auto increment |
| ユーザー名 | user_name | varchar(191) | False | |
| メールアドレス | email_address | varchar(191) | False | |
| パスワード | password | varchar(191) | False | |
| 管理者かどうか | is_admin | boolean | False | |
| ログイン可能か | can_login | boolean | False | |
| 登録日時 | created_at | datetime | False | default current_datetime |
| 更新日時 | updated_at | datetime | False | default current_datetime on update current_datetime | 

#### Meetingsテーブル

| 項目名 (日本語) | 項目名 (変数) | 型 | NULL許容か | 備考 | 
| ---- | ---- | ---- | ---- | ---- |
| ID | id | bigint unsigned(11) | False | primary key auto increment |
| ミーティング名 | title | varchar(191) | False | |
| 概要 | description | text | False | |
| 物理開催かどうか | is_onsite | boolean | False | |
| 集合場所 | place | varchar(191) | False | |
| ミーティングURL | url | varchar(191) | False | |
| 全員が回答したか | all_participants_responded | boolean | False | |
| 確定したか | is_confirmed | boolean | False | |
| 開始日時 | start_time | datetime | False | |
| 終了日時 | end_time | datetime | False | |
| ミーティング時間(分) | minutes | int | False | |
| 登録日時 | created_at | datetime | False | default current_datetime |
| 更新日時 | updated_at | datetime | False | default current_datetime on update current_datetime |

#### Participantテーブル

| 項目名 (日本語) | 項目名 (変数) | 型 | NULL許容か | 備考 | 
| ---- | ---- | ---- | ---- | ---- |
| ID | id  | bigint unsigned(11) | False | primary key auto increment |
| ユーザーID | user_id | bigint unsigned(11) | False | | 
| ミーティングID | meeting_id | bigint unsigned(11) | False | |
| 主催者かどうか | is_host | boolean | False | |
| 回答の有無 | has_responded | boolean | False | |
| 登録日時 | created_at | datetime | False | default current_datetime |
| 更新日時 | updated_at | datetime | False | default current_datetime on update current_datetime |

#### CandidateTimeテーブル

| 項目名 (日本語) | 項目名 (変数) | 型 | NULL許容か | 備考 | 
| ---- | ---- | ---- | ---- | ---- |
| ID | id  | bigint unsigned(11) | False | primary key auto increment |
| ユーザーID | user_id | bigint unsigned(11) | False | |
| ミーティングID | meeting_id | bigint unsigned(11) | False | |
| 開始時間 | start_time | datetime | False | |
| 終了時間 | end_time | datetime | False | |
| 登録日時 | created_at | datetime | False | default current_datetime |
| 更新日時 | updated_at | datetime | False | default current_datetime on update current_datetime |

#### キー一覧

##### Usersテーブル

| 種類 | キー名 | 備考 |
| ---- | ---- | ---- |
| プライマリーキー | id | |

##### Meetingsテーブル

| 種類 | キー名 | 備考 |
| ---- | ---- | ---- |
| プライマリーキー | id | |

##### Participantテーブル

| 種類 | キー名 | 備考 |
| ---- | ---- | ---- |
| プライマリーキー | id | |
| 外部キー | meeting_id | REFERENCES meeting(meeting_id) |

##### CandidateTimeテーブル

| 種類 | キー名 | 備考 |
| ---- | ---- | ---- |
| プライマリーキー | id | |
| 外部キー | user_id | REFERENCES user(user_id) |
| | meeting_id | REFERENCES meeting(meeting_id) |


## 非機能要件

### 可用性

#### SLA

今回は２台のEC2インスタンスの冗長構成を採用した。

インスタンス1台のSLAが99.5%であることから、サービスのSLAは

```
1 - (1 - 0.995)**2 = 0.999975
```

よって、SLAは99.9%は担保される。

### パフォーマンス

#### 同時接続数

同時接続数は250とする。

なお、根拠は以下の通りである。

研究室のメンバー (教員・学生含め) 112人 
合唱団のメンバー 38人
日常的に連絡を取っている友人 10人 
友人が参加する集まりの参加者数 10人程度

合計 112 + 38 + 10 * 10 = 250 

### セキュリティ

#### 証明書

WebサーバーとAPIサーバーには両方ともLet's Encryptのcertbotで取得した証明書を使用する。
なお、この場合の暗号化アルゴリズムはP-384のECDSAとする。

#### アクセス制御

##### アルゴリズム選定

ユーザー認証に使用するアルゴリズムを決定する際にBasic認証、JWT認証、OAuthの3つを比較して、JWT認証を採用する。
一時的な認証情報が提供できることと、使用するメールアドレスに制限がないことが選定理由である。

|  | Basic | JWT | OAuth | 
| ---- | ---- | ---- | ---- |
| 認証情報 | ユーザー名とパスワード | アクセストークン | アクセストークン |
| 有効期限 | 永続 | (有効期限を設定すれば)一時的 | 一時的 |
| 使用できるメールアドレス | 制限なし | 制限なし | 限定 |

##### 認証

![認証](./img/jwt-authentication.png)

##### 認可

![認可](./img/jwt-authorization.png)

### オブザーバビリティ

#### ホストVM

- アクセスログ: nginxのHTTPリクエストログ
- エラーログ: nginxのエラーログ
- アップストリームログ: webサーバーへのリダイレクトログ
- レスポンスログ: nginxのHTTPレスポンスログ

#### Webサーバー

- アクセスログ: nginxのアクセスログ
- エラーログ: nginxのエラーログ

#### APIサーバー

- アクセスログ: HTTPリクエストのログ
- レスポンスログ: HTTPレスポンスのログ
- クエリログ: DBで発行したクエリのログ

#### DBサーバー

- アクセスログ: DBサーバーのアクセスログ

### 保守性

#### インフラ構成図

![architecture](./img/architecture.png)

#### ネットワーク構成

##### VPC 

vpc  
- CIDR:20.0.0.0/16 (ap-northeast-1)

##### サブネット

subnet-public1 
- CIDR:20.0.1.0/24 
- Region: ap-northeast-1a

subnet-public2 
- CIDR: 20.0.2.0/24 
- Region: ap-northeast-1b

subnet-private1 
- CIDR: 20.0.3.0/24 
- Region: ap-northeast-1a

subnet-private2 
- CIDR: 20.0.4.0/24 
- Region: ap-northeast-1b

##### インターネットゲートウェイ
internet-gateway
- VPC:vpc

##### NATゲートウェイ
nat-gateway
サブネット: subnet-public1
Elastic IP: 新規に割り当て


#### セキュリティグループ

##### ALB

###### インバウンドルール

| プロトコル | ポート範囲 | ソース | 説明 | 
| ---- | ---- | ---- | ---- |
| TCP | 10000 | 0.0.0.0/0 | カスタムHTTPポート (公開アクセス)|
| TCP | 10001 | 0.0.0.0/0 | カスタムHTTPSポート (公開アクセス)|

###### アウトバウンドルール

| プロトコル | ポート範囲 | ソース |
| ---- | ---- | ---- |
| 全トラフィック | 全て | 0.0.0.0/0 | 

##### Webサーバー

###### インバウンドルール

| プロトコル | ポート範囲 | ソース | 説明 | 
| ---- | ---- | ---- | ---- |
| TCP | 10002 | 0.0.0.0/0 | カスタムHTTPポート (公開アクセス)|
| TCP | 10003 | 0.0.0.0/0 | カスタムHTTPSポート (公開アクセス)|
| TCP | 10004 | 0.0.0.0/0 | ALBのセキュリティグループ |

###### アウトバウンドルール

| プロトコル | ポート範囲 | ソース |
| ---- | ---- | ---- |
| 全トラフィック | 全て | 0.0.0.0/0 | 

##### APIサーバー

###### インバウンドルール

| プロトコル | ポート範囲 | ソース |
| ---- | ---- | ---- |
| TCP | 10005 | 0.0.0.0/0 |

###### アウトバウンドルール

| プロトコル | ポート範囲 | ソース |
| ---- | ---- | ---- |
| 全トラフィック | 全て | 0.0.0.0/0 | 


##### DBサーバー

###### インバウンドルール

| プロトコル | ポート範囲 | ソース |
| ---- | ---- | ---- |
| TCP | 10006 | 0.0.0.0/0 |

###### アウトバウンドルール

| プロトコル | ポート範囲 | ソース |
| ---- | ---- | ---- |
| 全トラフィック | 全て | 0.0.0.0/0 | 


## 使用技術

### フロントエンド

開発者が実務での使用経験のあるReact, Vue3のComposition API、Vue3のOption APIの中でVue3のOption APIを選定した。
選定基準は学習コストの低さと可読性の高さの2つで、両方を満たす点が選定理由である。


|  | React | Vue3 (Composition API) | Vue3 (Option API) | 
| ---- | ---- | ---- | ---- |
| 学習コスト | 高 | 中 | 低 |
| 可読性 | 低 | 中 | 高 |


### バックエンド

開発者が実務での使用経験のあるPython、Ruby、Golangの中でGolangを選定した。
選定基準は可読性の高さと静的型付けであるかの2つで、両方を満たす点が選定理由である。

|  | Python | Ruby | Golang | 
| ---- | ---- | ---- | ---- |
| 可読性 | 高 | 高 | 高 |
| 静的型付けかどうか | x | x | ◯ |


### DB

開発者が実務での使用経験のあるMySQL, MariaDB, PostgreSQLの中でMariaDBを選定した。
選定基準はRDSであることとdockerイメージがARMアーキテクチャのCPU(特に開発者の使用マシンであるM1チップ)で動作することである。

|  | MySQL | MariaDB | PostgreSQL | 
| ---- | ---- | ---- | ---- |
| DBの種類 | RDBMS | RDBMS | OBDBMS |
| ARM64対応 | x | o | o |



### インフラ

<!-- LambdaやECS・EKSなどのマネージドサービスとは違い、EC2インスタンスはVPC（Virtual Private Cloud）内で動作するため、ネットワーキングとセキュリティを細かく設定できる点が挙げられる。 -->

## 参考文献

Amazon Compute Service Level Agreement - Amazon AWS
https://aws.amazon.com/compute/sla/

https://dekh.medium.com/the-complete-guide-to-json-web-tokens-jwt-and-token-based-authentication-32501cb5125c
