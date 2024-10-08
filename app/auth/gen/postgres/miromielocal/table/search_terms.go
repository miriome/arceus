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

var SearchTerms = newSearchTermsTable("miromielocal", "search_terms", "")

type searchTermsTable struct {
	postgres.Table

	// Columns
	ID         postgres.ColumnInteger
	BaseTerm   postgres.ColumnString
	MappedTerm postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type SearchTermsTable struct {
	searchTermsTable

	EXCLUDED searchTermsTable
}

// AS creates new SearchTermsTable with assigned alias
func (a SearchTermsTable) AS(alias string) *SearchTermsTable {
	return newSearchTermsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new SearchTermsTable with assigned schema name
func (a SearchTermsTable) FromSchema(schemaName string) *SearchTermsTable {
	return newSearchTermsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new SearchTermsTable with assigned table prefix
func (a SearchTermsTable) WithPrefix(prefix string) *SearchTermsTable {
	return newSearchTermsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new SearchTermsTable with assigned table suffix
func (a SearchTermsTable) WithSuffix(suffix string) *SearchTermsTable {
	return newSearchTermsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newSearchTermsTable(schemaName, tableName, alias string) *SearchTermsTable {
	return &SearchTermsTable{
		searchTermsTable: newSearchTermsTableImpl(schemaName, tableName, alias),
		EXCLUDED:         newSearchTermsTableImpl("", "excluded", ""),
	}
}

func newSearchTermsTableImpl(schemaName, tableName, alias string) searchTermsTable {
	var (
		IDColumn         = postgres.IntegerColumn("id")
		BaseTermColumn   = postgres.StringColumn("base_term")
		MappedTermColumn = postgres.StringColumn("mapped_term")
		allColumns       = postgres.ColumnList{IDColumn, BaseTermColumn, MappedTermColumn}
		mutableColumns   = postgres.ColumnList{BaseTermColumn, MappedTermColumn}
	)

	return searchTermsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:         IDColumn,
		BaseTerm:   BaseTermColumn,
		MappedTerm: MappedTermColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
