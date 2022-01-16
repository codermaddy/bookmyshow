module www.example.com/v1

replace (
	example.com/authentication/authent => ./authentication
	example.com/models => ./models
)

go 1.17

require example.com/authentication/authent v0.0.0-00010101000000-000000000000

require (
	example.com/models v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
)
