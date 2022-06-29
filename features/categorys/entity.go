package category

type Core struct {
	ID   int
	Name string
}

type Business interface {
	GetAllCategory() (data []Core, err error)
}

type Data interface {
	SelectAllCategory() (data []Core, err error)
}
