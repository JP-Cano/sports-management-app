package utils

import (
	"github.com/google/uuid"
	"strconv"
	"time"
)

func ParseToInt(ageStr string) int {
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return 0
	}
	return age
}

func ParseDate(dateStr string) time.Time {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}
	}
	return date
}

func ParseUUID(uuidStr string) *uuid.UUID {
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		return nil
	}
	return &id
}
