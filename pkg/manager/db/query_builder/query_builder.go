package builder

import (
	"database/sql"
	"strings"
)

// QueryBuilderInterface -
type QueryBuilderInterface interface {
	PrepareStmt(tableName string) (*sql.Stmt, error)
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

// orderBy - Build "order by" expression specifying the direction in addition to the field name
func (b *Builder) orderBy(direction string, fields ...string) *Builder {

	if !b.isOrderBySet {

		b.OrderBy = " ORDER BY " + strings.Join(fields, " "+direction+", ") + " " + direction
		b.isOrderBySet = true
	} else {

		b.OrderBy = b.OrderBy + ", " + strings.Join(fields, " "+direction+", ") + " " + direction
	}

	return b
}

// BuildQuery - It takes care of building the query
func (b *Builder) BuildQuery(tableName string) string {

	var querySQL string

	if b.Select == "" {
		b.Select = " SELECT * "
	}

	querySQL = querySQL + b.Select

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

// GroupByField - Builds "group by" expression
// GroupByField - Costruisce una condizione di group by
func (b *Builder) GroupByField(fields ...string) *Builder {

	if !b.isGroupBySet {

		b.GroupBy = " GROUP BY " + strings.Join(fields, ", ")
		b.isGroupBySet = true

	} else {

		b.GroupBy = b.GroupBy + ", " + strings.Join(fields, ", ")
	}

	return b
}

// OrderByAsc - Build "order by" expression ASC
func (b *Builder) OrderByAsc(fields ...string) *Builder {
	return b.orderBy("Asc", fields...)
}

// OrderByDesc - Build "order by" expression DESC
func (b *Builder) OrderByDesc(fields ...string) *Builder {
	return b.orderBy("DESC", fields...)
}

// ResetStmt - Resets stmt fields with initial values
func (b *Builder) ResetStmt() {

	b.isGroupBySet = false
	b.isOrderBySet = false
	b.isSelectSet = false
	b.isWhereSet = false

	b.GroupBy = ""
	b.OrderBy = ""
	b.Select = ""
	b.Where = ""

	b.Params = make([]interface{}, 0)

}

// SelectField - Builds the n-fields passed in "select" expression
func (b *Builder) SelectField(fields ...string) *Builder {

	if !b.isSelectSet {

		b.Select = " SELECT "
		b.isSelectSet = true
	}

	if b.Select != " SELECT " {

		b.Select = b.Select + ", " + strings.Join(fields, ",")

	} else {

		b.Select = b.Select + strings.Join(fields, ",")
	}

	return b
}

// WhereEqual - Builds "where" expression with "=" operator
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

// WhereNull - Builds "where" expression on nullable field, if isNull is TRUE checks for "IS NULL" otherwise "IS NOT NULL"
func (b *Builder) WhereNull(field string, isNull bool) *Builder {

	if b.isWhereSet {

		if isNull {

			b.Where = b.Where + " AND " + field + " IS NULL "

		} else {

			b.Where = b.Where + " AND " + field + " IS NOT NULL "
		}

	} else {

		if isNull {

			b.Where = " WHERE " + field + " IS NULL "

		} else {

			b.Where = " WHERE " + field + " IS NOT NULL "
		}

		b.isWhereSet = true
	}

	return b
}

// WhereOperator - Builds "where" expression with passed operator, no checks if is valid operator
func (b *Builder) WhereOperator(field string, operator string, value interface{}) *Builder {

	if b.isWhereSet {

		b.Where = b.Where + " AND " + field + " " + operator + " " + " ? "

	} else {

		b.Where = " WHERE " + field + " " + operator + " " + " ? "
		b.isWhereSet = true
	}

	b.Params = append(b.Params, value)
	return b
}
