# BatteryDetector BackEnd Golang

## SignUp

- Register a new user with FirstName, LastName, Email and Password

POST: /api/v1/user/create

Request Body: 
    {
        "Email": "techbush.dev@gmail.com",
        "Password": "IvanP.9899",
        "FirstName": "Ttenochtchi",
        "LastName": "Bush"
    }
Response:
    {
        "payload": {
            "data": {
                "id": "64de3de3337679ed502c0b9b",
                "email": "tenochbush@gmail.com",
                "firstName": "Ttenochtchi",
                "lastName": "Bush",
                "password": "$2a$10$A89v.aV1QxF7tYdr.b936esckG4PVS2Kzd6TNX4aEtE/Cydx0G4pW"
            },
            "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlbm9jaGJ1c2hAZ21haWwuY29tIiwiZXhwIjoxNjkyMjkzNjM1LCJpYXQiOjE2OTIyODY0MzUsIm5iZiI6MTY5MjI4NjQzNX0.bnb3SD38yJ_9L71C5x9IiYi-vOpBcPqAtzbL9LE5DUixGsCeWu2q4YBDy1nWzKi1Tva-5PLYUZVb2LdKV3PiUQ"
        },
        "status": true
    }

## LogIn

- Login User with Email and Password

POST: /api/v1/user/signin
Request Body: 
    {
        "Email": "tenochbush@gmail.com",
        "Password": "IvanP.9899"
    }
Response:
    {
        "payload": {
            "data": {
                "id": "64de3de3337679ed502c0b9b",
                "email": "tenochbush@gmail.com",
                "firstName": "Ttenochtchi",
                "lastName": "Bush",
                "password": "$2a$10$A89v.aV1QxF7tYdr.b936esckG4PVS2Kzd6TNX4aEtE/Cydx0G4pW"
            },
            "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlbm9jaGJ1c2hAZ21haWwuY29tIiwiZXhwIjoxNjkyMjk2MzU3LCJpYXQiOjE2OTIyODkxNTcsIm5iZiI6MTY5MjI4OTE1N30.mwdbYmhyTVg7ND7LM3dbvNKce77TP5YV2d_vcBcVc1GsLYmaDiMt0OemZ6WXBiDm5Ui3sSBvMwNkHudedSZW3w"
        },
        "status": true
    }

## Add Charge (Admin)

- Send request to add a new charge with name, ipaddress, lat/lon

POST: /api/v1/charge/add
Request Body:
    {
        "Name" : "Charge - 1",
        "IPAddress" : "192.173.62.115",
        "LatLon" : "34:69"
    }
Response:
    {
        "payload": {
            "id": "64dedd697e386de67518a9c4",
            "name": "Charge - 1",
            "ipAddress": "192.173.62.115",
            "latlon": "34:69",
            "users": null
        },
        "status": true
    }

## Edit Charge (Admin)

- Send request to edit a charge with id

POST: /api/v1/charge/edit/:chargeId
Request Body:
    {
        "IPAddress" : "192.173.62.117",
    }
Response:
    {
        "payload": {
            "id": "64dedd697e386de67518a9c4",
            "name": "Charge - 1",
            "ipAddress": "192.173.62.117",
            "latlon": "34:69",
            "users": null
        },
        "status": true
    }
