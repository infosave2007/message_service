package models

import (
    "time"
    "gorm.io/gorm"
    "golang.org/x/crypto/bcrypt"
)

// Модель данных для роли
type Role struct {
    ID   int    `json:"id" gorm:"primaryKey;column:ID"`
    Name string `json:"name" gorm:"column:Name;unique;not null"`
}

// @Summary Краткое описание конечной точки
// @Description Полное описание конечной точки
// @Tags Название тега
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "ID параметр"
// @Success 200 {object} МодельУспеха
// @Failure 400 {object} МодельОшибки
// @Router /some_endpoint [get]

// Здесь начинается функция

// @Summary Краткое описание конечной точки
// @Description Полное описание конечной точки
// @Tags Название тега
// @Accept  json
// @Produce  json
// @Param   id     path    int     true        "ID параметр"
// @Success 200 {object} МодельУспеха
// @Failure 400 {object} МодельОшибки
// @Router /some_endpoint [get]

// Здесь начинается функция
func (Role) TableName() string {
    return "Roles"
}
// Модель данных для пользователя
type User struct {
    ID        int       `json:"id" gorm:"primaryKey;column:ID"`
    Username  string    `json:"username" gorm:"column:Username;unique;not null"`
    Password  string    `json:"password" gorm:"column:Password;not null"`
    Email     string    `json:"email" gorm:"column:Email;unique;not null"`
    RoleID    int       `json:"role_id" gorm:"column:RoleID;not null"`
    Role      Role      `gorm:"foreignKey:RoleID"` // Связь с моделью Role
    CreatedAt time.Time `json:"created_at" gorm:"column:CreatedAt;type:DATETIME"`
    UpdatedAt time.Time `json:"updated_at" gorm:"column:UpdatedAt;type:DATETIME"`
}
func (User) TableName() string {
    return "Users"
}
// Создание новой записи пользователя в базе данных
func (u *User) CreateUser(db *gorm.DB) error {
    // Сохраняем пользователя в базе данных
    return db.Create(u).Error
}

// Получение пользователя по ID
func GetUserByID(db *gorm.DB, id int) (*User, error) {
    var user User
    // Ищем пользователя по ID
    err := db.Preload("Role").First(&user, id).Error // Подгружаем связанную роль
    return &user, err
}

// Получение пользователя по имени пользователя (username)
func GetUserByUsername(db *gorm.DB, username string) (*User, error) {
    var user User
    // Ищем пользователя по имени пользователя
    err := db.Preload("Role").Where("username = ?", username).First(&user).Error // Подгружаем связанную роль
    return &user, err
}

// Обновление данных пользователя
func (u *User) UpdateUser(db *gorm.DB) error {
    // Обновляем данные пользователя
    return db.Save(u).Error
}

// Удаление пользователя
func DeleteUser(db *gorm.DB, id int) error {
    // Удаляем пользователя по ID
    return db.Delete(&User{}, id).Error
}

// ComparePassword проверяет хеш пароля с введенным паролем
func ComparePassword(hashedPassword, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
