package router

type GetRouter interface {
	Get()
}

type ListRouter interface {
	List()
}

type CreateRouter interface {
	Create()
}

type UpdateRouter interface {
	Update()
}

type DeleteRouter interface {
	Delete()
}

type Router interface {
	GetRouter
	ListRouter
	CreateRouter
	UpdateRouter
	DeleteRouter
}

type Registrator interface {
	Register()
}

func RegisterAll(registrators ...Registrator) {
	for _, registrator := range registrators {
		registrator.Register()
	}
}
