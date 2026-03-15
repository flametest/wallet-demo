### How to run the code

#### 1. Start with command line

##### Make sure the 3306 & 8080 & 9002 port is available and run the following command to start the server
a. start the mysql server.
```
docker compose up mysql -d
```

b. run the code
```
go run cmd/wallet-demo/main.go
```

or

```
make run
```

#### 2. Using Docker Compose

```
docker compose up -d
```

### API Example

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
--header 'Content-Type: application/json' 
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

### How to test the code

#### 1. http service is available on port: 8080, you can test with test/http_test.go.
```
go test ./test -v -run TestHttpService
go test ./test -bench=BenchmarkCreateWalletHTTP -benchmem
```

#### 2. grpc service is also available on port: 9002, you can also test with test/grpc_test.go.
```
go test ./test -v -run TestGrpcService
go test ./test -bench=BenchmarkCreateWallet -benchmem
```