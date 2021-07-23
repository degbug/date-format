package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"strconv"
	"time"
)

const TIME_LAYOUT = "2006-01-02 15:04:05"

var LOC, _ = time.LoadLocation("Asia/Shanghai")

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
	items := &Items{}
	if e != nil {
		items.Items = []Item{
			{
				Uid:   "com.degbug-0",
				Title: "error input ",
				Arg:   e.Error(),
				Icon:  "icon.png",
			},
		}
	} else {
		items.Items = []Item{
			{
				Uid:   "com.degbug-0",
				Title: t.Format(TIME_LAYOUT),
				Arg:   t.Format(TIME_LAYOUT),
				Icon:  "icon.png",
			},
			{
				Uid:      "com.degbug-1",
				Title:    strconv.Itoa(int(t.UnixNano() / 1000000)),
				Arg:      strconv.Itoa(int(t.UnixNano() / 1000000)),
				Subtitle: "毫秒",
				Icon:     "icon.png",
			},
			{
				Uid:      "com.degbug-2",
				Title:    strconv.Itoa(int(t.UnixNano() / 1000000000)),
				Arg:      strconv.Itoa(int(t.UnixNano() / 1000000000)),
				Subtitle: "秒",
				Icon:     "icon.png",
			},
		}
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
		time1, err := time.ParseInLocation(TIME_LAYOUT, s, LOC)
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
