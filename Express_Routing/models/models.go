package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type API struct {
	Id int
	A  string
	B  int64
	C  string
	D  time.Time
}

func init() {
	orm.RegisterModel(new(API))
}
