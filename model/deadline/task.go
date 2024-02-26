package deadline

import "time"

var Task_Status = map[int]string{
	1: "Unknown",
	2: "Queued",
	3: "Suspended",
	4: "Rendering",
	5: "Completed",
	6: "Failed",
	8: "Pending",
}

type Tasks struct {
	ID      string `json:"ID"`
	PreTask any    `json:"PreTask"`
	Tasks   []struct {
		JobID      string    `json:"JobID"`
		TaskID     int       `json:"TaskID"`
		Frames     string    `json:"Frames"`
		Slave      string    `json:"Slave"`
		Stat       int       `json:"Stat"`
		Prog       string    `json:"Prog"`
		RndStat    string    `json:"RndStat"`
		Errs       int       `json:"Errs"`
		Start      time.Time `json:"Start"`
		StartRen   time.Time `json:"StartRen"`
		Comp       time.Time `json:"Comp"`
		WtgStrt    bool      `json:"WtgStrt"`
		NormMult   float64   `json:"NormMult"`
		Size       int       `json:"Size"`
		RAMPeak    int64     `json:"RamPeak"`
		RAMPeakPer int       `json:"RamPeakPer"`
		CPU        int       `json:"Cpu"`
		CPUPer     int       `json:"CpuPer"`
		RAMAvg     int       `json:"RamAvg"`
		RAMAvgPer  int       `json:"RamAvgPer"`
		SwapAvg    int       `json:"SwapAvg"`
		SwapPeak   int       `json:"SwapPeak"`
		UsedClock  int       `json:"UsedClock"`
		TotalClock int       `json:"TotalClock"`
		Props      struct {
			Ex0   string `json:"Ex0"`
			Ex1   string `json:"Ex1"`
			Ex2   string `json:"Ex2"`
			Ex3   string `json:"Ex3"`
			Ex4   string `json:"Ex4"`
			Ex5   string `json:"Ex5"`
			Ex6   string `json:"Ex6"`
			Ex7   string `json:"Ex7"`
			Ex8   string `json:"Ex8"`
			Ex9   string `json:"Ex9"`
			ExDic struct {
			} `json:"ExDic"`
		} `json:"Props"`
		ConcurrencyToken any    `json:"ConcurrencyToken"`
		ID               string `json:"_id"`
		ExtraElements    any    `json:"ExtraElements"`
	} `json:"Tasks"`
	PostTask any `json:"PostTask"`
}
