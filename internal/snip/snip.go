package snip

import (
	"errors"
	"strings"

	"gapi/internal/db"
)

type Manager interface {
	InsertItem(item db.Item) error
	GetAllItems() ([]db.Item, error)
}

type Service struct {
	m Manager
}

func NewService(m Manager) *Service {
	return &Service{
		m: m,
	}
}

func (svc *Service) GetAll() ([]db.Item, error) {
	var result []db.Item
	items, err := svc.m.GetAllItems()
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		result = append(result, db.Item{
			Snippet: item.Snippet,
		})
	}
	return result, nil
}

func (svc *Service) AddItem(snip db.Item) error {
	items, err := svc.m.GetAllItems()
	if err != nil {
		return err
	}
	for _, item := range items {
		if item.Snippet == snip.Snippet {
			return errors.New("Snippet already exists")
		}
	}

	if err := svc.m.InsertItem(snip); err != nil {
		return err
	}
	return nil
}

func (svc *Service) SearchItem(query string) ([]string, error) {
	items, err := svc.m.GetAllItems()
	if err != nil {
		return nil, err
	}
	var result []string
	for _, item := range items {
		if strings.Contains(strings.ToLower(item.Snippet), strings.ToLower(query)) {
			result = append(result, item.Snippet)
		}
	}
	return result, nil
}
