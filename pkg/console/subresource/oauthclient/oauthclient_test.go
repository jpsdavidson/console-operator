package oauthclient

import (
	"github.com/go-test/deep"
	"github.com/openshift/console-operator/pkg/api"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"

	oauthv1 "github.com/openshift/api/oauth/v1"
	routev1 "github.com/openshift/api/route/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestDeRegisterConsoleFromOAuthClient(t *testing.T) {
	type args struct {
		client *oauthv1.OAuthClient
	}
	tests := []struct {
		name string
		args args
		want *oauthv1.OAuthClient
	}{
		{
			name: "Test - Deregister Console from Oauth Client",
			args: args{
				client: &oauthv1.OAuthClient{
					TypeMeta:                            v1.TypeMeta{},
					ObjectMeta:                          v1.ObjectMeta{},
					Secret:                              "TestSecret1",
					AdditionalSecrets:                   nil,
					RespondWithChallenges:               false,
					RedirectURIs:                        []string{"Test"},
					GrantMethod:                         "",
					ScopeRestrictions:                   nil,
					AccessTokenMaxAgeSeconds:            nil,
					AccessTokenInactivityTimeoutSeconds: nil,
				},
			},
			want: &oauthv1.OAuthClient{
				TypeMeta:                            v1.TypeMeta{},
				ObjectMeta:                          v1.ObjectMeta{},
				Secret:                              "shouldBeRandom",
				AdditionalSecrets:                   nil,
				RespondWithChallenges:               false,
				RedirectURIs:                        []string{},
				GrantMethod:                         "",
				ScopeRestrictions:                   nil,
				AccessTokenMaxAgeSeconds:            nil,
				AccessTokenInactivityTimeoutSeconds: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DeRegisterConsoleFromOAuthClient(tt.args.client)
			if diff := deep.Equal(got.RedirectURIs, tt.want.RedirectURIs); got.Secret == tt.want.Secret || diff != nil {
				t.Error(diff)
			}
		})
	}
}

func TestStub(t *testing.T) {
	tests := []struct {
		name string
		want *oauthv1.OAuthClient
	}{
		{
			name: "Test Stub oauth client",
			want: &oauthv1.OAuthClient{
				TypeMeta: v1.TypeMeta{},
				ObjectMeta: metav1.ObjectMeta{
					Name: api.OAuthClientName,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := deep.Equal(Stub(), tt.want); diff != nil {
				t.Error(diff)
			}
		})
	}
}

func TestSetRedirectURI(t *testing.T) {
	type args struct {
		client *oauthv1.OAuthClient
		route  *routev1.Route
	}
	tests := []struct {
		name string
		args args
		want *oauthv1.OAuthClient
	}{
		{
			name: "Test set redirect URI",
			args: args{
				client: &oauthv1.OAuthClient{
					TypeMeta:                            metav1.TypeMeta{},
					ObjectMeta:                          metav1.ObjectMeta{},
					Secret:                              "",
					AdditionalSecrets:                   nil,
					RespondWithChallenges:               false,
					RedirectURIs:                        nil,
					GrantMethod:                         "",
					ScopeRestrictions:                   nil,
					AccessTokenMaxAgeSeconds:            nil,
					AccessTokenInactivityTimeoutSeconds: nil,
				},
				route: &routev1.Route{
					TypeMeta:   metav1.TypeMeta{},
					ObjectMeta: metav1.ObjectMeta{},
					Spec: routev1.RouteSpec{
						Host: "example.com",
					},
					Status: routev1.RouteStatus{},
				},
			},
			want: &oauthv1.OAuthClient{
				TypeMeta:                            metav1.TypeMeta{},
				ObjectMeta:                          metav1.ObjectMeta{},
				Secret:                              "",
				AdditionalSecrets:                   nil,
				RespondWithChallenges:               false,
				RedirectURIs:                        []string{"https://example.com/auth/callback"},
				GrantMethod:                         "",
				ScopeRestrictions:                   nil,
				AccessTokenMaxAgeSeconds:            nil,
				AccessTokenInactivityTimeoutSeconds: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := deep.Equal(SetRedirectURI(tt.args.client, tt.args.route), tt.want); diff != nil {
				t.Error(diff)
			}
		})
	}
}
