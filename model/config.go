package model

type Config struct {
	System systemCfg
	Log    logCfg
	MySQL  mySQLCfg
	Redis  redisCfg
	Server serverCfg
}

type systemCfg struct {
	Model string
}

type logCfg struct {
	Path      string
	Level     int
	AddSource bool
}

type mySQLCfg struct {
	User   string
	Pass   string
	IP     string
	Port   int
	DBName string
}

type redisCfg struct {
	Addr     string
	Password string
	DB       int
}

type serverCfg struct {
	Addr string
}
