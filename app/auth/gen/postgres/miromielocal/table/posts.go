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

var Posts = newPostsTable("miromielocal", "posts", "")

type postsTable struct {
	postgres.Table

	// Columns
	ID          postgres.ColumnInteger
	Image       postgres.ColumnString
	Caption     postgres.ColumnString
	ChatEnabled postgres.ColumnBool
	Hashtag     postgres.ColumnString
	Hypertext   postgres.ColumnString
	Hyperlink   postgres.ColumnString
	AddedBy     postgres.ColumnInteger
	Likes       postgres.ColumnInteger
	Deleted     postgres.ColumnBool
	CreatedAt   postgres.ColumnInteger
	UpdatedAt   postgres.ColumnInteger
	DeletedAt   postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type PostsTable struct {
	postsTable

	EXCLUDED postsTable
}

// AS creates new PostsTable with assigned alias
func (a PostsTable) AS(alias string) *PostsTable {
	return newPostsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new PostsTable with assigned schema name
func (a PostsTable) FromSchema(schemaName string) *PostsTable {
	return newPostsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new PostsTable with assigned table prefix
func (a PostsTable) WithPrefix(prefix string) *PostsTable {
	return newPostsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new PostsTable with assigned table suffix
func (a PostsTable) WithSuffix(suffix string) *PostsTable {
	return newPostsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newPostsTable(schemaName, tableName, alias string) *PostsTable {
	return &PostsTable{
		postsTable: newPostsTableImpl(schemaName, tableName, alias),
		EXCLUDED:   newPostsTableImpl("", "excluded", ""),
	}
}

func newPostsTableImpl(schemaName, tableName, alias string) postsTable {
	var (
		IDColumn          = postgres.IntegerColumn("id")
		ImageColumn       = postgres.StringColumn("image")
		CaptionColumn     = postgres.StringColumn("caption")
		ChatEnabledColumn = postgres.BoolColumn("chat_enabled")
		HashtagColumn     = postgres.StringColumn("hashtag")
		HypertextColumn   = postgres.StringColumn("hypertext")
		HyperlinkColumn   = postgres.StringColumn("hyperlink")
		AddedByColumn     = postgres.IntegerColumn("added_by")
		LikesColumn       = postgres.IntegerColumn("likes")
		DeletedColumn     = postgres.BoolColumn("deleted")
		CreatedAtColumn   = postgres.IntegerColumn("created_at")
		UpdatedAtColumn   = postgres.IntegerColumn("updated_at")
		DeletedAtColumn   = postgres.TimestampzColumn("deleted_at")
		allColumns        = postgres.ColumnList{IDColumn, ImageColumn, CaptionColumn, ChatEnabledColumn, HashtagColumn, HypertextColumn, HyperlinkColumn, AddedByColumn, LikesColumn, DeletedColumn, CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn}
		mutableColumns    = postgres.ColumnList{ImageColumn, CaptionColumn, ChatEnabledColumn, HashtagColumn, HypertextColumn, HyperlinkColumn, AddedByColumn, LikesColumn, DeletedColumn, CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn}
	)

	return postsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		Image:       ImageColumn,
		Caption:     CaptionColumn,
		ChatEnabled: ChatEnabledColumn,
		Hashtag:     HashtagColumn,
		Hypertext:   HypertextColumn,
		Hyperlink:   HyperlinkColumn,
		AddedBy:     AddedByColumn,
		Likes:       LikesColumn,
		Deleted:     DeletedColumn,
		CreatedAt:   CreatedAtColumn,
		UpdatedAt:   UpdatedAtColumn,
		DeletedAt:   DeletedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
