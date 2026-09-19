package account

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dooray-go/dooray-sdk/utils"
)

const getMeSuccessBody = `{
    "header": {
        "isSuccessful": true,
        "resultCode": 0,
        "resultMessage": ""
    },
    "result": {
        "id": "member-1",
        "idProviderType": "sso",
        "idProviderUserId": "provider-user-1",
        "name": "Dooray User",
        "userCode": "dooray.user",
        "externalEmailAddress": "dooray.user@example.com",
        "defaultOrganization": {
            "id": "organization-1"
        },
        "locale": "ko_KR",
        "timezoneName": "Asia/Seoul",
        "englishName": "Dooray User",
        "nativeName": "두레이 사용자",
        "nickname": "Dooray",
        "displayMemberId": "dooray.user"
    }
}`

func TestAccount_GetMe(t *testing.T) {
	const apiKey = "personal-auth-token"

	var gotAuthorization, gotAccept string
	mux := http.NewServeMux()
	mux.HandleFunc("GET /common/v1/members/me", func(w http.ResponseWriter, r *http.Request) {
		gotAuthorization = r.Header.Get("Authorization")
		gotAccept = r.Header.Get("Accept")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(getMeSuccessBody))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	response, err := NewAccount(server.URL).GetMe(apiKey)
	if err != nil {
		t.Fatalf("GetMe returned an error: %v", err)
	}

	if gotAuthorization != "dooray-api "+apiKey {
		t.Errorf("Authorization: want %q, got %q", "dooray-api "+apiKey, gotAuthorization)
	}
	if gotAccept != "application/json" {
		t.Errorf("Accept: want %q, got %q", "application/json", gotAccept)
	}
	if response.RawJSON != getMeSuccessBody {
		t.Error("RawJSON did not match")
	}
	if response.Result.ID != "member-1" {
		t.Errorf("result.id: want %q, got %q", "member-1", response.Result.ID)
	}
	if response.Result.DefaultOrganization.ID != "organization-1" {
		t.Errorf("result.defaultOrganization.id: want %q, got %q", "organization-1", response.Result.DefaultOrganization.ID)
	}
	if response.Result.TimezoneName != "Asia/Seoul" {
		t.Errorf("result.timezoneName: want %q, got %q", "Asia/Seoul", response.Result.TimezoneName)
	}
	if !response.Header.IsSuccessful {
		t.Error("header.isSuccessful: want true, got false")
	}
}

func TestAccount_GetMe_HTTPError(t *testing.T) {
	for _, statusCode := range []int{http.StatusUnauthorized, http.StatusInternalServerError} {
		t.Run(http.StatusText(statusCode), func(t *testing.T) {
			mux := http.NewServeMux()
			mux.HandleFunc("GET /common/v1/members/me", func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(statusCode)
			})
			server := httptest.NewServer(mux)
			defer server.Close()

			_, err := NewAccount(server.URL).GetMe("invalid-token")
			if err == nil {
				t.Fatalf("expected an error for HTTP %d", statusCode)
			}

			var statusErr utils.StatusCodeError
			if !errors.As(err, &statusErr) {
				t.Fatalf("error type: want utils.StatusCodeError, got %T", err)
			}
			if statusErr.HTTPStatusCode() != statusCode {
				t.Errorf("HTTP status: want %d, got %d", statusCode, statusErr.HTTPStatusCode())
			}
		})
	}
}

func TestAccount_GetMe_InvalidJSON(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /common/v1/members/me", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"header":`))
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	_, err := NewAccount(server.URL).GetMe("personal-auth-token")
	if err == nil {
		t.Fatal("expected an error for invalid JSON")
	}
}
