# authenticator
An authentication server.

## Table of Contents  
1. [Technologies](#1)  
2. [Authentication](#2)  
3. [System API](#3)

   3.1. [Create a system user](#31)
   
   3.2. [Get a system user](#32)
   
   3.3. [Update a system user](#33)
   
   3.4. [Delete a system user](#34)
   
   3.5. [List system users](#35)
   
   3.6. [List applications](#36)
   
4. [Application API](#4)

   4.1. [Create a application](#41)
   
   4.2. [Get a application](#42)
   
   4.3. [Update a application](#43)
   
   4.4. [Delete a application](#44)
   
   4.5. [List application users](#45)
   
   4.6. [Create a application user](#46)
   
   4.7. [Get a application user](#47)
   
   4.8. [Update a application user](#48)
   
   4.9. [Delete a application user](#49)

<a name="1"/>

## Technologies
* GoLang 1.11.4
* MySQL 5.7.25

<a name="2"/>

## Authentication
* HTTP Basic (RFC 7617)
* OAuth2 (Coming soon)

<a name="3"/>

## System API
This application exposes a restful API that you can use to interface with it. All requests must be 
authenticated by including the HTTP Basic header `Authentication: bearer 1231ewqw...`. The bearer token is
the username and password separated by a colon, base 64 encoded.

The system API can be used to manage the users that can be authorized to use the application API.

When the application is started a temporary user is available with the username `temp` and password `temp`. 
Use this temporary login to create your own user account, after which you should delete the `temp` user.

<a name="31"/>

### Create a system user
A system user is a user that can configure the application. At least one system user should be created to
replace the temporary `temp` user that is automatically created when the application is first started.

Endpoint
```text
POST /system/user/create
```

Payload
```json
{
  "username": "user",
  "password": "password",
  "email": "user@email.com"
}
```

Response
```json
{
  "id": 2,
  "username": "user",
  "password": "password",
  "application_id": 0,
  "email": "user@email.com",
  "active": true,
  "deleted": false,
  "api": true,
  "created": "2018-11-07 16:59:06",
  "created_by": "temp",
  "last_updated": "2018-11-07 16:59:06",
  "last_updated_by": "temp"
}
```

<a name="32"/>

### Get a system user

Endpoint
```text
GET /system/user/{id}
```

Response
```json
{
  "id": 1,
  "username": "temp",
  "password": "password",
  "application_id": 0,
  "email": "user@email.com",
  "active": true,
  "deleted": false,
  "api": true,
  "created": "2018-11-07 16:59:06",
  "created_by": "system",
  "last_updated": "2018-11-07 16:59:06",
  "last_updated_by": "system"
}
```

<a name="33"/>

### Update a system user

Endpoint
```text
PUT /system/user/{id}
```

Payload
```json
{
  "username": "user",
  "password": "password",
  "application_id": 0,
  "email": "user@email.com",
  "active": true
}
```

Response
```json
{
  "id": 2,
  "username": "user",
  "password": "changed",
  "application_id": 0,
  "email": "user@email.com",
  "active": true,
  "deleted": false,
  "api": true,
  "created": "2018-11-07 16:59:06",
  "created_by": "temp",
  "last_updated": "2018-11-07 16:59:06",
  "last_updated_by": "user"
}
```

<a name="34"/>

### Delete a system user

Endpoint
```text
DELETE /system/user/{id}
```

Response
```text
200 OK
```

<a name="35"/>

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
      "created": "2018-11-07 16:59:06",
      "created_by": "temp",
      "last_updated": "2018-11-07 16:59:06",
      "last_updated_by": "user"
    }
]
```

<a name="35"/>

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

<a name="4"/>

## Application API
The application API can be used to manage the applications that will be using this service to authenticate its users.
Once an application has been registered, users belonging to that application can then be managed within the context of
that application.

<a name="41"/>

### Create a application

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

<a name="42"/>

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

<a name="43"/>

### Update a application

Endpoint
```text
PUT /application/{id}
```

Payload
```json
{
  "name": "changed",
  "url": "https://changed.com/"
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

<a name="44"/>

### Delete a application

Endpoint
```text
DELETE /application/{id}
```

Response
```text
200 OK
```

<a name="45"/>

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
      "created": "2018-11-07 16:59:06",
      "created_by": "temp",
      "last_updated": "2018-11-07 16:59:06",
      "last_updated_by": "user"
    }
]
```

<a name="46"/>

### Create a application user

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
  "created": "2018-11-07 16:59:06",
  "created_by": "temp",
  "last_updated": "2018-11-07 16:59:06",
  "last_updated_by": "user"
}
```

<a name="47"/>

### Get a application user

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
  "created": "2018-11-07 16:59:06",
  "created_by": "temp",
  "last_updated": "2018-11-07 16:59:06",
  "last_updated_by": "user"
}
```

<a name="48"/>

### Update a application user

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
  "created": "2018-11-07 16:59:06",
  "created_by": "temp",
  "last_updated": "2018-11-07 16:59:06",
  "last_updated_by": "user"
}
```

<a name="49"/>

### Delete a application user

```text
DELETE /user/{id}
```

Response
```text
200 OK
```