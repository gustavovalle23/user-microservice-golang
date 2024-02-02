package main

import (
	"fmt"

	"github.com/cucumber/godog"
)

type loginContext struct {
	username string
	password string
	err      error
}

func (l *loginContext) aUserWithUsernameAndPassword(username, password string) error {
	l.username = username
	l.password = password
	return nil
}

func (l *loginContext) theUserLogsIn() error {
	if l.username == "john" && l.password == "password" {
		return nil
	}
	return fmt.Errorf("invalid username or password")
}

func (l *loginContext) theUserShouldBeLoggedInSuccessfully() error {
	if l.err != nil {
		return fmt.Errorf("expected successful login, but got an error: %v", l.err)
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	loginCtx := &loginContext{}

	ctx.Given(`^a user with username "([^"]*)" and password "([^"]*)"$`, loginCtx.aUserWithUsernameAndPassword)
	ctx.When(`^the user logs in$`, loginCtx.theUserLogsIn)
	ctx.Then(`^the user should be logged in successfully$`, loginCtx.theUserShouldBeLoggedInSuccessfully)
}
