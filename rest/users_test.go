package rest

import (
	"errors"
	"testing"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
	"github.com/stretchr/testify/assert"
)

func TestRestService_Login(t *testing.T) {
	type fields struct {
		myDoer Doer
	}
	type args struct {
		credentials *models.UserCredentials
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "OK",
			fields: fields{
				testDoer{
					responseCode: 200,
					response: `{
						"status": "success",
						"data": {
							"authToken": "9HqLlyZOugoStsXCUfD_0YdwnNnunAJF8V47U3QHXSq",
							"userId": "aobEdbYhXfu5hkeqG",
							"me": {
								  "_id": "aYjNnig8BEAWeQzMh",
								  "name": "Rocket Cat",
								  "emails": [
									  {
										"address": "rocket.cat@rocket.chat",
										"verified": false
									  }
								  ],
								  "status": "offline",
								  "statusConnection": "offline",
								  "username": "rocket.cat",
								  "utcOffset": -3,
								  "active": true,
								  "roles": [
									  "admin"
								  ],
								  "settings": {
									  "preferences": {}
									}
							  }
						 }
					  }`,
				},
			},
			args: args{
				credentials: &models.UserCredentials{Email: "fred", Password: "smith"},
			},
			wantErr: nil,
		},
		{
			name: "Token",
			fields: fields{
				testDoer{
					responseCode: 200,
					response: `{
						"status": "success",
						"data": {
							"authToken": "9HqLlyZOugoStsXCUfD_0YdwnNnunAJF8V47U3QHXSq",
							"userId": "aobEdbYhXfu5hkeqG",
							"me": {
								  "_id": "aYjNnig8BEAWeQzMh",
								  "name": "Rocket Cat",
								  "emails": [
									  {
										"address": "rocket.cat@rocket.chat",
										"verified": false
									  }
								  ],
								  "status": "offline",
								  "statusConnection": "offline",
								  "username": "rocket.cat",
								  "utcOffset": -3,
								  "active": true,
								  "roles": [
									  "admin"
								  ],
								  "settings": {
									  "preferences": {}
									}
							  }
						 }
					  }`,
				},
			},
			args: args{
				credentials: &models.UserCredentials{ID: "aobEdbYhXfu5hkeqG", Token: "9HqLlyZOugoStsXCUfD_0YdwnNnunAJF8V47U3QHXSq"},
			},
			wantErr: nil,
		},
		{
			name: "Err",
			fields: fields{
				testDoer{
					responseCode: 200,
					response: `{
						"success": false
					  }`,
				},
			},
			args: args{
				credentials: &models.UserCredentials{Email: "", Password: ""},
			},
			wantErr: errors.New("got false response"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := CreateTestRestClient(tt.fields.myDoer)
			err := rt.Login(tt.args.credentials)

			assert.Equal(t, err, tt.wantErr, "Unexpected error")

		})
	}
}

func TestRestService_Logout(t *testing.T) {
	type fields struct {
		myDoer Doer
	}
	tests := []struct {
		name    string
		fields  fields
		authed  bool
		want    string
		wantErr error
	}{
		{
			name: "Logged in",
			fields: fields{
				testDoer{
					responseCode: 200,
					response: `{
						"status": "success",
						"data": {
						  "message": "You've been logged out!"
						}
					 }`,
				},
			},
			authed:  true,
			want:    "You've been logged out!",
			wantErr: nil,
		},
		{
			name: "Was not logged in",
			fields: fields{
				testDoer{
					responseCode: 200,
					response:     ``,
				},
			},
			authed:  false,
			want:    "Was not logged in",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := CreateTestRestClient(tt.fields.myDoer)
			if tt.authed {
				rt.auth = &authInfo{id: "HAS A TOKEN"}
			}
			got, err := rt.Logout()

			assert.Equal(t, err, tt.wantErr, "Unexpected error")

			if err == nil {
				assert.Equal(t, got, tt.want, "ID did not match")
			}

		})
	}
}
