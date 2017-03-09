package todo

type Todo struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type List map[int]Todo

func (l *List) Add(t string) *List {
	key := getMaxKey(l) + 1
	(*l)[key] = Todo{Title: t, Completed: false}
	return l
}

func (l *List) Remove(id int) *List {
	_, ok := (*l)[id]
	if ok {
		delete(*l, id)
	}

	return l
}

func (l *List) Edit(id int, t string) *List {
	td, ok := (*l)[id]
	if ok {
		td.Title = t
		(*l)[id] = td
	}

	return l
}

func (l *List) ChangeStatus(id int) *List {
	td, ok := (*l)[id]
	if ok {
		td.Completed = !td.Completed
		(*l)[id] = td
	}

	return l
}

// auto_incrementもどき
func getMaxKey(l *List) int {
	maxKey := 0

	for key := range *l {
		if key > maxKey {
			maxKey = key
		}
	}

	return maxKey
}
