package task

type Task struct {
	Id    string `datastore:"-" goon:"id"`
	Name  string
}
