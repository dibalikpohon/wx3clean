package repo

// DaoT = DAO Type

type GetAllAction[EntityT any] interface {
	GetAll() ([]EntityT, error)
}

type GetByIdAction[EntityT any] interface {
	GetById(id string) (EntityT, error)
}

type InsertAction[EntityT any] interface {
	Insert(dao EntityT) error
}

type UpdateAction[EntityT any] interface {
	Update(id string, dao EntityT) error
}

type DeleteAction[EntityT any] interface {
	Delete(id string) error
}

type BasicAction[DaoT any] interface {
	GetAllAction[DaoT]
	GetByIdAction[DaoT]
	InsertAction[DaoT]
	UpdateAction[DaoT]
	DeleteAction
}
