//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type Reports struct {
	ID         uint32 `sql:"primary_key"`
	ReportType string
	TypeID     int32
	Reason     string
	ReportedBy int32
}
