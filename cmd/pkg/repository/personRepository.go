package repository

import (
	"EuroVote/cmd/models"

	"gorm.io/gorm"
)

type PersonRepository app.App

func NewPersonRepository(db *gorm.DB) *PersonRepository {
	return &PersonRepository{db: db}
}

func (r *PersonRepository) CreatePerson(person *models.Person) error {
	return r.db.Create(person).Error
}

func (r *PersonRepository) GetPerson(id int) (*models.Person, error) {
	var person models.Person
	err := r.db.First(&person, id).Error
	return &person, err
}

func (r *PersonRepository) UpdatePerson(person *models.Person) error {
	return r.db.Save(person).Error
}

func (r *PersonRepository) DeletePerson(id int) error {
	return r.db.Delete(&models.Person{}, id).Error
}
