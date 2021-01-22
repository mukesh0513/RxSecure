package sqlData

type ISqlDatabase interface {
	Find(model interface{}, args interface{}) (interface{}, error)
	Create(model interface{}) error
	Delete(model interface{}, args interface{}) error
}
