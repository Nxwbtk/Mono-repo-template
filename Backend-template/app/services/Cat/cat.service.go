package cat

import catschemas "github.com/Nxwbtk/Mono-repo-template/Backend-template/services/Cat/cat-schemas"

type CatService interface {
	GetCatsService() ([]catschemas.TGetCat, error)
	GetCatsByIDService(id string) (catschemas.TGetCat, error)
	CreateCatService(cat catschemas.TPostCat) (catschemas.TGetCat, error)
	UpdateCatService(id string, cat catschemas.TPostCat) (catschemas.TGetCat, error)
	DeleteCatService(id string) error
}

type catService struct {
	catRepository CatRepo
}

func NewCatService(catRepository CatRepo) CatService {
	return &catService{
		catRepository: catRepository,
	}
}

func (s *catService) GetCatsService() ([]catschemas.TGetCat, error) {
	return s.catRepository.GetCatsRepo()
}

func (s *catService) CreateCatService(cat catschemas.TPostCat) (catschemas.TGetCat, error) {
	return s.catRepository.CreateCatRepo(cat)
}

func (s *catService) GetCatsByIDService(id string) (catschemas.TGetCat, error) {
	return s.catRepository.GetCatsByIDRepo(id)
}

func (s *catService) UpdateCatService(id string, cat catschemas.TPostCat) (catschemas.TGetCat, error) {
	return s.catRepository.UpdateCatRepo(id, cat)
}

func (s *catService) DeleteCatService(id string) error {
	return s.catRepository.DeleteCatRepo(id)
}
