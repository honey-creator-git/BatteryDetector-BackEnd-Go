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
