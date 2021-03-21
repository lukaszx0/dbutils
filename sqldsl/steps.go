package sqldsl

type Query interface {
	String() string
}

type SelectFromStep interface {
	Query
	From(Table) SelectJoinStep
}

type SelectJoinStep interface {
	Query
	Join(Table) SelectOnStep
	SelectWhereStep
	SelectGroupByStep
	SelectHavingStep
	SelectOrderByStep
	SelectLimitStep
	SelectOffsetStep
}

type SelectOnStep interface {
	Query
	On(...Condition) SelectJoinStep
}

type SelectWhereStep interface {
	Query
	Where(...Condition) SelectGroupByStep
	SelectGroupByStep
	SelectHavingStep
	SelectOrderByStep
	SelectLimitStep
	SelectOffsetStep
}

type SelectGroupByStep interface {
	Query
	GroupBy(...Field) SelectHavingStep
	SelectHavingStep
	SelectOrderByStep
	SelectLimitStep
	SelectOffsetStep
}

type SelectHavingStep interface {
	Query
	Having(...Condition) SelectOrderByStep
	SelectOrderByStep
	SelectLimitStep
	SelectOffsetStep
}

type SelectOrderByStep interface {
	Query
	OrderBy(...Order) SelectLimitStep
	SelectLimitStep
	SelectOffsetStep
}

type SelectLimitStep interface {
	Query
	Limit(int) SelectOffsetStep
	SelectOffsetStep
}

type SelectOffsetStep interface {
	Query
	Offset(int) Query
}
