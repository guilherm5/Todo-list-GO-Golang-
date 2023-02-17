# TodoList 

![Exemplo de como usar](http://g.recordit.co/ixoCGah8to.gif)

## Functionality:
- [x] Create Task
- [x] Post Task
- [x] Update Task
- [x] Delete Task
- [x] Get Task
- [x] Routes authenticated with oauth2
- [x] MVC pattern project


## Technologies used:
- [x] Golang
- [x] Gin Framework
- [x] ORM Gorm
- [x] Postgresql
- [x] Oauth2
- [x] .env



## :rocket: install and run
### prerequisites

> Updated version of GO installed
> Client API of your choice
> Install dependencies below:

### :wrench: Install dependencies

```
git clone https://github.com/guilherm5/ListaDeTarefasGO.git
```

```
go get -u gorm.io/gorm
```

```
 go get -u github.com/gin-gonic/gin
```

```
go get github.com/joho/godotenv
```

```
go mod tidy
```

```
go run main.go 
```

## How to use the project

<p>First you must obtain the credentials in the /Credentials route, and as shown in the video above, you must populate the /Token route with the credentials generated in /Credentials. access for you, which will give you access to the other existing routes in the application, follow the video because it is very detailed :) </p>

## :crossed_flags: Routes

```
GET - localhost:8089/Todo
```

```
POST - localhost:8089/Todo
```

```
PUT - localhost:8089/Todo
```

```
DELETE - localhost:8089/Todo
```

```
GET - localhost:8089/Credentials
```

```
GET - localhost:8089/Token
```

#### 	:warning: In the .env file, replace the bank information with your local bank information



