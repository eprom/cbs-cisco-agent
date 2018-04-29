package sms

import (
	"github.com/jakhog/cbs-cisco-agent/cisco/sms"
	"github.com/jakhog/cbs-cisco-agent/log"
	"github.com/jakhog/cbs-cisco-agent/storage"
)

type IncomingTask struct {
	allocated int
	used      int
	stored    int
	deleted   int
}

func NewIncomingTask() *IncomingTask {
	return &IncomingTask{}
}

func (*IncomingTask) Name() string {
	return "Incoming SMS"
}

func (t *IncomingTask) Run(logger *log.Logger) {
	// FIXME: First, delete any messages that we have successfully stored

	// Get current SMS stats from the server
	// That should be cheaper and easier than parsing the actual messages
	status, err := sms.GetSMSStatus()
	if err != nil {
		panic(err)
	}
	// If the stats have changed, we might have new messages to handle
	if t.allocated != status.Allocated || t.used != status.Used || t.stored != status.Stored || t.deleted != status.Deleted {
		// Get the new messages
		smses, err := sms.GetAllSMSes()
		if err != nil {
			panic(err)
		}
		// Store all these messages
		storage.StoreIncomingSMS(smses) // TODO: Delete the ones persisted
		// Save updated stats
		t.allocated = status.Allocated
		t.used = status.Used
		t.stored = status.Stored
		t.deleted = status.Deleted
	}
}
