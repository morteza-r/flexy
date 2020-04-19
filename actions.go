package flexy

func (q *Query) Add() (err error) {
	q.QType = "add"
	err = q.CallApi()
	if err != nil {
		return
	}

	return
}

func (q *Query) Get() (err error) {
	q.QType = "get"
	err = q.CallApi()
	if err != nil {
		return
	}

	return
}

func (q *Query) Exist() (err error) {
	q.QType = "exist"
	err = q.CallApi()
	if err != nil {
		return
	}

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
