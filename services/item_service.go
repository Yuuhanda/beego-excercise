package services

import (
    "github.com/beego/beego/v2/client/orm"
    "myproject/models"
    "time"
    "fmt"
    "github.com/google/uuid"
    "mime/multipart"
    "path/filepath"
    "os"
    "io"
)

type ItemService struct {
    ormer orm.Ormer
}

func NewItemService() *ItemService {
    return &ItemService{
        ormer: orm.NewOrm(),
    }
}

func (s *ItemService) List(page, pageSize int, filters map[string]interface{}) ([]models.Item, int64, error) {
    var items []models.Item
    o := orm.NewOrm()
    
    qs := o.QueryTable(new(models.Item))
    
    // Apply filters using exact field names from Item model
    if name, ok := filters["name"].(string); ok && name != "" {
        qs = qs.Filter("ItemName__icontains", name)
    }
    if category, ok := filters["category"].(string); ok && category != "" {
        qs = qs.Filter("Category__CategoryName__icontains", category)
    }
    if sku, ok := filters["sku"].(string); ok && sku != "" {
        qs = qs.Filter("SKU__icontains", sku)
    }
    if status, ok := filters["status"].(uint); ok && status > 0 {
        qs = qs.Filter("Status", status)
    }
    
    total, _ := qs.Count()
    _, err := qs.RelatedSel("Category").Limit(pageSize).Offset((page - 1) * pageSize).All(&items)
    
    return items, total, err
}

func (s *ItemService) ListWithCategories(page, pageSize int) ([]models.Item, int64, error) {
    var items []models.Item
    var total int64
    
    offset := (page - 1) * pageSize
    
    // Join items with categories table
    qs := s.ormer.QueryTable("item").RelatedSel("Category")
    
    total, err := qs.Count()
    if err != nil {
        return nil, 0, err
    }
    
    _, err = qs.Limit(pageSize, offset).All(&items)
    if err != nil {
        return nil, 0, err
    }
    
    return items, total, nil
}

func (s *ItemService) GetByID(id uint) (*models.Item, error) {
    var item models.Item
    o := orm.NewOrm()
    
    qs := o.QueryTable(new(models.Item))
    err := qs.Filter("IdItem", id).RelatedSel("Category").One(&item)
    
    return &item, err
}


func (s *ItemService) Create(item *models.Item) error {
    o := orm.NewOrm()
    // Create the item
    _, err := o.Insert(item)
    if err != nil {
        return err
    }
    // Load the complete category data
    o.LoadRelated(item, "Category")
    return nil
}


func (s *ItemService) Update(item *models.Item) error {
    _, err := s.ormer.Update(item)
    return err
}

func (s *ItemService) Delete(id uint) error {
    _, err := s.ormer.Delete(&models.Item{IdItem: id})
    return err
}

func (s *ItemService) SearchDashboard(page, pageSize int, filters map[string]interface{}, warehouseId int) ([]map[string]interface{}, int64, error) {
    o := orm.NewOrm()
    qb, _ := orm.NewQueryBuilder("mysql")

    // Base query
    qb.Select("i.item_name",
        "i.SKU",
        "COUNT(CASE WHEN TRIM(iu.status) = '1' AND iu.condition != 4 AND iu.condition != 5 THEN 1 END) as available",
        "COUNT(CASE WHEN TRIM(iu.status) = '2' THEN 1 END) as in_use",
        "COUNT(CASE WHEN TRIM(iu.status) = '3' THEN 1 END) as in_repair",
        "COUNT(CASE WHEN TRIM(iu.status) = '4' THEN 1 END) as lost",
        "i.id_item",
        "i.imagefile",
        "ic.category_name as category").
        From("item i").
        LeftJoin("item_unit iu ON i.id_item = iu.id_item").
        LeftJoin("item_category ic ON i.id_category = ic.id_category")

    // Add warehouse filter if specified
    if warehouseId > 0 {
        qb.Where("iu.id_wh = ?")
    }

    // Add filters
    if name, ok := filters["item_name"].(string); ok && name != "" {
        qb.And("i.item_name LIKE ?")
    }
    if sku, ok := filters["SKU"].(string); ok && sku != "" {
        qb.And("i.SKU LIKE ?")
    }
    if category, ok := filters["category"].(string); ok && category != "" {
        qb.And("ic.category_name LIKE ?")
    }

    qb.GroupBy("i.id_item").
        Limit(pageSize).
        Offset((page - 1) * pageSize)

    sql := qb.String()
    var maps []orm.Params
    var args []interface{}

    // Add filter values to args
    if warehouseId > 0 {
        args = append(args, warehouseId)
    }
    if name, ok := filters["item_name"].(string); ok && name != "" {
        args = append(args, "%"+name+"%")
    }
    if sku, ok := filters["SKU"].(string); ok && sku != "" {
        args = append(args, "%"+sku+"%")
    }
    if category, ok := filters["category"].(string); ok && category != "" {
        args = append(args, "%"+category+"%")
    }

    num, err := o.Raw(sql, args...).Values(&maps)
    if err != nil {
        return nil, 0, err
    }

    // Convert orm.Params to []map[string]interface{}
    result := make([]map[string]interface{}, len(maps))
    for i, m := range maps {
        result[i] = make(map[string]interface{})
        for k, v := range m {
            result[i][k] = v
        }
    }

    return result, num, nil
}


//auto generate SKU
func (s *ItemService) GenerateSKU(item *models.Item) error {
    o := orm.NewOrm()
    
    category := &models.ItemCategory{IdCategory: item.Category.IdCategory}
    if err := o.Read(category); err != nil {
        return err
    }
    
    year := time.Now().Year() % 100
    
    // Keep trying until we find an unused SKU
    var sku string
    for i := 1; ; i++ {
        sku = fmt.Sprintf("%s%02d-%04d", category.CatCode, year, i)
        
        // Check if SKU exists
        exists := o.QueryTable(new(models.Item)).Filter("SKU", sku).Exist()
        if !exists {
            break
        }
    }
    
    item.SKU = sku
    item.Category = category
    
    return nil
}

func (s *ItemService) UploadImage(file multipart.File, handler *multipart.FileHeader) (string, error) {
    uploadDir := "static/uploads/items"
    if err := os.MkdirAll(uploadDir, 0755); err != nil {
        return "", err
    }

    // Generate unique filename while preserving extension
    filename := uuid.New().String() + filepath.Ext(handler.Filename)
    filepath := filepath.Join(uploadDir, filename)

    dst, err := os.Create(filepath)
    if err != nil {
        return "", err
    }
    defer dst.Close()

    if _, err := io.Copy(dst, file); err != nil {
        return "", err
    }

    return filename, nil
}


