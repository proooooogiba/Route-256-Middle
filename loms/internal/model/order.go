package model

import "time"

type Status string

func (s Status) String() string {
	return string(s)
}

const (
	New             Status = "new"
	AwaitingPayment Status = "awaiting_payment"
	Failed          Status = "failed"
	Payed           Status = "payed"
	Cancelled       Status = "cancelled"
)

type Order struct {
	ID     int64  `db:"id"`
	UserID int64  `db:"user_id"`
	Status string `db:"status"`
	Items  []*Item
}

type OrderEvent struct {
	ID              int64     `json:"id"`
	EventType       Status    `json:"event"`
	OperationMoment time.Time `json:"moment"`
	ExtraInfo       string    `json:"extra"`
}
