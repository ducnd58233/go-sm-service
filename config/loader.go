package config

type Loader interface {
	Load() error
}
