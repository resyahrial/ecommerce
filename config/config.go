package config

type Configuration struct {
	Db struct {
		User     string `yaml:"user"`
		Pass     string `yaml:"pass"`
		Port     int    `yaml:"port"`
		Host     string `yaml:"host"`
		Name     string `yaml:"db"`
		LogLevel int    `yaml:"loglevel"`

		MaxIdleConns    int `yaml:"maxidlecon"`
		MaxOpenConns    int `yaml:"maxopencon"`
		ConnMaxLifetime int `yaml:"conmaxlifetime"`
	} `yaml:"db"`
	Jwt struct {
		KeyAccess       string `yaml:"keyaccess"`
		KeyRefresh      string `yaml:"keyrefresh"`
		ExpiryAgeAccess struct {
			Value int    `yaml:"value"`
			Unit  string `yaml:"unit"`
		} `yaml:"expiryageaccess"`
		ExpiryAgeRefresh struct {
			Value int    `yaml:"value"`
			Unit  string `yaml:"unit"`
		} `yaml:"expiryagerefresh"`
	} `yaml:"jwt"`
}
