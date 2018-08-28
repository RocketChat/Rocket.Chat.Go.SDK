package goRocket

import (
	"errors"
	"testing"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
	"github.com/stretchr/testify/assert"
)

func TestRocket_Send(t *testing.T) {

	type fields struct {
		myDoer  Doer
		channel *models.Channel
		msg     string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			name: "Send OK",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
						"ts": 1481748965123,
						"channel": "general",
						"message": {
						  "alias": "",
						  "msg": "This is a test!",
						  "parseUrls": true,
						  "groupable": false,
						  "ts": "2016-12-14T20:56:05.117Z",
						  "u": {
							"_id": "y65tAmHs93aDChMWu",
							"username": "graywolf336"
						  },
						  "rid": "GENERAL",
						  "_updatedAt": "2016-12-14T20:56:05.119Z",
						  "_id": "jC9chsFddTvsbFQG7"
						},
						"success": true
					  }`,
				},
				channel: &models.Channel{ID: "GENERAL"},
				msg:     "This is a test!",
			},
			wantErr: nil,
		},
		{
			name: "Send Err",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
						"status": "error",
						"message": "got false response"
					  }`,
				},
				channel: &models.Channel{ID: "GENERAL"},
			},
			wantErr: errors.New("got false response"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := CreateTestRestClient(tt.fields.myDoer)
			err := rt.Rest.Send(tt.fields.channel, tt.fields.msg)

			assert.Equal(t, err, tt.wantErr, "Unexpected error")

		})
	}
}

func TestRocket_PostMessage(t *testing.T) {

	type fields struct {
		myDoer Doer
		msg    *models.PostMessage
	}
	tests := []struct {
		name    string
		fields  fields
		want    *MessageResponse
		wantErr error
	}{
		{
			name: "PostMessage OK",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
						"ts": 1481748965123,
						"channel": "general",
						"message": {
						  "alias": "",
						  "msg": "This is a test!",
						  "parseUrls": true,
						  "groupable": false,
						  "ts": "2016-12-14T20:56:05.117Z",
						  "u": {
							"_id": "y65tAmHs93aDChMWu",
							"username": "graywolf336"
						  },
						  "rid": "GENERAL",
						  "_updatedAt": "2016-12-14T20:56:05.119Z",
						  "_id": "jC9chsFddTvsbFQG7"
						},
						"success": true
					  }`,
				},
				msg: &models.PostMessage{
					RoomID:  "",
					Channel: "",
					Text:    "",
				},
			},
			want: &MessageResponse{
				Status:  Status{Success: true, Status: "", Message: ""},
				Message: models.Message{RoomID: "GENERAL", Msg: "This is a test!"},
			},
			wantErr: nil,
		},
		{
			name: "PostMessage Err",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
						"status": "error",
						"message": "got false response"
					  }`,
				},
				msg: &models.PostMessage{},
			},
			want:    &MessageResponse{},
			wantErr: errors.New("got false response"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := CreateTestRestClient(tt.fields.myDoer)
			got, err := rt.Rest.PostMessage(tt.fields.msg)

			assert.Equal(t, err, tt.wantErr, "Unexpected error")

			if err == nil {
				assert.NotNil(t, got)
				assert.Equal(t, got.Status.Message, tt.want.Status.Message, "Status Message did not match")
				assert.Equal(t, got.Status.Status, tt.want.Status.Status, "Status Status did not match")
				assert.Equal(t, got.Message.Msg, tt.want.Message.Msg, "Message Msg did not match")
				assert.Equal(t, got.Message.RoomID, tt.want.Message.RoomID, "Message RoomID did not match")

			}

		})
	}
}

func TestRocket_GetMessages(t *testing.T) {

	type fields struct {
		myDoer  Doer
		channel *models.Channel
		page    *models.Pagination
	}
	tests := []struct {
		name    string
		fields  fields
		want    []models.Message
		wantErr error
	}{
		{
			name: "GetMessages OK",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
						"messages": [
						  {
							"_id": "AkzpHAvZpdnuchw2a",
							"rid": "ByehQjC44FwMeiLbX",
							"msg": "hi",
							"ts": "2016-12-09T12:50:51.555Z",
							"u": {
							  "_id": "y65tAmHs93aDChMWu",
							  "username": "testing"
							},
							"_updatedAt": "2016-12-09T12:50:51.562Z"
						  },
						  {
							"_id": "vkLMxcctR4MuTxreF",
							"t": "uj",
							"rid": "ByehQjC44FwMeiLbX",
							"ts": "2016-12-08T15:41:37.730Z",
							"msg": "testing2",
							"u": {
							  "_id": "bRtgdhzM6PD9F8pSx",
							  "username": "testing2"
							},
							"groupable": false,
							"_updatedAt": "2016-12-08T16:03:25.235Z"
						  },
						  {
							"_id": "bfRW658nEyEBg75rc",
							"t": "uj",
							"rid": "ByehQjC44FwMeiLbX",
							"ts": "2016-12-07T15:47:49.099Z",
							"msg": "testing",
							"u": {
							  "_id": "nSYqWzZ4GsKTX4dyK",
							  "username": "testing1"
							},
							"groupable": false,
							"_updatedAt": "2016-12-07T15:47:49.099Z"
						  },
						  {
							"_id": "pbuFiGadhRZTKouhB",
							"t": "uj",
							"rid": "ByehQjC44FwMeiLbX",
							"ts": "2016-12-06T17:57:38.635Z",
							"msg": "testing",
							"u": {
							  "_id": "y65tAmHs93aDChMWu",
							  "username": "testing"
							},
							"groupable": false,
							"_updatedAt": "2016-12-06T17:57:38.635Z"
						  }
						],
						"success": true
					  }`,
				},
				channel: &models.Channel{ID: "GENERAL"},
				page:    &models.Pagination{Count: 1, Offset: 0, Total: 1},
			},
			want: []models.Message{
				models.Message{ID: "AkzpHAvZpdnuchw2a", RoomID: "ByehQjC44FwMeiLbX", Msg: "hi"},
				models.Message{ID: "vkLMxcctR4MuTxreF", RoomID: "ByehQjC44FwMeiLbX", Msg: "testing2"},
				models.Message{ID: "bfRW658nEyEBg75rc", RoomID: "ByehQjC44FwMeiLbX", Msg: "testing"},
				models.Message{ID: "pbuFiGadhRZTKouhB", RoomID: "ByehQjC44FwMeiLbX", Msg: "testing"},
			},
			wantErr: nil,
		},
		{
			name: "GetMessages Err",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
						"status": "error",
						"message": "got false response"
					  }`,
				},
				channel: &models.Channel{ID: "GENERAL"},
				page:    &models.Pagination{Count: 1, Offset: 0, Total: 1},
			},
			want:    []models.Message{},
			wantErr: errors.New("status: error, message: got false response"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := CreateTestRestClient(tt.fields.myDoer)
			got, err := rt.Rest.GetMessages(tt.fields.channel, tt.fields.page)

			assert.Equal(t, err, tt.wantErr, "Unexpected error")

			if err == nil {
				assert.NotNil(t, got)

				for i := 0; i < len(got); i++ {
					assert.Equal(t, got[i].ID, tt.want[i].ID, "ID did not match")
					assert.Equal(t, got[i].RoomID, tt.want[i].RoomID, "RoomID did not match")
					assert.Equal(t, got[i].Msg, tt.want[i].Msg, "Msg did not match")
				}

			}

		})
	}
}
