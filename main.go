package main

import (
    "fmt"
    "time"
    "net/http"
    "log"
    "io/ioutil"
    "encoding/json"
    "flag"
    "github.com/plutov/paypal/v4"
    "os"
    "context"
    "strconv"
)



func get_digitalocean_balance() string {
    auth := flag.String("auth", "xxxxxxxxxxxxx", "DigitalOcean API Key")
    flag.Parse()
    auth_value := "Bearer " + *auth

    client := &http.Client{}
    req, err := http.NewRequest("GET", "https://api.digitalocean.com/v2/customers/my/balance", nil)
    req.Header.Set("Authorization", auth_value)
    res, err := client.Do(req)

    if err != nil {
		log.Fatal( err )
	}
        
    resp, _ := ioutil.ReadAll( res.Body )
    res.Body.Close()
    //fmt.Printf("%s\n\n", resp) // Uncomment to see the complete API response

    type Result struct {
        Balance     string  `json:"month_to_date_balance"`
        Generated   string  `json:"generated_at"`
    }
    
    var result Result
    json.Unmarshal(resp, &result)

    return result.Balance
}

func bill_with_paypal() {
    // Create a client instance
    c, err := paypal.NewClient("clientID", "secretID", paypal.APIBaseSandBox)
    c.SetLog(os.Stdout) // Set log to terminal stdout

    accessToken, err := c.GetAccessToken(context.Background())
    fmt.Println(accessToken)

    if err != nil {
		log.Fatal( err )
	}
    
    return
}

func main() {
    fmt.Println("Starting DigitalOcean billing...\n")

    // Understand the billing cycle
    now := time.Now()
    year, month, _ := now.Date()
    firstDayOfThisMonth := time.Date(year, month, 1, 0, 0, 0, 0, now.Location())
    endOfThisMonth := time.Date(year, month+1, 0, 0, 0, 0, 0, now.Location())
    fmt.Printf("Billing Period: %s thru %s\n\n", firstDayOfThisMonth, endOfThisMonth)
    // To Do: Check if it is the end of the month, etc.

    // Get outstanding DigitalOcean balance
    total_balance := get_digitalocean_balance()
    fmt.Printf("Total Balance: $%s\n", total_balance)

    // Divide out the bill (total due / total users)
    var balance_due, _ = strconv.ParseFloat(total_balance, 10) // TO DO: This probably shouldn't be hardcoded to 2
    balance_due = (balance_due / 2)
    fmt.Printf("Balance Due: $%s\n\n", fmt.Sprintf("%.2f", balance_due))

    // Send bill(s) for the outstanding balance with PayPal
    bill_with_paypal()
    fmt.Printf("Done\n\n")
}