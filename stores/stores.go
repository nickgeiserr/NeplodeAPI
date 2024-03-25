package stores

type Stores struct {
	User    UserStore
	Chapter ChapterStore
}

func New() *Stores {
	return &Stores{
		User:    &userStore{},
		Chapter: &chapterStore{},
	}
}
