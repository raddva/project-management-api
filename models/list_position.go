package models

import (
	"github.com/raddva/project-management-api/models/types"

	"github.com/google/uuid"
)

type ListPosition struct {
	InternalID    int64     `json:"internal_id" db:"internal_id" gorm:"primaryKey;autoIncrement"`
	PublicID	  uuid.UUID `json:"public_id" db:"public_id"`
	BoardID 	int64     `json:"board_internal_id" db:"board_internal_id"`
	ListOrder	types.UUIDArray `json:"list_order" db:"list_order" gorm:"type:uuid[]"`
}