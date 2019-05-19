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
   
4. [Application API](#4)

   4.1. [Register an application](#41)

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
  "email": "user@email.com",
  "active": true,
  "deleted": false
}
```

Response
```json
{
  "id": 2,
  "username": "user",
  "password": "changed",
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

<a name="4"/>

## Application API
The application API can be used to manage the applications that will be using this service to authenticate its users.
Once an application has been registered, users belonging to that application can then be managed within the context of
that application.

<a name="41"/>

### Register an application

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
