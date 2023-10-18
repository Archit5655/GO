package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "trace",
	Short: "trace ip",
	Long:  `trace ip`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			for _, ip := range args {
				fmt.Println(ip)
				showdata(ip)
			}

		} else {
			fmt.Println("Please write the ip to trace")
		}
	},
}

func getdatA(url string) []byte {
	// url:="https://ipinfo.io/112.196.62.6/geo"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

type ip struct {
	Ip       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
	Readme   string `json:"readme"`
}

func showdata(ipp string) {
	url := "https://ipinfo.io/" + ipp + "/geo"
	data := getdatA(url)
	savedata := ip{}
	err := json.Unmarshal(data, &savedata)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data found")
	fmt.Printf("IP is %s\n  City is %s\n Region is %s\n Country is %s\n Postal code is %s\n Time zone is %s\n  ", savedata.Ip, savedata.City, savedata.Region, savedata.Country, savedata.Postal, savedata.Timezone)
}
