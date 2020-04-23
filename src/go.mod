module github.com/big-gate-0219/workrecord-go

go 1.14

replace (
	databases => ./databases
	middlewares => ./middlewares
	models => ./models
	routes => ./routes
	web/api => ./web/api

)

require (
	databases v0.0.0-00010101000000-000000000000 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/howeyc/fsnotify v0.9.0 // indirect
	github.com/jinzhu/gorm v1.9.12 // indirect
	github.com/joho/godotenv v1.3.0
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.3.0 // indirect
	github.com/pilu/config v0.0.0-20131214182432-3eb99e6c0b9a // indirect
	github.com/pilu/fresh v0.0.0-20190826141211-0fa698148017 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.5.0
	github.com/valyala/fasthttp v1.9.0 // indirect
	golang.org/x/crypto v0.0.0-20200406173513-056763e48d71 // indirect
	middlewares v0.0.0-00010101000000-000000000000
	models v0.0.0-00010101000000-000000000000 // indirect
	routes v0.0.0-00010101000000-000000000000
	rsc.io/quote v1.5.2
	web/api v0.0.0-00010101000000-000000000000 // indirect
)
