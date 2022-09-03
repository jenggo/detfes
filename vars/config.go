package vars

var Config configs

type configs struct {
	Listen  string `yaml:"listen" env:"LISTEN" env-default:":2804"`
	Verbose bool   `yaml:"verbose" env:"VERBOSE" env-default:"false"`

	Detect struct {
		Multiple bool   `yaml:"multiple" env:"DETECT_MULTIPLE" env-default:"false"`
		FormName string `yaml:"form_name" env:"DETECT_FORM_NAME" env-default:"image"`
	} `yaml:"detect"`

	Key struct {
		API string `yaml:"api" env:"KEY_API"`
	} `yaml:"key"`

	Path struct {
		Models string `yaml:"models" env:"PATH_MODELS" env-default:"./models"`
		Temp   string `yaml:"temp" env:"PATH_TEMP" env-default:"./.tmp"`
	} `yaml:"path"`

	Memcached struct {
		Host string `yaml:"host" env:"MEMCACHED_HOST" env-default:"memcached"`
		Port string `yaml:"port" env:"MEMCACHED_PORT" env-default:"11211"`
	} `yaml:"memcached"`
}
