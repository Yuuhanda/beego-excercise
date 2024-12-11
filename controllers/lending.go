package controllers

import (
    "encoding/json"
    "strconv"
    "github.com/beego/beego/v2/server/web"
    "myproject/models"
    "myproject/services"
)

type LendingController struct {
    web.Controller
    lendingService *services.LendingService
}

func (c *LendingController) Prepare() {
    c.lendingService = services.NewLendingService()
}

// Create handles POST request to create new lending
func (c *LendingController) Create() {
    var lending models.Lending
    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &lending); err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": "Invalid request body",
        }
        c.ServeJSON()
        return
    }

    if err := c.lendingService.Create(&lending); err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = lending
    c.ServeJSON()
}

// Get handles GET request to fetch lending by ID
func (c *LendingController) Get() {
    idStr := c.Ctx.Input.Param(":id")
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": "Invalid ID format",
        }
        c.ServeJSON()
        return
    }

    lending, err := c.lendingService.GetByID(uint(id))
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = lending
    c.ServeJSON()
}

// List handles GET request to fetch paginated lendings
func (c *LendingController) List() {
    page, _ := c.GetInt("page", 1)
    pageSize, _ := c.GetInt("pageSize", 10)

    lendings, total, err := c.lendingService.List(page, pageSize)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]interface{}{
        "data":  lendings,
        "total": total,
        "page":  page,
    }
    c.ServeJSON()
}

// Update handles PUT request to update lending
func (c *LendingController) Update() {
    var lending models.Lending
    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &lending); err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": "Invalid request body",
        }
        c.ServeJSON()
        return
    }

    if err := c.lendingService.Update(&lending); err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = lending
    c.ServeJSON()
}

// Delete handles DELETE request to remove lending
func (c *LendingController) Delete() {
    idStr := c.Ctx.Input.Param(":id")
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": "Invalid ID format",
        }
        c.ServeJSON()
        return
    }

    if err := c.lendingService.Delete(uint(id)); err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]interface{}{
        "message": "Lending deleted successfully",
    }
    c.ServeJSON()
}

// GetActiveLoans handles GET request to fetch active loans
func (c *LendingController) GetActiveLoans() {
    lendings, err := c.lendingService.GetActiveLoans()
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = lendings
    c.ServeJSON()
}

// GetReturnedLoans handles GET request to fetch returned loans
func (c *LendingController) GetReturnedLoans() {
    lendings, err := c.lendingService.GetReturnedLoans()
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "error": err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = lendings
    c.ServeJSON()
}
