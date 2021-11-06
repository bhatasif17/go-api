# This project contains 4 APIs. The users are stored in PostgreSql Database.

1.  Registering the user

2.  User Login

3.  Get Block Details from https://sochain.com/api/v2/get_block/{NETWORK}/{BLOCKHASH}

4.  Get Transaction Details from https://sochain.com/api/v2/get_tx/{NETWORK}/{TXID}

# API PATHS:

# Api for registration

1.  POST http://localhost:8000/api/user/new
    Content-Type: application/json
    {
    "email" : "a@a.com",
    "password" : "Asif@123"
    }

# Api for login

2.  POST http://localhost:8000/api/user/login
    Content-Type: application/json
    {
    "email" : "a@a.com",
    "password" : "password"
    }

# Api for getting transaction details

3. GET http://localhost:8000/api/transaction/{network}/{txid}
   Authorization: Bearer <token>

# Api for getting block details

4. GET http://localhost:8000/api/block/{network}/{blockhash}
   Authorization: Bearer <token>

# Packages used in this project

1. gorilla/mux — A powerful URL router and dispatcher. We use this package to match URL paths with their handlers.
2. jinzhu/gorm — ORM library for Golang. We use this ORM(Object relational mapper) package to interact smoothly with our database.
3. dgrijalva/jwt-go — Used to sign and verify JWT tokens
4. joho/godotenv — Used to load .env files into the project
