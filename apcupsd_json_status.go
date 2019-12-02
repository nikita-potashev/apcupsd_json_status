package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"time"

	"github.com/mdlayher/apcupsd"
)

var (
	daemonMode      = flag.Bool("d", false, "Use daemon mode")
	collectInterval = flag.Int("i", 20, "Collect interval in seconds")
	apcupdsHostPort = flag.String("c", "localhost:3551", "Connection string for apcupsd NIS server")
)

func init() {
	flag.Parse()
}

// UPSCollector has the values
type UPSCollector struct {
	Timestamp              time.Time `json:"timestamp"`
	UPSName                string    `json:"USBName"`
	UPSMode                string    `json:"UPSMode"`
	UPSModel               string    `json:"UPSModel"`
	UPSLoadPercent         float64   `json:"LoadPercent"`
	BatteryChargePercent   float64   `json:"BatteryChargePercent"`
	LineVolts              float64   `json:"LineVolts"`
	NominalInputVoltage    float64   `json:"NominalInputVoltage"`
	BatteryVoltage         float64   `json:"BatteryVoltage"`
	NominalBatteryVoltage  float64   `json:"NominalBatteryVoltage"`
	NumberTransfers        float64   `json:"BatteryNumberTransfersTotal"`
	BatteryTimeLeftSeconds float64   `json:"BatteryTimeLeftSeconds"`
	BatteryTimeOnSeconds   float64   `json:"BatteryTimeOnSeconds"`
	LastTransferOnBattery  time.Time `json:"LastTransferOnBattery"`
	LastTransferOffBattery time.Time `json:"LastTransferOffBattery"`
	LastSelftest           time.Time `json:"LastSelftest"`
	NominalPowerWatts      float64   `json:"NominalPowerWatts"`
	Status                 string    `json:"status"`
	Error                  bool      `json:"error"`
}

func collect() *UPSCollector {
	client, err := apcupsd.Dial("tcp4", *apcupdsHostPort)
	if err != nil {
		log.Print("Errored: ", err)
		return &UPSCollector{Timestamp: time.Now(), Error: true}
	}

	s, err := client.Status()
	if err != nil {
		log.Print("Client collection error:", err)
		return &UPSCollector{Timestamp: time.Now(), Error: true}
	}

	res := &UPSCollector{
		Timestamp:              time.Now(),
		UPSName:                s.UPSName,
		UPSMode:                s.UPSMode,
		UPSModel:               s.Model,
		UPSLoadPercent:         s.LoadPercent,
		BatteryChargePercent:   s.BatteryChargePercent,
		LineVolts:              s.LineVoltage,
		NominalInputVoltage:    s.NominalInputVoltage,
		BatteryVoltage:         s.BatteryVoltage,
		NominalBatteryVoltage:  s.NominalBatteryVoltage,
		NumberTransfers:        float64(s.NumberTransfers),
		BatteryTimeLeftSeconds: s.TimeLeft.Seconds(),
		BatteryTimeOnSeconds:   s.TimeOnBattery.Seconds(),
		LastTransferOnBattery:  s.XOnBattery,
		LastTransferOffBattery: s.XOffBattery,
		LastSelftest:           s.LastSelftest,
		NominalPowerWatts:      float64(s.NominalPower),
		Status:                 s.Status,
		Error:                  false,
	}

	// Jearh, we can still talk to apcupsd, but things are not OK.
	if res.Status == "COMMLOST" {
		res.Error = true
	}

	return res
}

// Output json from the data collected from the client
func output(enc json.Encoder) {
	res := collect()
	enc.Encode(res)
}

func main() {
	enc := json.NewEncoder(os.Stdout)

	if *daemonMode {
		for {
			output(*enc)
			time.Sleep(time.Duration(*collectInterval) * time.Second)
		}
	} else {
		output(*enc)
	}

}
