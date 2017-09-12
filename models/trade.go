package models

import (
	"Tradeoff/core"
	"log"

	"github.com/google/uuid"
)

// type Trade struct {
// 	Id     int64
// 	Uid    int64   `xorm:"notnull default(0)"`
// 	Price  float64 `xorm:"decimal(20,8) notnull default(0.0)"`
// 	Amount float64 `xorm:"decimal(20,8) notnull default(0.0)"`
// 	Remain float64 `xorm:"decimal(20,8) notnull default(0.0)"`
// 	Rate   float64 `xorm:"decimal(10,5) notnull default(0.0)"`
// 	Stat   int     `xorm:"tinyint(4) notnull default(0)"` //0无效默认值，1新记录，2未成交 3部分成交 4全部成交 5撤销
// 	Cip    uint32  `xorm:"notnull default(0)"`
// 	CTime  int64   `xorm:"created notnull"`
// 	MTime  int64   `xorm:"updated notnull"`
// 	Tid    string  `xorm:"char(36) index notnull default ''"`
// }

type TradeBuy struct {
	Id     int64
	Uid    int64   `xorm:"notnull default(0)"`
	Price  float64 `xorm:"decimal(20,8) notnull default(0.0)"`
	Amount float64 `xorm:"decimal(20,8) notnull default(0.0)"`
	Remain float64 `xorm:"decimal(20,8) notnull default(0.0)"`
	Rate   float64 `xorm:"decimal(10,5) notnull default(0.0)"`
	Stat   int     `xorm:"tinyint(4) notnull default(0)"` //0无效默认值，1新记录，2未成交 3部分成交 4全部成交 5撤销
	Cip    uint32  `xorm:"notnull default(0)"`
	CTime  int64   `xorm:"created notnull"`
	MTime  int64   `xorm:"updated notnull"`
	Tid    string  `xorm:"char(36) index notnull default ''"`
}

type TradeSale struct {
	Id     int64
	Uid    int64   `xorm:"notnull default(0)"`
	Price  float64 `xorm:"decimal(20,8) notnull default(0.0)"`
	Amount float64 `xorm:"decimal(20,8) notnull default(0.0)"`
	Remain float64 `xorm:"decimal(20,8) notnull default(0.0)"`
	Rate   float64 `xorm:"decimal(10,5) notnull default(0.0)"`
	Stat   int     `xorm:"tinyint(4) notnull default(0)"` //0无效默认值，1新记录，2未成交 3部分成交 4全部成交 5撤销
	Cip    uint32  `xorm:"notnull default(0)"`
	CTime  int64   `xorm:"created notnull"`
	MTime  int64   `xorm:"updated notnull"`
	Tid    string  `xorm:"char(36) index notnull default ''"`
}

func init() {
	if err := core.SyncModels(&TradeBuy{}); err != nil {
		log.Fatalln("sync model TradeBuy failed,", err)
		return
	}
	if err := core.SyncModels(&TradeSale{}); err != nil {
		log.Fatalln("sync model TradeSale failed,", err)
		return
	}
}

func (tb *TradeBuy) TableName() string {
	return "buy"
}

func (ts *TradeSale) TableName() string {
	return "sale"
}

func Show() {
	uu := uuid.New()
	tb := &TradeBuy{Tid: uu.String()}
	core.DBEngine.InsertOne(tb)
	uu = uuid.New()
	ts := &TradeSale{Tid: uu.String()}
	core.DBEngine.InsertOne(ts)
}
