package test

import (
	"testing"

	"go.uber.org/mock/gomock"
)

func TestDeleteUser(t *testing.T) {
	control := gomock.NewController(t)
	defer control.Finish()

}
