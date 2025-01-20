package userModel

import (
	"crud-golang/src/database"
	"time"
)

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func Create(name, email, password string) User {
	return User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func Save(user User) error {
	con := database.GetConnection()
	r := con.Save(&user)

	if r.Error != nil {
		return r.Error
	}

	return nil
}

type UserPage struct {
	Page         uint    `json:"page"`
	ItemsPerPage uint    `json:"items_per_page"`
	Total        int64   `json:"total"`
	Data         []*User `json:"data"`
}

func FindAllPaginated(page, limit uint) UserPage {
	if limit <= 0 {
		limit = 10
	}

	if limit > 100 {
		limit = 100
	}

	con := database.GetConnection()

	offset := page * limit

	data := make([]*User, 0)
	con.Limit(int(limit)).Offset(int(offset)).Find(&data)

	var count int64
	con.Model(&User{}).Count(&count)

	return UserPage{
		Page:         page,
		ItemsPerPage: limit,
		Total:        count,
		Data:         data,
	}
}

func FindById(id uint) User {
	con := database.GetConnection()

	var user User

	con.Where("id = ?", id).First(&user)

	return user
}

func Delete(id uint) {
	con := database.GetConnection()

	con.Where("id = ?", id).Delete(&User{})
}
