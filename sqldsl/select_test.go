package sqldsl

import "testing"

type BookTable struct {
	ID    IntField
	Title StringField
}

func (t *BookTable) TableName() string {
	return "books"
}

var Book = &BookTable{
	ID: IntField{"id", "INT"},
	Title: StringField{"title", "VARCHAR(400)"},
}

func TestSelect(t *testing.T) {
	tests := []struct {
		name, gen, sql string
	}{
		{
			"select",
			Select(Book.Title).From(Book).String(),
			"SELECT books.title FROM books",
		},
		{
			"select_where",
			Select(Book.Title).From(Book).Where(Book.ID.Eq(123)).String(),
			"SELECT books.title FROM books WHERE books.id = 123",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.gen != test.sql {
				t.Errorf("\nexpected:  %s\ngenerated: %s", test.sql, test.gen)
			}
		})
	}
}
