package tests

import (
	"testing"

	"github.com/gustavovalle23/user-microservice-golang/pkg/user/domain"
)

func TestUserFactory(t *testing.T) {

	wantedName := "User1"
	wantedEmail := "email@gmail.com"
	wantedPassword := "password"

	user := domain.UserFactory(wantedName, wantedEmail, wantedPassword)

	t.Log(user)

	if user.GetName() != wantedName {
		t.Errorf("Got %q, wanted %q", user.GetName(), wantedName)
	}

	if user.GetEmail() != wantedEmail {
		t.Errorf("Got %q, wanted %q", user.GetEmail(), wantedEmail)
	}

	if user.IsActive() != true {
		t.Errorf("User is not active by default")
	}

	if user.GetPoins() != 0 {
		t.Errorf("User has not 0 points by default")
	}
}

func TestUserActivateAndDeactivate(t *testing.T) {

	wantedName := "User1"
	wantedEmail := "email@gmail.com"
	wantedPassword := "password"

	user := domain.UserFactory(wantedName, wantedEmail, wantedPassword)

	t.Log(user)

	user.Deactivate()

	if user.IsActive() != false {
		t.Errorf("User should be inactive")
	}

	user.Activate()

	if user.IsActive() != true {
		t.Errorf("User should be active")
	}
}

func TestUserIncreasePoints(t *testing.T) {

	wantedName := "User1"
	wantedEmail := "email@gmail.com"
	wantedPassword := "password"

	user := domain.UserFactory(wantedName, wantedEmail, wantedPassword)
	pointsToIncrease := 10

	t.Log(user)

	user.IncreasePoints(pointsToIncrease)

	if user.GetPoins() != pointsToIncrease {
		t.Errorf("Points should be %d after increase points and receive: %d", pointsToIncrease, user.GetPoins())
	}
}
