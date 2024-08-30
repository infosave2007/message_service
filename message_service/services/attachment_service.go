package services

import (
    "image"
    "image/jpeg"
    "image/png"
    "message_service/models"
    "message_service/config"
    "github.com/nfnt/resize"
    "gorm.io/gorm"
    "mime/multipart"
    "os"
    "path/filepath"
    "time"
    "fmt"
    "bytes"
    "errors"
    "io"
)

// Путь для хранения временных файлов
const tempUploadPath = "temp_uploads/"

// Сохранение файла на сервере с последующим удалением временного файла

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
func SaveFile(file *multipart.FileHeader, cfg *config.Config) (string, error) {
    fileType := file.Header.Get("Content-Type")

    // Проверяем размер документа
    if fileType == "application/pdf" || fileType == "application/vnd.openxmlformats-officedocument.wordprocessingml.document" {
        if file.Size > int64(cfg.MaxDocSize*1024*1024) {
            return "", errors.New("размер документа превышает установленный лимит")
        }
    }

    // Обрабатываем изображение
    if fileType == "image/jpeg" || fileType == "image/png" {
        // Открываем файл
        src, err := file.Open()
        if err != nil {
            return "", fmt.Errorf("не удалось открыть файл: %v", err)
        }
        defer src.Close()

        // Декодируем изображение
        var img image.Image
        if fileType == "image/jpeg" {
            img, err = jpeg.Decode(src)
        } else {
            img, err = png.Decode(src)
        }
        if err != nil {
            return "", fmt.Errorf("не удалось декодировать изображение: %v", err)
        }

        // Изменяем размер изображения
        img = resize.Resize(0, 0, img, resize.Lanczos3)

        // Проверяем размер
        buffer := new(bytes.Buffer)
        if fileType == "image/jpeg" {
            err = jpeg.Encode(buffer, img, nil)
        } else {
            err = png.Encode(buffer, img)
        }
        if err != nil {
            return "", fmt.Errorf("не удалось закодировать изображение: %v", err)
        }

        if buffer.Len() > cfg.MaxImageSize*1024 {
            return "", fmt.Errorf("размер изображения превышает %d KB после изменения размера", cfg.MaxImageSize)
        }

        // Сохраняем временный файл
        tempFilePath := filepath.Join(tempUploadPath, fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename))
        if err := os.MkdirAll(tempUploadPath, os.ModePerm); err != nil {
            return "", fmt.Errorf("не удалось создать директорию: %v", err)
        }

        outFile, err := os.Create(tempFilePath)
        if err != nil {
            return "", fmt.Errorf("не удалось создать временный файл: %v", err)
        }
        defer outFile.Close()

        _, err = buffer.WriteTo(outFile)
        if err != nil {
            return "", fmt.Errorf("не удалось сохранить изображение: %v", err)
        }

        // Чтение временного файла для загрузки в базу данных
        fileData, err := os.ReadFile(tempFilePath)
        if err != nil {
            return "", fmt.Errorf("не удалось прочитать временный файл: %v", err)
        }

        // Удаляем временный файл после использования
        defer os.Remove(tempFilePath)

        return string(fileData), nil
    }

    // Обычное сохранение файла
    tempFilePath := filepath.Join(tempUploadPath, fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename))
    if err := os.MkdirAll(tempUploadPath, os.ModePerm); err != nil {
        return "", fmt.Errorf("не удалось создать директорию: %v", err)
    }

    // Используем функцию для сохранения файла
    if err := SaveMultipartFile(file, tempFilePath); err != nil {
        return "", fmt.Errorf("не удалось сохранить временный файл: %v", err)
    }

    // Чтение временного файла для сохранения его в базу данных
    fileData, err := os.ReadFile(tempFilePath)
    if err != nil {
        return "", fmt.Errorf("не удалось прочитать временный файл: %v", err)
    }

    // Удаляем временный файл после использования
    defer os.Remove(tempFilePath)

    return string(fileData), nil
}

// Функция для сохранения файла из *multipart.FileHeader
func SaveMultipartFile(file *multipart.FileHeader, dst string) error {
    src, err := file.Open()
    if err != nil {
        return err
    }
    defer src.Close()

    out, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer out.Close()

    _, err = io.Copy(out, src)
    return err
}

// Сохранение информации о вложении в базе данных
func SaveAttachment(attachment *models.Attachment, db *gorm.DB) error {
    return db.Create(attachment).Error
}

// Получение вложения по ID
func GetAttachmentByID(id int, db *gorm.DB) (*models.Attachment, error) {
    var attachment models.Attachment
    if err := db.First(&attachment, id).Error; err != nil {
        return nil, err
    }
    return &attachment, nil
}
