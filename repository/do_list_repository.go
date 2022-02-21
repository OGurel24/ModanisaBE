package repository

type DoList struct {
	Items []string
}

type DoListInterface interface {
	CreateNewList() DoList
	AddItem(item string) DoList
}

func CreateNewList() *DoList {
	doList := DoList{}
	return &doList
}

func (l *DoList) AddItem(item string) {
	l.Items = append(l.Items, item)
}
