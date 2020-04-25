package db

import (
	"base/utils"
	"fmt"
)

func (q *Query) Insert() error {
	_, err := q.Session.Insert(q.Table)
	if nil != err {
		sqlStr, param := q.Session.LastSQL()
		utils.SqlError(fmt.Sprintf("创建 %s 记录错误: %s", q.Table.TableName(), err.Error()), sqlStr, param)
	}

	return err
}

func (q *Query) Update() error {
	_, err := q.Session.Update(q.Table)
	if nil != err {
		sqlStr, param := q.Session.LastSQL()
		utils.SqlError(fmt.Sprintf("更新 %s 记录错误: %s", q.Table.TableName(), err.Error()), sqlStr, param)
	}

	return err
}

func (q *Query) Read() (bool, error) {
	hasRecord, err := q.Session.Get(q.Table)
	if nil != err {
		sqlStr, param := q.Session.LastSQL()
		utils.SqlError(fmt.Sprintf("读取 %s 记录错误: %s", q.Table.TableName(), err.Error()), sqlStr, param)
	}

	return hasRecord, err
}

func (q *Query) List() (int64, error) {
	count, err := q.Session.FindAndCount(q.Table)
	if nil != err {
		sqlStr, param := q.Session.LastSQL()
		utils.SqlError(fmt.Sprintf("读取 %s 列表错误: %s", q.Table.TableName(), err.Error()), sqlStr, param)
	}

	return count, err
}
