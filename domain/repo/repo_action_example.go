package repo

type daodummy struct{}

// There are ways to create the repo
// Both does not change the function receiver

// ---------==========----------
// type 1:
// Create the basic repo generic struct that takes T
// that should be a DAO

// type basicRepo[T any] struct {
// 	db int
// }

// // alias the repo to match the requirement
// type dummyRepo basicRepo[daodummy]

// ---------==========----------

// ---------==========----------
// type 2:
// Just create a single repo struct
type dummyRepo struct {
	db int
}

// ---------==========----------

func newDummyRepo(db int) *dummyRepo {
	return &dummyRepo{
		db: db,
	}
}

// implement the functions
func (dr *dummyRepo) GetAll() (dao []daodummy, err error) {
	return
}

func (dr *dummyRepo) GetById(id string) (dao daodummy, err error) {
	return
}

func (dr *dummyRepo) Insert(dao daodummy) (err error) {
	return
}

// use the repo, e.g. in usecases
// so to speak
type GetAllDummyUsecase struct {
	repoAction GetAllAction[daodummy]
}

type GetByIdDummyUsecase struct {
	repoAction GetByIdAction[daodummy]
}

type InsertDummyUsecase struct {
	repoAction InsertAction[daodummy]
}

func NewGetAllDummyUsecase(repoAction GetAllAction[daodummy]) *GetAllDummyUsecase {
	return &GetAllDummyUsecase{repoAction}
}

func (uc GetAllDummyUsecase) execute() {
	uc.repoAction.GetAll()
}

func NewGetByIdDummyUsecase(repoAction GetByIdAction[daodummy]) *GetByIdDummyUsecase {
	return &GetByIdDummyUsecase{repoAction}
}

func NewInsertDummyUsecase(repoAction InsertAction[daodummy]) *InsertDummyUsecase {
	return &InsertDummyUsecase{repoAction}
}

func example() {
	repo := newDummyRepo(10)
	_ = NewGetAllDummyUsecase(repo)
	_ = NewGetByIdDummyUsecase(repo)
	// This yeets error if the repo does not implement insert method
	_ = NewInsertDummyUsecase(repo)
}
