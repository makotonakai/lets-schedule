package controller

import (
    "fmt"
    "net/http"
    "api/model"
    "github.com/labstack/echo"
    connect "github.com/MakotoNakai/lets-schedule/api/connect"
)

func Create() echo.HandlerFunc{
    return func(c echo.Context) error {
        //Project/api/database/connect.goで定義したやつ。
        db := connect.Connect()
        //この記述は絶対に必要でdeferを書くことでメソッド終了後に発動し、DBをCloseしてくれる。そのためこの下にメソッドを書いても問題ありません。
        defer db.Close()

        result := new(model.User)
        //c.Bind()でリクエストボディから更新データを取得。
        //err変数にc.Bindを入れてエラーがnilでなければエラーを返す。
        if err := c.Bind(result); err != nil {
            return err
        }
        // 基本的にCreate(),Update(),Delete()などは値ではなくアドレス(&XXX)を渡す。
        // &を使うことで変数のアドレスを参照することができる。
        db.Create(&result)    
        return c.JSON(http.StatusOK, result)
    }
}