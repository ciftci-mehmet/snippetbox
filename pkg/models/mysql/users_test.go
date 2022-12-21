package mysql

import (
	"reflect"
	"testing"
	"time"

	"github.com/ciftci-mehmet/snippetbox/pkg/models"
)

func TestUserModelGet(t *testing.T) {
	// skip test if `-short` flag
	if testing.Short() {
		t.Skip("mysql: skipping integration test")
	}

	// set up table drivern tests
	tests := []struct {
		name      string
		userID    int
		wantUser  *models.User
		wantError error
	}{
		{
			name:   "Valid ID",
			userID: 1,
			wantUser: &models.User{
				ID:      1,
				Name:    "Alice Jones",
				Email:   "alice@example.com",
				Created: time.Date(2018, 12, 23, 17, 25, 22, 0, time.UTC),
				Active:  true,
			},
			wantError: nil,
		},
		{
			name:      "Zero ID",
			userID:    0,
			wantUser:  nil,
			wantError: models.ErrNoRecord,
		},
		{
			name:      "Non-existent ID",
			userID:    2,
			wantUser:  nil,
			wantError: models.ErrNoRecord,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// init connection and defer teardown
			db, teardown := newTestDB(t)
			defer teardown()

			// create new instance of UserModel
			m := UserModel{db}

			user, err := m.Get(tt.userID)

			if err != tt.wantError {
				t.Errorf("want %v; got %s", tt.wantError, err)
			}

			// reflect.DeepEqual() is an effective way to check for equality between arbitrarily complex custom types.
			if !reflect.DeepEqual(user, tt.wantUser) {
				t.Errorf("want %v; got %v", tt.wantUser, user)
			}
		})
	}
}
