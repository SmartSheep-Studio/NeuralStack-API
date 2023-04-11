package models

import "time"

const (
	PunishCannotPost = iota
	PunishCannotOperation
	PunishCannotVisit
)

type Punish struct {
	Model

	Reason     string     `json:"reason"`
	Level      int        `json:"level"`
	EffectedAt *time.Time `json:"effected_at"`
	ExpiredAt  *time.Time `json:"expired_at"`
	UserID     uint       `json:"user_id"`
}
