package sqldsl

import "testing"

type AuthorTable struct {
	ID        IntField
	FirstName StringField
	LastName  StringField
}

func (t *AuthorTable) Name() string {
	return "authors"
}

var Author = &AuthorTable{
	ID:        IntField{"authors", "id", "INT"},
	FirstName: StringField{"authors", "first_name", "VARCHAR(50)"},
	LastName:  StringField{"authors", "last_name", "VARCHAR(50) NOT NULL"},
}

type BookTable struct {
	ID       IntField
	Title    StringField
	AuthorID IntField
}

func (t *BookTable) Name() string {
	return "books"
}

var Book = &BookTable{
	ID:       IntField{"books", "id", "INT"},
	Title:    StringField{"books", "title", "VARCHAR(400)"},
	AuthorID: IntField{"books", "author_id", "INT"},
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
			"select_join",
			Select(Book.Title).From(Book).Join(Author).On(Author.ID.IsEq(Book.AuthorID)).String(),
			"SELECT books.title FROM books JOIN authors ON authors.id = books.author_id",
		},
		{
			"select_where",
			Select(Book.Title).From(Book).Where(Book.ID.Eq(123)).String(),
			"SELECT books.title FROM books WHERE books.id = 123",
		},
		{
			"select_group",
			Select(Book.Title).From(Book).GroupBy(Book.ID).String(),
			"SELECT books.title FROM books GROUP BY books.id",
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
