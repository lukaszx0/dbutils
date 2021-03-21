package sqldsl

import (
	"fmt"
	"strings"
)

type Table interface {
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

var predicates = map[Predicate]string{
	EqPredicate:  "=",
	NeqPredicate: "!=",
	GtPredicate:  ">",
	LtPredicate:  "<",
	GePredicate:  ">=",
	LePredicate:  "<=",
	InPredicate:  "IN",
}

type Selector interface {
	TableName() string
}

type Condition struct {
	Predicate    Predicate
	FieldBinding FieldBinding
}

type Join struct {
	Table      Table
	Conditions []Condition
}

type selection struct {
	table      Table
	projection []Field
	joins      []Join
	predicates []Condition
	grouping   []Field
	having     []Condition
	ordering   []Order
	limit      int
	offset     int
}

func Select(f ...Field) SelectFromStep {
	return &selection{projection: f}
}

func (s *selection) From(t Table) SelectJoinStep {
	s.table = t
	return s
}

func (s *selection) Join(t Table) SelectOnStep {
	s.joins = append(s.joins, Join{Table: t})
	return s
}

func (s *selection) On(c ...Condition) SelectJoinStep {
	var j Join
	j, s.joins = s.joins[len(s.joins)-1], s.joins[:len(s.joins)-1]
	j.Conditions = c
	s.joins = append(s.joins, j)
	return s
}

func (s *selection) Where(c ...Condition) SelectGroupByStep {
	s.predicates = c
	return s
}

func (s *selection) GroupBy(f ...Field) SelectHavingStep {
	s.grouping = f
	return s
}

func (s *selection) Having(c ...Condition) SelectOrderByStep {
	s.having = c
	return s
}

func (s *selection) OrderBy(f ...Order) SelectLimitStep {
	s.ordering = f
	return s
}

func (s *selection) Limit(l int) SelectOffsetStep {
	s.limit = l
	return s
}

func (s *selection) Offset(o int) Query {
	s.offset = o
	return s
}

func (s *selection) String() string {
	var fields []string
	for _, f := range s.projection {
		fields = append(fields, fmt.Sprintf("%s.%s", s.table.Name(), f.Name()))
	}
	q := fmt.Sprintf("SELECT %s FROM %s", strings.Join(fields, ", "), s.table.Name())
	// JOIN
	if len(s.joins) > 0 {
		var joins []string
		for _, j := range s.joins {
			join := fmt.Sprintf("JOIN %s", j.Table.Name())
			if len(j.Conditions) > 0 {
				var conds []string
				for _, c := range j.Conditions {
					var v interface{}
					switch f := c.FieldBinding.Value.(type) {
					case Field:
						v = fmt.Sprintf("%s.%s", f.TableName(), f.Name())
					default:
						v = f
					}
					conds = append(conds, fmt.Sprintf("%s.%s %s %v", c.FieldBinding.Field.TableName(), c.FieldBinding.Field.Name(), predicates[c.Predicate], v))
				}
				join = fmt.Sprintf("%s ON %s", join, strings.Join(conds, " "))
			}
			joins = append(joins, join)
		}
		q = fmt.Sprintf("%s %s", q, strings.Join(joins, " "))
	}
	// WHERE
	if len(s.predicates) > 0 {
		var w []string
		for _, p := range s.predicates {
			w = append(w, fmt.Sprintf("%s.%s %s %v", s.table.Name(), p.FieldBinding.Field.Name(), predicates[p.Predicate], p.FieldBinding.Value))
		}
		q = fmt.Sprintf("%s WHERE %s", q, strings.Join(w, " AND "))
	}
	// GROUP BY
	if len(s.grouping) > 0 {
		var groups []string
		for _, g := range s.grouping {
			groups = append(groups, fmt.Sprintf("%s.%s", g.TableName(), g.Name()))
		}
		q = fmt.Sprintf("%s GROUP BY %s", q, strings.Join(groups, ", "))
	}
	// ORDER BY
	if len(s.ordering) > 0 {
		var order []string
		for _, o := range s.ordering {
			order = append(order, fmt.Sprintf("%s.%s %s", o.Field.TableName(), o.Field.Name(), orders[o.Direction]))
		}
		q = fmt.Sprintf("%s ORDER BY %s", q, strings.Join(order, ", "))
	}
	return q
}
