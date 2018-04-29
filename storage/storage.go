package storage

import (
	"github.com/jakhog/cbs-cisco-agent/cisco/sms"
	"github.com/satori/go.uuid"
)

type Storage struct {
	Incoming []sms.IncomingSMS `json="incoming"`
	Outgoing []sms.OutgoingSMS `json="outgoing"`
}

func StoreIncomingSMS(smses []sms.IncomingSMS) []int {
	storageLock.Lock()
	defer storageLock.Unlock()

	// Messages stored in persisted storage are safe to remove from the Cisco device
	toDelete := make([]int, 0, len(smses))

	// Map the UUIDs of SMSes to those already in storage, or create new ones for new messages
	for _, message := range smses {
		found := false
		// Check if the new message is already in the memory storage
		for i, stored := range inMemoryStorage.Incoming {
			if sms.InEquals(message, stored) {
				if message.Id != stored.Id {
					// Message has moved on Cisco device, reflect in storage
					inMemoryStorage.Incoming[i].Id = message.Id
					inMemoryStorageDirty = true
				}
				found = true
				break
			}
		}
		// If not, assign an UUID and add it to the list
		if !found {
			newUUID, err := uuid.NewV4()
			if err != nil {
				panic(err)
			}
			message.UUID = newUUID
			inMemoryStorage.Incoming = append(inMemoryStorage.Incoming, message)
		}
		// Check to see if the SMS has been persisted, so that it can be deleted
		for _, persisted := range persistedStorage.Incoming {
			if sms.InEquals(message, persisted) {
				toDelete = append(toDelete, message.Id)
				break
			}
		}
	}

	return toDelete
}
