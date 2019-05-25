# authenticator
An authentication server.

## Table of Contents  
1. [Technologies](#1)  
2. [Authentication](#2)  
3. [API](#3)
   
   3.1. [List system users](#31)
   
   3.2. [List applications](#32)
   
   3.3. [Create a application](#33)

   3.4. [Create a user](#34)
   
   3.5. [Get a user](#35)
   
   3.6. [Update a user](#36)
   
   3.7. [Delete a user](#37)
   
   3.8. [Get a application](#38)
   
   3.9. [Update a application](#39)
   
   4.0. [Delete a application](#40)
   
   4.1. [List application users](#41)
   
   4.2. [Login](#42)

<a name="1"/>

## Technologies
* GoLang 1.11.4
* MySQL 5.7.25

<a name="2"/>

## Authentication
* HTTP Basic (RFC 7617)
* OAuth2 (Coming soon)

<a name="3"/>

## API
This application exposes a restful API that you can use to interface with it. All requests must be 
authenticated by including the HTTP Basic header `Authentication: bearer 1231ewqw...`. The bearer token is
the username and password separated by a colon, base 64 encoded.

When the application is started a temporary user is available with the username `temp` and password `temp`. 
Use this temporary login to create your own user account, after which you should delete the `temp` user.

<a name="31"/>

## List system users

Endpoint
```text
GET /system/users
```

Response
```json
[
    {
      "id": 1,
      "username": "user1",
      "password": "password",
      "email": "user1@email.com",
      "application_id": 0,
      "active": true,
      "deleted": false,
      "api": true,
      "admin": false,
      "created": "2018-11-07 16:59:06",
      "created_by": "temp",
      "last_updated": "2018-11-07 16:59:06",
      "last_updated_by": "user"
    },
    {
      "id": 2,
      "username": "user2",
      "password": "password",
      "email": "user2@email.com",
      "application_id": 0,
      "active": true,
      "deleted": false,
      "api": true,
      "admin": false,
      "created": "2018-11-07 16:59:06",
      "created_by": "temp",
      "last_updated": "2018-11-07 16:59:06",
      "last_updated_by": "user"
    }
]
```

<a name="32"/>

## List applications

Endpoint
```text
GET /system/applications
```

Response
```json
[
  {
    "id": "1",
    "name": "example",
    "description": "description",
    "url": "https://example.com/",
    "active": true,
    "deleted": false,
    "created": "2018-11-07 16:59:06",
    "created_by": "user1",
    "last_updated": "2018-11-07 16:59:06",
    "last_updated_by": "user1"
  },
  {
    "id": "2",
    "name": "another",
    "description": "description",
    "url": "https://another.com/",
    "active": true,
    "deleted": false,
    "created": "2018-11-07 16:59:06",
    "created_by": "user1",
    "last_updated": "2018-11-07 16:59:06",
    "last_updated_by": "user1"
  }
]
```

<a name="33"/>

### Create a application

Rules
1. Only API users can create applications.

Parameters

| Name        | Type           | Required  |
| ------------- |:-------------:| -----:|
| name      | string | yes | |
| description     | string      |   yes |
| email | string      |    yes |

Endpoint
```text
POST /application/create
```

Payload
```json
{
  "name": "example",
  "description": "description",
  "url": "https://example.com/"
}
```

Response
```json
{
  "id": "1",
  "name": "example",
  "description": "description",
  "url": "https://example.com/",
  "active": true,
  "deleted": false,
  "created": "2018-11-07 16:59:06",
  "created_by": "user1",
  "last_updated": "2018-11-07 16:59:06",
  "last_updated_by": "user1"
}
```

<a name="34"/>

### Create a user

Rules
1. Only admins can create other admins.
2. Only API users and admins can create other API users.

Parameters

| Name        | Type           | Required  | Default |
| ------------- |:-------------:| -----:|  -----:|
| username      | string | yes | |
| password     | string      |   yes | |
| email | string      |    yes | |
| application_id | int      |    yes | |
| api | bool      |    no | false|
| admin | bool      |    no | false|
| active | bool      |    no | false|

Endpoint
```text
POST /user
```

Payload
```json
{
  "username": "user",
  "password": "password",
  "email": "user@email.com",
  "api": false,
  "admin": false,
  "active": true,
  "application_id": 1
}
```

Response
```json
{
  "id": 1,
  "username": "user",
  "password": "password",
  "email": "user@email.com",
  "application_id": 1,
  "active": true,
  "deleted": false,
  "api": false,
  "admin": false,
  "created": "2018-11-07 16:59:06",
  "created_by": "temp",
  "last_updated": "2018-11-07 16:59:06",
  "last_updated_by": "user"
}
```

<a name="35"/>

### Get a user

Endpoint
```text
GET /user/{id}
```

Response
```json
{
  "id": 1,
  "username": "user",
  "password": "password",
  "application_id": 1,
  "email": "user@email.com",
  "active": true,
  "deleted": false,
  "api": false,
  "admin": false,
  "created": "2018-11-07 16:59:06",
  "created_by": "temp",
  "last_updated": "2018-11-07 16:59:06",
  "last_updated_by": "user"
}
```

<a name="36"/>

### Update a user

Endpoint
```text
PUT /user/{id}
```

Payload
```json
{
  "username": "user",
  "password": "password",
  "application_id": 1,
  "email": "user@email.com",
  "api": false,
  "admin": false,
  "active": true
}
```

Response
```json
{
  "id": 1,
  "username": "user",
  "password": "password",
  "application_id": 1,
  "email": "user@email.com",
  "active": true,
  "deleted": false,
  "api": false,
  "admin": false,
  "created": "2018-11-07 16:59:06",
  "created_by": "temp",
  "last_updated": "2018-11-07 16:59:06",
  "last_updated_by": "user"
}
```

<a name="37"/>

### Delete a user

```text
DELETE /user/{id}
```

Response
```text
200 OK
```

<a name="38"/>

### Get a application

Endpoint
```text
GET /application/{id}
```

Response
```json
{
  "id": "1",
  "name": "example",
  "description": "description",
  "url": "https://example.com/",
  "active": true,
  "deleted": false,
  "created": "2018-11-07 16:59:06",
  "created_by": "user1",
  "last_updated": "2018-11-07 16:59:06",
  "last_updated_by": "user1"
}
```

<a name="39"/>

### Update a application

Endpoint
```text
PUT /application/{id}
```

Payload
```json
{
  "name": "changed",
  "description": "description",
  "url": "https://changed.com/",
  "active": true
}
```

Response
```json
{
  "id": "1",
  "name": "changed",
  "description": "description",
  "url": "https://changed.com/",
  "active": true,
  "deleted": false,
  "created": "2018-11-07 16:59:06",
  "created_by": "user1",
  "last_updated": "2018-11-07 16:59:06",
  "last_updated_by": "user1"
}
```

<a name="40"/>

### Delete a application

Endpoint
```text
DELETE /application/{id}
```

Response
```text
200 OK
```

<a name="41"/>

### List application users

Endpoint
```text
GET /application/users/{id}
```

Response
```json
[
    {
      "id": 1,
      "username": "user1",
      "password": "password",
      "email": "user1@email.com",
      "application_id": 1,
      "active": true,
      "deleted": false,
      "api": false,
      "admin": false,
      "created": "2018-11-07 16:59:06",
      "created_by": "temp",
      "last_updated": "2018-11-07 16:59:06",
      "last_updated_by": "user"
    },
    {
      "id": 2,
      "username": "user2",
      "password": "password",
      "email": "user2@email.com",
      "application_id": 1,
      "active": true,
      "deleted": false,
      "api": false,
      "admin": false,
      "created": "2018-11-07 16:59:06",
      "created_by": "temp",
      "last_updated": "2018-11-07 16:59:06",
      "last_updated_by": "user"
    }
]
```

<a name="42"/>

### Login

Endpoint
```text
POST /login
```

Payload
```json
{
  "username": "username",
  "password": "password"
}
```

Response
```text
200 OK
```