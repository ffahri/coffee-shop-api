package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go.opentelemetry.io/otel/api/trace"
)

type Util struct {
	DB     *sqlx.DB
	Tracer trace.Tracer
}

func (u *Util) ListAllDrinks(context *gin.Context) {
	ctx, span := u.Tracer.Start(context.Request.Context(), "listAllDrinks")
	defer span.End()
	context.JSON(200, u.List(ctx))
}

func (u *Util) OrderNewDrink(context *gin.Context) {
	ctx, span := u.Tracer.Start(context.Request.Context(), "orderNewDrink")
	defer span.End()
	status, message := u.Order(ctx, context.Request.Body)
	context.JSON(status, message)
}
