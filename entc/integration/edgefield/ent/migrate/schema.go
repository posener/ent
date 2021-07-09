// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CarsColumns holds the columns for the "cars" table.
	CarsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "number", Type: field.TypeString, Nullable: true},
	}
	// CarsTable holds the schema information for the "cars" table.
	CarsTable = &schema.Table{
		Name:       "cars",
		Columns:    CarsColumns,
		PrimaryKey: []*schema.Column{CarsColumns[0]},
	}
	// CardsColumns holds the columns for the "cards" table.
	CardsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "number", Type: field.TypeString, Nullable: true},
		{Name: "owner_id", Type: field.TypeInt, Unique: true, Nullable: true, SchemaType: map[string]string{"sqlite3": "integer"}},
	}
	// CardsTable holds the schema information for the "cards" table.
	CardsTable = &schema.Table{
		Name:       "cards",
		Columns:    CardsColumns,
		PrimaryKey: []*schema.Column{CardsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cards_users_card",
				Columns:    []*schema.Column{CardsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "card_number_owner_id",
				Unique:  false,
				Columns: []*schema.Column{CardsColumns[1], CardsColumns[2]},
			},
		},
	}
	// InfosColumns holds the columns for the "infos" table.
	InfosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "content", Type: field.TypeJSON},
	}
	// InfosTable holds the schema information for the "infos" table.
	InfosTable = &schema.Table{
		Name:       "infos",
		Columns:    InfosColumns,
		PrimaryKey: []*schema.Column{InfosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "infos_users_user",
				Columns:    []*schema.Column{InfosColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MetadataColumns holds the columns for the "metadata" table.
	MetadataColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "age", Type: field.TypeInt, Default: 0},
	}
	// MetadataTable holds the schema information for the "metadata" table.
	MetadataTable = &schema.Table{
		Name:       "metadata",
		Columns:    MetadataColumns,
		PrimaryKey: []*schema.Column{MetadataColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "metadata_users_metadata",
				Columns:    []*schema.Column{MetadataColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PetsColumns holds the columns for the "pets" table.
	PetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "owner_id", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"sqlite3": "integer"}},
	}
	// PetsTable holds the schema information for the "pets" table.
	PetsTable = &schema.Table{
		Name:       "pets",
		Columns:    PetsColumns,
		PrimaryKey: []*schema.Column{PetsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "pets_users_pets",
				Columns:    []*schema.Column{PetsColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "text", Type: field.TypeString},
		{Name: "author_id", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"sqlite3": "integer"}},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "posts_users_author",
				Columns:    []*schema.Column{PostsColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "post_author_id_text",
				Unique:  false,
				Columns: []*schema.Column{PostsColumns[2], PostsColumns[1]},
			},
		},
	}
	// RentalsColumns holds the columns for the "rentals" table.
	RentalsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "date", Type: field.TypeTime},
		{Name: "car_id", Type: field.TypeInt, Nullable: true},
		{Name: "user_id", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"sqlite3": "integer"}},
	}
	// RentalsTable holds the schema information for the "rentals" table.
	RentalsTable = &schema.Table{
		Name:       "rentals",
		Columns:    RentalsColumns,
		PrimaryKey: []*schema.Column{RentalsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "rentals_cars_rentals",
				Columns:    []*schema.Column{RentalsColumns[2]},
				RefColumns: []*schema.Column{CarsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "rentals_users_rentals",
				Columns:    []*schema.Column{RentalsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "rental_car_id_user_id",
				Unique:  true,
				Columns: []*schema.Column{RentalsColumns[2], RentalsColumns[3]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, SchemaType: map[string]string{"sqlite3": "integer"}},
		{Name: "parent_id", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"sqlite3": "integer"}},
		{Name: "spouse_id", Type: field.TypeInt, Unique: true, Nullable: true, SchemaType: map[string]string{"sqlite3": "integer"}},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_users_children",
				Columns:    []*schema.Column{UsersColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "users_users_spouse",
				Columns:    []*schema.Column{UsersColumns[2]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CarsTable,
		CardsTable,
		InfosTable,
		MetadataTable,
		PetsTable,
		PostsTable,
		RentalsTable,
		UsersTable,
	}
)

func init() {
	CardsTable.ForeignKeys[0].RefTable = UsersTable
	InfosTable.ForeignKeys[0].RefTable = UsersTable
	MetadataTable.ForeignKeys[0].RefTable = UsersTable
	PetsTable.ForeignKeys[0].RefTable = UsersTable
	PostsTable.ForeignKeys[0].RefTable = UsersTable
	RentalsTable.ForeignKeys[0].RefTable = CarsTable
	RentalsTable.ForeignKeys[1].RefTable = UsersTable
	UsersTable.ForeignKeys[0].RefTable = UsersTable
	UsersTable.ForeignKeys[1].RefTable = UsersTable
}
