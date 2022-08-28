package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var Host_address string

func get_app_version() string {
	resp, err := http.Get(Host_address + "/api/v2/app/version")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	return string(body)
}

func get_webapiVersion() string {
	resp, err := http.Get(Host_address + "/api/v2/app/webapiVersion")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	return string(body)
}

func get_buildInfo() string {
	resp, err := http.Get(Host_address + "/api/v2/app/buildInfo")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	return string(body)
}

func get_torrent_info() []torrent_info {
	resp, err := http.Get(Host_address + "/api/v2/torrents/info")
	if err != nil {
		log.Fatalln(err)
	}
	var torrents []torrent_info
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&torrents)
	if err != nil {
		log.Fatalln(err)
	}

	for i, val := range torrents {
		fmt.Printf("%v Values are %v\t%v\n", i, val.Name, val.Downloaded)
	}

	return torrents
}

func search_torrents(search string) {
	url := fmt.Sprintf("%v/api/v2/search/start?pattern=%v&plugins=all&category=all",Host_address,search)
	method := "GET"
	client := &http.Client {}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	  }

	  req.Header.Add("Cookie", "SID=3aRrI2rKhoQ+1jNLF9eYY3myNvXzx6nl")
	  res, err := client.Do(req)
	  if err != nil {
		fmt.Println(err)
		return
	  }
	  defer res.Body.Close()

	  body, err := io.ReadAll(res.Body)
	  if err != nil {
		fmt.Println(err)
		return
	  }
	  fmt.Println(string(body))

}

type Config struct {
	Host_ip   string `json:"host_ip"`
	Host_port string `json:"host_port"`
}

type torrent_info struct {
	Added_on           int     `json:"added_on"`
	Amount_left        int     `json:"amount_left"`
	Auto_tmm           bool    `json:"auto_tmm"`
	Availability       float64 `json:"availability"`
	Category           string  `json:"category"`
	Completed          int     `json:"completed"`
	Completed_on       int     `json:"completed_on"`
	Content_path       string  `json:"content_path"`
	Dl_limit           int     `json:"dl_limit"`
	Dlspeed            int     `json:"dlspeed"`
	Downloaded         int     `json:"downloaded"`
	Downloaded_session int     `json:"downloaded_session"`
	Eta                int     `json:"eta"`
	F_l_piece_prio     bool    `json:"f_l_piece_prio"`
	Force_start        bool    `json:"force_start"`
	Hash               string  `json:"hash"`
	Last_activity      int     `json:"last_activity"`
	Magnet_uri         string  `json:"magnet_uri"`
	Max_ratio          float64 `json:"max_ratio"`
	Max_seeding_time   int     `json:"max_seeding_time"`
	Name               string  `json:"name"`
	Num_complete       int     `json:"num_complete"`
	Num_incomplete     int     `json:"num_incomplete"`
	Num_leechs         int     `json:"num_leechs"`
	Num_seeds          int     `json:"num_seeds"`
	Priority           int     `json:"priority"`
	Progress           float64 `json:"progress"`
	Ratio              float64 `json:"ratio"`
	Ratio_limit        float64 `json:"ratio_limit"`
	Save_path          string  `json:"save_path"`
	Seeding_time       int     `json:"seeding_time"`
	Seeding_time_limit int     `json:"seeding_time_limit"`
	Seen_complete      int     `json:"seen_complete"`
	Seq_dl             bool    `json:"seq_dl"`
	Size               int     `json:"size"`
	State              string  `json:"state"`
	Super_seeding      bool    `json:"super_seeding"`
	Tags               string  `json:"tags"`
	Time_active        int     `json:"time_active"`
	Total_size         int     `json:"total_size"`
	Tracker            string  `json:"tracker"`
	Up_limit           int     `json:"up_limit"`
	Uploaded           int     `json:"uploaded"`
	Uploaded_session   int     `json:"uploaded_session"`
	Upspeed            int     `json:"upspeed"`
}

type search_result struct {
	DescrLink  string      `json:"descrLink"`
	FileName   string      `json:"fileName"`
	FileSize   json.Number `json:"fileSize"`
	FileUrl    string      `json:"fileUrl"`
	NbLeechers json.Number `json:"nbLeechers"`
	NbSeeders  json.Number `json:"nbSeeders"`
	SiteUrl    string      `json:"siteUrl"`
}

func main() {
	configFile, err := os.ReadFile("config.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened config.json")
	var config Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		fmt.Printf("Could not unmarshal config.json file\n")
		log.Panic(err)
	}

	Host_address = config.Host_ip + ":" + config.Host_port
	fmt.Printf("The Host Ip address is: %v\n\n", Host_address)

	fmt.Printf("Application version:\t%v\n", get_app_version())
	fmt.Printf("Web api version:\t%v\n", get_webapiVersion())
	fmt.Printf("BuildInfo:\t%v\n", get_buildInfo())
	fmt.Print("\n######## Torrent information\n")

	get_torrent_info()

	search_torrents("House%20of%20Dragon")
}
