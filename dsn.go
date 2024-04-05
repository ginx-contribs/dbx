package dbx

import (
	"fmt"
	_ "github.com/glebarez/go-sqlite"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/microsoft/go-mssqldb"

	"net/netip"
)

const (
	Mysql     = "mysql"
	Postgres  = "postgres"
	Sqlite    = "sqlite"
	Sqlserver = "sqlserver"
)

func MysqlDsn(opt Options) string {
	return fmt.Sprintf(`%s:%s@tcp(%s)/%s?%s`, opt.User, opt.Password, opt.Address, opt.Database, opt.Params)
}

func PostgresDsn(opt Options) string {
	addrPort, _ := netip.ParseAddrPort(opt.Address)
	dsn := fmt.Sprintf("host=%s, port=%d user=%s password=%s %s", addrPort.Addr(), addrPort.Port(), opt.User, opt.Password, opt.Params)
	return dsn
}

func SQLiteDsn(opt Options) string {
	dsn := fmt.Sprintf("%s?%s", opt.Database, opt.Params)
	return dsn
}

func SQLServerDsn(opt Options) string {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s %s", opt.User, opt.Password, opt.Address, opt.Database, opt.Params)
	return dsn
}
