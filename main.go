package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/inhies/go-bytesize"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
	datastar "github.com/starfederation/datastar/sdk/go"
)

func main() {
	r := chi.NewMux()

	r.Get("/live/stats", func(w http.ResponseWriter, r *http.Request) {
		sse := datastar.NewSSE(w, r)
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-r.Context().Done():
				return
			case <-ticker.C:
				stats := collectStats()
				sse.MergeFragmentTempl(Stat(stats))
			}
		}
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		stats := collectStats()
		Main(stats).Render(r.Context(), w)
	})

	log.Printf("Starting server on :8088")
	http.ListenAndServe(":8088", r)
}

func collectStats() Stats {
	v, _ := mem.VirtualMemory()
	cp, _ := cpu.Times(false)
	var stats Stats
	stats.MemUsed = bytesize.New(float64(v.Used)).String()
	stats.MemTotal = bytesize.New(float64(v.Total)).String()
	stats.MemFree = bytesize.New(float64(v.Free)).String()
	stats.MemPercent = v.UsedPercent

	for _, c := range cp {
		stats.CpuSystem += int(c.System)
		stats.CpuUser += int(c.User)
		stats.CpuIdle += int(c.Idle)
	}

	stats.CpuSystem /= len(cp)
	stats.CpuUser /= len(cp)
	stats.CpuIdle /= len(cp)
	return stats
}
