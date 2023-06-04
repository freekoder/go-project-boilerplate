# Blueprint/Boilerplate for golang service project

A boilerplate for fast golang service development. 
This project can be useful for golang beginners as a collection of service recipes.

> **Warning**
> This project is on the development mode now

## Targets
- [ ] Create template for rapid golang service development
- [ ] Collect in one place best practices for golang service development
- [ ] Highlight the best tools for beginner golang developers

## Project layout
https://github.com/golang-standards/project-layout

## Makefile
The Makefile was taken from the article 
["Quick tip: A time-saving Makefile for your Go projects"](https://www.alexedwards.net/blog/a-time-saving-makefile-for-your-go-projects)
by Alex Edwards.

```shell
$ make help                                                            
Usage:
  help                print this help message
  tidy                format code and tidy modfile
  audit               run quality control checks
  lint                run linter
  test                run all tests
  test/cover          run all tests and display coverage
  build               build the application
  run                 run the  application
  run/live            run the application with reloading on file changes
  push                push changes to the remote Git repository
  production/deploy   deploy the application to production
```
This Makefile contains **run/live** target. It allows to run application with live reloads.
**run/live** implemented on base of 
[Air - Live reload for Go apps](https://github.com/cosmtrek/air). More information can be found at 
["Using Air with Go to implement live reload"](https://blog.logrocket.com/using-air-go-implement-live-reload/).

Unfortunately, this Makefile is missing linting target, based on 
[golangci-lint](https://github.com/golangci/golangci-lint).

## Graceful shutdown
Shutdown of the service is based on the article 
["Implementing Graceful Shutdown in Go"](https://www.rudderstack.com/blog/implementing-graceful-shutdown-in-go/) 
by Leonidas Vrachnis.

Graceful shutdown uses:
* [signal.NotifyContext() from os/signal package](https://pkg.go.dev/os/signal#NotifyContext) 
* [errgroup.WithContext() from golang.org/x/sync/errgroup package](https://pkg.go.dev/golang.org/x/sync/errgroup#WithContext)

> **Warning**
> signal.NotifyContext was added only in golang 1.16

> **Note**
> errgroup use cases:
> * ["Why you should be using errgroup.WithContext() in your Golang server handlers"](https://www.fullstory.com/blog/why-errgroup-withcontext-in-golang-server-handlers/)
> by Scott Blum
> * ["LEARNING GO: CONCURRENCY PATTERNS USING ERRGROUP PACKAGE"](https://mariocarrion.com/2021/09/03/learning-golang-concurrency-patterns-errgroup-package.html)
> by Mario Carrion

## Web framework
https://echo.labstack.com/

## Clean template
https://evrone.com/blog/go-clean-template

## Configuration
* https://github.com/ilyakaznacheev/cleanenv
* https://dev.to/ilyakaznacheev/clean-configuration-management-in-golang-1c89

## Logging
* https://github.com/uber-go/zap
* https://github.com/snovichkov/zap-gelf

## REST-API
https://betterprogramming.pub/intro-77f65f73f6d3

## Swagger
* https://github.com/swaggo/swag
* https://github.com/go-swagger/go-swagger

## TODO
- [ ] Project structure
- [ ] Makefile
- [ ] Logging
- [ ] Graceful shutdown
- [ ] Graylog
- [ ] Command line args
- [ ] Configuration
- [ ] Profiling
- [ ] Dependency injection
- [ ] Web framework
- [ ] Templates
- [ ] REST-API
- [ ] JWT
- [ ] Swagger
- [ ] Websockets
- [ ] gRPC
- [ ] CI
- [ ] Sentry
- [ ] Docker
- [ ] Clean architecture
- [ ] Linter
- [ ] Testing
- [ ] Database
- [ ] Migrations
- [ ] Redis
- [ ] RabbitMQ/AMQP
