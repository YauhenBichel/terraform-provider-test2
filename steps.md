- go mod edit -module terraform-provider-test2
- go mod tidy
- mkdir docker_compose
- go run main.go
- terraform init
- go install .
- terraform plan
- go get github.com/hashicorp-demoapp/hashicups-client-go@v0.1.0
- go mod tidy
- go install .
- cd docker_compose
- docker-compose up
- curl localhost:19090/health/readyz
    kill postgres if it is running locally
- curl -X POST localhost:19090/signup -d '{"username":"education", "password":"test123"}'
- {"UserID":1,"Username":"education","token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTY5MTMxNDMsInRva2VuX2lkIjoxLCJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImVkdWNhdGlvbiJ9.6yVNw_hYe1oPTiK5e0EmddkzTrxan0q-dMD4F070xSM"}
    export HASHICUPS_TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTY5MTMxNDMsInRva2VuX2lkIjoxLCJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImVkdWNhdGlvbiJ9.6yVNw_hYe1oPTiK5e0EmddkzTrxan0q-dMD4F070xSM

- Build and install the updated provider.
$ go install .

- terraform plan