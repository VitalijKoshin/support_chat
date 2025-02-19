package database

type MemoryDB struct{}

func NewMemoryDB() *MemoryDB {
	return &MemoryDB{}
}
