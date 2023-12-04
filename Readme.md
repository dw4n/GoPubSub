# Go Pub Sub
A go Publisher and Subscriber implementation using Azure Service Bus

## How to Start
1. Go Mod Tidy and Go Mod Vendor
2. create .env variable, paste your service bus connection string
```
CONN_STR_SUBSCRIBER=your_subscriber_connection_string
CONN_STR_PUBLISHER=your_publisher_connection_string
TOPIC_NAME=your_topic_name

```
3. Go Run .
4. Hit api / to check if its working
5. Hit api /create to publish user data
```
{
  "id": "1",
  "fullname": "Tommy Vercetti",
  "username": "tommy_vercetti",
  "email": "tommy@example.com",
  "isEmailVerified": true,
  "emailLastLogin": "2023-12-04T10:00:00Z",
  "mobilePhone": "+1234567890",
  "mobileLastLogin": "2023-12-04T11:00:00Z",
  "HashPassword": "hashed_password_here",
  "isDeleted": false,
  "isLocked": false,
  "lockLimitUtc": "2023-12-05T12:00:00Z",
  "invalidPasswordCounter": 0,
  "forgotPasswordCounter": 0,
  "signUpDate": "2023-12-01T08:00:00Z"
}

```
6. See the console log if it receive the data. For example
```
Sending -----
{1 Tommy Vercetti tommy_vercetti tommy@example.com true 2023-12-04 10:00:00 +0000 UTC +1234567890 2023-12-04 11:00:00 +0000 UTC hashed_password_here false false 2023-12-05 12:00:00 +0000 UTC 0 0 2023-12-01 08:00:00 +0000 UTC}
Success -----
17:22:57 | 200 |    940.5416ms |       127.0.0.1 | POST    | /create         
------------------This is from the cloud----------------------------
&{map[] [123 34 105 100 34 58 34 49 34 44 34 102 117 108 108 110 97 109 101 34 58 34 84 111 109 109 121 32 86 101 114 99 101 116 116 105 34 44 34 117 115 101 114 110 97 109 101 34 58 34 116 111 109 109 121 95 118 101 114 99 101 116 116 105 34 44 34 101 109 97 105 108 34 58 34 116 111 109 109 121 64 101 120 97 109 112 108 101 46 99 111 109 34 44 34 105 115 69 109 97 105 108 86 101 114 105 102 105 101 100 34 58 116 114 117 101 44 34 101 109 97 105 108 76 97 115 116 76 111 103 105 110 34 58 34 50 48 50 51 45 49 50 45 48 52 84 49 48 58 48 48 58 48 48 90 34 44 34 109 111 98 105 108 101 80 104 111 110 101 34 58 34 43 49 50 51 52 53 54 55 56 57 48 34 44 34 109 111 98 105 108 101 76 97 115 116 76 111 103 105 110 34 58 34 50 48 50 51 45 49 50 45 48 52 84 49 49 58 48 48 58 48 48 90 34 44 34 72 97 115 104 80 97 115 115 119 111 114 100 34 58 34 104 97 115 104 101 100 95 112 97 115 115 119 111 114 100 95 104 101 114 101 34 44 34 105 115 68 101 108 101 116 101 100 34 58 102 97 108 115 101 44 34 105 115 76 111 99 107 101 100 34 58 102 97 108 115 101 44 34 108 111 99 107 76 105 109 105 116 85 116 99 34 58 34 50 48 50 51 45 49 50 45 48 53 84 49 50 58 48 48 58 48 48 90 34 44 34 105 110 118 97 108 105 100 80 97 115 115 119 111 114 100 67 111 117 110 116 101 114 34 58 48 44 34 102 111 114 103 111 116 80 97 115 115 119 111 114 100 67 111 117 110 116 101 114 34 58 48 44 34 115 105 103 110 85 112 68 97 116 101 34 58 34 50 48 50 51 45 49 50 45 48 49 84 48 56 58 48 48 58 48 48 90 34 125] <nil> <nil> <nil> <nil> <nil> 1 0xc00000a4a0 2023-12-04 09:22:54.771 +0000 UTC 2023-12-06 09:22:54.771 +0000 UTC 2023-12-04 09:23:54.818 +0000 UTC [44 217 163 50 223 153 71 143 145 44 38 69 0 50 0 57] 4cc0f53a-bd8b-4130-be24-fc55df9238a5 0xc00021ea20 <nil> <nil> <nil> 0xc00000a488 <nil> 0 0xc0000882a0 48h0m0s <nil> 0xc00006b040 false}
------------------This is the readable version----------------------------
Message ID: 4cc0f53a-bd8b-4130-be24-fc55df9238a5
Body: {"id":"1","fullname":"Tommy Vercetti","username":"tommy_vercetti","email":"tommy@example.com","isEmailVerified":true,"emailLastLogin":"2023-12-04T10:00:00Z","mobilePhone":"+1234567890","mobileLastLogin":"2023-12-04T11:00:00Z","HashPassword":"hashed_password_here","isDeleted":false,"isLocked":false,"lockLimitUtc":"2023-12-05T12:00:00Z","invalidPasswordCounter":0,"forgotPasswordCounter":0,"signUpDate":"2023-12-01T08:00:00Z"}
Delivery Count: 1
Enqueued Time: 2023-12-04 09:22:54.771 +0000 UTC
Expires At: 2023-12-06 09:22:54.771 +0000 UTC
To: <nil>
Subject: create
```