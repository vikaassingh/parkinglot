package model

import "time"

type Invoice struct {
	ID          uint
	Number      uint
	Payment     *Payment
	Token       *Token
	Date        time.Time
	GeneratedBY *Operator
}
