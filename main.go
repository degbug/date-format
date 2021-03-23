package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"
)

const TIME_LAYOUT = "2006-01-02 15:04:05"

type Item struct {
	XMLName  xml.Name `xml:"item"`
	Uid      string   `xml:"uid,attr"`
	Arg      string   `xml:"arg,attr"`
	Title    string   `xml:"title"`
	Subtitle string   `xml:"subtitle"`
	Icon     string   `xml:"icon"`
}

type Items struct {
	XMLName xml.Name `xml:"items"`
	Items   []Item   `xml:"item"`
}

func main() {
	s := flag.String("dt", "dt", "dt")
	flag.Parse()

	t, e := parseTime(*s)
	if e != nil {
		log.Panic(e)
	}

	items := &Items{}
	items.Items = []Item{
		{
			Uid:   "1",
			Title: t.Format(TIME_LAYOUT),
			Arg:   t.Format(TIME_LAYOUT),
			Icon:  "icon.png",
		},
		{
			Uid:   "2",
			Title: strconv.Itoa(int(t.UnixNano() / 1000000)),
			Arg:   strconv.Itoa(int(t.UnixNano() / 1000000)),
			Icon:  "icon.png",
		},
	}

	out, _ := xml.MarshalIndent(items, "", "")
	fmt.Println(string(out))
}

func parseTime(s string) (time.Time, error) {
	if s == "now" {
		return time.Now(), nil
	}

	timestamp, err := strconv.Atoi(s)
	if err != nil {
		time1, err := time.Parse(TIME_LAYOUT, s)
		if err != nil {
			return time.Now(), err
		} else {
			return time1, nil
		}
	} else {
		if len(s) == 13 {
			return time.Unix(int64(timestamp)/1000, int64(timestamp)%1000*1000000), nil
		} else if len(s) == 12 {
			return time.Unix(int64(timestamp)/100, int64(timestamp)%100*10000000), nil
		} else if len(s) == 11 {
			return time.Unix(int64(timestamp)/10, int64(timestamp)%10*100000000), nil
		} else {
			return time.Unix(int64(timestamp), 0), nil
		}
	}
}
