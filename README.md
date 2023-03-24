# godotenvsafe ![CI](https://github.com/uulwake/godotenvsafe/workflows/CI/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/uulwake/godotenvsafe)](https://goreportcard.com/report/github.com/uulwake/godotenvsafe)

A package to help you to check missing environment variable before running your app and getting unnecessary runtime error.

It uses [godotenv](https://github.com/joho/godotenv) under the hood to load your environment variables. This package will make sure that you don't miss any env vars.

## Installation

```shell
go get github.com/uulwake/godotenvsafe
```

## Usage
For example below is `.env` file.
```
DB_NAME=db-name
DB_USER=user
DB_PASSWORD=pass
```

Then create `.env.template` with the key only like below.
```
DB_NAME=
DB_USER=
DB_PASSWORD=
```

Then you can load your environment variables safely like this.
```go
import (
    "log"

    "github.com/uulwake/godotenvsafe"
)

func main() {
    err := godotenvsafe.Load(".env")
    if err != nil {
        log.Fatal("Error loading env variables", err.Error())
    }
}
```

## Conventions
This package will automatically search for env template by adding suffix `.template` in your env file. Some examples like below:

| Env file              | Env Template File         |
|-----------------------|---------------------------|
|`.env`                 |`.env.template`            |
|`.env.local`           |`.env.local.template`      |
|`.env.development`     |`.env.development.template`|
|`.env.testing`         |`.env.testing.template`    |
|`.env.prod`            |`.env.prod.template`       |



