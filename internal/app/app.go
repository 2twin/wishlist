package app

type App struct {
	Users  []User
	Wishes []Wish
}

func NewApp() *App {
	return &App{}
}
