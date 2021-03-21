package sqldsl


type SelectFromStep interface {
	Query
	From(Table) SelectJoinStep
}

type SelectWhereStep interface {
	Query
	Where(...Condition) Query
}

type SelectJoinStep interface {
	Query
	SelectWhereStep
	Join(Table) SelectOnStep
}

type SelectOnStep interface {
	Query
	On(...Condition) SelectJoinStep
}

