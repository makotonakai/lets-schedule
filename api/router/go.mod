module github.com/MakotoNakai/lets-schedule/api/router

go 1.17

replace github.com/MakotoNakai/lets-schedule/api/controller => ../controller

replace github.com/MakotoNakai/lets-schedule/api/dbconnection => ../db

replace github.com/MakotoNakai/lets-schedule/api/model => ../model

require (
	github.com/MakotoNakai/lets-schedule/api/controller v0.0.0-00010101000000-000000000000
	github.com/labstack/echo v3.3.10+incompatible
)

require (
	github.com/MakotoNakai/lets-schedule/api/dbconnection v0.0.0-00010101000000-000000000000 // indirect
	github.com/MakotoNakai/lets-schedule/api/model v0.0.0-00010101000000-000000000000 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/jinzhu/gorm v1.9.16 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/labstack/gommon v0.3.1 // indirect
	github.com/mattn/go-colorable v0.1.11 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	golang.org/x/crypto v0.0.0-20191205180655-e7c4368fe9dd // indirect
	golang.org/x/net v0.0.0-20200324143707-d3edc9973b7e // indirect
	golang.org/x/sys v0.0.0-20211103235746-7861aae1554b // indirect
	golang.org/x/text v0.3.0 // indirect
)
