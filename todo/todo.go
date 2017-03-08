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

// auto_incrementもどき
func getMaxKey(l *List) int {
	maxKey := 0

	for key, _ := range *l {
		if key > maxKey {
			maxKey = key
		}
	}

	return maxKey
}

func (l *List) Remove(id int) *List {
	_, ok := (*l)[id]
	if ok {
		delete(*l, id)
	}

	return l
}
