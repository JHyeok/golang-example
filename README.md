# golang example

## 설치 및 개발 환경

1. [Go 설치](https://go.dev/dl)

2. VS Code에서 [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.go) 설치

3. VS Code setting
```json
{
  "[go]": {
    "editor.defaultFormatter": "golang.go",
    "editor.formatOnSave": true
  },
}
```

## 메모

- 명령 팔레트(ctrl + shift + p)에서 `Go: test`라고 입력하면 테스트 관련 명령어들이 많다.
- [learn-go-with-tests](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world)
- [Testing in Go: Naming Conventions](https://ieftimov.com/post/testing-in-go-naming-conventions/)
- `main.go:6:2: no required module provides package github.com/JHyeok/golang-example/learngo/banking: go.mod file not found in current directory or any parent directory; see 'go help modules'` 에러가 발생하면 터미널에서 `go env -w GO111MODULE=auto` 실행
- `go get` 명령어는 패키지 및 관련 종속성(dependency)을 다운로드 및 설치하는 명령어이다.
- `go get github.com/gofiber/fiber/v2`로 fiber v2 패키지 설치를 진행하는데 `cannot find package "github.com/gofiber/fiber/v2" in any of`라는 오류가 발생하면, 아래와 같은 방식으로 해결할 수 있다.
```go
import "github.com/gofiber/fiber/v2"
go mod init <app_name>
go mod tidy
```
- [Fiber의 Prefork](https://github.com/gofiber/fiber/issues/180)


## 참고
https://medium.com/@gauravsingharoy/asynchronous-programming-with-go-546b96cd50c1

https://nomadcoders.co/go-for-beginners