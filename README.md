# golang example

## 설치 및 개발환경

1. [Go 설치](https://go.dev/dl)

2. VSCode에서 [Go extension](https://marketplace.visualstudio.com/items?itemName=golang.go) 설치

3. VSCode setting
```json
{
  "[go]": {
    "editor.defaultFormatter": "golang.go",
    "editor.formatOnSave": true
  },
}
```

## 메모

- go mod init
- 명령 팔레트(ctrl + shift + p)에서 `Go: test`라고 입력하면 테스트 관련 명령어들이 많다.
- [learn-go-with-tests](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world)
- [Testing in Go: Naming Conventions](https://ieftimov.com/post/testing-in-go-naming-conventions/)