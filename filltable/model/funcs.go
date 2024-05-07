package model

import (
	"database/sql"
	"errors"
	"kra/data"

	_ "github.com/lib/pq"
)

type TableFunc struct {
	TableName string
	TableStatus string
	Efunc interface{}
}

type FuncsHolder struct {
	db *sql.DB
	data []data.TrackInfo

	funcs []TableFunc
}

func InitFuncHolder(db *sql.DB, data []data.TrackInfo, funcs []TableFunc) FuncsHolder {
	return FuncsHolder{
		db: db,
		data: data,
		funcs: funcs,
	}
}

func (f FuncsHolder) GetOptionDescription() []string {
	res := make([]string, 0, len(f.funcs))
	for _, efunc := range f.funcs {
		res = append(res, efunc.TableName + ", " + efunc.TableStatus)
	}
	return res
}

func (f FuncsHolder) DoFunc(i int) error {
	doFunc := f.funcs[i].Efunc
	switch doFunc := doFunc.(type) {
	case func(*sql.DB) error:
		return doFunc(f.db)

	case func([]data.TrackInfo, *sql.DB) error:
		return doFunc(f.data, f.db)
	}
	return errors.New("not implement func interface")
}


