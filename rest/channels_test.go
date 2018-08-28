package rest

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
	"github.com/stretchr/testify/assert"
)

func TestRestService_ChannelArchivel(t *testing.T) {

	type fields struct {
		myDoer  Doer
		channel *models.Channel
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			name: "ChannelArchive OK",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
						"success": true
					 }`,
				},
				channel: &models.Channel{ID: "GENERAL"},
			},
			wantErr: nil,
		},
		{
			name: "ChannelArchive Err",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
						"status": "error",
						"message": "you must be logged in to do this"
					  }`,
				},
				channel: &models.Channel{ID: "GENERAL"},
			},
			wantErr: errors.New("status: error, message: you must be logged in to do this"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := CreateTestRestClient(tt.fields.myDoer)
			err := rt.Rest.ChannelArchive(tt.fields.channel)

			assert.Equal(t, err, tt.wantErr, "Unexpected error")
		})
	}
}

func TestRestService_GetPublicChannels(t *testing.T) {

	type fields struct {
		myDoer Doer
	}
	tests := []struct {
		name    string
		fields  fields
		want    ChannelsResponse
		wantErr error
	}{
		{
			name: "GetPublicChannels OK",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
						"channels": [
							{
								"_id": "ByehQjC44FwMeiLbX",
								"name": "test-test",
								"t": "c",
								"usernames": [
									"testing1"
								],
								"msgs": 0,
								"u": {
									"_id": "aobEdbYhXfu5hkeqG",
									"username": "testing1"
								},
								"ts": "2016-12-09T15:08:58.042Z",
								"ro": false,
								"sysMes": true,
								"_updatedAt": "2016-12-09T15:22:40.656Z"
							},
							{
								"_id": "t7qapfhZjANMRAi5w",
								"name": "testing",
								"t": "c",
								"usernames": [
									"testing2"
								],
								"msgs": 0,
								"u": {
									"_id": "y65tAmHs93aDChMWu",
									"username": "testing2"
								},
								"ts": "2016-12-01T15:08:58.042Z",
								"ro": false,
								"sysMes": true,
								"_updatedAt": "2016-12-09T15:22:40.656Z"
							}
						],
						"success": true
					}`,
				},
			},
			want: ChannelsResponse{
				Channels: []models.Channel{
					models.Channel{ID: "ByehQjC44FwMeiLbX", Name: "test-test"},
					models.Channel{ID: "t7qapfhZjANMRAi5w", Name: "testing"},
				},
			},
			wantErr: nil,
		},
		{
			name: "GetPublicChannels Err",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
						"status": "error",
						"message": "you must be logged in to do this"
					  }`,
				},
			},
			want: ChannelsResponse{
				Channels: []models.Channel{
					models.Channel{ID: "ByehQjC44FwMeiLbX", Name: "invite-me"},
				},
			},
			wantErr: errors.New("status: error, message: you must be logged in to do this"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := CreateTestRestClient(tt.fields.myDoer)
			got, err := rt.GetPublicChannels()

			assert.Equal(t, err, tt.wantErr, "Unexpected error")

			if err == nil {
				//ToDO : needs a better test !
				assert.NotNil(t, got)
				assert.Equal(t, got.Channels[0].ID, tt.want.Channels[0].ID, "ID did not match")
				assert.Equal(t, got.Channels[0].Name, tt.want.Channels[0].Name, "Name did not match")
			}
		})
	}
}

func TestRestService_GetJoinedChannels(t *testing.T) {

	type fields struct {
		myDoer Doer
		params url.Values
	}
	tests := []struct {
		name    string
		fields  fields
		want    ChannelsResponse
		wantErr error
	}{
		{
			name: "GetJoinedChannels OK",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
						"channels": [
							{
								"_id": "ByehQjC44FwMeiLbX",
								"name": "invite-me",
								"t": "c",
								"usernames": [
									"testing1"
								],
								"msgs": 0,
								"u": {
									"_id": "aobEdbYhXfu5hkeqG",
									"username": "testing1"
								},
								"ts": "2016-12-09T15:08:58.042Z",
								"ro": false,
								"sysMes": true,
								"_updatedAt": "2016-12-09T15:22:40.656Z"
							}
						],
						"offset": 0,
						"count": 2,
						"total": 2,
						"success": true
					}`,
				},
				params: url.Values{},
			},
			want: ChannelsResponse{
				Channels: []models.Channel{
					models.Channel{ID: "ByehQjC44FwMeiLbX", Name: "invite-me"},
				},
			},
			wantErr: nil,
		},
		{
			name: "GetJoinedChannels Err",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
						"status": "error",
						"message": "you must be logged in to do this"
					  }`,
				},
				params: url.Values{},
			},
			want: ChannelsResponse{
				Channels: []models.Channel{
					models.Channel{ID: "ByehQjC44FwMeiLbX", Name: "invite-me"},
				},
			},
			wantErr: errors.New("status: error, message: you must be logged in to do this"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := CreateTestRestClient(tt.fields.myDoer)
			got, err := rt.GetJoinedChannels(tt.fields.params)

			assert.Equal(t, err, tt.wantErr, "Unexpected error")

			if err == nil {
				//ToDO : needs a better test !
				assert.NotNil(t, got)
				assert.Equal(t, got.Channels[0].ID, tt.want.Channels[0].ID, "ID did not match")
				assert.Equal(t, got.Channels[0].Name, tt.want.Channels[0].Name, "Name did not match")
			}
		})
	}
}

func TestRestService_LeaveChannel(t *testing.T) {

	type fields struct {
		myDoer  Doer
		channel *models.Channel
	}
	tests := []struct {
		name    string
		fields  fields
		want    ChannelsResponse
		wantErr error
	}{
		{
			name: "LeaveChannel OK",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
						"channel": {
						  "_id": "ByehQjC44FwMeiLbX",
						  "name": "invite-me",
						  "t": "c",
						  "usernames": [
							"testing2"
						  ],
						  "msgs": 0,
						  "u": {
							"_id": "aobEdbYhXfu5hkeqG",
							"username": "testing1"
						  },
						  "ts": "2016-12-09T15:08:58.042Z",
						  "ro": false,
						  "sysMes": true,
						  "_updatedAt": "2016-12-09T15:22:40.656Z"
						},
						"success": true
					  }`,
				},
				channel: &models.Channel{ID: "GENERAL"},
			},
			want: ChannelsResponse{
				Channels: []models.Channel{
					models.Channel{ID: "ByehQjC44FwMeiLbX", Name: "invite-me"},
				},
			},
			wantErr: nil,
		},
		{
			name: "LeaveChannel Err",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
						"status": "error",
						"message": "you must be logged in to do this"
					  }`,
				},
				channel: &models.Channel{ID: "GENERAL"},
			},
			want: ChannelsResponse{
				Channels: []models.Channel{
					models.Channel{ID: "ByehQjC44FwMeiLbX", Name: "invite-me"},
				},
			},
			wantErr: errors.New("status: error, message: you must be logged in to do this"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := CreateTestRestClient(tt.fields.myDoer)
			err := rt.LeaveChannel(tt.fields.channel)

			assert.Equal(t, err, tt.wantErr, "Unexpected error")
		})
	}
}

type testDoer struct {
	response     string
	responseCode int
	http.Header
}

func (nd testDoer) Do(*http.Request) (*http.Response, error) {
	return &http.Response{
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(nd.response))),
		StatusCode: nd.responseCode,
		Header:     nd.Header,
	}, nil
}

func CreateTestRestClient(d Doer) *Client {
	rockerServer := &url.URL{Host: "rocketchat.localhost", Scheme: "https"}
	client := NewClient(rockerServer, false)
	client.myDoer = d
	return client
}

func TestRestService_GetMembersList(t *testing.T) {

	type fields struct {
		myDoer Doer
		roomID string
	}
	tests := []struct {
		name    string
		fields  fields
		want    MembersResponse
		wantErr error
	}{
		{
			name: "GetMembersList OK",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
						"members": [
							{
								"_id": "Loz7qh9ChSqHMPymx",
								"username": "customField_apiuser.test.1529436896005",
								"name": "customField_apiuser.test.1529436896005",
								"status": "offline"
							},
							{
								"_id": "Zc3Y3cRW7ZtS7Y8Hk",
								"username": "customField_apiuser.test.1529436997563",
								"name": "customField_apiuser.test.1529436997563",
								"status": "offline"
							}
						],
						"count": 2,
						"offset": 0,
						"total": 35,
						"success": true
					}`,
				},
				roomID: "GENERAL",
			},
			want: MembersResponse{
				Status:     Status{Status: "OK"},
				Pagination: models.Pagination{Count: 2, Offset: 0, Total: 35},
				MembersList: []models.Member{
					models.Member{
						ID:       "Loz7qh9ChSqHMPymx",
						Username: "customField_apiuser.test.1529436896005",
						Name:     "customField_apiuser.test.1529436896005",
						Status:   "offline"},
					models.Member{
						ID:       "Zc3Y3cRW7ZtS7Y8Hk",
						Username: "customField_apiuser.test.1529436997563",
						Name:     "customField_apiuser.test.1529436997563",
						Status:   "offline"},
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := CreateTestRestClient(tt.fields.myDoer)
			got, err := rt.GetMembersList(tt.fields.roomID)

			assert.Equal(t, tt.wantErr, err, "Unexpected error")
			assert.Equal(t, tt.want.Count, got.Count, "Member count dose not match")
			assert.Equal(t, got.MembersList, tt.want.MembersList, "MembersList did not match")

		})
	}
}

func TestRestService_GetChannelInfo(t *testing.T) {

	type fields struct {
		myDoer  Doer
		channel *models.Channel
	}
	tests := []struct {
		name    string
		fields  fields
		want    models.Channel
		wantErr error
	}{
		{
			name: "GetChannelInfo OK",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
							"channel": {
							  "_id": "ByehQjC44FwMeiLbX",
							  "ts": "2016-11-30T21:23:04.737Z",
							  "t": "c",
							  "name": "testing",
							  "usernames": [
								"testing",
								"testing1",
								"testing2"
							  ],
							  "msgs": 1,
							  "default": true,
							  "_updatedAt": "2016-12-09T12:50:51.575Z",
							  "lm": "2016-12-09T12:50:51.555Z"
							},
							"success": true
						  }`,
				},
				channel: &models.Channel{ID: "GENERAL"},
			},
			want: models.Channel{
				ID:   "ByehQjC44FwMeiLbX",
				Name: "testing",
				Type: "c",
			},
			wantErr: nil,
		},
		{
			name: "GetChannelInfo Err",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
							"status": "error",
							"message": "you must be logged in to do this"
						  }`,
				},
				channel: &models.Channel{ID: "GENERAL"},
			},
			want:    models.Channel{},
			wantErr: errors.New("status: error, message: you must be logged in to do this"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := CreateTestRestClient(tt.fields.myDoer)
			got, err := rt.GetChannelInfo(tt.fields.channel)

			assert.Equal(t, err, tt.wantErr, "Unexpected error")

			if err == nil {
				assert.NotNil(t, got)
				assert.Equal(t, got.ID, tt.want.ID, "ID did not match")
				assert.Equal(t, got.Name, tt.want.Name, "Name did not match")
				assert.Equal(t, got.Type, tt.want.Type, "Name did not match")
			}
		})
	}
}

func TestRestService_KickChannel(t *testing.T) {

	type fields struct {
		myDoer  Doer
		channel *models.Channel
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			name: "KickChannel OK",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
						"channel": {
						  "_id": "ByehQjC44FwMeiLbX",
						  "name": "invite-me",
						  "t": "c",
						  "usernames": [
							"testing1"
						  ],
						  "msgs": 0,
						  "u": {
							"_id": "aobEdbYhXfu5hkeqG",
							"username": "testing1"
						  },
						  "ts": "2016-12-09T15:08:58.042Z",
						  "ro": false,
						  "sysMes": true,
						  "_updatedAt": "2016-12-09T15:22:40.656Z"
						},
						"success": true
					  }`,
				},
				channel: &models.Channel{
					ID:   "GENERAL",
					User: &models.User{ID: "nSYqWzZ4GsKTX4dyK"},
				},
			},
			wantErr: nil,
		},
		{
			name: "KickChannel Err",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
						"status": "error",
						"message": "you must be logged in to do this"
					  }`,
				},
				channel: &models.Channel{
					ID:   "GENERAL",
					User: &models.User{ID: "nSYqWzZ4GsKTX4dyK"},
				},
			},
			wantErr: errors.New("status: error, message: you must be logged in to do this"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := CreateTestRestClient(tt.fields.myDoer)
			err := rt.KickChannel(tt.fields.channel)

			assert.Equal(t, err, tt.wantErr, "Unexpected error")
		})
	}
}

func TestRestService_ChannelsCreate(t *testing.T) {

	type fields struct {
		myDoer  Doer
		channel *models.Channel
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			name: "ChannelsCreate OK",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
						"channel": {
						   "_id": "ByehQjC44FwMeiLbX",
						   "name": "channelname",
						   "t": "c",
						   "usernames": [
							  "example"
						   ],
						   "msgs": 0,
						   "u": {
							  "_id": "aobEdbYhXfu5hkeqG",
							  "username": "example"
						   },
						   "ts": "2016-05-30T13:42:25.304Z"
						},
						"success": true
					 }`,
				},
				channel: &models.Channel{Name: "channelname"},
			},
			wantErr: nil,
		},
		{
			name: "ChannelsCreate Err",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
						"status": "error",
						"message": "you must be logged in to do this"
					  }`,
				},
				channel: &models.Channel{ID: "GENERAL"},
			},
			wantErr: errors.New("status: error, message: you must be logged in to do this"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := CreateTestRestClient(tt.fields.myDoer)
			err := rt.ChannelsCreate(tt.fields.channel)

			assert.Equal(t, err, tt.wantErr, "Unexpected error")

		})
	}
}

func TestRestService_ChannelClose(t *testing.T) {

	type fields struct {
		myDoer  Doer
		channel *models.Channel
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			name: "ChannelClose OK",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
						"success": true
					 }`,
				},
				channel: &models.Channel{ID: "ByehQjC44FwMeiLbX"},
			},
			wantErr: nil,
		},
		{
			name: "ChannelClose Err",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
						"status": "error",
						"message": "you must be logged in to do this"
					  }`,
				},
				channel: &models.Channel{ID: "ByehQjC44FwMeiLbX"},
			},
			wantErr: errors.New("status: error, message: you must be logged in to do this"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := CreateTestRestClient(tt.fields.myDoer)
			err := rt.ChannelClose(tt.fields.channel)

			assert.Equal(t, err, tt.wantErr, "Unexpected error")

		})
	}
}
