package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type thing struct {
    company_id string
    region string
    cars string
    order_date string
}

func count(idx []string) map[string]int {
    ordersPerCompany := make(map[string]int);
    for _, v := range idx {
        val, ok := ordersPerCompany[v];
        if ok {
            val++;
            ordersPerCompany[v] = val;
        }
        if !ok {
            ordersPerCompany[v] = 1;
        }
    }
    return ordersPerCompany;
}

func sortMap(orders map[string]int) []string{
    sortedVals := []string{}

    for k := range orders {
        sortedVals = append(sortedVals, k)
    }

    sort.SliceStable(sortedVals, func(i, j int) bool {
        return orders[sortedVals[i]] > orders[sortedVals[j]]
    })
    return sortedVals;
}
func main() {
    f, err := os.Open("car_orders.csv")
    if err != nil {
        log.Fatal(err)
    }

    defer f.Close();

    csvReader := csv.NewReader(f);

    data, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }
    idx := parseId(data);
    cars := parseCar(data)
    result := count(idx);
    res := count(cars)
    sortRes := sortMap(res)
    sortedRes := sortMap(result);
    for _, k := range sortedRes{
        fmt.Printf("%q %v\n",k, result[k]);
    }
    for _, k := range sortRes{
        fmt.Printf("%q %v\n",k, res[k]);
    }
}

func parseId(data [][]string) []string {
    parts := []string{};
    for i, line := range data {
        if i > 0 {
            var record []string;
            for _, field := range line {
                record = strings.Split(field, ";")
            }
            parts = append(parts, record[0]);
        }
    }
    return parts;
}

func parseCar(data [][]string) []string {
    parts := []string{};
    for i, line := range data {
        if i > 0 {
            var record []string;
            for _, field := range line {
                record = strings.Split(field, ";")
            }
            parts = append(parts, record[2]);
        }
    }
    return parts;
}
