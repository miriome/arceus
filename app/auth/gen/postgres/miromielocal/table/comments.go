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

var Comments = newCommentsTable("miromielocal", "comments", "")

type commentsTable struct {
	postgres.Table

	// Columns
	ID               postgres.ColumnInteger
	UserID           postgres.ColumnInteger
	PostID           postgres.ColumnInteger
	Comment          postgres.ColumnString
	CreatedAt        postgres.ColumnTimestampz
	CreatedTimestamp postgres.ColumnInteger
	DeletedTimestamp postgres.ColumnTimestampz
	IsDeleted        postgres.ColumnBool

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type CommentsTable struct {
	commentsTable

	EXCLUDED commentsTable
}

// AS creates new CommentsTable with assigned alias
func (a CommentsTable) AS(alias string) *CommentsTable {
	return newCommentsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new CommentsTable with assigned schema name
func (a CommentsTable) FromSchema(schemaName string) *CommentsTable {
	return newCommentsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new CommentsTable with assigned table prefix
func (a CommentsTable) WithPrefix(prefix string) *CommentsTable {
	return newCommentsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new CommentsTable with assigned table suffix
func (a CommentsTable) WithSuffix(suffix string) *CommentsTable {
	return newCommentsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newCommentsTable(schemaName, tableName, alias string) *CommentsTable {
	return &CommentsTable{
		commentsTable: newCommentsTableImpl(schemaName, tableName, alias),
		EXCLUDED:      newCommentsTableImpl("", "excluded", ""),
	}
}

func newCommentsTableImpl(schemaName, tableName, alias string) commentsTable {
	var (
		IDColumn               = postgres.IntegerColumn("id")
		UserIDColumn           = postgres.IntegerColumn("user_id")
		PostIDColumn           = postgres.IntegerColumn("post_id")
		CommentColumn          = postgres.StringColumn("comment")
		CreatedAtColumn        = postgres.TimestampzColumn("created_at")
		CreatedTimestampColumn = postgres.IntegerColumn("created_timestamp")
		DeletedTimestampColumn = postgres.TimestampzColumn("deleted_timestamp")
		IsDeletedColumn        = postgres.BoolColumn("is_deleted")
		allColumns             = postgres.ColumnList{IDColumn, UserIDColumn, PostIDColumn, CommentColumn, CreatedAtColumn, CreatedTimestampColumn, DeletedTimestampColumn, IsDeletedColumn}
		mutableColumns         = postgres.ColumnList{UserIDColumn, PostIDColumn, CommentColumn, CreatedAtColumn, CreatedTimestampColumn, DeletedTimestampColumn, IsDeletedColumn}
	)

	return commentsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:               IDColumn,
		UserID:           UserIDColumn,
		PostID:           PostIDColumn,
		Comment:          CommentColumn,
		CreatedAt:        CreatedAtColumn,
		CreatedTimestamp: CreatedTimestampColumn,
		DeletedTimestamp: DeletedTimestampColumn,
		IsDeleted:        IsDeletedColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
