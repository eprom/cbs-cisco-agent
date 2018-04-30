package sms

import (
	"github.com/jakhog/cbs-cisco-agent/cisco/sms/status"
	"github.com/jakhog/cbs-cisco-agent/cisco/sms/table"
	"github.com/satori/go.uuid"
)

type OutgoingSMSStatus int

const ( // NOTE to self, don't change these numbers
	PENDING OutgoingSMSStatus = 0
	SENT    OutgoingSMSStatus = 1
	FAILED  OutgoingSMSStatus = 2
	UNKNOWN OutgoingSMSStatus = 3
)

type OutgoingSMS struct {
	UUID   uuid.UUID         `json:"uuid"`
	To     string            `json:"to"`
	Text   string            `json:"text"`
	Status OutgoingSMSStatus `json:"status"`
}

type IncomingSMS = table.SMS

type Status = status.Status

func InEquals(a, b IncomingSMS) bool {
	return a.From == b.From && a.Received == b.Received && a.Size == b.Size && a.Text == b.Text
}

func OutEquals(a, b OutgoingSMS) bool {
	return uuid.Equal(a.UUID, b.UUID) && a.To == b.To && a.Text == b.Text
}
