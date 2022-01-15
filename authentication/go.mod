module example.com/v1

replace example.com/models => ../models/

go 1.17

require (
	example.com/models v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.6.0
)
