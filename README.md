### How to run the code

### 1. Start with commanline

#### Make sure the 8080 &9002 port is available and run the following command to start the server

```
go run cmd/wallet-demo/main.go
```

or

```
make run
```

### 2. Using Docker Compose

```
docker compose up -d
```

### How to test the code

#### 1. create the wallet

```
curl --location 'http://localhost:8080/wallets' \
--header 'Content-Type: application/json' \
--data '{
    "name": "test1"
}'
```

#### 2. get the wallet detail

```
curl --location --request GET 'http://localhost:8080/wallets/8748ea6a-1a11-4cfd-9a42-d77d700c9ef5' \
--header 'Content-Type: application/json' \
--data '{
    "name": "test"
}'
```

#### 3. wallet transfer

```
curl --location 'http://localhost:8080/wallets/transfer' \
--header 'Content-Type: application/json' \
--data '{
    "from_display_id":  "b90adfae-5b16-42fd-9291-689dd7c4e3d6",
    "to_display_id": "8748ea6a-1a11-4cfd-9a42-d77d700c9ef5",
    "amount": "1"
}'
```