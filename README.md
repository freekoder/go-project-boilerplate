# Blueprint/Boilerplate for golang service project

Boilerplate for fast golang service development. 
This project can be useful for golang beginners as a collection of service recipes.

## Project layout
https://github.com/golang-standards/project-layout

## Makefile
https://www.alexedwards.net/blog/a-time-saving-makefile-for-your-go-projects

## Graceful shutdown
Shutdown of the service is based on the article "Implementing Graceful Shutdown in Go" by Leonidas Vrachnis:
https://www.rudderstack.com/blog/implementing-graceful-shutdown-in-go/

Shutdown uses signal.NotifyContext(https://pkg.go.dev/os/signal#NotifyContext) and 
errgroup.WithContext(https://pkg.go.dev/golang.org/x/sync/errgroup#WithContext)


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
