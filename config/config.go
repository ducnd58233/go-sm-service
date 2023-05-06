package config

type Config struct {
	DB
	Redis
}

func (c *Config) Load() error {
	err := c.DB.Load()
	if err != nil {
		return err
	}

	// err = c.Redis.Load()
	// if err != nil {
	// 	return err
	// }

	return nil
}
