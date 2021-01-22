package redis

type PayloadService struct{}

func (service PayloadService) GetEncryptedData(string) string {
	return ""
}

func (service PayloadService) SetEncryptedData(string, string) string {
	return ""
}

func NewPayloadService() PayloadService {
	return PayloadService{}
}