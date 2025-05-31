package models

type CategoryCreate struct {
	Name string
}

type CategoryRetrieve struct {
	ID   int
	Name string
}

type CategoryUpdate struct {
	Name string
}
