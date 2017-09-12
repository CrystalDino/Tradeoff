package models

import (
	"Tradeoff/core"
	"log"
)

//TradeConfig support multipul configures at one trade type, by different time cycle
type TradeConfig struct {
	Id           int64
	SettleId     int64   `xorm:"notnull defualt(0)"`
	TradeId      int64   `xorm:"notnull defualt(0)"`
	TradeName    string  `xorm:"char(128) notnull defualt ''"`
	OpenTime     int64   `xorm:"notnull defualt(0)"` //second-day
	CloseTime    int64   `xorm:"notnull defualt(0)"`
	MinBuy       float64 `xorm:"decimal(20,8) notnull default(0.0)"`
	MaxBuy       float64 `xorm:"decimal(20,8) notnull default(0.0)"`
	MinSale      float64 `xorm:"decimal(20,8) notnull default(0.0)"`
	MaxSale      float64 `xorm:"decimal(20,8) notnull default(0.0)"`
	Brate        float64 `xorm:"decimal(10,5) notnull default(0.0)"` //buy rate
	Srate        float64 `xorm:"decimal(10,5) notnull default(0.0)"` //sale rate
	PercentCycle int64   `xorm:"notnull defualt(0)"`                 //hour
	PercentLimit float64 `xorm:"decimal(10,5) notnull default(0.0)"`
	FirstPrice   float64 `xorm:"decimal(20,8) notnull default(0.0)"`
	CTime        int64   `xorm:"created notnull"`
	MTime        int64   `xorm:"updated notnull"`
	Valid        bool
}

func init() {
	if err := core.SyncModels(&TradeConfig{}); err != nil {
		log.Fatalln("sync model TradeConfig failed,", err)
		return
	}
}

func (tc *TradeConfig) TableName() string {
	return "configure"
}
