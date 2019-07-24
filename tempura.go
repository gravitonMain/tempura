package main

import (
    "fmt"
    "log"
    "flag"
    "time"
    "github.com/oltoko/go-am2320"
)

func main() {
    endless, count, sleep, verbose, separater, header, logHeader := args()
    
    sensor := am2320.Create(am2320.DefaultI2CAddr)
    sensor.Read() // trash data in first time
    
    if header {
        if logHeader {
            fmt.Print("time", separater)
        }
        fmt.Printf("Temperature%sHumidity\n", separater)
    }
    counter := 0
    for count != counter{
        value, err := sensor.Read()
        if err != nil {
            log.Fatalln("Failed to read from Sensor", err)
        }
	if logHeader {
            fmt.Print(time.Now().Unix(), separater)
        }
        fmt.Println(tempuraMessage(value, verbose, separater))
        time.Sleep(time.Duration(sleep) * time.Second)
	if !endless {
            counter++
        }
    }
}

func args()(bool, int, int, bool, string, bool, bool){
    var (
        endless = flag.Bool("e", false, "endless flag (default false)")
        count = flag.Int("c", 1, "request count (default 1)")
        sleep = flag.Int("i", 1, "request interval second (default 1)")
        verbose = flag.Bool("v", false, "verbose (default false)")
        separater = flag.String("s", "\t", "separater (default tab)")
        header = flag.Bool("h", false, "show header (default false)")
        logHeader = flag.Bool("l", false, "show timestamp (default false)")
    )
    flag.Parse()
    return *endless, *count, *sleep, *verbose, *separater, *header, *logHeader
}

func tempuraMessage(value *am2320.SensorValues, verbose bool, separater string) string {
    if verbose {
        return fmt.Sprintf("%.1f℃%s%.1f％", value.Temperature, separater, value.Humidity)
    } else {
        return fmt.Sprintf("%.1f%s%.1f", value.Temperature, separater, value.Humidity)
    }
}
