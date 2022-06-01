# Study gRPC Video
[gRPC [Golang] Master Class: Build Modern API & Microservices](https://www.udemy.com/course/grpc-golang/) 강의를 구매하여서 듣고 있습니다.

> 이제서야 본격적으로 gRPC 공부를 하고자 한다.  
이 강의를 산지 약 1년 정도 된 것 같은데 말만 하지 말고 하자...

## Project Set Up
[gRPC - Go QuickStart](https://grpc.io/docs/languages/go/quickstart/)를 참고하여서 설치하였습니다.
```shell
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
$ go mod init
$ go get -u google.golang.org/grpc
```

## To-Do
- Session1: gRPC Course Overview
    - [X] 1.gRPC Introduction - 2022.05.25 완료
    - [X] 2. Course Objective - 2022.05.25 완료
    - [X] 3. About your instructor - 2022.05.25 완료
    - [X] 4. Important Message - 2022.05.25 완료
- Session2: Code Download
    - [X] 5. Code Download - 2022.05.25 완료
- Session3: [Theory] gRPC Internals Deep Dive
    - [X] 6. Why this section? - 2022.05.25 완료
    - [X] 7. Protocol Buffers & Language Interoperability - 2022.05.25 완료
    - [X] 8. HTTP/2 - 2022.05.25 완료
    - [X] 9. 4 Types of gRPC APIs - 2022.05.25 완료
    - [X] 10. Scalability in gRPC - 2022.05.25 완료
    - [X] 11. Security in gRPC(SSL) - 2022.05.25 완료
    - [X] 12. gRPC vs REST - 2022.05.25 완료
    - [X] 13. Section Summary - why use gRPC - 2022.05.25 완료
- Session4: 섹션 4: [Hands-On] gRPC Project Overview & Setup
    - [X] 14. Project Setup (Golang + VSCode + Protoc) - 2022.05.26 완료
    - [X] 15. Go Dependencies Setup - 2022.05.26 완료
    - [X] *16. Makefile(Window)* - 2022.05.26 완료
    - [X] 17. Makefile - 2022.05.26 완료
    - [X] 18. Server Setup Boilerpate Code - 2022.05.26 완료
    - [X] 19. Client Setup Boilerplate Code - 2022.05.26 완료
- Sesion5: [Hand-On] gRPC Unary
    - [X] 20. Unary API Server Impiementation - 2022.05.26 완료
    - [X] 21. Unary API Client Implementaion - 2022.05.26 완료
    - [X] 22. [Exercise] Sum API - 2022.05.26 완료
    - [X] 23. [Solution] Sum API - 2022.05.26 완료
- Session6: [Hand-On] gRPC Server Streaming
    - [X] 24. Server Streaming API Server Implementaion - 2022.05.26 완료
    - [X] 25. Server Streaming API Client Implementaion - 2022.05.26 완료
    - [X] 26. [Exercise] Primes API  - 2022.05.27 완료
    - [X] 27. [Solution] Primes API - 2022.05.27 완료
- Session7 [Hand-On] gRPC Client Streaming
    - [X] 28. Client Streaming API Server Implementaion - 2022.06.01 완료
    - [X] 29. Client Streaming API Client Implementaion - 2022.06.01 완료
    - [ ] 30. [Exercise] Avg API
    - [ ] 31. [Solution] Avg API