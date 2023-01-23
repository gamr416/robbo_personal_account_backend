package integration_tests

import (
	"context"
	"github.com/skinnykaen/robbo_student_personal_account.git/app"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	app, cleanerContainer := app.TestApp()
	ctx := context.Background()
	app.Start(ctx)
	code := m.Run()
	app.Stop(ctx)
	cleanerContainer()
	os.Exit(code)
}
