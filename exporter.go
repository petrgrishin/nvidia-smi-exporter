package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

const nvidiaSmiCmd = "nvidia-smi"

func metrics(response http.ResponseWriter, request *http.Request) {

	out, err := exec.Command(
		nvidiaSmiCmd,
		"--query-gpu=name,index,fan.speed,temperature.gpu,clocks.gr,clocks.sm,clocks.mem,power.draw,utilization.gpu,utilization.memory,memory.total,memory.free,memory.used",
		"--format=csv,noheader,nounits").Output()

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	csvReader := csv.NewReader(bytes.NewReader(out))
	csvReader.TrimLeadingSpace = true
	records, err := csvReader.ReadAll()

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	metricList := []string{
		"nvidia_fan_speed", "nvidia_temperature_gpu", "nvidia_clocks_gr", "nvidia_clocks_sm", "nvidia_clocks_mem", "nvidia_power_draw",
		"nvidia_utilization_gpu", "nvidia_utilization_memory", "nvidia_memory_total", "nvidia_memory_free", "nvidia_memory_used"}

	result := ""
	for _, row := range records {
		name := fmt.Sprintf("%s", row[0])
		index := fmt.Sprintf("%s", row[1])
		for idx, value := range row[2:] {
			result = fmt.Sprintf(
				"%s%s{gpu=\"%s\", name=\"%s\"} %s\n", result,
				metricList[idx], index, name, value)
		}
	}

	fmt.Fprintf(response, result)
}

func home(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "<html><head><title>Nvidia SMI Exporter</title></head><body><h1>Nvidia SMI Exporter</h1><p><a href=\"/metrics\">Metrics</a></p></body></html>")
}

func main() {
	addr := ":9102"
	if len(os.Args) > 1 {
		addr = ":" + os.Args[1]
	}

	http.HandleFunc("/metrics", metrics)
	http.HandleFunc("/", home)

	log.Printf("Listen port http://localhost%s\n", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
