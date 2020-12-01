package main

import (
	"database/sql"
	"fmt"
	pkgerr "github.com/pkg/errors"
)

func DaoGet(id int) (string, error){
	var retV string
	db := MockDB{}
	strSQL := fmt.Sprintf("select col1 from tmp where id %d", id)
	row, err := db.Query(strSQL)
	if err != nil {
		//将sql执行语句封装添加到原始错误中
		retErr := pkgerr.Wrapf(err,"strSQL: %s", strSQL)
		return retV, retErr
	}
	retV = row.GetColValue("col1")
	return retV, nil
}


// mock DB
type MockDB struct {
}

type MockRow struct {
}

func (MockRow) GetColValue(colName string) string{
	return "mockValue"
}

func (db *MockDB)Query(strSQL string ) (*MockRow, error){
	return nil, sql.ErrNoRows
}
