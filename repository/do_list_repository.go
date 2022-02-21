package repository

type DoList struct {
	Items []string
}

type DoListInterface interface {
	CreateNewList() DoList
	AddItem(item string) DoList
}

func CreateNewList() DoList {
	doList := DoList{}
	doList.Items = append(doList.Items, "Achieve Modanisa Bootcamp assignment")
	return doList
}

func (l *DoList) AddItem(item string) DoList {
	l.Items = append(l.Items, item)
	return *l
}
