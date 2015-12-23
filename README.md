# timerapp
A fitness timer app.

# Contents
+ [Register User](#register-user)
+ [Login User](#login-user)
+ [List Timers](#list-timers)
+ [Retrieve Timer](#retrieve-timer)
+ [Update Timer](#update-timer)
+ [Delete Timer](#delete-timer)

## Register User
**POST** `/v1/user/register/`

**Request:**
```json
{
    "email": "bruh@bruh.io",
    "password": "password1"
}
```

**Response:**
```json
{
    "id": 1,
    "token": "sup3rl3g1tT0ken"
}
```

**Status Codes:**
+ `201`
+ `400`
+ `409`


## Login User
**POST** `/v1/user/login/`

**Request:**
```json
{
    "email": "bruh@bruh.io",
    "password": "password1"
}
```

**Response:**
```json
{
    "id": 1,
    "token": "sup3rl3g1tT0ken"
}
```

**Status Codes:**
+ `200`
+ `400`
+ `401`

## List Timers
**GET:** `/v1/timer/`

**Note(s):**
+ User must be authenticated.

**Response**
```json
{
    "count": 3,
    "previous": null,
    "next": null,
    "results": [
        {
            "id": 1,
            "user": 1,
            "time": 75,
            "created": "2015-12-23T06:57:53.150609Z"
        },
        ...
    ]
}
```

**Status Codes:**
+ `200`
+ `401`

## Retrieve Timer
**GET:** `/v1/timer/:timer_id/`

**Note(s):**
+ User must be authenticated.

**Response**
```json

{
    "id": 1,
    "user": 1,
    "time": 75,
    "created": "2015-12-23T06:57:53.150609Z"
}
```

**Status Codes:**
+ `200`
+ `401`
+ `404`

## Update Timer
**PATCH:** `/v1/timer/:timer_id/`

**Note(s):**
+ User must be authenticated.

**Request:**
```json
{
    "time": 120
}
```

**Response**
```json

{
    "id": 1,
    "user": 1,
    "time": 120,
    "created": "2015-12-23T06:57:53.150609Z"
}
```

**Status Codes:**
+ `200`
+ `400`
+ `401`
+ `404`

## Delete Timer
**DELETE:** `/v1/timer/:timer_id/`

**Note(s):**
+ User must be authenticated.

**Status Codes:**
+ `204`
+ `401`
+ `404`
