package pggorm

import "fmt"

type Pgsql struct {
	Host              string
	ReplicaHost       string
	Port              string
	Config            string
	Dbname            string
	Username          string
	Password          string
	MaxIdleCons       int
	MaxOpenCons       int
	MaxLifeTimeMinute int
	LogLevel          int
}

func (p *Pgsql) Dsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Ulaanbaatar%s", p.Host, p.Username, p.Password, p.Dbname, p.Port, p.Config)
}

func (p *Pgsql) DsnReplica() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Ulaanbaatar%s", p.ReplicaHost, p.Username, p.Password, p.Dbname, p.Port, p.Config)
}
