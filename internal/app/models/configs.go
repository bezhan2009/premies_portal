package models

import "time"

type Configs struct {
	LogParams       LogParams       `json:"log_params"`
	AppParams       AppParams       `json:"app_params"`
	ServerParams    ServerParams    `json:"server_params"`
	AppLogicParams  AppLogicParams  `json:"app_logic_params"`
	PostgresParams  PostgresParams  `json:"postgres_params"`
	RedisParams     RedisParams     `json:"redis_params"`
	ProvidersParams ProvidersConfig `json:"providers"`
	Clients         ClientsConfig   `json:"clients"`
	Cors            Cors            `json:"cors"`
	AuthParams      AuthParams      `json:"auth_params"`
}

type LogParams struct {
	LogDirectory     string `json:"log_directory"`
	LogInfo          string `json:"log_info"`
	LogError         string `json:"log_error"`
	LogWarn          string `json:"log_warn"`
	LogDebug         string `json:"log_debug"`
	MaxSizeMegabytes int    `json:"max_size_megabytes"`
	MaxBackups       int    `json:"max_backups"`
	MaxAge           int    `json:"max_age"`
	Compress         bool   `json:"compress"`
	LocalTime        bool   `json:"local_time"`
}

type AppParams struct {
	GinMode    string `json:"gin_mode"`
	ServerURL  string `json:"server_url"`
	ServerName string `json:"server_name"`
	PortRun    string `json:"port_run"`
}

type ServerParams struct {
	Addr         string `json:"addr"`
	MaxHeaderMBs int    `json:"max_header_mbs"`
	ReadTimeout  int    `json:"read_timeout"`
	WriteTimeout int    `json:"write_timeout"`
}

type PostgresParams struct {
	User     string `json:"user"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	SSLMode  string `json:"sslmode"`
}

type RedisParams struct {
	Host       string `json:"host"`
	Port       int    `json:"port"`
	Password   string `json:"password"`
	DB         int    `json:"db"`
	TTLMinutes int    `json:"ttl_minutes"`
}

type AuthParams struct {
	JwtTtlMinutes int `json:"jwt_ttl_minutes"`
	JwtTtlHours   int `json:"jwt_ttl_hours"`
}

type GoogleProvider struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Redirect     string `json:"redirect"`
}

type ProvidersConfig struct {
	GoogleProvider GoogleProvider `json:"google_provider"`
}

type Client struct {
	ClientAddress string        `json:"address"`
	Timeout       time.Duration `json:"timeout"`
	RetriesCount  int           `json:"retries_count"`
	Insecure      bool          `json:"insecure"`
}

type ClientsConfig struct {
	Premies Client `json:"premies"`
}

type Cors struct {
	AllowOrigins     []string `json:"allow_origins"`
	AllowMethods     []string `json:"allow_methods"`
	AllowHeaders     []string `json:"allow_headers"`
	ExposeHeaders    []string `json:"expose_headers"`
	AllowCredentials bool     `json:"allow_credentials"`
}

type PaginationParams struct {
	Limit int `json:"limit"`
}

type AppLogicParams struct {
	PaginationParams PaginationParams `json:"pagination_params"`
}
