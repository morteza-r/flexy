package flexy

import (
	"reflect"
)

func (q *Query) Add() (err error) {
	q.QType = "add"
	err = q.CallApi()
	if err != nil {
		return
	}

	return
}

func (q *Query) All() (err error) {
	q.QType = "all"
	err = q.CallApi()
	if err != nil {
		return
	}

	return
}

func (q *Query) Get() (err error) {
	s := reflect.Indirect(reflect.ValueOf(q.QModel))
	if s.Kind() == reflect.Slice {
		q.QType = "mget"
	} else {
		q.QType = "get"
	}
	err = q.CallApi()
	if err != nil {
		return
	}

	return
}

func (q *Query) Exists() (err error, ok bool) {
	ok = false
	q.QModel = &ok
	q.QType = "exists"
	err = q.CallApi()
	if err != nil {
		ok = false
		return
	}
	ok = *q.QModel.(*bool)

	return
}

func (q *Query) Update() (err error) {
	q.QType = "update"
	err = q.CallApi()
	if err != nil {
		return
	}

	return
}

func (q *Query) Replace() (err error) {
	q.QType = "replace"
	err = q.CallApi()
	if err != nil {
		return
	}

	return
}

func (q *Query) Delete() (err error) {
	q.QType = "delete"
	err = q.CallApi()
	if err != nil {
		return
	}

	return
}
