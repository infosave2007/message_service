package controllers

import (
    "message_service/models"
    "message_service/services"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
    "strconv"
)

// UploadAttachmentResponse структура для ответа при загрузке вложения
type UploadAttachmentResponse struct {
    Message      string `json:"message"`
    AttachmentID uint   `json:"attachment_id"`
    FileName     string `json:"filename"`
    FileType     string `json:"filetype"`
    FileSize     int64  `json:"filesize"`
    UploadTime   string `json:"upload_time"`
}

// UploadAttachment загружает вложение
// @Summary Загрузка вложения
// @Description Этот эндпоинт загружает файл как вложение и сохраняет его в базе данных.
// @Tags Вложения
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Файл для загрузки"
// @Success 200 {object} UploadAttachmentResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /attachment [post]
func UploadAttachment(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)

    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Не удалось получить файл"})
        return
    }

    fileContent, err := file.Open()
    if err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Не удалось открыть файл"})
        return
    }
    defer fileContent.Close()

    fileData := make([]byte, file.Size)
    _, err = fileContent.Read(fileData)
    if err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Не удалось прочитать файл"})
        return
    }

    attachment := models.Attachment{
        FileName: file.Filename,
        FileType: file.Header.Get("Content-Type"),
        FileData: fileData,
        FileSize: file.Size,
    }

    if err := services.SaveAttachment(&attachment, db); err != nil {
        c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Ошибка сохранения вложения"})
        return
    }

    c.JSON(http.StatusOK, UploadAttachmentResponse{
        Message:      "Вложение успешно загружено",
        AttachmentID: uint(attachment.ID), // Преобразуем ID в uint
        FileName:     attachment.FileName,
        FileType:     attachment.FileType,
        FileSize:     attachment.FileSize,
        UploadTime:   attachment.CreatedAt.Format("2006-01-02 15:04:05"), // Форматируем время загрузки
    })
}

// DownloadAttachment скачивает вложение по ID
// @Summary Скачивание вложения
// @Description Этот эндпоинт скачивает вложение по его ID.
// @Tags Вложения
// @Produce application/octet-stream
// @Param id path int true "ID вложения"
// @Success 200 {file} file "Вложение"
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /attachment/{id} [get]
func DownloadAttachment(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)

    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Неверный ID вложения"})
        return
    }

    attachment, err := models.GetAttachmentByID(db, id)
    if err != nil {
        c.JSON(http.StatusNotFound, ErrorResponse{Error: "Вложение не найдено"})
        return
    }

    c.Header("Content-Disposition", "attachment; filename="+attachment.FileName)
    c.Header("Content-Type", attachment.FileType)
    c.Data(http.StatusOK, attachment.FileType, attachment.FileData)
}
