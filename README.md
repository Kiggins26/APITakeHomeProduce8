# APITakeHomeProduce8
## APIs
- `POST: /calculate-payment` - a POST api for getting the amount for the schedule payment of a mortgage from this formula [^1] 
## How to run locally
### Install Go
1. Run `wget https://go.dev/dl/go1.21.3.linux-amd64.tar.gz`
2. Run `rm -rf /usr/local/go && tar -C /usr/local -xzf go1.21.3.linux-amd64.tar.gz `
3. Run `export PATH=$PATH:/usr/local/go/bin`
4. Verify installation with this command `go version`
   
*for non-linux operating systems go to [Download Link](https://go.dev/doc/install)*

### Clone repo
1. Run `git clone https://github.com/Kiggins26/APITakeHomeProduce8/tree/main`

### Setup local dependencies and run server
1. Run `go mod tidy`
2. Run `go run main.go`

*To run tests `cd ./pkg` and run `go test`*

### Testing Live 

`curl -X POSY https://apitakehomeproduce8-production.up.railway.app/calculate-payment -H "Content-Type: application/json" -d '{"propertyPrice": 10000.00, "downPayment": 100.00, "annualInterestRate": 1.1, "amortizationPeriod": 30, "paymentSchedule": "monthly" }'`

[^1]: ![formula](https://github.com/Kiggins26/APITakeHomeProduce8/assets/30563055/50664071-b52b-46a4-b822-a7804c394470)

