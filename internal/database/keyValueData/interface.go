package keyValueData

type IKeyValueDatabase interface {
	FindByKey(table string, key string) interface{}
	Create(table string, model interface{}) interface{}
	Delete()
}
