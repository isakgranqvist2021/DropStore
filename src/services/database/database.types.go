package database

type Read interface {
	ReadMany(options ReadOptions)
	ReadOne(optons ReadOptions)
}

type Create interface {
	CreateOne(options CreateOptions)
}

type Update interface {
	UpdateOneByID(options UpdateOptions)
}

type ReadOptions struct {
	Collection string
	Filter     interface{}
}

type CreateOptions struct {
	Collection string
	Payload    interface{}
}

type UpdateOptions struct {
	Collection string
	Payload    interface{}
}
