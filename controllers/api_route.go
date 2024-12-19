package controllers

import (
    "github.com/beego/beego/v2/core/logs"
    "github.com/beego/beego/v2/server/web"
    "myproject/services"
    "fmt"
    "go/ast"
    "go/parser"
    "go/token"
    "strings"
    "strconv"
)

type APIRouteController struct {
    web.Controller
    routeService *services.APIRouteService
}

func init() {
    // Enable console logging
    logs.SetLogger(logs.AdapterConsole)
    // Set log level to debug
    logs.SetLevel(logs.LevelDebug)
}

func (c *APIRouteController) Prepare() {
    c.routeService = services.NewAPIRouteService()
}

func (c *APIRouteController) ListRoutes() {
    routes, err := c.routeService.List()
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to retrieve routes",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "data":    routes,
        }
    }
    c.ServeJSON()
}

func (c *APIRouteController) ScanRoutes() {
    routeData := make([]map[string]string, 0)
    
    // Read and parse router.go file
    fset := token.NewFileSet()
    routerFile := "routers/router.go"
    node, err := parser.ParseFile(fset, routerFile, nil, parser.ParseComments)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to parse router file",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    // Find InitRoutes function
    ast.Inspect(node, func(n ast.Node) bool {
        if funcDecl, ok := n.(*ast.FuncDecl); ok && funcDecl.Name.Name == "InitRoutes" {
            // Analyze function body for web.Router calls
            ast.Inspect(funcDecl.Body, func(n ast.Node) bool {
                if callExpr, ok := n.(*ast.CallExpr); ok {
                    if selExpr, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
                        if ident, ok := selExpr.X.(*ast.Ident); ok && ident.Name == "web" && selExpr.Sel.Name == "Router" {
                            // Extract route information from Router call
                            if len(callExpr.Args) >= 3 {
                                path := extractStringValue(callExpr.Args[0])
                                controller := extractControllerName(callExpr.Args[1])
                                methodAction := extractStringValue(callExpr.Args[2])
                                
                                // Parse method and action from string like "post:Create"
                                parts := strings.Split(methodAction, ":")
                                if len(parts) == 2 {
                                    method := strings.ToUpper(parts[0])
                                    action := parts[1]
                                    
                                    routeData = append(routeData, map[string]string{
                                        "path":        path,
                                        "method":      method,
                                        "controller":  controller,
                                        "action":      action,
                                        "description": fmt.Sprintf("API endpoint for %s %s", method, path),
                                    })
                                }
                            }
                        }
                    }
                }
                return true
            })
        }
        return true
    })

    // Save only new routes to database
    if err := c.routeService.ScanAndSaveRoutes(routeData); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to save routes",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "message": fmt.Sprintf("Successfully scanned and saved %d routes", len(routeData)),
            "data":    routeData,
        }
    }
    c.ServeJSON()
}

// Helper functions to extract values from AST nodes
func extractStringValue(expr ast.Expr) string {
    if lit, ok := expr.(*ast.BasicLit); ok && lit.Kind == token.STRING {
        return strings.Trim(lit.Value, "\"")
    }
    return ""
}

func extractControllerName(expr ast.Expr) string {
    if unary, ok := expr.(*ast.UnaryExpr); ok {
        if comp, ok := unary.X.(*ast.CompositeLit); ok {
            if sel, ok := comp.Type.(*ast.SelectorExpr); ok {
                return sel.Sel.Name
            }
        }
    }
    return ""
} 

func (c *APIRouteController) DeleteRoute() {
    idStr := c.Ctx.Input.Param(":id")
    id, _ := strconv.Atoi(idStr)
    
    if err := c.routeService.DeleteRoute(id); err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Failed to delete route",
            "error":   err.Error(),
        }
    } else {
        c.Data["json"] = map[string]interface{}{
            "success": true,
            "message": "Route deleted successfully",
        }
    }
    c.ServeJSON()
}

//get route by id
func (c *APIRouteController) Get() {
    idStr := c.Ctx.Input.Param(":id")
    id, err := strconv.ParseUint(idStr, 10, 32)
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Invalid ID format",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    service := services.NewAPIRouteService()
    employee, err := service.GetRoute(int(id))
    if err != nil {
        c.Data["json"] = map[string]interface{}{
            "success": false,
            "message": "Route not found",
            "error":   err.Error(),
        }
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]interface{}{
        "success": true,
        "data":    employee,
    }
    c.ServeJSON()
}
