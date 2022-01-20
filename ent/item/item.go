// Code generated by entc, DO NOT EDIT.

package item

import (
	"time"
)

const (
	// Label holds the string label denoting the item type in the database.
	Label = "item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDesc holds the string denoting the desc field in the database.
	FieldDesc = "desc"
	// FieldTags holds the string denoting the tags field in the database.
	FieldTags = "tags"
	// FieldCount holds the string denoting the count field in the database.
	FieldCount = "count"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeGroup holds the string denoting the group edge name in mutations.
	EdgeGroup = "group"
	// EdgeTag holds the string denoting the tag edge name in mutations.
	EdgeTag = "tag"
	// Table holds the table name of the item in the database.
	Table = "items"
	// GroupTable is the table that holds the group relation/edge. The primary key declared below.
	GroupTable = "group_group_items"
	// GroupInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupInverseTable = "groups"
	// TagTable is the table that holds the tag relation/edge. The primary key declared below.
	TagTable = "tag_tag_items"
	// TagInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagInverseTable = "tags"
)

// Columns holds all SQL columns for item fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldDesc,
	FieldTags,
	FieldCount,
	FieldCreatedAt,
	FieldUpdatedAt,
}

var (
	// GroupPrimaryKey and GroupColumn2 are the table columns denoting the
	// primary key for the group relation (M2M).
	GroupPrimaryKey = []string{"group_id", "item_id"}
	// TagPrimaryKey and TagColumn2 are the table columns denoting the
	// primary key for the tag relation (M2M).
	TagPrimaryKey = []string{"tag_id", "item_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCount holds the default value on creation for the "count" field.
	DefaultCount int
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)
