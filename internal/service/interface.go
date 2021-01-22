package service

type IKeys interface {
	GetEncryptedKey(int64) string
	SetEncryptedKey(int64, string) string
}

type IPayload interface {
	GetEncryptedData(string) string
	SetEncryptedData(string, string) string
}
