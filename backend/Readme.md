Live reloading:
`nodemon -e go --signal SIGTERM --exec 'go' run main.go`

After importing some package:
`go mod tidy` // auto imports the package looking at the dependencies
`go get packageName` // imports the package 

Create a `.env` file at root of the backend to 