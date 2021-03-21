package sqldsl

type OrderDirection int

const (
	OrderASC OrderDirection = iota
	OrderDESC
)

var orders = map[OrderDirection]string{
	OrderASC:  "ASC",
	OrderDESC: "DESC",
}

type Order struct {
	Field     Field
	Direction OrderDirection
}

type Field interface {
	TableName() string
	Name() string
}

type FieldBinding struct {
	Field Field
	Value interface{}
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

type StringField struct {
	table  string
	name   string
	dbtype string
}

func (f StringField) Name() string {
	return f.name
}

func (f StringField) TableName() string {
	return f.table
}

func (f StringField) Eq(v string) Condition {
	return Condition{Predicate: EqPredicate, FieldBinding: FieldBinding{f, v}}
}

func (f StringField) IsEq(v StringField) Condition {
	return Condition{Predicate: EqPredicate, FieldBinding: FieldBinding{f, v}}
}

func (f StringField) ASC() Order {
	return Order{f, OrderASC}
}

func (f StringField) DESC() Order {
	return Order{f, OrderDESC}
}

type IntField struct {
	table  string
	name   string
	dbtype string
}

func (f IntField) Name() string {
	return f.name
}

func (f IntField) TableName() string {
	return f.table
}

func (f IntField) Eq(v int) Condition {
	return Condition{Predicate: EqPredicate, FieldBinding: FieldBinding{f, v}}
}

func (f IntField) IsEq(v IntField) Condition {
	return Condition{Predicate: EqPredicate, FieldBinding: FieldBinding{f, v}}
}

func (f IntField) ASC() Order {
	return Order{f, OrderASC}
}

func (f IntField) DESC() Order {
	return Order{f, OrderDESC}
}
