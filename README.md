# Better Dev Link

A weekly/daily news letter of resources/articles around the web which is
language/framework agnostic to help readers build a better
understanding of programming in general.

# Submit Content

Create a PR and add a new issue

# Workflow

We use Go itself as a task runner. Since it's so fast, no need to
compile go. Just use `go run`

## build

```
go run cmd/compile.go
```

## run internal server

```
go run cmd/server.go
```

## deploy

```
go run cmd/deploy.go
```

## publish a new issue

```
go run cmd/publish.go issue-number

## Official sendout email
go run cmd/publish.go issue-number --send
```
