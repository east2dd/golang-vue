# Golang Vue Boilerplate (golang, mysql, vue)
##### Golang - REST API
  - Install project with go get
  - Config project configuration
  - Seed sampe data to database
  - Run project
##### Vue - Frontend
  - Project run
  
# REST API
Installing project with go get
```sh
go get -u github.com/xyingsoft/golang-vue
```

Editting .env file for configuration
```sh
export APP_URL=localhost:3000
export APP_ENV=development
export DB_USERNAME=root
export DB_PASSWORD=
export DB_HOST=localhost
export DB_PORT=3306
export DB_DATABASE=store
export token_password=tokenGoesRightHere
export PORT=3000
```

Seeding sample data to database
```sh
$ cd $GOPATH/src/path/to/project/root
go run seeds/main.go
```

Running project
```sh
$ cd $GOPATH/src/path/to/project/root
go run main.go
```

# Frontend
Setting API Endpoint. Edit path-to-project-root/client/src/config.js
```sh
export const API_BASE_URL = 'http://localhost:3000/'
```

Running project
```sh
$ cd $GOPATH/src/path/to/project/root
cd client
npm install
npm run start
```
bring browser up and browse http://localhost:8080