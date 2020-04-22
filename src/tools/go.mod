module app/tools

go 1.14

replace (
	databases => ./databases
	models => ./models
)

require (
	databases v0.0.0-00010101000000-000000000000 // indirect
	github.com/jinzhu/gorm v1.9.12 // indirect
	github.com/joho/godotenv v1.3.0 // indirect
	github.com/sirupsen/logrus v1.5.0 // indirect
	models v0.0.0-00010101000000-000000000000 // indirect
)
