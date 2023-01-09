package credential_storage

import "errors"

type ConfStorage struct {
	kv map[string]string
}

var NotExist = errors.New("not exist")

func ConfStorageInit(m map[string]string) ConfStorage {

	return ConfStorage{
		kv: m,
	}
}

func (c ConfStorage) GetPassword(appID string) (string, error) {
	v, ok := c.kv[appID]
	if !ok {
		return "", NotExist
	}

	return v, nil
}
