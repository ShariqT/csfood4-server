package core

type ModelResource struct {
	name string
}

func (m ModelResource) String() string {
	return m.name
}

var (
	ProductResouce = ModelResource{"product"}
	UserResource   = ModelResource{"user"}
	MarketResource = ModelResource{"markets"}
)

type DB interface {
	QuerySingle(query string) (Modelable, error)
	QueryMany(query string) ([]Modelable, error)
	ListAll(resource ModelResource) ([]Modelable, error)
	AuthUser(username string, password string) (*User, error)
	GetById(mr ModelResource, uuid string) (Modelable, error)
}

type Modelable interface {
	GenerateID() string
}
