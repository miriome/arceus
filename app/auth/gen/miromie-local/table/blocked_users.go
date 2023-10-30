//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/mysql"
)

var BlockedUsers = newBlockedUsersTable("miromie-local", "blocked_users", "")

type blockedUsersTable struct {
	mysql.Table

	// Columns
	ID        mysql.ColumnInteger
	BlockedBy mysql.ColumnInteger
	UserID    mysql.ColumnInteger

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type BlockedUsersTable struct {
	blockedUsersTable

	NEW blockedUsersTable
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
		NEW:               newBlockedUsersTableImpl("", "new", ""),
	}
}

func newBlockedUsersTableImpl(schemaName, tableName, alias string) blockedUsersTable {
	var (
		IDColumn        = mysql.IntegerColumn("id")
		BlockedByColumn = mysql.IntegerColumn("blocked_by")
		UserIDColumn    = mysql.IntegerColumn("user_id")
		allColumns      = mysql.ColumnList{IDColumn, BlockedByColumn, UserIDColumn}
		mutableColumns  = mysql.ColumnList{BlockedByColumn, UserIDColumn}
	)

	return blockedUsersTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		BlockedBy: BlockedByColumn,
		UserID:    UserIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
