module coffee-shop-api

go 1.14

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jmoiron/sqlx v1.2.0
	github.com/rs/zerolog v1.19.0
	go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin v0.11.0
	go.opentelemetry.io/otel v0.11.0
	go.opentelemetry.io/otel/exporters/trace/jaeger v0.11.0
	go.opentelemetry.io/otel/sdk v0.11.0
)
