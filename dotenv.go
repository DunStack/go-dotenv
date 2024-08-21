package dotenv

var defaultConfig = &Config{}

func Load(names ...string) {
	if len(names) == 0 {
		defaultConfig.Names = []string{".env"}
	} else {
		defaultConfig.Names = names

	}
	defaultConfig.Load()
}

func Get(key string, defaultValue ...string) string {
	return defaultConfig.Get(key, defaultValue...)
}
