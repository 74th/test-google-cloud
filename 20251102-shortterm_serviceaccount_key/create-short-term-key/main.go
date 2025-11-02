package main

// https://docs.cloud.google.com/iam/docs/create-short-lived-credentials-direct?hl=ja#create-access

import (
	"context"
	"fmt"
	"os"
	"time"

	flag "github.com/spf13/pflag"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/impersonate"
	"google.golang.org/api/option"
)

// サービスアカウントの生成
func getAccessTokenFromImpersonatedCredentials(impersonatedServiceAccount, scope string) (string, error) {
	ctx := context.Background()

	// このプログラムを実行しているGoogle Cloudの認証情報を取得
	credentials, err := google.FindDefaultCredentials(ctx, scope)
	if err != nil {
		return "", fmt.Errorf("failed to generate default credentials: %w", err)
	}

	ts, err := impersonate.CredentialsTokenSource(ctx, impersonate.CredentialsConfig{
		TargetPrincipal: impersonatedServiceAccount,
		Scopes:          []string{scope},
		Lifetime:        300 * time.Second,
		// delegates: The chained list of delegates required to grant the final accessToken.
		// For more information, see:
		// https://cloud.google.com/iam/docs/create-short-lived-credentials-direct#sa-credentials-permissions
		// Delegates is NOT USED here.
		Delegates: []string{},
	}, option.WithCredentials(credentials))
	if err != nil {
		return "", fmt.Errorf("CredentialsTokenSource error: %w", err)
	}

	t, err := ts.Token()
	if err != nil {
		return "", fmt.Errorf("failed to receive token: %w", err)
	}

	return t.AccessToken, nil
}

func main() {
	var (
		flagHelp               bool
		flagSearviceAccount    string
		flagScopes             string
		flagTermDurationMinute int
	)

	flag.BoolVarP(&flagHelp, "help", "p", false, "show help message")
	flag.StringVarP(&flagSearviceAccount, "service-account-email", "s", "", "service account email to impersonate")
	flag.StringVarP(&flagScopes, "scopes", "c", "", "comma-separated list of additional scopes, default: https://www.googleapis.com/auth/cloud-platform ")
	flag.IntVarP(&flagTermDurationMinute, "term", "t", 5, "The duration (in minute) for the short-term credentials. default: 5 minute")
	flag.Parse()

	if flagHelp {
		flag.PrintDefaults()
		return
	}

	if flagSearviceAccount == "" {
		fmt.Println("service-account-email is required")
		flag.PrintDefaults()
		return
	}

	scopes := "https://www.googleapis.com/auth/cloud-platform"
	if flagScopes != "" {
		scopes += "," + flagScopes
	}

	token, err := getAccessTokenFromImpersonatedCredentials(flagSearviceAccount, scopes)
	if err != nil {
		fmt.Printf("Error: %\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Access Token: %s\n", token)
}
