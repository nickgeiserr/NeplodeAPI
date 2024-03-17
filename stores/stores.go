package stores

type Stores struct {
	User UserStore
}

func New() *Stores {
	return &Stores{
		User: &userStore{},
	}
}
