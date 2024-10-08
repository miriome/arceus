//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Follow = newFollowTable("miromielocal", "follow", "")

type followTable struct {
	postgres.Table

	// Columns
	ID       postgres.ColumnInteger
	UserID   postgres.ColumnInteger
	TargetID postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type FollowTable struct {
	followTable

	EXCLUDED followTable
}

// AS creates new FollowTable with assigned alias
func (a FollowTable) AS(alias string) *FollowTable {
	return newFollowTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FollowTable with assigned schema name
func (a FollowTable) FromSchema(schemaName string) *FollowTable {
	return newFollowTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FollowTable with assigned table prefix
func (a FollowTable) WithPrefix(prefix string) *FollowTable {
	return newFollowTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FollowTable with assigned table suffix
func (a FollowTable) WithSuffix(suffix string) *FollowTable {
	return newFollowTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFollowTable(schemaName, tableName, alias string) *FollowTable {
	return &FollowTable{
		followTable: newFollowTableImpl(schemaName, tableName, alias),
		EXCLUDED:    newFollowTableImpl("", "excluded", ""),
	}
}

func newFollowTableImpl(schemaName, tableName, alias string) followTable {
	var (
		IDColumn       = postgres.IntegerColumn("id")
		UserIDColumn   = postgres.IntegerColumn("user_id")
		TargetIDColumn = postgres.IntegerColumn("target_id")
		allColumns     = postgres.ColumnList{IDColumn, UserIDColumn, TargetIDColumn}
		mutableColumns = postgres.ColumnList{UserIDColumn, TargetIDColumn}
	)

	return followTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:       IDColumn,
		UserID:   UserIDColumn,
		TargetID: TargetIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
