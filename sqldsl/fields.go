package sqldsl

type Field interface {
	TableName() string
	Name() string
}

type FieldBinding struct {
	Field Field
	Value interface{}
}

type StringField struct {
	table string
	name string
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

type IntField struct {
	table string
	name string
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
