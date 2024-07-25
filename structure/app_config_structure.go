package structure

type BaseConfig struct {
	ID          int           `json:"id"`
	ConfigValue ConfigValue `json:"config_value"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
type ConfigValue struct {
	Home     struct {
		Icon      string `json:"icon"`
	} `json:"home"`
	User    []BaseConfigUser `json:"user"`
	Slideshow []string     `json:"slideshow"`
}

type BaseConfigUser struct {
	Title string `json:"title"`
	Type string `json:"type"`
	Desc string `json:"desc"`
	Icon string `json:"icon"`
}