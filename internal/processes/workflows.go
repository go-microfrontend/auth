package processes

import (
	"time"

	"github.com/pkg/errors"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"

	repo "github.com/go-microfrontend/auth/internal/repository"
)

var itemsActivityOptions = workflow.ActivityOptions{
	StartToCloseTimeout: time.Minute,
	RetryPolicy: &temporal.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2.0,
		MaximumInterval:    10 * time.Second,
		MaximumAttempts:    5,
	},
}

var Workflows = []any{CreateUser, GetUser}

func CreateUser(ctx workflow.Context,
	input userInput,
) (string, error) {
	ctx = workflow.WithActivityOptions(ctx, itemsActivityOptions)

	var user repo.User
	err := workflow.ExecuteActivity(ctx, "CreateUser", input).Get(ctx, &user)
	if err != nil {
		return "", errors.Wrap(err, "executing CreateUser activity")
	}

	return user.ID.String(), nil
}

func GetUser(
	ctx workflow.Context,
	input userInput,
) (bool, error) {
	ctx = workflow.WithActivityOptions(ctx, itemsActivityOptions)

	var user repo.User
	err := workflow.ExecuteActivity(ctx, "GetUserByEmail", input.Email).Get(ctx, &user)
	if err != nil {
		return false, errors.Wrap(err, "executing GetUserByEmail activity")
	}

	var check bool
	err = workflow.ExecuteActivity(ctx, "CheckHash",
		hashInput{
			Hash:     user.PasswordHash,
			Password: input.Password,
		}).
		Get(ctx, &check)

	return check, nil
}
