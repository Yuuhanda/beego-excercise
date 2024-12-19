package services

import (
    "time"
    "github.com/beego/beego/v2/client/orm"
    "myproject/models"
    "github.com/beego/beego/v2/core/logs"
)

type APIRouteService struct {
    ormer orm.Ormer
}

func NewAPIRouteService() *APIRouteService {
    return &APIRouteService{
        ormer: orm.NewOrm(),
    }
}

func (s *APIRouteService) Create(route *models.ApiRoute) error {
    route.CreatedAt = time.Now()
    route.UpdatedAt = time.Now()
    _, err := s.ormer.Insert(route)
    return err
}

func (s *APIRouteService) List() ([]*models.ApiRoute, error) {
    var routes []*models.ApiRoute
    _, err := s.ormer.QueryTable(new(models.ApiRoute)).All(&routes)
    return routes, err
}

func (s *APIRouteService) ScanAndSaveRoutes(routes []map[string]string) error {
    // Get all existing routes from DB
    var existingRoutes []struct {
        ID     int
        Path   string
        Method string
        Controller string
        Action string
        Description string
    }
    _, err := s.ormer.Raw("SELECT id, path, method FROM api_route").QueryRows(&existingRoutes)
    if err != nil {
        return err
    }

    // Track which routes still exist in router.go
    currentRoutes := make(map[string]bool)
    
    // Add new routes
    for _, route := range routes {
        routeKey := route["path"] + "-" + route["method"]
        currentRoutes[routeKey] = true
        
        exists, err := s.routeExists(route["path"], route["method"])
        if err != nil {
            logs.Error("Error checking route existence:", err)
            continue
        }

        if !exists {
            _, err := s.ormer.Raw("INSERT INTO api_route (path, method, controller, action, description, created_at, updated_at) VALUES (?, ?, ?, ?, ?, NOW(), NOW())",
                route["path"], route["method"], route["controller"], route["action"], route["description"]).Exec()
            if err != nil {
                logs.Error("Error saving route:", err)
                continue
            }
            logs.Info("Added new route:", route["method"], route["path"])
        }
    }

    // Delete routes that no longer exist in router.go
    for _, existingRoute := range existingRoutes {
        found := false
        for _, newRoute := range routes {
            if existingRoute.Path == newRoute["path"] && 
               existingRoute.Method == newRoute["method"] && 
               existingRoute.Controller == newRoute["controller"] &&
               existingRoute.Action == newRoute["action"] &&
               existingRoute.Description == newRoute["description"] {
                found = true
                break
            }
        }
        
        if !found {
            _, err := s.ormer.Raw("DELETE FROM api_route WHERE id = ?", existingRoute.ID).Exec()
            if err != nil {
                logs.Error("Error deleting route:", err)
                continue
            }
            logs.Info("Deleted removed route:", existingRoute.Method, existingRoute.Path)
        }
    }

    return nil
}

func (s *APIRouteService) GetRoute(id int) (*models.ApiRoute, error) {
    route := &models.ApiRoute{Id: id}
    err := s.ormer.Read(route)
    return route, err
}

func (s *APIRouteService) routeExists(path, method string) (bool, error) {
    var count int
    err := s.ormer.Raw("SELECT COUNT(*) FROM api_route WHERE path = ? AND method = ?", path, method).QueryRow(&count)
    return count > 0, err
}

func (s *APIRouteService) DeleteRoute(id int) error {
    _, err := s.ormer.Raw("DELETE FROM api_route WHERE id = ?", id).Exec()
    if err != nil {
        logs.Error("Error deleting route:", err)
        return err
    }
    logs.Info("Successfully deleted route with ID:", id)
    return nil
}
