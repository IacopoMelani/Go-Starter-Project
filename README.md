# Go-Starter-Project
Template of a generic project written in **GO** ![Go](https://github.com/IacopoMelani/Go-Starter-Project/workflows/Go/badge.svg) [![Build Status](https://travis-ci.org/IacopoMelani/Go-Starter-Project.svg?branch=master)](https://travis-ci.org/IacopoMelani/Go-Starter-Project) [![codecov](https://codecov.io/gh/IacopoMelani/Go-Starter-Project/branch/master/graph/badge.svg)](https://codecov.io/gh/IacopoMelani/Go-Starter-Project) [![Go Report Card](https://goreportcard.com/badge/github.com/IacopoMelani/Go-Starter-Project)](https://goreportcard.com/report/github.com/IacopoMelani/Go-Starter-Project) [![Maintainability](https://api.codeclimate.com/v1/badges/31257bde8ba9f709dd65/maintainability)](https://codeclimate.com/github/IacopoMelani/Go-Starter-Project/maintainability)

The project uses **Echo**
- https://github.com/labstack/echo

For reading the .env file is used **gotenv**
- https://github.com/subosito/gotenv

## Functionality 🍹

- Definition of controllers
- Definition of controllers **GET** and **POST**, in any case it is possible to define further methods
- Data manipulation system retrieved from database **MYSQL**
- Timed remote data recovery system
- Queue request system
- Configuration cache system via .env file
- Database migration system

## Installation

* Clone or download the project in your own **GOPATH**
* Download all dependecies with
```shell
go get -d ./...
```
* Build the application
```shell
go build main.go
```
* Ready to Go!

## Configuration
In the root directory is present a .env.sample
```env
export APP_NAME="Go-Starter-Project"
export STRING_CONNECTION="root:root@tcp(127.0.0.1:3306)/test?parseTime=true"
export APP_PORT=:8888
export USER_TIME_TO_REFRESH=30
```
Copy and rename in *.env* and fill with whatever you like

## Useful commands
To interact with the application you need to **FIRE!**

Use the following command:
```shell
./Go-Starter-Project -fire
```
This will show the following message

```shell
commands:
  -fire go!         -> start the Server 
  -fire go-migrate  -> migrate database 
  -fire go-rollback -> rollback database 
  -fire go-config   -> show the current environment 
```
