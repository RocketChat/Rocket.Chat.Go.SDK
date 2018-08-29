package rest

import (
	"errors"
	"io"
	"net/url"
	"testing"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/common_testing"
	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
	"github.com/RocketChat/Rocket.Chat.Go.SDK/realtime"
	"github.com/stretchr/testify/assert"
)

var (
	testUserName  string   // nolint
	testUserEmail string   // nolint
	testPassword  = "test" // nolint
	testClient    *Client  // nolint
)

// nolint - deadcode !
func getDefaultClient(t *testing.T) (c *Client) {

	if testClient == nil {
		testUserEmail = common_testing.GetRandomEmail()
		testUserName = common_testing.GetRandomString()
		c = getAuthenticatedClient(t, testUserName, testUserEmail, testPassword)
	}

	return c
}

// nolint is unused
func getAuthenticatedClient(t *testing.T, name, email, password string) *Client {
	client := Client{Protocol: common_testing.Protocol, Host: common_testing.Host, Port: common_testing.Port}
	credentials := &models.UserCredentials{Name: name, Email: email, Password: password}

	rtClient, err := realtime.NewClient(&url.URL{Host: common_testing.Host + ":" + common_testing.Port}, true)
	assert.Nil(t, err)
	_, regErr := rtClient.RegisterUser(credentials)
	assert.Nil(t, regErr)

	loginErr := client.Login(credentials)
	assert.Nil(t, loginErr)

	return &client
}

// nolint - deadcode !
func findMessage(messages []models.Message, user string, msg string) *models.Message {
	var m *models.Message
	for i := range messages {
		m = &messages[i]
		if m.User.UserName == user && m.Msg == msg {
			return m
		}
	}

	return nil
}

// nolint - deadcode !
func getChannel(channels []models.Channel, name string) *models.Channel {
	for _, r := range channels {
		if r.Name == name {
			return &r
		}
	}

	return nil
}

func TestRestService_getURL(t *testing.T) {
	type fields struct {
		myDoer   Doer
		protocol string
		host     string
		port     string
		version  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Full",
			fields: fields{
				myDoer:   testDoer{},
				host:     "rocketchat.localhost",
				protocol: "https",
				port:     "443",
				version:  "v1",
			},
			want: "https://rocketchat.localhost:443/api/v1",
		},
		{
			name: "host",
			fields: fields{
				myDoer: testDoer{},
				host:   "rocketchat.localhost",
			},
			want: "://rocketchat.localhost:/api/v1",
		},
		{
			name: "protocol",
			fields: fields{
				myDoer:   testDoer{},
				protocol: "https",
			},
			want: "https://:/api/v1",
		},
		{
			name: "port",
			fields: fields{
				myDoer: testDoer{},
				port:   "443",
			},
			want: "://:443/api/v1",
		},
		{
			name: "version",
			fields: fields{
				myDoer:  testDoer{},
				version: "v1",
			},
			want: "://:/api/v1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := CreateTestRestClient(tt.fields.myDoer)
			rt.Host = tt.fields.host
			rt.Protocol = tt.fields.protocol
			rt.Port = tt.fields.port
			rt.Version = tt.fields.version
			got := rt.getURL()

			assert.Equal(t, got, tt.want, "Unexpected error")
		})
	}
}

func TestRestService_Get(t *testing.T) {
	type fields struct {
		myDoer Doer
	}
	type args struct {
		api      string
		params   url.Values
		response Response
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "Ok",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response:     `{"success": true}`,
				},
			},
			args: args{
				api:      "",
				params:   url.Values{"query": []string{`#foobar`}},
				response: new(InfoResponse),
			},
			wantErr: nil,
		},
		{
			name: "Err 200",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response:     `{"success": false}`,
				},
			},
			args: args{
				api:      "",
				params:   url.Values{"query": []string{`#foobar`}},
				response: new(InfoResponse),
			},
			wantErr: errors.New("got false response"),
		},
		{
			name: "Err 500",
			fields: fields{
				myDoer: testDoer{
					responseCode: 500,
					response:     ``,
				},
			},
			args: args{
				api:      "",
				params:   url.Values{"query": []string{`#foobar`}},
				response: new(InfoResponse),
			},
			wantErr: errors.New("Request error: "),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := CreateTestRestClient(tt.fields.myDoer)
			err := rt.Get(tt.args.api, tt.args.params, tt.args.response)

			assert.Equal(t, err, tt.wantErr, "Unexpected error")
		})
	}
}

func TestRestService_Post(t *testing.T) {
	type fields struct {
		myDoer Doer
	}
	type args struct {
		api      string
		body     io.Reader
		response Response
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "Ok",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response:     `{"success": true}`,
				},
			},
			args: args{
				api:      "",
				body:     nil,
				response: new(InfoResponse),
			},
			wantErr: nil,
		},
		{
			name: "Err 200",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response:     `{"success": false}`,
				},
			},
			args: args{
				api:      "",
				body:     nil,
				response: new(InfoResponse),
			},
			wantErr: errors.New("got false response"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := CreateTestRestClient(tt.fields.myDoer)
			err := rt.Post(tt.args.api, tt.args.body, tt.args.response)

			assert.Equal(t, err, tt.wantErr, "Unexpected error")
		})
	}
}

func TestRestService_PostForm(t *testing.T) {
	type fields struct {
		myDoer Doer
	}
	type args struct {
		api      string
		params   url.Values
		response Response
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "Ok",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response:     `{"success": true}`,
				},
			},
			args: args{
				api:      "",
				params:   url.Values{"query": []string{`#foobar`}},
				response: new(InfoResponse),
			},
			wantErr: nil,
		},
		{
			name: "Err 200",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response:     `{"success": false}`,
				},
			},
			args: args{
				api:      "",
				params:   url.Values{"query": []string{`#foobar`}},
				response: new(InfoResponse),
			},
			wantErr: errors.New("got false response"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := CreateTestRestClient(tt.fields.myDoer)
			err := rt.PostForm(tt.args.api, tt.args.params, tt.args.response)

			assert.Equal(t, err, tt.wantErr, "Unexpected error")
		})
	}
}

func TestRestService_doRequest(t *testing.T) {
	type fields struct {
		myDoer Doer
	}
	type args struct {
		method   string
		api      string
		params   url.Values
		body     io.Reader
		response Response
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name: "Ok",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response:     `{"success": true}`,
				},
			},
			args: args{
				method:   "GET",
				api:      "",
				params:   url.Values{"query": []string{`#foobar`}},
				body:     nil,
				response: new(InfoResponse),
			},
			wantErr: nil,
		},
		{
			name: "Err 200",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response:     `{"success": false}`,
				},
			},
			args: args{
				method:   "GET",
				api:      "",
				params:   url.Values{"query": []string{`#foobar`}},
				body:     nil,
				response: new(InfoResponse),
			},
			wantErr: errors.New("got false response"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := CreateTestRestClient(tt.fields.myDoer)
			err := rt.doRequest(tt.args.method, tt.args.api, tt.args.params, tt.args.body, tt.args.response)
			assert.Equal(t, err, tt.wantErr, "Unexpected error")
		})
	}
}
