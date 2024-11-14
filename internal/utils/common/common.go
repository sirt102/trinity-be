package common

import (
	"trinity-be/global"

	"github.com/google/uuid"
)

type RegisterSubscriptionStatus string

const (
	StatusRegistered   RegisterSubscriptionStatus = "registered"
	StatusUnregistered RegisterSubscriptionStatus = "unregistered"
	StatusExpired      RegisterSubscriptionStatus = "expired"
	StatusCancelled    RegisterSubscriptionStatus = "cancelled"
	StatusFailed       RegisterSubscriptionStatus = "failed"
	StatusPending      RegisterSubscriptionStatus = "pending"
	StatusError        RegisterSubscriptionStatus = "error"
	StatusUnknown      RegisterSubscriptionStatus = "unknown"
	StatusProcessing   RegisterSubscriptionStatus = "processing"
	StatusCompleted    RegisterSubscriptionStatus = "completed"
	StatusRefunded     RegisterSubscriptionStatus = "refunded"
	StatusVoided       RegisterSubscriptionStatus = "voided"
)

func StringToUUID(str string) uuid.UUID {
	uuid, err := uuid.Parse(str)
	if err != nil {
		global.LogError(err, "")
	}

	return uuid
}
