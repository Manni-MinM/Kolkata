package database

type Database interface {
    Set(key string, val string) (string, error)
    Get(key string) (string, error)
    Delete(key string) (string, error)
}

func New(database string) (Database, error) {
    switch database {
    case "redis":
        return NewRedis()
    default:
        return nil, &ErrUnimplementedDatabase{database}
    }
}

