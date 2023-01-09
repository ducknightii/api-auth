package credential_storage

type Storage interface {
	GetPassword(appID string) (secret string, err error)
}
