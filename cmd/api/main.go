package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rajanyadav1162/go-microservice-system/internal/db"
	"github.com/rajanyadav1162/go-microservice-system/internal/event"
	"github.com/rajanyadav1162/go-microservice-system/internal/model"
	"github.com/rajanyadav1162/go-microservice-system/internal/msg"
	_ "github.com/rajanyadav1162/go-microservice-system/internal/msg"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"time"
)

var natsConn *msg.NATS

func init() {
	zerolog.TimeFieldFormat = time.RFC3339
	log.Logger = log.Output(os.Stdout).With().Timestamp().Logger()
	gin.SetMode(gin.ReleaseMode) // quieter console
}

type createOrderReq struct {
	UserID    string  `json:"user_id" binding:"required"`
	ConcertID string  `json:"concert_id" binding:"required"`
	Qty       int     `json:"qty" binding:"required,min=1"`
	Amount    float64 `json:"amount" binding:"required,gt=0"`
}

func createOrder(c *gin.Context) {
	var req createOrderReq
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Warn().Err(err).Msg("bad request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order := model.TicketOrder{
		UserID:    req.UserID,
		ConcertID: req.ConcertID,
		Qty:       req.Qty,
		Amount:    req.Amount,
		Status:    "pending",
	}
	if err := db.DB.Create(&order).Error; err != nil {
		log.Error().Err(err).Msg("db create failed")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}
	log.Info().Uint("orderID", order.ID).Msg("order created")
	c.JSON(http.StatusCreated, order)

	//push event to jetstream................................
	ev := event.OrderCreated{
		OrderID:   fmt.Sprintf("%d", order.ID),
		UserID:    order.UserID,
		ConcertID: order.ConcertID,
		Qty:       order.Qty,
		Amount:    order.Amount,
		Status:    order.Status,
		CreatedAt: order.CreatedAt,
	}
	seq, err := natsConn.PublishEvent(context.Background(), "ORDERS.created", ev)
	if err != nil {
		log.Error().Err(err).Msg("publish failed")
		// we still return 201 to caller; publisher is best-effort for now
	} else {
		log.Info().Uint64("seq", seq).Msg("event published to ORDERS.created")
	}
}

func main() {
	dsn := "host=localhost user=tickets password=tickets dbname=ticketsdb port=5432 sslmode=disable"
	if err := db.Init(dsn); err != nil {
		log.Fatal().Err(err).Msg("db init failed")
	}
	log.Info().Msg("db connected & migrated")

	//connecting jetstream
	var err error
	natsConn, err = msg.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatal().Err(err).Msg("nats connect failed")
	}
	defer natsConn.Close()
	log.Info().Msg("nats jetstream connected")
	defer natsConn.Close()
	defer natsConn.Close()
	log.Info().Msg("nats jetstream connected")

	log.Info().Msg("starting api service")
	r := gin.New()
	r.Use(gin.Recovery(), requestLogger())
	r.GET("/ping", pingHandler)
	r.POST("/orders", createOrder)
	if err := r.Run(":8080"); err != nil {
		log.Fatal().Err(err).Msg("server crashed")
	}
}

func requestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.Info().
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Int("status", c.Writer.Status()).
			Dur("latency", time.Since(start)).
			Msg("request")
	}
}

func pingHandler(c *gin.Context) {
	log.Debug().Msg("ping called")
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
