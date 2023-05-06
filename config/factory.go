package config

func New() (*Config, error) {
	c := &Config{}
	err := c.Load()
	if err != nil {
		return nil, err
	}

	return c, nil
}
