package twilio_test

import (
	"context"
	"fmt"

	twilio "github.com/andybar2/twilio-go"
)

func Example_wireless() {
	client := twilio.NewClient("AC123", "123", nil)
	sim, _ := client.Wireless.Sims.Get(context.TODO(), "DE123")
	fmt.Println(sim.Status)
}
