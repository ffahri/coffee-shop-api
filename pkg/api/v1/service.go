package v1

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"io"
	"io/ioutil"
)

func (u *Util) List(ctx context.Context) []drinks {
	spanContext, span := u.Tracer.Start(ctx, "database-list")
	defer span.End()
	drinkSlice := []drinks{}
	err := u.DB.SelectContext(ctx, &drinkSlice, "SELECT * FROM Drinks")
	if err != nil {
		span.RecordError(spanContext, err)
		log.Error().Err(err).Msg("select statement error in list function")
	}
	return drinkSlice
}

func (u *Util) Order(ctx context.Context, body io.ReadCloser) (int, gin.H) {
	spanContext, span := u.Tracer.Start(ctx, "order")
	defer span.End()
	byteArr, err := ioutil.ReadAll(body)
	if err != nil {
		span.RecordError(spanContext, err)
		log.Error().Err(err).Msg("body read error in order function")
		return 500, gin.H{"Message": "Internal Error (I/O)"}
	}
	order := OrderRequest{}
	err = json.Unmarshal(byteArr, &order)
	if err != nil {
		span.RecordError(spanContext, err)
		log.Error().Err(err).Msg("json unmarshall error in order function")
		return 422, gin.H{"Message": "UnprocessableEntity"}
	}
	if !validateOrder(order) {
		return 422, gin.H{"Message": "UnprocessableEntity"}
	}
	stmt, err := u.DB.PreparexContext(ctx, "SELECT Stock from Drinks where Id = ?")
	if err != nil {
		span.RecordError(spanContext, err)
		log.Error().Err(err).Msg("prepare statement error in order function")
		return 500, gin.H{"Message": "Internal Server Error"}
	}
	var stock int
	err = stmt.Get(&stock, order.Id)
	if err != nil {
		span.RecordError(spanContext, err)
		log.Error().Err(err).Msg("json unmarshall error in order function")
		return 404, gin.H{"Message": "Drink not found"}
	}
	if order.Quantity > stock {
		return 400, gin.H{"Message": "We don't have enough stock for this order"}
	}

	stock = stock - order.Quantity
	tx, err := u.DB.BeginTx(ctx, nil)
	if err != nil {
		span.RecordError(spanContext, err)
		log.Error().Err(err).Msg("tx error in order function")
		return 500, gin.H{"Message": "Internal Server Error"}
	}
	transactionStmt, err := tx.PrepareContext(ctx, "UPDATE Drinks SET Stock = ? where Id = ?")
	if err != nil {
		span.RecordError(spanContext, err)
		log.Error().Err(err).Msg("tx prepare statement error in order function")
		return 500, gin.H{"Message": "Internal Server Error"}
	}
	res, err := transactionStmt.Exec(stock, order.Id)
	if err != nil {
		span.RecordError(spanContext, err)
		log.Error().Err(err).Msg("tx exec statement error in order function")
		return 500, gin.H{"Message": "Internal Server Error"}
	}
	rCount, err := res.RowsAffected()
	if rCount > 0 {
		err = tx.Commit()
		if err != nil {
			span.RecordError(spanContext, err)
			log.Error().Err(err).Msg("tx commit error in order function")
		}
		return 200, gin.H{"Message": "Order Placed!"}
	}
	err = tx.Rollback()
	if err != nil {
		span.RecordError(spanContext, err)
		log.Error().Err(err).Msg("tx rollback error in order function")
	}
	log.Err(err).Msg("order could not be placed")
	return 422, gin.H{"Message": "Order could not be placed !"}
}

func validateOrder(request OrderRequest) bool {
	if request.Quantity < 0 || request.Id == "" {
		return false
	}
	return true
}
