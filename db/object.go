package db

import (
	"strings"
	"xorm.io/builder"
	"xorm.io/xorm"
)

type Table interface {
	TableName() string
}

type Query struct {
	Session  *xorm.Session
	Table    Table
	Cond     builder.Cond
	Cols     []string
	Order    []string
	Paginate *Paginate
}

type Paginate struct {
	Offset int
	Limit  int
}

func GetQuery() *Query {
	query := new(Query)
	query.Session = GetDB().NewSession()

	return query
}

func (q *Query) SetTable(table Table) *Query {
	q.Table = table
	if nil != q.Table {
		q.Session.Table(q.Table.TableName())
	}
	return q
}

func (q *Query) SetCond(cond builder.Cond) *Query {
	q.Cond = cond
	if nil != q.Cond {
		q.Session.Where(q.Cond)
	}
	return q
}

func (q *Query) SetCols(cols ...string) *Query {
	q.Cols = cols
	if nil != q.Cols {
		q.Session.Select(strings.Join(q.Cols, ","))
	}
	return q
}

func (q *Query) SetOrder(order ...string) *Query {
	q.Order = order
	if nil != q.Order {
		q.Session.OrderBy(strings.Join(q.Order, ","))
	}
	return q
}

func (q *Query) SetPage(page, size int) *Query {
	if 0 < page && 0 < size {
		q.Paginate = new(Paginate)
		q.Paginate.Offset = (page - 1) * size
		q.Paginate.Limit = size
		q.Session.Limit(q.Paginate.Limit, q.Paginate.Offset)
	}
	return q
}
