package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "time"
    "os"
    "path/filepath"
    "github.com/beego/beego/v2/server/web"
    "myproject/models"
    "myproject/services"
)

type DocUploadedController struct {
    web.Controller
    docUploadedService *services.DocUploadedService
}

func (c *DocUploadedController) Prepare() {
    c.docUploadedService = services.NewDocUploadedService()
}

// Create creates a new document upload record
// @router /api/docs [post]
func (c *DocUploadedController) Create() {
    var request struct {
        FileName string `json:"FileName"`
        UserId   int    `json:"UserId"`  // Changed from uint to int
    }

    if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&request); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid request body",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    doc := &models.DocUploaded{
        FileName: request.FileName,
        UserId:   &models.User{Id: request.UserId},
        Datetime: time.Now(),
    }


    if err := c.docUploadedService.Create(doc); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to create document record",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]interface{}{
        "success": true,
        "message": "Document record created successfully",
        "data":    doc,
    }
    c.ServeJSON()
}

// Get retrieves a document record by ID
// @router /api/docs/:id [get]
func (c *DocUploadedController) Get() {
    idStr := c.Ctx.Input.Param(":id")
    id, _ := strconv.Atoi(idStr)

    doc, err := c.docUploadedService.GetByID(id)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Document record not found",
            "error":   err.Error(),
        }
        c.Ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]interface{}{
        "success": true,
        "data":    doc,
    }
    c.ServeJSON()
}

// List retrieves document records with pagination
// @router /api/docs [get]
func (c *DocUploadedController) List() {
    page, _ := strconv.Atoi(c.GetString("page", "1"))
    pageSize, _ := strconv.Atoi(c.GetString("pageSize", "10"))

    filters := make(map[string]string)
    filters["file_name"] = c.GetString("file_name")
    filters["user_id"] = c.GetString("user_id")
    filters["start_date"] = c.GetString("start_date")
    filters["end_date"] = c.GetString("end_date")

    docs, total, err := c.docUploadedService.List(page, pageSize, filters)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to retrieve documents",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]interface{}{
        "success": true,
        "data": map[string]interface{}{
            "docs":  docs,
            "total": total,
            "page":  page,
            "size":  pageSize,
        },
    }
    c.ServeJSON()
}

// Delete deletes a document record
// @router /api/docs/:id [delete]
func (c *DocUploadedController) Delete() {
    idStr := c.Ctx.Input.Param(":id")
    id, _ := strconv.Atoi(idStr)

    if err := c.docUploadedService.Delete(id); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to delete document record",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]interface{}{
        "success": true,
        "message": "Document record deleted successfully",
    }
    c.ServeJSON()
}

// DownloadTemplate downloads the unit template file
// @router /api/docs/template/download [get]
func (c *DocUploadedController) DownloadTemplate() {
    templatePath := "static/doc-template/add-unit-template.xlsx"
    
    c.Ctx.Output.Download(templatePath, "add-unit-template.xlsx")
}

// Upload handles file upload with timestamp-based filename
// @router /api/docs/upload [post]
func (c *DocUploadedController) Upload() {
    file, header, err := c.GetFile("file")
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to get file",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }
    defer file.Close()

    userId, _ := strconv.Atoi(c.GetString("user_id"))
    
    // Generate new filename with timestamp
    fileExt := filepath.Ext(header.Filename)
    timestamp := time.Now()
    newFileName := "BULK_UPLOAD-" + timestamp.Format("02-01-2006-150405") + fileExt
    
    // Create upload directory if it doesn't exist
    uploadDir := "static/uploads/documents"
    if err := os.MkdirAll(uploadDir, 0755); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to create upload directory",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    // Save file with new name
    filePath := filepath.Join(uploadDir, newFileName)
    if err := c.SaveToFile("file", filePath); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to save file",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    // Create document record with new filename
    doc := &models.DocUploaded{
        FileName: newFileName,
        UserId:   &models.User{Id: userId},
        Datetime: timestamp,
    }

    if err := c.docUploadedService.Create(doc); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to create document record",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    // Before returning the response, create a clean response structure
    cleanDoc := struct {
        IdDoc    int       `json:"IdDoc"`
        FileName string    `json:"FileName"`
        Datetime time.Time `json:"Datetime"`
        UserId   struct {
            Id       int    `json:"Id"`
            Username string `json:"Username"`
            Email    string `json:"Email"`
        } `json:"UserId"`
    }{
        IdDoc:    doc.IdDoc,
        FileName: doc.FileName,
        Datetime: doc.Datetime,
    }
    
    // Copy only the required user fields
    cleanDoc.UserId.Id = doc.UserId.Id
    cleanDoc.UserId.Username = doc.UserId.Username
    cleanDoc.UserId.Email = doc.UserId.Email

    c.Data["json"] = map[string]interface{}{
        "success": true,
        "message": "File uploaded successfully",
        "data": map[string]interface{}{
            "filename": newFileName,
            "doc":     cleanDoc,
        },
    }
    c.ServeJSON()
}