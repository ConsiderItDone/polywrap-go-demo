![Polywrap Go Demo Application](https://user-images.githubusercontent.com/1008882/186656507-c448651d-bf75-40fb-91fc-b13f799af38e.png)

[**Polywrap**](https://polywrap.io/) is a developer tool that enables easy integration of Web3 protocols into any application. It makes it possible for applications on any platform, written in any language, to read and write data to Web3 protocols.

# Getting Started

This is a demo calculator application written as Polywrap Wrapper on Golang.

## Pre-reqs

You'll need the following installed before building your wrapper:

`nvm`

`yarn`

`docker`

`docker-compose`

`golang`

## Installation

To ensure all of your project's dependencies are installed, from inside your project's directory, simply run:

```
nvm install && nvm use
yarn
```

Then generate golang code for your wrapper

```
go generate module/calc.go 
```