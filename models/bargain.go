package models

import (
	"Tradeoff/core"
	"log"
)

type Bargain struct {
	Id     int64
	Buid   int64   `xorm:"notnull defualt(0)"`                 //buyer user id
	Suid   int64   `xorm:"notnull default(0)"`                 //sale user id
	Btid   string  `xorm:"char(36) index notnull default ''"`  //trust buy id
	Stid   string  `xorm:"char(36) index notnull default ''"`  //trust sale id
	Brate  float64 `xorm:"decimal(10,5) notnull default(0.0)"` //buy rate
	Srate  float64 `xorm:"decimal(10,5) notnull default(0.0)"` //sale rate
	Price  float64 `xorm:"decimal(20,8) notnull default(0.0)"`
	Amount float64 `xorm:"decimal(20,8) notnull default(0.0)"`
	Ctime  int64   `xorm:"created"`
	Tp     int8    `xorm:"tinyint(1) notnull default(0)"` //trade type 0:unset 1:sale 2:buy
}

func init() {
	if err := core.SyncModels(&Bargain{}); err != nil {
		log.Fatalln("sync model Bargain failed,", err)
		return
	}
}

func (bg *Bargain) TableName() string {
	return "bargain"
}
