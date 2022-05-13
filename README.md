# Echozap

> Middleware for Golang [Echo](https://echo.labstack.com/) framework that provides integration with Uber´s [Zap](https://github.com/uber-go/zap)  logging library for logging HTTP requests.


[![Build Status](https://travis-ci.com/diptomondal007/bdstockexchange.svg?branch=master)](https://travis-ci.com/github/diptomondal007/bdstockexchange)
[![Coverage Status](https://coveralls.io/repos/github/diptomondal007/echozap/badge.svg?branch=master)](https://coveralls.io/github/diptomondal007/echozap?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/diptomondal007/echozap)](https://goreportcard.com/report/github.com/diptomondal007/echozap)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

## Pre-requisites

*  Go with Go modules enabled.
*  [Echo v4](https://echo.labstack.com/)
*  [Zap](https://github.com/uber-go/zap)

## Usage

```go
package main

import (
	"net/http"

	"github.com/diptomondal007/echozap"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func main() {
	e := echo.New()

	logger, _ := zap.NewProduction()

	e.Use(echozap.ZapLogger(echozap.WrapSugared(logger.Sugar())))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
```

## Logged details

The following information is logged:

*  Status Code
*  Time
*  Uri
*  Method
*  Hostname
*  Remote IP Address

## 🤝 Contributing

Contributions, issues and feature requests are welcome!

## Show your support

If this project have been useful for you, I would be grateful to have your support.

Give a ⭐️ to the project, or just:

<a href="https://www.buymeacoffee.com/Z1Bu6asGV" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: auto !important;width: auto !important;" ></a>

## 📝 License

Copyright © 2019 [Dipto Mondal](https://github.com/diptomondal007).

This project is [MIT](LICENSE) licensed.
