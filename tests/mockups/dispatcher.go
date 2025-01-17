package mockups

import (
	"fmt"

	"github.com/pushbits/server/internal/model"
)

// MockDispatcher is a dispatcher used for testing - it does not need any storage interface
type MockDispatcher struct {
}

func (d *MockDispatcher) RegisterApplication(id uint, name, token, user string) (string, error) {
	return fmt.Sprintf("%d-%s", id, name), nil
}

func (d *MockDispatcher) DeregisterApplication(a *model.Application, u *model.User) error {
	return nil
}

func (d *MockDispatcher) UpdateApplication(a *model.Application) error {
	return nil
}

func (d *MockDispatcher) SendNotification(a *model.Application, n *model.Notification) (id string, err error) {
	return randStr(15), nil
}

func (d *MockDispatcher) DeleteNotification(a *model.Application, n *model.DeleteNotification) error {
	return nil
}
