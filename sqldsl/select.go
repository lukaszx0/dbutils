package sqldsl

import (
	"fmt"
	"strings"
)

type Field interface {
	Name() string
}

type Predicate int

const (
	EqPredicate Predicate = iota
	NeqPredicate
	GtPredicate
	LtPredicate
	GePredicate
	LePredicate
	InPredicate
)

var predicates = map[Predicate]string {
	EqPredicate: "=",
	NeqPredicate: "!=",
	GtPredicate: ">",
	LtPredicate: "<",
	GePredicate: ">=",
	LePredicate: "<=",
	InPredicate: "IN",
}

type StringField struct {
	name string
	dbtype string
}

func (f StringField) Name() string {
	return f.name
}

func (f StringField) Eq(v string) Condition {
	return Condition{Predicate: EqPredicate, FieldBinding: FieldBinding{f, v}}
}

type IntField struct {
	name string
	dbtype string
}

func (f IntField) Name() string {
	return f.name
}

func (f IntField) Eq(v int) Condition {
	return Condition{Predicate: EqPredicate, FieldBinding: FieldBinding{f, v}}
}

type Selector interface {
	TableName() string
}

type Query interface {
	String() string
}

type SelectFromStep interface {
	Query
	From(Selector) SelectWhereStep
}

type Condition struct {
	Predicate Predicate
	FieldBinding FieldBinding
}

type FieldBinding struct {
	Field Field
	Value interface{}
}

type selection struct {
	selection Selector
	projection []Field
	predicates  []Condition
}

type SelectWhereStep interface {
	Query
	Where(...Condition) Query
}

func (s *selection) From(sel Selector) SelectWhereStep {
	s.selection = sel
	return s
}

func (s *selection) Where(c ...Condition) Query {
	s.predicates = c
	return s
}

func (s *selection) String() string {
	var fields []string
	for _, f := range s.projection {
		fields = append(fields, fmt.Sprintf("%s.%s", s.selection.TableName(), f.Name()))
	}
	q := fmt.Sprintf("SELECT %s FROM %s", strings.Join(fields, ", "), s.selection.TableName())
	if len(s.predicates) > 0 {
		var w []string
		for _, p := range s.predicates {
			w = append(w, fmt.Sprintf("%s.%s %s %v", s.selection.TableName(), p.FieldBinding.Field.Name(), predicates[p.Predicate], p.FieldBinding.Value))
		}
		q = fmt.Sprintf("%s WHERE %s", q, strings.Join(w, " AND "))
	}
	return q
}

func Select(f ...Field) SelectFromStep {
	return &selection{projection: f}
}
