package config

type PostgresConfig interface {
	URI() string
	MigrationDir() string
}
