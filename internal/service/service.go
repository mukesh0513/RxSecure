package service

import (
	"github.com/mukesh0513/RxSecure/internal/database/keyValueData/redis"
	"github.com/mukesh0513/RxSecure/internal/database/sqlData/gormSupported"
)

func IKeyFactory(dbType string) IKeys {
	switch dbType {
	case "mysql":
		return gormSupported.NewKeyService()
	case "redis":
		return redis.NewKeyService()
	}
	return  nil
}

func IPayloadFactory(dbType string) IPayload {
	switch dbType {
	case "mysql":
		return gormSupported.NewPayloadService()
	case "redis":
		return redis.NewPayloadService()
	}
	return  nil
}
