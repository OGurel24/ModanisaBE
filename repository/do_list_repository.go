package repository

type doList struct {
	Items []string
}

func CreateNewList() doList {
	doList := doList{}
	doList.Items = append(doList.Items, "Achieve Modanisa Bootcamp assignment")
	return doList
}

func (l *doList) AddItem(item string) doList {
	l.Items = append(l.Items, item)
	return *l
}
