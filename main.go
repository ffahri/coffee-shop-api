package main

import (
	v1 "coffee-shop-api/pkg/api/v1"
	"coffee-shop-api/pkg/db"
	"github.com/gin-gonic/gin"
	gintrace "go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin"
	otelglobal "go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/exporters/stdout"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"log"
)

func initTracer() {
	exporter, err := stdout.NewExporter(stdout.WithPrettyPrint())
	if err != nil {
		log.Fatal(err)
	}
	cfg := sdktrace.Config{
		DefaultSampler: sdktrace.AlwaysSample(),
	}
	tp, err := sdktrace.NewProvider(
		sdktrace.WithConfig(cfg),
		sdktrace.WithSyncer(exporter),
	)
	if err != nil {
		log.Fatal(err)
	}
	otelglobal.SetTraceProvider(tp)
}

func main() {
	initTracer()
	u := v1.Util{
		DB:     db.Start(),
		Tracer: otelglobal.Tracer("coffee-shop"),
	}
	r := gin.New()
	r.Use(gintrace.Middleware("coffee-shop-api"))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/list", u.ListAllDrinks)
	r.POST("/order", u.OrderNewDrink)
	//management
	r.POST("/management/store") //TODO AUTH

	r.Run()
}
