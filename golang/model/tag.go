package model

type Tag struct {
	ID int64 `datastore:"-" goon:"id"`
	Name  string
}
