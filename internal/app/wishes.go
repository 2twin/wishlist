package app

import (
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
)

type Wish struct {
	ID        uuid.UUID
	User      string
	Title     string
	Link      string
	Completed bool
	CreatedAt time.Time
}

var (
	ErrWishNotFound = errors.New("wish id not found")
)

func (a *App) AddWish(title, link, user string) {
	wish := Wish{
		ID:        uuid.New(),
		Title:     title,
		Link:      link,
		User:      user,
		Completed: false,
		CreatedAt: time.Now(),
	}

	a.Wishes = append(a.Wishes, wish)
	log.Printf("Added wish to user %s", user)
}

func (a *App) RemoveWish(id uuid.UUID) error {
	i := 0
	for i < len(a.Wishes) {
		if a.Wishes[i].ID == id {
			log.Printf("Removed wish \"%s\" from user %s", a.Wishes[i].Title, a.Wishes[i].User)
			a.Wishes = append(a.Wishes[:i], a.Wishes[i+1:]...)
			return nil
		}
		i++
	}

	log.Printf("<Error> Wish with id %d not found", id)
	return ErrWishNotFound
}

func (a *App) EditWish(id uuid.UUID, title, link string) error {
	idx := -1
	for i, wish := range a.Wishes {
		if wish.ID == id {
			idx = i
			break
		}
	}

	if idx == -1 {
		return ErrWishNotFound
	}

	if title != "" {
		a.Wishes[idx].Title = title
	}

	if link != "" {
		a.Wishes[idx].Link = link
	}

	return nil
}

func (a *App) ToggleWishStatus(id uuid.UUID) error {
	idx := -1
	for i, wish := range a.Wishes {
		if wish.ID == id {
			idx = i
			break
		}
	}

	if idx == -1 {
		return ErrWishNotFound
	}

	a.Wishes[idx].Completed = !a.Wishes[idx].Completed
	return nil
}

func (a *App) GetWishes() []Wish {
	return a.Wishes
}
