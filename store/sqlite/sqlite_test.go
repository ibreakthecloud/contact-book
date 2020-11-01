package sqlite

import (
	"github.com/ibreakthecloud/contact-book/store"
	"os"
	"reflect"
	"testing"
)

func TestSQLite_AddContact(t *testing.T) {

	// initialize the store with dummy DB
	store.NewStore = New("temp-db.db")

	defer os.Remove("temp-db.db")

	type args struct {
		name  string
		email string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Add New Contact",
			args:    args{
				name:  "test",
				email: "test@test.io",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := store.NewStore
			if err := s.AddContact(tt.args.name, tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("AddContact() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSQLite_DeleteContact(t *testing.T) {
	store.NewStore = New("temp-db.db")

	defer os.Remove("temp-db.db")

	// Create One entry
	store.NewStore.AddContact("test", "test@test.io")

	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Delete Contact",
			args:    args{
				email: "test@test.io",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := store.NewStore
			if err := s.DeleteContact(tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("DeleteContact() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSQLite_Get(t *testing.T) {

	store.NewStore = New("temp-db.db")

	defer os.Remove("temp-db.db")

	// Create entry
	store.NewStore.AddContact("test", "test@test.io")
	store.NewStore.AddContact("test2", "test2@test.io")

	type args struct {
		name  string
		email string
		page  int
	}
	tests := []struct {
		name    string
		args    args
		want    []store.Result
		wantErr bool
	}{
		{
			name:    "List all",
			args:    args{
				name:  "",
				email: "",
				page:  1,
			},
			want:    []store.Result{
				{Id: 1, Name: "test", Email: "test@test.io"},
				{Id: 2, Name: "test2", Email: "test2@test.io"},
			},
			wantErr: false,
		},
		{
			name:    "List one",
			args:    args{
				name:  "",
				email: "test@test.io",
				page:  1,
			},
			want:    []store.Result{
				{Id: 1, Name: "test", Email: "test@test.io"},
			},
			wantErr: false,
		},
		{
			name:    "List none",
			args:    args{
				name:  "",
				email: "wrong@email.io",
				page:  1,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := store.NewStore
			got, err := s.Get(tt.args.name, tt.args.email, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSQLite_UpdateContact(t *testing.T) {

	store.NewStore = New("temp-db.db")

	defer os.Remove("temp-db.db")

	// Create One entry
	store.NewStore.AddContact("test", "test@test.io")


	type args struct {
		name  string
		email string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Update contact",
			args:    args{
				name:  "test updated",
				email: "test@test.io",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := store.NewStore
			if err := s.UpdateContact(tt.args.name, tt.args.email); (err != nil) != tt.wantErr {
				t.Errorf("UpdateContact() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
