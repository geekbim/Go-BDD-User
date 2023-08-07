package main_test

import (
	"context"
	"testing"

	"go-bdd-user/internal/user"
	"go-bdd-user/pkg/passwordhashing"

	"github.com/cucumber/godog"
)

type (
	appContext struct{}
	errContext struct{}
)

var ctx context.Context

func register(username, password string) (context.Context, error) {
	userService := ctx.Value(appContext{}).(*user.UserService)
	ctx = context.WithValue(ctx, errContext{}, userService.Register(username, password))

	return ctx, nil
}

func login(username, password string) (context.Context, error) {
	userService := ctx.Value(appContext{}).(*user.UserService)
	ctx = context.WithValue(ctx, errContext{}, userService.Login(username, password))

	return ctx, nil
}

func check(status string) error {
	if status == "failed" && ctx.Value(errContext{}) != nil {
		return nil
	}
	if ctx.Value(errContext{}) == nil {
		return nil
	}
	return ctx.Value(errContext{}).(error)
}

func initializeScenario(sCtx *godog.ScenarioContext) {
	sCtx.Step(`^I register a new account with username ([\da-zA-Z0-9]+) and password ([\da-zA-Z0-9]+)$`, register)
	sCtx.Step(`^The registration (succeeded|failed)$`, check)
	sCtx.Step(`^I log in to the app using username ([\da-zA-Z0-9]+) and password ([\da-zA-Z0-9]+)$`, login)
	sCtx.Step(`^The logging is (succeeded|failed)$`, check)
}

func initializeSuite(sCtx *godog.TestSuiteContext) {
	sCtx.BeforeSuite(func() {
		userRepository := user.NewMapUserRepository()
		passwordHashingService := passwordhashing.NewBcryptPasswordHashingService()

		ctx = context.WithValue(context.Background(), appContext{}, user.NewUserService(userRepository, passwordHashingService))
	})
}

func TestFeatures(t *testing.T) {
	godog.TestSuite{
		TestSuiteInitializer: initializeSuite,
		ScenarioInitializer:  initializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"../../test/features/app"},
			TestingT: t,
		},
	}.Run()
}
