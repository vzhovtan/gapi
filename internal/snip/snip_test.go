package snip_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/vzhovtan/gapi/internal/db"
	"github.com/vzhovtan/gapi/internal/snip"
)

type MockDb struct {
	items []db.Item
}

func (m *MockDb) InsertItem(item db.Item) error {
	m.items = append(m.items, item)
	return nil
}

func (m *MockDb) GetAllItems() ([]db.Item, error) {
	return m.items, nil
}

func TestService_Search(t *testing.T) {
	tests := []struct {
		name      string
		snipToAdd []string
		query     string
		want      []string
	}{
		{
			name:      "given a snippet of 'done' and search of 'ne', getting 'done' back",
			snipToAdd: []string{"done"},
			query:     "ne",
			want:      []string{"done"}},
		{
			name:      "given a snippet of 'Done' and search of 'done', getting 'done' back",
			snipToAdd: []string{"Done"},
			query:     "done",
			want:      []string{"Done"}},
		{
			name:      "given a snippet of ' Done' and search of 'done', getting ' Done' back",
			snipToAdd: []string{" Done"},
			query:     "done",
			want:      []string{" Done"}},
		{
			name:      "given a snippet of ' done ' and search of 'done', getting 'done' back",
			snipToAdd: []string{" done "},
			query:     " done",
			want:      []string{" done "}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockDb{}
			svc := snip.NewService(m)
			for _, toAdd := range tt.snipToAdd {
				err := svc.AddItem(db.Item{Snippet: toAdd})
				if err != nil {
					t.Error(err)
				}
			}
			fmt.Printf("%v \n", m)
			got, err := svc.SearchItem(tt.query)
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_AddNoError(t *testing.T) {
	tests := []struct {
		name      string
		snipToAdd []string
		wantErr   error
	}{
		{
			name:      "given a snippet of 'done', there is no error expected",
			snipToAdd: []string{"done"},
			wantErr:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockDb{}
			svc := snip.NewService(m)
			for _, toAdd := range tt.snipToAdd {
				err := svc.AddItem(db.Item{Snippet: toAdd})
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("Add() returned error %v, expected error is %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestService_AddError(t *testing.T) {
	tests := []struct {
		name      string
		snipToAdd []string
		wantErr   string
	}{
		{
			name:      "given a snippet of 'done, done', there is error expected",
			snipToAdd: []string{"done", "done"},
			wantErr:   "Snippet already exists",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockDb{}
			svc := snip.NewService(m)
			var err error
			for _, toAdd := range tt.snipToAdd {
				err = svc.AddItem(db.Item{Snippet: toAdd})
			}
			if err.Error() != tt.wantErr {
				t.Errorf("Add() returned error = %v, expected error = %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_GetAll(t *testing.T) {
	tests := []struct {
		name      string
		snipToAdd []string
		want      []db.Item
	}{
		{
			name:      "given a snippet of 'one, two, three', there is no error expected",
			snipToAdd: []string{"one", "two", "three"},
			want: []db.Item{{"one"},
				{"two"},
				{"three"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockDb{}
			svc := snip.NewService(m)
			for _, toAdd := range tt.snipToAdd {
				err := svc.AddItem(db.Item{Snippet: toAdd})
				if err != nil {
					t.Error(err)
				}
			}
			fmt.Println(svc.GetAll())
			got, err := svc.GetAll()
			if err != nil {
				t.Error(err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
