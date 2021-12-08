package auth

import (
	"context"
	firebase "firebase.google.com/go/v4"
)

func IsUserLoggedIn(ctx context.Context, idToken string) (bool, error) {
	app, err := firebase.NewApp(ctx, nil)

	if err != nil {
		return false, err
	}

	a, err := app.Auth(ctx)

	if err != nil {
		return false, err
	}

	_, err = a.VerifyIDToken(ctx, idToken)

	if err != nil {
		return false, nil
	}

	return true, nil
}
