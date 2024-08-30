package models

import (
    "gorm.io/gorm"
    "time"
    "errors"
)

// Структура модели для хранения вложений
type Attachment struct {
    ID        int       `json:"id" gorm:"primaryKey;column:ID"`
    FileName  string    `json:"file_name" gorm:"column:FileName;not null"`
    FileType  string    `json:"file_type" gorm:"column:FileType;not null"`
    FileData  []byte    `json:"file_data" gorm:"column:FileData;not null"`
    FileSize  int64     `json:"file_size" gorm:"column:FileSize;not null"`
    CreatedAt time.Time `json:"created_at" gorm:"column:CreatedAt"`
    UpdatedAt time.Time `json:"updated_at" gorm:"column:UpdatedAt"`
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
func (Attachment) TableName() string {
    return "Attachments"
}
// Сохранение информации о вложении в базе данных
func SaveAttachment(db *gorm.DB, attachment *Attachment) error {
    return db.Create(attachment).Error
}

// Получение вложения по ID
func GetAttachmentByID(db *gorm.DB, id int) (*Attachment, error) {
    var attachment Attachment
    if err := db.First(&attachment, id).Error; err != nil {
        return nil, err
    }
    return &attachment, nil
}

// Функция для загрузки файла по его ID
func DownloadAttachment(db *gorm.DB, id int) ([]byte, string, error) {
    attachment, err := GetAttachmentByID(db, id)
    if err != nil {
        return nil, "", err
    }

    // Проверка наличия файла
    if len(attachment.FileData) == 0 {
        return nil, "", errors.New("файл не найден")
    }

    return attachment.FileData, attachment.FileName, nil
}
