# learn-async-with-go

- part1.go : 각 웹사이트를 하나씩 확인한다.
```
❯ go run .\part1\part1.go
https://github.com/jhyeok is up!
https://jhyeok.com is up!
https://djqtsmstkdlxm.co is down!
```

- part2.go : 아무 것도 출력하지 않음, 모든 go 루틴이 실행되기 전에 main() 이 종료가 되기 때문이다.
```
❯ go run .\part2\part2.go
```

- part3.go : 출력의 순서는 urls에 있는 순서와 다르다. 비동기적으로 실행되었기 때문이다.
```
❯ go run .\part3\part3.go
https://djqtsmstkdlxm.co is down!
https://github.com/jhyeok is up!
https://jhyeok.com is up!
```

- part4.go : 고루틴을 사용해서 구현한다. (코드에 주석으로 전체적인 설명)
```
❯ go run .\part4\part4.go
https://djqtsmstkdlxm.co is down!
https://github.com/jhyeok is up!
https://jhyeok.com is up!
```

- part5.go : 채널을 사용해서 구현한다. (코드에 주석으로 전체적인 설명)
```
❯ go run .\part5\part5.go
https://djqtsmstkdlxm.co is down!
https://github.com/jhyeok is up!
https://jhyeok.com is up!
```

---
#### Reference

https://medium.com/@gauravsingharoy/asynchronous-programming-with-go-546b96cd50c1