package main

import (
	"fmt"
	"log"
	"os"
  "time"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func main() {
  defer timer("main")()
	csvfile, err := os.Open("data/202206-divvy-tripdata.csv")
	if err != nil {
		log.Fatal(err)
	}
	df := dataframe.ReadCSV(csvfile)
	members_df := df.Filter(dataframe.F{Colname: "member_casual", Comparator: series.Eq, Comparando: "member"})
	station_groups := members_df.GroupBy("start_station_name")
	station_rides := station_groups.Aggregation([]dataframe.AggregationType{dataframe.Aggregation_COUNT}, []string{"ride_id"})
	sorted := station_rides.Arrange(dataframe.RevSort("ride_id_COUNT"))
	fmt.Println("df: ", sorted)
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
