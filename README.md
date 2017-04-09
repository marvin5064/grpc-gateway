# grpc-gateway
play around with grpc-gateway with the swagger documentation

# Set up
1. make sure you have your local go-compiler set up already. If not, please set it up: https://golang.org/project/
2. make sure you have google protobuf compiler >= 3.0.0 set up already. If not, please set it up:
```
mkdir tmp
cd tmp
git clone https://github.com/google/protobuf
cd protobuf
./autogen.sh
./configure
make
make check
sudo make install
```

3. update the package below:
```
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
```
4. in `grpc-client`, `grpc-server`, `grpc-gateway`:
```
go build
./{name above}
```
if you go build in `grpc-client`, then you shall `./grpc-client`

5. If you have the `grpc-gateway` running, please checkout url:
```
localhost:8080/v1/student/{id} < you will get return in id: 1, 2, 3
localhost:8080/v1/course/{id} < you will get return in id: 1, 2
localhost:8080/v1/professor/{id} < you will get return in id: 1, 2
```
Please try `localhost:8080/v1/student/1`


PS: If you want to execute `grpc-client` or `grpc-gateway`, make sure you have `grpc-server` already runnning


- The UI of generated documentation of API on swaggerhub: `https://app.swaggerhub.com/api/marvin5064/university-university_proto/0.0.1`