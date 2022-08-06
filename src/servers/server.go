package servers

type DataBaseServer interface{
	AddItem(id,quantity int64, name,description string, price, vat float64) error
	UpdateItem(id int64, quantity int64) error
	DeleteItem(id int64) (error)
	// ListAll() ([]entities.Entities, error)
	// SelectItem(id int64) (entities.Entities, error)
}