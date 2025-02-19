package database

type Database struct {
	Redis  *RedisDB
	Mysql  *MySQLDB
	Memory *MemoryDB
}

func NewDatabase(redis *RedisDB, mysql *MySQLDB, memory *MemoryDB) *Database {
	return &Database{
		Redis:  redis,
		Mysql:  mysql,
		Memory: memory,
	}
}
