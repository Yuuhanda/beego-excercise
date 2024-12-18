package services

import (
    "time"
    "github.com/beego/beego/v2/client/orm"
    "myproject/models"
	"strings"
)

type UserVisitLogService struct {
    ormer orm.Ormer
}

func NewUserVisitLogService() *UserVisitLogService {
    return &UserVisitLogService{
        ormer: orm.NewOrm(),
    }
}

func (s *UserVisitLogService) LogVisit(userID int, ip, userAgent, language string) error {
    visitLog := &models.UserVisitLog{
        Token:     generateToken(),
        IP:        ip,
        Language:  language,
        UserAgent: userAgent,
        VisitTime: int(time.Now().Unix()),
        Browser:   detectBrowser(userAgent),
        OS:        detectOS(userAgent),
    }
    
    if userID > 0 {
        visitLog.UserId = &models.User{Id: userID}
    }
    
    _, err := s.ormer.Insert(visitLog)
    return err
}

func (s *UserVisitLogService) GetUserVisits(userID int, page, pageSize int) ([]*models.UserVisitLog, int64, error) {
    var logs []*models.UserVisitLog
    offset := (page - 1) * pageSize
    
    qs := s.ormer.QueryTable(new(models.UserVisitLog)).Filter("user_id", userID)
    
    total, err := qs.Count()
    if err != nil {
        return nil, 0, err
    }
    
    _, err = qs.OrderBy("-visit_time").Offset(offset).Limit(pageSize).All(&logs)
    return logs, total, err
}

func generateToken() string {
    return time.Now().Format("060102150405")
}

func detectBrowser(userAgent string) string {
    userAgent = strings.ToLower(userAgent)
    
    switch {
    case strings.Contains(userAgent, "firefox"):
        return "Firefox"
    case strings.Contains(userAgent, "edge"):
        return "Edge"
    case strings.Contains(userAgent, "chrome"):
        return "Chrome"
    case strings.Contains(userAgent, "safari"):
        return "Safari"
    case strings.Contains(userAgent, "opera"):
        return "Opera"
    default:
        return "Unknown"
    }
}

func detectOS(userAgent string) string {
    userAgent = strings.ToLower(userAgent)
    
    switch {
    case strings.Contains(userAgent, "windows"):
        return "Windows"
    case strings.Contains(userAgent, "mac"):
        return "MacOS"
    case strings.Contains(userAgent, "linux"):
        return "Linux"
    case strings.Contains(userAgent, "android"):
        return "Android"
    case strings.Contains(userAgent, "iphone") || strings.Contains(userAgent, "ipad"):
        return "iOS"
    default:
        return "Unknown"
    }
}