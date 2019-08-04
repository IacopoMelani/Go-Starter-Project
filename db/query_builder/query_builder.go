package builder

import (
	"database/sql"
	"strings"
)

// QueryBuilderInterface -
type QueryBuilderInterface interface {
	Query(tableName string) (*sql.Stmt, error)
}

// Builder -
type Builder struct {
	isGroupBySet bool
	isOrderBySet bool
	isSelectSet  bool
	isWhereSet   bool
	GroupBy      string
	OrderBy      string
	Params       []interface{}
	Select       string
	Where        string
}

// BuildQuery - Si occupa di costrutire la query
func (b *Builder) BuildQuery(tableName string) string {

	var querySQL string

	if b.Select == "" {
		b.Select = " SELECT * "
	}

	querySQL = querySQL + " FROM " + tableName

	if b.Where != "" {
		querySQL = querySQL + b.Where
	}

	if b.GroupBy != "" {
		querySQL = querySQL + b.GroupBy
	}

	if b.OrderBy != "" {
		querySQL = querySQL + b.OrderBy
	}

	return querySQL
}

// SelectField - Costrutisce gli n campi passati in select
func (b *Builder) SelectField(fields ...string) *Builder {

	if !b.isSelectSet {

		b.Select = " SELECT "
		b.isSelectSet = true
	}

	b.Select = b.Select + strings.Join(fields, ",")

	return b
}

// WhereEqual - Costruisce una condizione di where con operatore "="
func (b *Builder) WhereEqual(field string, value interface{}) *Builder {

	if b.isWhereSet {

		b.Where = b.Where + " AND " + field + " = ? "

	} else {

		b.Where = " WHERE " + field + " = ? "
		b.isWhereSet = true
	}

	b.Params = append(b.Params, value)
	return b
}

// WhereNull - Construisce una condizione di where sulla presenza del valore
func (b *Builder) WhereNull(field string, isNull bool) *Builder {

	if b.isWhereSet {

		if isNull {

			b.Where = " WHERE " + field + " IS NULL "

		} else {

			b.Where = " WHERE " + field + " IS NOT NULL "
		}

	} else {

		if isNull {

			b.Where = " AND " + field + " IS NULL "

		} else {

			b.Where = " AND " + field + " IS NOT NULL "
		}

		b.isWhereSet = true
	}

	return b
}

// WhereOperator - Construisce una condizione di where con un operatore specifico, al momento non viene fatto nessun controllo sull'operatore, è compito dell'utilizzatore accertarsi della validità dell'operatore passato
func (b *Builder) WhereOperator(field string, operator string, value interface{}) *Builder {

	if b.isWhereSet {

		b.Where = b.Where + " AND " + field + operator + " ? "

	} else {

		b.Where = " WHERE " + field + operator + " ? "
		b.isWhereSet = true
	}

	b.Params = append(b.Params, value)
	return b
}
