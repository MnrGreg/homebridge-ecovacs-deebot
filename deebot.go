package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	vacbot "github.com/skburgart/go-vacbot"
)

const (
	CLEAN_MODE_AUTO        = "auto"
	CLEAN_MODE_EDGE        = "edge"
	CLEAN_MODE_SPOT        = "spot"
	CLEAN_MODE_SINGLE_ROOM = "single_room"
	CLEAN_MODE_STOP        = "stop"
)

func main() {

	deebotjson := `{
	"email": "",
	"password_hash": "",
	"device_id": "",
	"country": "us",
	"continent": "na",
	"lang": "en",
	"app_code": "i_eco_e",
	"app_version": "1.3.5",
	"channel": "c_googleplay",
	"device_type": "1",
	"timezone": "GMT-5",
	"realm": "ecouser.net"
}
`

	if len(os.Args) == 1 {
		fmt.Println("Usage: deebot [charge, clean-auto, clean-auto-strong, clean-border, clean-spot, clean-singleroom, status]\n")
		fmt.Println("Create /usr/local/etc/deebot.json with the following contents")
		fmt.Printf("%s", deebotjson)
		return
	}

	switch os.Args[1] {
	case "charge":
		v := vacbot.NewFromConfigFile("/usr/local/etc/deebot.json")
		v.Charge()
	case "clean-auto":
		v := vacbot.NewFromConfigFile("/usr/local/etc/deebot.json")
		v.CleanAuto()
	case "clean-auto-strong":
		v := vacbot.NewFromConfigFile("/usr/local/etc/deebot.json")
		v.CleanAutoStrong()
	case "clean-border":
		v := vacbot.NewFromConfigFile("/usr/local/etc/deebot.json")
		v.CleanBorder()
	case "clean-spot":
		v := vacbot.NewFromConfigFile("/usr/local/etc/deebot.json")
		v.CleanSpot()
	case "clean-singleroom":
		v := vacbot.NewFromConfigFile("/usr/local/etc/deebot.json")
		v.CleanSingleroom()
	case "status":
		v := vacbot.NewFromConfigFile("/usr/local/etc/deebot.json")
		callback := func(result interface{}, err error) {
			if result != nil {
				resultValue := reflect.ValueOf(result)
				_, ok := resultValue.Type().FieldByName("Query")
				if ok {
					status := reflect.ValueOf(result).FieldByName("Query").Bytes()
					statusString := string(status[:])
					if statusString != "" {
						if strings.Contains(statusString, CLEAN_MODE_STOP) {
							log.Println("DEEBOT STOPPED")
						} else {
							log.Println("DEEBOT RUNNING")
						}
					}
				}
			}
		}
		v.RecvHandler(callback)
		v.FetchCleanState()
		time.Sleep(2000 * time.Millisecond)
	default:
		fmt.Printf("%q is not valid command.\n", os.Args[1])
		os.Exit(2)
	}
}
