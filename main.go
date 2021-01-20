package main

import (
	"time"

	"github.com/jrallison/go-workers"
)

type LowBalance struct {
	Package
	UUID               string    `json:"uuid"`
	PricePlan          int       `json:"price_plan"`
	ThresholdDirection string    `json:"threshold_direction"`
	Trigger            string    `json:"trigger"`
	ThresholdID        string    `json:"threshold_id"`
	ThresholdLimit     float64   `json:"threshold_limit"`
	EventTime          time.Time `json:"event_time"`
	AccountGroupID     string    `json:"account_group_id"`
	Carrier            string    `json:"carrier"`
	IngestTime         time.Time `json:"ingest_time"`
	EventType          string    `json:"event_type"`
}

type Package struct {
	ExpirationDate time.Time `json:"expiration_date"`
	Duration       int       `json:"duration"`
	Offset         int       `json:"offset"`
}

func main() {
	workers.Configure(map[string]string{
		// location of redis instance
		"server": "localhost:6379",
		// instance of the database
		"database": "0",
		// number of connections to keep open with redis
		"pool": "30",
		// unique process id for this instance of workers (for proper recovery of inprogress jobs on crash)
		"process": "1",
	})

	pkg := Package{
		Duration: 1,
	}

	lbe := LowBalance{
		Carrier:   "tim_brazil",
		UUID:      "sdjfjsndjfknasdjf",
		EventType: "EVENT_TYPE_LOW_BALANCE",
		Package:   pkg,
	}

	workers.Enqueue("low-balance-events", "LowBalanceNotificationProcessorJob", []LowBalance{lbe})
}
