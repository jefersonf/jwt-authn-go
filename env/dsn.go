package env

import (
	"fmt"
	"os"
	"strings"
)

type DSNBuilder struct {
	DSN strings.Builder
}

func (b *DSNBuilder) WithParam(key, value string) {
	b.DSN.WriteString(fmt.Sprintf(" %s=%s", key, value))
}

func (b *DSNBuilder) String() string {
	return strings.TrimSpace(b.DSN.String())
}

func buildDSN() (string, error) {
	// TODO check env variables first
	var dsn DSNBuilder
	dsn.WithParam("host", os.Getenv("PG_HOSTNAME"))
	dsn.WithParam("user", os.Getenv("PG_USERNAME"))
	dsn.WithParam("password", os.Getenv("PG_PASSWORD"))
	dsn.WithParam("dbname", os.Getenv("PG_DBNAME"))
	dsn.WithParam("port", os.Getenv("PG_PORT"))
	dsn.WithParam("sslmode", os.Getenv("PG_SSLMODE"))

	return dsn.String(), nil
}
