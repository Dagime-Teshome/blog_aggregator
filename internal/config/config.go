package config

type Config struct {
	db_url            string
	current_user_name string
}

const configFileName = ".gatorconfig.json"

func (*Config) SetUser(user string) {
	// set the user to the config
	// write the new value to the json config file
}

func Read() *Config {
	// read values from home director
	// write those value to a config struct
	// return the config struct

	return nil
}

func getConfigFilePath() (string, error) {
	return "", nil
}

func write(cfg Config) error {
	return nil
}
