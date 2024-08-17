package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	binance_dto "shared/binance/dto"
	message_broker "shared/broker"
	"shared/config"
	db "shared/db/sqlc"
	"shared/utils"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	origin = "http://localhost"
)

func main() {
	time.Sleep(time.Second * 3)

	config.SetEnv()
	db.Connect()
	// store := db.NewStore(db.GetConnection())

	conf := config.Getenv()

	u := url.URL{Scheme: "ws", Host: fmt.Sprintf("%s:%s", conf.BrokerHost, conf.BrokerPort), Path: "/socket/ws"}

	var wg sync.WaitGroup
	wg.Add(1)

	mb_conn := message_broker.NewConn(u, origin)

	mb_conn.SubscribeTopics(conf.Tickers)

	mb_conn.GoHandleJobs(jobHandler)

	wg.Wait()
}

func jobHandler(message []byte) bool {
	store := db.DbStore

	reverseBinanceAggTradeDto := binance_dto.ReverseBinanceAggTradeDto{}
	err := json.Unmarshal(message, &reverseBinanceAggTradeDto)
	if err != nil {
		logrus.Error(err)
	}

	stockParam := ConvertAggToStockParam(reverseBinanceAggTradeDto)
	utils.PrettyDisplay("params : ", stockParam)
	ctx := context.Background()
	_, err = (*store).CreateStock(ctx, stockParam)
	if err != nil {
		logrus.Errorf("cannot create stock", err)
	}
	fmt.Println("--- GOT MESSAGE : ", string(message))
	return true
}

func ConvertAggToStockParam(aggT binance_dto.ReverseBinanceAggTradeDto) db.CreateStockParams {
	eventTimestamp := time.Unix(aggT.EventTime/1000, (aggT.EventTime%1000)*int64(time.Millisecond))
	tradeTimestamp := time.Unix(aggT.TradeTime/1000, (aggT.TradeTime%1000)*int64(time.Millisecond))
	logrus.Warn("-----", eventTimestamp.GoString())
	return db.CreateStockParams{
		EventType:                  aggT.EventType,
		EventTime:                  eventTimestamp,
		Symbol:                     aggT.Symbol,
		PriceChange:                aggT.PriceChange,
		LastTradeID:                aggT.LastTradeID,
		TotalTradedQuotAssetVolume: aggT.TotalTradedQuotAssetVolume,
		AggregateTradeID:           aggT.AggregateTradeID,
		IsTheBuyerTheMarkerMaker:   aggT.IsTheBuyerTheMarkerMaker,
		IsIgnore:                   aggT.Ignore,
		TradeTime:                  tradeTimestamp,
	}
}
