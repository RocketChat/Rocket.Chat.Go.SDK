package rest

import (
	"errors"
	"net/url"
	"testing"

	"github.com/RocketChat/Rocket.Chat.Go.SDK/models"
	"github.com/stretchr/testify/assert"
)

func TestRocket_GetServerInfo(t *testing.T) {

	type fields struct {
		myDoer Doer
	}
	tests := []struct {
		name    string
		fields  fields
		want    models.Info
		wantErr error
	}{
		{
			name: "Version match Ok",
			fields: fields{
				testDoer{
					responseCode: 200,
					response: `{
						"success": true,
						"info": {
						  "version": "0.47.0-develop",
						  "build": {
							"nodeVersion": "v4.6.2",
							"arch": "x64",
							"platform": "linux",
							"cpus": 4
						  },
						  "commit": {
							"hash": "5901cc7270e3587101631ee222def950d705c611",
							"date": "Thu Dec 1 19:08:01 2016 -0200",
							"author": "Gabriel Engel",
							"subject": "Merge branch 'develop' into experimental",
							"tag": "0.46.0",
							"branch": "experimental"
						  }
						}
					  }`,
				},
			},
			want: models.Info{
				Version: "0.47.0-develop",
			},
			wantErr: nil,
		},
		{
			name: "Version err",
			fields: fields{
				testDoer{
					responseCode: 200,
					response: `{
						"success": false
					  }`,
				},
			},
			want: models.Info{
				Version: "0.47.0-develop",
			},
			wantErr: errors.New("got false response"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := CreateTestRestClient(tt.fields.myDoer)
			got, err := rt.GetServerInfo()

			assert.Equal(t, err, tt.wantErr, "Unexpected error")
			if err == nil {
				assert.Equal(t, tt.want.Version, got.Version, "Version not matching")
			}

		})
	}
}

func TestRocket_GetDirectory(t *testing.T) {

	type fields struct {
		myDoer Doer
		params url.Values
	}
	tests := []struct {
		name    string
		fields  fields
		want    models.Directory
		wantErr error
	}{
		{
			name: "channels OK",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
						"result": [
						  {
							"_id": "GENERAL",
							"ts": "2018-08-23T01:39:00.881Z",
							"name": "general",
							"usersCount": 11,
							"lastMessage": {
							  "alias": null,
							  "msg": "",
							  "attachments": [
								{
								  "title": "Gify",
								  "image_url": "https://media0.giphy.com/media/3ohc11UljvpPKWeNva/giphy.gif?cid=e1bb72ff5b7fa6964c576d2f6b464d03",
								  "color": "#225159"
								}
							  ],
							  "parseUrls": false,
							  "bot": {
								"i": "h3m3p69QjBtgDgABp"
							  },
							  "groupable": false,
							  "ts": "2018-08-24T06:32:55.154Z",
							  "u": {
								"_id": "z5fLnYsiYZzREohkE",
								"username": "bradb",
								"name": null
							  },
							  "rid": "GENERAL",
							  "mentions": [],
							  "channels": [],
							  "_updatedAt": "2018-08-24T06:32:55.156Z",
							  "_id": "Pt7RsaNSDAE7tJH6t",
							  "sandstormSessionId": null
							}
						  }
						],
						"count": 1,
						"offset": 0,
						"total": 1,
						"success": true
					  }`,
				},
				params: url.Values{"query": []string{`{"text": "gene", "type": "channels"}`}},
			},
			want: models.Directory{
				Pagination: models.Pagination{Count: 1, Offset: 0, Total: 1},
			},
			wantErr: nil,
		},
		{
			name: "users OK",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
						"result": [
						  {
							"_id": "rocket.cat",
							"createdAt": "2018-08-23T01:39:00.958Z",
							"name": "Rocket.Cat",
							"username": "rocket.cat"
						  },
						  {
							"_id": "7foMpeuHzfW5Npoqg",
							"createdAt": "2018-08-23T01:53:05.715Z",
							"emails": [
							  {
								"address": "rocketchat@localhost",
								"verified": false
							  }
							],
							"name": "Rocket Admin",
							"username": "rocketadmin"
						  }
						],
						"count": 2,
						"offset": 0,
						"total": 2,
						"success": true
					  }`,
				},
				params: url.Values{"query": []string{`{"text": "rocket", "type": "users"}`}},
			},
			want: models.Directory{
				Pagination: models.Pagination{Count: 2, Offset: 0, Total: 2},
			},
			wantErr: nil,
		},
		{
			name: "users err",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
						"result": [
						  {
							"_id": "rocket.cat",
							"createdAt": "2018-08-23T01:39:00.958Z",
							"name": "Rocket.Cat",
							"username": "rocket.cat"
						  },
						  {
							"_id": "7foMpeuHzfW5Npoqg",
							"createdAt": "2018-08-23T01:53:05.715Z",
							"emails": [
							  {
								"address": "rocketchat@localhost",
								"verified": false
							  }
							],
							"name": "Rocket Admin",
							"username": "rocketadmin"
						  }
						],
						"count": 2,
						"offset": 0,
						"total": 2,
						"success": false
					  }`,
				},
				params: url.Values{"query": []string{`{"text": "rocket", "type": "users"}`}},
			},
			want:    models.Directory{},
			wantErr: errors.New("got false response"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := CreateTestRestClient(tt.fields.myDoer)
			got, err := rt.GetDirectory(tt.fields.params)

			assert.Equal(t, err, tt.wantErr, "Unexpected error")

			if err == nil {
				assert.Equal(t, got.Pagination.Count, tt.want.Pagination.Count, "Incorrect Pagination Count")
				assert.Equal(t, got.Pagination.Offset, tt.want.Pagination.Offset, "Incorrect Pagination Offset")
				assert.Equal(t, got.Pagination.Total, tt.want.Pagination.Total, "Incorrect Pagination Total")
				assert.NotNil(t, got)
				//ToDO : needs a better test !
				//	assert.Equal(t, got, tt.want, "Unexpected error")
			}
		})
	}
}

func TestRocket_GetSpotlight(t *testing.T) {

	type fields struct {
		myDoer Doer
		params url.Values
	}
	tests := []struct {
		name    string
		fields  fields
		want    models.Spotlight
		wantErr error
	}{
		{
			name: "GetSpotlight OK",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
						"users": [
						  {
							"_id": "rocket.cat",
							"name": "Rocket.Cat",
							"username": "rocket.cat",
							"status": "online"
						  }
						],
						"rooms": [],
						"success": true
					  }`,
				},
				params: url.Values{"query": []string{`#foobar`}},
			},
			want: models.Spotlight{
				Users: []models.User{
					models.User{
						ID:           "rocket.cat",
						Name:         "Rocket.Cat",
						UserName:     "rocket.cat",
						Status:       "online",
						Token:        "",
						TokenExpires: 0},
				},
				Rooms: []models.Channel{},
			},
			wantErr: nil,
		},
		{
			name: "GetSpotlight err",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
						"users": [
						  {
							"_id": "rocket.cat",
							"name": "Rocket.Cat",
							"username": "rocket.cat",
							"status": "online"
						  }
						],
						"rooms": [],
						"success": false
					  }`,
				},
				params: url.Values{"query": []string{`#foobar`}},
			},
			want:    models.Spotlight{},
			wantErr: errors.New("got false response"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := CreateTestRestClient(tt.fields.myDoer)
			got, err := rt.GetSpotlight(tt.fields.params)

			assert.Equal(t, err, tt.wantErr, "Unexpected error")

			if err == nil {
				assert.NotNil(t, got)
				assert.Equal(t, got.Users[0].ID, tt.want.Users[0].ID, "ID did not match")
				assert.Equal(t, got.Users[0].Name, tt.want.Users[0].Name, "Name did not match")
				assert.Equal(t, got.Users[0].UserName, tt.want.Users[0].UserName, "UserName did not match")
				assert.Equal(t, got.Users[0].Status, tt.want.Users[0].Status, "Name did not match")
			}
		})
	}
}

func TestRocket_GetStatistics(t *testing.T) {

	type fields struct {
		myDoer Doer
	}
	tests := []struct {
		name    string
		fields  fields
		want    models.StatisticsInfo
		wantErr error
	}{
		{
			name: "Statistics OK",
			fields: fields{
				testDoer{
					responseCode: 200,
					response: `{
						"statistics": {
						  "_id":"wufRdmSrjmSMhBdTN",
						  "uniqueId":"wD4EP3M7FeFzJZgk9",
						  "installedAt":"2018-02-18T19:40:45.369Z",
						  "version":"0.61.0-develop",
						  "totalUsers":88,
						  "activeUsers":88,
						  "nonActiveUsers":0,
						  "onlineUsers":0,
						  "awayUsers":1,
						  "offlineUsers":87,
						  "totalRooms":81,
						  "totalChannels":41,
						  "totalPrivateGroups":37,
						  "totalDirect":3,
						  "totlalLivechat":0,
						  "totalMessages":2408,
						  "totalChannelMessages":730,
						  "totalPrivateGroupMessages":1869,
						  "totalDirectMessages":25,
						  "totalLivechatMessages":0,
						  "lastLogin":"2018-02-24T12:44:45.045Z",
						  "lastMessageSentAt":"2018-02-23T18:14:03.490Z",
						  "lastSeenSubscription":"2018-02-23T17:58:54.779Z",
						  "os": {
							"type":"Linux",
							"platform":"linux",
							"arch":"x64",
							"release":"4.13.0-32-generic",
							"uptime":76242,
							"loadavg": [
							  0.0576171875,0.04638671875,0.00439453125
							],
							"totalmem":5787901952,
							"freemem":1151168512,
							"cpus": [
							  {
								"model":"Intel(R) Xeon(R) CPU           E5620  @ 2.40GHz",
								"speed":2405,
								"times": {
								  "user":6437000,
								  "nice":586500,
								  "sys":1432200,
								  "idle":750117500,
								  "irq":0
								}
							  },
							  {
								"model":"Intel(R) Xeon(R) CPU           E5620  @ 2.40GHz",
								"speed":2405,
								"times": {
								  "user":7319700,
								  "nice":268800,
								  "sys":1823600,
								  "idle":747642700,
								  "irq":0
								}
							  },
							  {
								"model":"Intel(R) Xeon(R) CPU           E5620  @ 2.40GHz",
								"speed":2405,
								"times": {
								  "user":7484600,
								  "nice":1003500,
								  "sys":1446000,
								  "idle":748873400,
								  "irq":0
								}
							  },
							  {
								"model":"Intel(R) Xeon(R) CPU           E5620  @ 2.40GHz",
								"speed":2405,
								"times": {
								  "user":8378200,
								  "nice":548500,
								  "sys":1443200,
								  "idle":747053300,
								  "irq":0
								}
							  }
							]
						  },
						  "process": {
							"nodeVersion":"v8.9.4",
							"pid":11736,
							"uptime":16265.506
						  },
						  "deploy": {
							"method":"tar",
							"platform":"selfinstall"
						  },
						  "migration": {
							"_id":"control",
							"version":106,
							"locked":false,
							"lockedAt":"2018-02-23T18:13:13.948Z",
							"buildAt":"2018-02-18T17:22:51.212Z"
						  },
						  "instanceCount":1,
						  "createdAt":"2018-02-24T13:13:00.236Z",
						  "_updatedAt":"2018-02-24T13:13:00.236Z"
						},
						"success":true
					  }`,
				},
			},
			want: models.StatisticsInfo{
				Statistics: models.Statistics{
					Version:    "0.61.0-develop",
					TotalUsers: 88,
				},
			},
			wantErr: nil,
		},
		{
			name: "Statistics Err",
			fields: fields{
				testDoer{
					responseCode: 500,
					response: `{
						"success":false,
						"error":"Not allowed [error-not-allowed]",
						"errorType":"error-not-allowed"
					}`,
				},
			},
			want: models.StatisticsInfo{
				Statistics: models.Statistics{},
			},
			wantErr: errors.New("Not allowed [error-not-allowed]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := CreateTestRestClient(tt.fields.myDoer)
			got, err := rt.GetStatistics()

			assert.Equal(t, err, tt.wantErr, "Unexpected error")
			if err == nil {
				assert.Equal(t, tt.want.Statistics.Version, got.Statistics.Version, "Version not matching")
				assert.Equal(t, tt.want.Statistics.TotalUsers, got.Statistics.TotalUsers, "TotalUsers not matching")
			}
		})
	}
}

// func TestRocket_GetStatisticsList(t *testing.T) {
// 	rocket := Client{Protocol: common_testing.Protocol, Host: common_testing.Host, Port: common_testing.Port}
// 	// TODO admin user
// 	rocket.auth = &authInfo{id: "4SicoW2wDjcAaRh4M", token: "4nI24JkcHTsqTtOUeJYbbrhum5T_Y6IdjAwyj72qDBu"}

// 	statistics, err := rocket.GetStatisticsList(url.Values{"query": []string{`{"_id" : "zT26ye8RAM7MaEN7S"}`}})
// 	assert.Nil(t, err)
// 	assert.NotNil(t, statistics)
// }

func TestRocket_GetStatisticsList(t *testing.T) {

	type fields struct {
		myDoer Doer
		params url.Values
	}
	tests := []struct {
		name    string
		fields  fields
		want    models.StatisticsList
		wantErr error
	}{
		{
			name: "GetStatisticsList OK",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
						"statistics": [
							{
								"_id":"v3D4mvobwfznKozH8",
								"uniqueId":"wD4EP3M7FeFzJZgk9",
								"installedAt":"2018-02-18T19:40:45.369Z",
								"version":"0.61.0-develop",
								"totalUsers":88,
								"activeUsers":88,
								"nonActiveUsers":0,
								"onlineUsers":0,
								"awayUsers":1,
								"offlineUsers":87,
								"totalRooms":81,
								"totalChannels":41,
								"totalPrivateGroups":37,
								"totalDirect":3,
								"totlalLivechat":0,
								"totalMessages":2408,
								"totalChannelMessages":730,
								"totalPrivateGroupMessages":1869,
								"totalDirectMessages":25,
								"totalLivechatMessages":0,
								"lastLogin":"2018-02-24T12:44:45.045Z",
								"lastMessageSentAt":"2018-02-23T18:14:03.490Z",
								"lastSeenSubscription":"2018-02-23T17:58:54.779Z",
								"instanceCount":1,
								"createdAt":"2018-02-24T15:13:00.312Z",
								"_updatedAt":"2018-02-24T15:13:00.312Z"
							}
						],
						"count":1,
						"offset":0,
						"total":1,
						"success":true
					}`,
				},
				params: url.Values{"query": []string{`{"_id" : "zT26ye8RAM7MaEN7S"}`}},
			},
			want: models.StatisticsList{
				Statistics: []models.Statistics{
					models.Statistics{
						ID:          "v3D4mvobwfznKozH8",
						UniqueID:    "wD4EP3M7FeFzJZgk9",
						ActiveUsers: 88,
						TotalRooms:  81,
					},
				},
				Pagination: models.Pagination{Count: 2, Offset: 0, Total: 2},
			},
			wantErr: nil,
		},
		{
			name: "GetStatisticsList err",
			fields: fields{
				myDoer: testDoer{
					responseCode: 200,
					response: `{
						"statistics": [
							{
								"_id":"v3D4mvobwfznKozH8",
								"uniqueId":"wD4EP3M7FeFzJZgk9",
								"installedAt":"2018-02-18T19:40:45.369Z",
								"version":"0.61.0-develop",
								"totalUsers":88,
								"activeUsers":88,
								"nonActiveUsers":0,
								"onlineUsers":0,
								"awayUsers":1,
								"offlineUsers":87,
								"totalRooms":81,
								"totalChannels":41,
								"totalPrivateGroups":37,
								"totalDirect":3,
								"totlalLivechat":0,
								"totalMessages":2408,
								"totalChannelMessages":730,
								"totalPrivateGroupMessages":1869,
								"totalDirectMessages":25,
								"totalLivechatMessages":0,
								"lastLogin":"2018-02-24T12:44:45.045Z",
								"lastMessageSentAt":"2018-02-23T18:14:03.490Z",
								"lastSeenSubscription":"2018-02-23T17:58:54.779Z",
								"instanceCount":1,
								"createdAt":"2018-02-24T15:13:00.312Z",
								"_updatedAt":"2018-02-24T15:13:00.312Z"
							}
						],
						"count":1,
						"offset":0,
						"total":1,
						"success":false
					}`,
				},
				params: url.Values{"query": []string{`{"_id" : "zT26ye8RAM7MaEN7S"}`}},
			},
			want:    models.StatisticsList{},
			wantErr: errors.New("got false response"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := CreateTestRestClient(tt.fields.myDoer)
			got, err := rt.GetStatisticsList(tt.fields.params)

			assert.Equal(t, err, tt.wantErr, "Unexpected error")

			if err == nil {
				assert.NotNil(t, got)
				assert.Equal(t, got.Statistics[0].ID, tt.want.Statistics[0].ID, "ID did not match")
				assert.Equal(t, got.Statistics[0].UniqueID, tt.want.Statistics[0].UniqueID, "UniqueID did not match")
				assert.Equal(t, got.Statistics[0].ActiveUsers, tt.want.Statistics[0].ActiveUsers, "ActiveUsers did not match")
				assert.Equal(t, got.Statistics[0].TotalRooms, tt.want.Statistics[0].TotalRooms, "TotalRooms did not match")
			}
		})
	}
}
