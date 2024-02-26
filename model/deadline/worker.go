package deadline

import (
	"time"
)

var Worker_Status = map[int]string{
	0: "Unknown",
	1: "Rendering",
	2: "Idle",
	3: "Offline",
	4: "Stalled",
	8: "StartingJob",
}

type Workers []struct {
	Info struct {
		Name           string    `json:"Name"`
		Config         string    `json:"Config"`
		Stat           int       `json:"Stat"`
		LastRenderTime any       `json:"LastRenderTime"`
		StatDate       time.Time `json:"StatDate"`
		OnTskComp      string    `json:"OnTskComp"`
		CompletedTasks int       `json:"CompletedTasks"`
		TskFail        int       `json:"TskFail"`
		Limits         []any     `json:"Limits"`
		RndTime        float64   `json:"RndTime"`
		Msg            string    `json:"Msg"`
		Port           int       `json:"Port"`
		Serv           bool      `json:"Serv"`
		WorkerFleetID  string    `json:"WorkerFleetId"`
		JobPlug        string    `json:"JobPlug"`
		JobUser        string    `json:"JobUser"`
		JobName        string    `json:"JobName"`
		JobID          string    `json:"JobId"`
		JobPri         int       `json:"JobPri"`
		JobPool        string    `json:"JobPool"`
		JobGrp         string    `json:"JobGrp"`
		TskName        string    `json:"TskName"`
		TskID          string    `json:"TskId"`
		TskProg        string    `json:"TskProg"`
		TskTime        string    `json:"TskTime"`
		TskStat        string    `json:"TskStat"`
		LicFree        bool      `json:"LicFree"`
		LicPerm        bool      `json:"LicPerm"`
		Ubl            bool      `json:"UBL"`
		LicMode        int       `json:"LicMode"`
		LicEx          int       `json:"LicEx"`
		LicErr         string    `json:"LicErr"`
		Lic            string    `json:"Lic"`
		Pools          string    `json:"Pools"`
		Grps           string    `json:"Grps"`
		Reg            string    `json:"Reg"`
		RAMFree        int64     `json:"RAMFree"`
		Disk           int64     `json:"Disk"`
		DiskStr        string    `json:"DiskStr"`
		ProcSpd        int       `json:"ProcSpd"`
		CPU            int       `json:"CPU"`
		NetSent        int       `json:"NetSent"`
		NetRecv        int       `json:"NetRecv"`
		Swap           int       `json:"Swap"`
		DiskRead       int       `json:"DiskRead"`
		DiskWrit       int       `json:"DiskWrit"`
		UpTime         int       `json:"UpTime"`
		Host           string    `json:"Host"`
		User           string    `json:"User"`
		IP             string    `json:"IP"`
		Mac            string    `json:"MAC"`
		Procs          int       `json:"Procs"`
		RAM            int64     `json:"RAM"`
		Arch           string    `json:"Arch"`
		Os             string    `json:"OS"`
		Ver            string    `json:"Ver"`
		Vid            string    `json:"Vid"`
		Awsp           bool      `json:"AWSP"`
		AWSInfo        struct {
		} `json:"AWSInfo"`
		ConcurrencyToken any    `json:"ConcurrencyToken"`
		ID               string `json:"_id"`
		ExtraElements    any    `json:"ExtraElements"`
	} `json:"Info"`
	Settings struct {
		Name   string   `json:"Name"`
		Pools  []string `json:"Pools"`
		Grps   []string `json:"Grps"`
		OptIns struct {
		} `json:"OptIns"`
		EventOI              []any   `json:"EventOI"`
		Mode                 string  `json:"Mode"`
		UsrNms               []any   `json:"UsrNms"`
		OvrSchd              bool    `json:"OvrSchd"`
		IdlStrt              bool    `json:"IdlStrt"`
		IdlMin               int     `json:"IdlMin"`
		IdlStp               bool    `json:"IdlStp"`
		IdlStpIfStrt         bool    `json:"IdlStpIfStrt"`
		FinIdlStp            bool    `json:"FinIdlStp"`
		CPUThreshOn          bool    `json:"CpuThreshOn"`
		CPUThresh            int     `json:"CpuThresh"`
		RAMPerThreshOn       bool    `json:"RamPerThreshOn"`
		RAMPerThresh         int     `json:"RamPerThresh"`
		RAMThreshOn          bool    `json:"RamThreshOn"`
		RAMThresh            int     `json:"RamThresh"`
		ProcOn               bool    `json:"ProcOn"`
		Proc                 []any   `json:"Proc"`
		IdleUserOn           bool    `json:"IdleUserOn"`
		IdleUsers            []any   `json:"IdleUsers"`
		Enable               bool    `json:"Enable"`
		WorkerStopType       int     `json:"WorkerStopType"`
		Cmmt                 string  `json:"Cmmt"`
		Desc                 string  `json:"Desc"`
		TskLmt               int     `json:"TskLmt"`
		NrmTime              float64 `json:"NrmTime"`
		NrmTimeout           float64 `json:"NrmTimeout"`
		AffinOvr             bool    `json:"AffinOvr"`
		Affin                []any   `json:"Affin"`
		GpuAffinOvr          bool    `json:"GpuAffinOvr"`
		GpuAffin             []any   `json:"GpuAffin"`
		Mac                  string  `json:"Mac"`
		HostMachineIPAddress string  `json:"HostMachineIPAddress"`
		UseListPort          bool    `json:"UseListPort"`
		ListPort             int     `json:"ListPort"`
		Cloud                bool    `json:"Cloud"`
		RegName              string  `json:"RegName"`
		GrpMapID             string  `json:"GrpMapID"`
		TmpDataPath          bool    `json:"TmpDataPath"`
		Ex0                  string  `json:"Ex0"`
		Ex1                  string  `json:"Ex1"`
		Ex2                  string  `json:"Ex2"`
		Ex3                  string  `json:"Ex3"`
		Ex4                  string  `json:"Ex4"`
		Ex5                  string  `json:"Ex5"`
		Ex6                  string  `json:"Ex6"`
		Ex7                  string  `json:"Ex7"`
		Ex8                  string  `json:"Ex8"`
		Ex9                  string  `json:"Ex9"`
		ExDic                struct {
		} `json:"ExDic"`
		ID string `json:"_id"`
	} `json:"Settings"`
}
