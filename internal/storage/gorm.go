package storage

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type GORMStore struct {
	db *gorm.DB
}

func NewGORMStore(dbPath string) (*GORMStore, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&Contact{}); err != nil {
		return nil, err
	}

	return &GORMStore{db: db}, nil
}

func (g *GORMStore) Add(contact *Contact) error {
	return g.db.Create(contact).Error
}

func (g *GORMStore) GetAll() map[int]*Contact {
	var contacts []Contact
	g.db.Find(&contacts)

	result := make(map[int]*Contact)
	for i := range contacts {
		result[contacts[i].ID] = &contacts[i]
	}
	return result
}

func (g *GORMStore) GetByID(id int) (*Contact, bool) {
	var contact Contact
	err := g.db.First(&contact, id).Error
	if err != nil {
		return nil, false
	}
	return &contact, true
}

func (g *GORMStore) Remove(id int) error {
	return g.db.Delete(&Contact{}, id).Error
}

func (g *GORMStore) Update(id int, name, email string) error {
	updates := make(map[string]interface{})
	if name != "" {
		updates["name"] = name
	}
	if email != "" {
		updates["email"] = email
	}
	return g.db.Model(&Contact{}).Where("id = ?", id).Updates(updates).Error
}

func (g *GORMStore) GetNextID() int {
	var count int64
	g.db.Model(&Contact{}).Count(&count)
	return int(count) + 1
}
