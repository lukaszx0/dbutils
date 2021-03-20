package sqldsl

import "testing"

type BookTable struct {
	Title Field
}

func (t *BookTable) TableName() string {
	return "books"
}

var Book = &BookTable{
	Title: Field{"title", "VARCHAR(400)"},
}

func TestSQLQueries(t *testing.T) {
	tests := []struct {
		name string
		gen string
		sql string
	}{
		{
			name: "select",
			gen: Select(Book.Title).From(Book).String(),
			sql: "SELECT books.title FROM books",
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
