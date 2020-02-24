package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	//serves the http endpoint
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

var listenerPort = os.Getenv("EXPORTERPORT")
var steamid = os.Getenv("STEAMID")
var appid = "730"
var steamkey = os.Getenv("STEAMKEY")
var api = "http://api.steampowered.com/ISteamUserStats/GetUserStatsForGame/v0002/?"

type CSGOstats struct {
	Playerstats Playerstats `json:"playerstats"`
}
type Stats struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
type Achievements struct {
	Name     string `json:"name"`
	Achieved int    `json:"achieved"`
}
type Playerstats struct {
	SteamID      string         `json:"steamID"`
	GameName     string         `json:"gameName"`
	Stats        []Stats        `json:"stats"`
	Achievements []Achievements `json:"achievements"`
}

func init() {
	log.SetLevel(log.DebugLevel)
}

func makeApiCall() (CSGOstats, error) {
	log.Info("Starting recieving JSON Data from API")
	url := fmt.Sprintf("%sappid=%s&key=%s&steamid=%s", api, appid, steamkey, steamid)
	response, err := http.Get(url)
	if err != nil {
		return CSGOstats{}, err
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return CSGOstats{}, err
	}
	var s CSGOstats

	err = json.Unmarshal(data, &s)
	if err != nil {
		return CSGOstats{}, err
	}
	return s, err
}

var (
	total_kills = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "CSGOstats_total_kills",
		Help: "The total number of Kills",
	})
)

var (
	total_deaths = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "CSGOstats_total_deaths",
		Help: "The total number of deaths",
	})
)

var (
	total_wins = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "CSGOstats_total_wins",
		Help: "The total number of wins",
	})
)

var (
	total_kills_deagle = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "CSGOstats_total_kills_deagle",
		Help: "The total number of kills with deagle",
	})
)

var (
	total_kills_awp = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "CSGOstats_total_kills_awp",
		Help: "The total number of kills with awp",
	})
)

var (
	total_kills_ak47 = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "CSGOstats_total_kills_ak47",
		Help: "The total number of Kills with an AK47",
	})
)

var (
	total_kills_aug = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "CSGOstats_total_kills_aug",
		Help: "The total number of kills with aug",
	})
)

var (
	total_kills_famas = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "CSGOstats_total_kills_famas",
		Help: "The total number of kills with famas",
	})
)

var (
	total_kills_headshot = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "CSGOstats_total_kills_headshot",
		Help: "The total number of kills with headshot",
	})
)

var (
	last_match_kills = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "CSGOstats_last_match_kills",
		Help: "The total number of kills in the last match",
	})
)

var (
	total_matches_won = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "CSGOstats_total_matches_won",
		Help: "The total number of matches won",
	})
)

var (
	total_matches_played = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "CSGOstats_total_matches_played",
		Help: "The total number of matches played",
	})
)

var (
	total_kills_sg556 = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "CSGOstats_total_kills_sg556",
		Help: "The total number of kills with sk556",
	})
)

var (
	total_kills_m4a1 = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "CSGOstats_total_kills_m4a1",
		Help: "The total number of kills with m4a1",
	})
)

func startEndpoint() {
	log.Info("Beginning to serve on port " + listenerPort)
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":"+listenerPort, nil)
}

func main() {
	go startEndpoint()
	for {
		stats, err := makeApiCall()
		if err != nil {
			log.Errorf("The HTTP request failed with error Code %s ", err) //TODO: error handling
			return
		}

		for _, stat := range stats.Playerstats.Stats {
			{
				if stat.Name == "total_kills" {
					total_kills.Set(float64(stat.Value))
					log.Debug(stat.Value)
				}
				if stat.Name == "total_deaths" {
					total_deaths.Set(float64(stat.Value))
					log.Debug(stat.Value)
				}
				if stat.Name == "total_wins" {
					total_wins.Set(float64(stat.Value))
					log.Debug(stat.Value)
				}
				if stat.Name == "total_kills_deagle" {
					total_kills_deagle.Set(float64(stat.Value))
					log.Debug(stat.Value)
				}
				if stat.Name == "total_kills_awp" {
					total_kills_awp.Set(float64(stat.Value))
					log.Debug(stat.Value)
				}
				if stat.Name == "total_kills_ak47" {
					total_kills_ak47.Set(float64(stat.Value))
					log.Debug(stat.Value)
				}
				if stat.Name == "total_kills_aug" {
					total_kills_aug.Set(float64(stat.Value))
					log.Debug(stat.Value)
				}
				if stat.Name == "total_kills_famas" {
					total_kills_famas.Set(float64(stat.Value))
					log.Debug(stat.Value)
				}
				if stat.Name == "total_kills_headshot" {
					total_kills_headshot.Set(float64(stat.Value))
					log.Debug(stat.Value)
				}
				if stat.Name == "last_match_kills" {
					last_match_kills.Set(float64(stat.Value))
					log.Debug(stat.Value)
				}
				if stat.Name == "total_matches_won" {
					total_matches_won.Set(float64(stat.Value))
					log.Debug(stat.Value)
				}
				if stat.Name == "total_matches_played" {
					total_matches_played.Set(float64(stat.Value))
					log.Debug(stat.Value)
				}
				if stat.Name == "total_kills_sg556" {
					total_kills_sg556.Set(float64(stat.Value))
					log.Debug(stat.Value)
				}
				if stat.Name == "total_kills_m4a1" {
					total_kills_m4a1.Set(float64(stat.Value))
					log.Debug(stat.Value)
				}
			}
		}
		time.Sleep(time.Minute)
		log.Debug("sleep 1 min")
	}
}
