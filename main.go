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
        Balance     string        `json:"month_to_date_balance"`
        Generated   string        `json:"generated_at"`
    }
    
    var result Result
    json.Unmarshal(resp, &result)

    return result.Balance
}

func bill_with_paypal() string {
    // Create a client instance
    c, err := paypal.NewClient("clientID", "secretID", paypal.APIBaseSandBox)
    c.SetLog(os.Stdout) // Set log to terminal stdout

    accessToken, err := c.GetAccessToken(context.Background())
    fmt.Println(accessToken)

    if err != nil {
		log.Fatal( err )
	}
    
    return "true"
}

func main() {
    fmt.Println("Starting DigitalOcean billing...\n")

    // Understand the billing cycle
    now := time.Now()
    year, month, _ := now.Date()
    firstDayOfThisMonth := time.Date(year, month, 1, 0, 0, 0, 0, now.Location())
    endOfThisMonth := time.Date(year, month+1, 0, 0, 0, 0, 0, now.Location())
    fmt.Printf("Billing Period: %s thru %s\n\n", firstDayOfThisMonth, endOfThisMonth)

    // Get outstanding DigitalOcean balance
    balance_due := get_digitalocean_balance()
    fmt.Printf("Current balance: $%s\n\n", balance_due)

    // To Do: Check if it is the end of the month, etc.

    // To Do: Who are we billing?

    // To Do: Divide out the bill (total due / total users)

    // Send bill(s) for the outstanding balance with PayPal
    bill_with_paypal()
    fmt.Printf("Done\n\n")
}