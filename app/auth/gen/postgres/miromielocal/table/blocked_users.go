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

var BlockedUsers = newBlockedUsersTable("miromielocal", "blocked_users", "")

type blockedUsersTable struct {
	postgres.Table

	// Columns
	ID        postgres.ColumnInteger
	BlockedBy postgres.ColumnInteger
	UserID    postgres.ColumnInteger

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type BlockedUsersTable struct {
	blockedUsersTable

	EXCLUDED blockedUsersTable
}

// AS creates new BlockedUsersTable with assigned alias
func (a BlockedUsersTable) AS(alias string) *BlockedUsersTable {
	return newBlockedUsersTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new BlockedUsersTable with assigned schema name
func (a BlockedUsersTable) FromSchema(schemaName string) *BlockedUsersTable {
	return newBlockedUsersTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new BlockedUsersTable with assigned table prefix
func (a BlockedUsersTable) WithPrefix(prefix string) *BlockedUsersTable {
	return newBlockedUsersTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new BlockedUsersTable with assigned table suffix
func (a BlockedUsersTable) WithSuffix(suffix string) *BlockedUsersTable {
	return newBlockedUsersTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newBlockedUsersTable(schemaName, tableName, alias string) *BlockedUsersTable {
	return &BlockedUsersTable{
		blockedUsersTable: newBlockedUsersTableImpl(schemaName, tableName, alias),
		EXCLUDED:          newBlockedUsersTableImpl("", "excluded", ""),
	}
}

func newBlockedUsersTableImpl(schemaName, tableName, alias string) blockedUsersTable {
	var (
		IDColumn        = postgres.IntegerColumn("id")
		BlockedByColumn = postgres.IntegerColumn("blocked_by")
		UserIDColumn    = postgres.IntegerColumn("user_id")
		allColumns      = postgres.ColumnList{IDColumn, BlockedByColumn, UserIDColumn}
		mutableColumns  = postgres.ColumnList{BlockedByColumn, UserIDColumn}
	)

	return blockedUsersTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		BlockedBy: BlockedByColumn,
		UserID:    UserIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
