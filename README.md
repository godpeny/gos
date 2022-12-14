# GOS
## Architecture
golang 1.19.1 or above
grpc
codegen(deepmap)
ent
generic

clean archtecture (https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

- entity (== model)
- (entity_handler == <repository>)
- usecase : 실제 비즈니스 로직이 일어나는 곳. entity(or entity_handler), 다른 usecase들을 사용.
- delivery(+ middleware) : grpc, restAPI, MQ. 유저에게 인풋을 받아서 usecase가 인식할 수 있는 형태로 가공해서 전달.

usecase 인터페이스 구현 <-> delivery와 커뮤니케이션
entity_handler 인터페이스 구현 <-> usecase 와 커뮤니케이션 (아마도?)

## Dev Env
### db
```shell
docker run --name mysql -e MYSQL_ROOT_PASSWORD=toor -d -p 3306:3306 mysql:latest
```

### ent
```shell
go get -d entgo.io/ent/cmd/ent
go run -mod=mod entgo.io/ent/cmd/ent init User

```

### run protoc

```shell
protoc --go_out=. --go_opt=paths=source_relative \
   --go-grpc_out=. --go-grpc_opt=paths=source_relative \
   delivery/grpc/_.proto
```


## reference
https://go.dev/doc/tutorial/generics#prerequisites
https://github.com/improbable-eng/grpc-web/blob/master/client/grpc-web-react-example/go/exampleserver/exampleserver.go
https://github.com/easyCZ/grpc-web-hacker-news/blob/master/server/main.go
https://github.com/floydjones1/grpc-chatApp-react
https://medium.com/geekculture/writing-a-simple-grpc-application-in-golang-from-scratch-6b70e8688f14

https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
https://github.com/golang-standards/project-layout
https://medium.com/easyread/golang-clean-archithecture-efd6d7c43047
https://github.com/bxcodec/go-clean-arch
https://github.com/evrone/go-clean-template