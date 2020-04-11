# API
- ルーム作成
- ルーム更新
- point更新
- ルーム参加
- ルームGET

## エンティティ
- room
- Participants
### room
|key|type|description|note|  
|:---|:---|:---|:----|
|ID|string|自動採番||
|No|string|roomNo(自動採番)||
|Title|string|title||
|Requirements|[Requirement]|要件||
|Participants|[String]|参加者IDs||
|CurrentRequirement|string|現在取り扱う要件ID||
|ModeratorID|string|主催者のID||

### Requirement
|key|type|description|note|  
|:---|:---|:---|:----|
|ID|string|自動採番||
|Title|string|title||
|Description|string|description||
|IsFinish|bool|完了したか?||

### Participant
|key|type|description|note|  
|:---|:---|:---|:----|
|ID|string|自動採番||
|JoinedRoom|string|参加しているroom||
|Name|string|名前||
|Points|[point]||
|IsModerator|bool|主催者か？||

### Point
|key|type|description|note|  
|:---|:---|:---|:----|
|ID|string|要件ID||
|point|string|ポイント||


## ルーム作成
### request
POST /api/v1/room
X-PL-TOKEN:UPsDPB5uz3QkZAPgjwYFZCgpagEASNK3
```json
{
    "Title":"room name",
    "Name":"name",
}
```
### response
```json
{
    "room":{
        "ID":"id",
        "No":"No",
        "Title":"title",
        "Requirements":[
        ],
        "Participants":[
            {
                "ID":"id",
                "Name":"name",
                "points":[
                    {}
                ]
            }
        ]        
     }
}
```

## ルーム更新
主催者しか呼ばない想定だが、ガードなぞない
### request
PUT /api/v1/room?"room_id"="aaaa"
X-PL-TOKEN:UPsDPB5uz3QkZAPgjwYFZCgpagEASNK3
X-PL-PARTICIPANT-ID:"Participant id"
```json
{
}
```

## Participant更新
ポイントの更新もこちらで行う
### request
PUT /api/v1/room/participant
X-PL-TOKEN:UPsDPB5uz3QkZAPgjwYFZCgpagEASNK3
X-PL-PARTICIPANT-ID:"Participant id"
```json
{
}
```

## ルーム参加
### request
POST /api/v1/room/join?"room_no"="1111"
X-PL-TOKEN:UPsDPB5uz3QkZAPgjwYFZCgpagEASNK3
```json
{
    "room_no":"1111"
}
```
### response
```JSON
{
  "room":{
    "room情報",
  }
}
```

## ルームGET
### request
GET /api/v1/room/participant?"room_no"="1111"
X-PL-TOKEN:UPsDPB5uz3QkZAPgjwYFZCgpagEASNK3
X-PL-PARTICIPANT-ID:"Participant id"
``` json
```
### response
```JSON
{
  "room":{
    "room情報",
  }, 
}
```
