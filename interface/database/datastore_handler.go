package database

// IDataStore has basic cred
type IDataStore interface {
	Create(interface{}) (string, error)
	Update(interface{}) (string, error)
	Read(interface{}) (string, error)
	Close()
}
