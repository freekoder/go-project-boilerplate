# Blueprint/Boilerplate for golang service project

A boilerplate for fast golang service development. 
This project can be useful for golang beginners as a collection of service recipes.

## Project layout
https://github.com/golang-standards/project-layout

## Makefile
https://www.alexedwards.net/blog/a-time-saving-makefile-for-your-go-projects

## Graceful shutdown
Shutdown of the service is based on the article 
["Implementing Graceful Shutdown in Go"](https://www.rudderstack.com/blog/implementing-graceful-shutdown-in-go/) 
by Leonidas Vrachnis.

Graceful shutdown uses:
* [signal.NotifyContext from os/signal package](https://pkg.go.dev/os/signal#NotifyContext) 
* [errgroup.WithContext from golang.org/x/sync/errgroup package](https://pkg.go.dev/golang.org/x/sync/errgroup#WithContext)

> **Warning**
> signal.NotifyContext was added only in golang 1.16

> **Note**
> 
> [Why you should be using errgroup.WithContext() in your Golang server handlers](https://www.fullstory.com/blog/why-errgroup-withcontext-in-golang-server-handlers/)
> [LEARNING GO: CONCURRENCY PATTERNS USING ERRGROUP PACKAGE](https://mariocarrion.com/2021/09/03/learning-golang-concurrency-patterns-errgroup-package.html)

## Web framework
https://echo.labstack.com/

## Clean template
https://evrone.com/blog/go-clean-template

## Configuration
https://github.com/ilyakaznacheev/cleanenv

## Logging
https://github.com/uber-go/zap
https://github.com/rs/zerolog
https://github.com/snovichkov/zap-gelf

## REST-API
https://betterprogramming.pub/intro-77f65f73f6d3

## Swagger
https://github.com/swaggo/swag
https://github.com/go-swagger/go-swagger

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
