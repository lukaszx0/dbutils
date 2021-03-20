package sqldsl

import (
	"fmt"
	"strings"
)

type Field struct {
	Name string
	Type string
}

type Selector interface {
	TableName() string
}

type Query interface {
	String() string
}

type SelectFromStep interface {
	From(Selector) Query
}

type selection struct {
	selection Selector
	projection []Field
}

func (s *selection) From(sel Selector) Query {
	s.selection = sel
	return s
}

func (s *selection) String() string {
	var fields []string
	for _, f := range s.projection {
		fields = append(fields, fmt.Sprintf("%s.%s", s.selection.TableName(), f.Name))
	}
	return fmt.Sprintf("SELECT %s FROM %s", strings.Join(fields, ", "), s.selection.TableName())
}

func Select(f ...Field) SelectFromStep {
	return &selection{projection: f}
}
