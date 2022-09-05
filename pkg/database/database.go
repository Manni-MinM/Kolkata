package database

type Database interface {
    Set(key string, val string) (string, error)
    Get(key string) (string, error)
    Delete(key string) (string, error)
}

