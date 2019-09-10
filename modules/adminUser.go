package modules

type AdminUser struct {
	Username    string  `json:"username"`
	EnableHours float32 `json:"enable_hours"`
}
