//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type Notifications struct {
	ID               int32 `sql:"primary_key"`
	UserID           int32
	PostID           int32
	NotificationType string
	Content          string
	SentBy           int32
	CreatedAt        time.Time
}
