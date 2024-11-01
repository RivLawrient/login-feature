package users

import "gorm.io/gorm"

type UsersRepository struct {
}

func NewUsersRepository() *UsersRepository {
	return &UsersRepository{}
}

func (r *UsersRepository) Create(db *gorm.DB, user *Users) error {
	return db.Create(user).Error
}
func (r *UsersRepository) Update(db *gorm.DB, user *Users) error {
	return db.Save(user).Error
}
func (r *UsersRepository) UpdateByEmail(db *gorm.DB, user *Users, email string, token string) error {
	// return db.Save(user).Error
	return db.Model(user).Where("email =? ", email).Update("token", token).Error
}

func (r *UsersRepository) FindByEmail(db *gorm.DB, user *Users, email string) error {
	return db.Where("email = ?", email).First(user).Error
}

func (r *UsersRepository) FindyByToken(db *gorm.DB, user *Users, token string) error {
	return db.Where("token = ?", token).First(user).Error
}
