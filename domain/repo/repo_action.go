package repo

// DaoT = DAO Type

type GetAllAction[DaoT any] interface {
	GetAll() ([]DaoT, error)
}

type GetByIdAction[DaoT any] interface {
	GetById(id string) (DaoT, error)
}

type InsertAction[DaoT any] interface {
	Insert(dao DaoT) error
}

type UpdateAction[DaoT any] interface {
	Update(id string, dao DaoT) error
}

type DeleteAction interface {
	Delete(id string) error
}

type BasicAction[DaoT any] interface {
	GetAllAction[DaoT]
	GetByIdAction[DaoT]
	InsertAction[DaoT]
	UpdateAction[DaoT]
	DeleteAction
}
