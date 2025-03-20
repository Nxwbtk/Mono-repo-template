package cat

import (
	"log"
	"time"

	"github.com/Nxwbtk/Mono-repo-template/Backend-template/model"
	catschemas "github.com/Nxwbtk/Mono-repo-template/Backend-template/services/Cat/cat-schemas"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CatRepo interface {
	GetCatsRepo() ([]catschemas.TGetCat, error)
	GetCatsByIDRepo(id string) (catschemas.TGetCat, error)
	CreateCatRepo(cat catschemas.TPostCat) (catschemas.TGetCat, error)
	UpdateCatRepo(id string, cat catschemas.TPostCat) (catschemas.TGetCat, error)
	DeleteCatRepo(id string) error
}

type catRepo struct {
	db *gorm.DB
}

func NewCatRepo(db *gorm.DB) CatRepo {
	return &catRepo{
		db: db,
	}
}

func (r *catRepo) GetCatsRepo() ([]catschemas.TGetCat, error) {
	var cats []catschemas.TGetCat = []catschemas.TGetCat{}
	err := r.db.Model(&model.Cat{}).Find(&cats).Error
	if err != nil {
		return nil, err
	}
	return cats, nil
}

func (r *catRepo) CreateCatRepo(cat catschemas.TPostCat) (catschemas.TGetCat, error) {
	var catBody model.Cat = model.Cat{
		ID:        uuid.New().String(),
		Name:      cat.Name,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
	}
	err := r.db.Model(&model.Cat{}).Create(&catBody).Error
	if err != nil {
		log.Fatal(err)
		return catschemas.TGetCat{}, err
	}
	return catschemas.TGetCat{
		ID:   catBody.ID,
		Name: catBody.Name,
	}, nil
}

func (r *catRepo) GetCatsByIDRepo(id string) (catschemas.TGetCat, error) {
	var cat model.Cat
	err := r.db.Model(&model.Cat{}).Where("id = ?", id).Find(&cat).Error
	if err != nil {
		return catschemas.TGetCat{}, err
	}
	return catschemas.TGetCat{
		ID:   cat.ID,
		Name: cat.Name,
	}, nil
}

func (r *catRepo) UpdateCatRepo(id string, cat catschemas.TPostCat) (catschemas.TGetCat, error) {
	var catBody model.Cat
	if err := r.db.Model(&model.Cat{}).Where("id = ?", id).First(&catBody).Error; err != nil {
		return catschemas.TGetCat{}, err
	}

	catBody.Name = cat.Name
	catBody.UpdatedAt = time.Now().String()

	if err := r.db.Save(&catBody).Error; err != nil {
		return catschemas.TGetCat{}, err
	}

	return catschemas.TGetCat{
		ID:   catBody.ID,
		Name: catBody.Name,
	}, nil
}

func (r *catRepo) DeleteCatRepo(id string) error {
	var cat model.Cat
	if err := r.db.Model(&model.Cat{}).Where("id = ?", id).First(&cat).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&cat).Error; err != nil {
		return err
	}

	return nil
}
