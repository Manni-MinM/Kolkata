package database

type mockDB struct {
    kvMap    map[string]string
}

func NewMock() Database {
    mp := map[string]string{}
    return &mockDB{kvMap: mp}
}

func (mdb *mockDB) Set(key string, val string) (string, error) {
    mdb.kvMap[key] = val
    return key, nil
}

func (mdb *mockDB) Get(key string) (string, error) {
    if _, prs := mdb.kvMap[key]; !prs {
        return "", &ErrOperation{"get"}
    }

    return mdb.kvMap[key], nil
}

func (mdb *mockDB) Delete(key string) (string, error) {
    if _, prs := mdb.kvMap[key]; !prs {
        return "", &ErrOperation{"delete"}
    }
    delete(mdb.kvMap, key)

    return key, nil
}

