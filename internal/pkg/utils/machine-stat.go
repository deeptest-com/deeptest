package commUtils

import (
	"github.com/aaronchen2k/deeptest/internal/pkg/domain"
	_floatUtils "github.com/aaronchen2k/deeptest/pkg/lib/float"
	logUtils "github.com/aaronchen2k/deeptest/pkg/lib/log"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	gonet "net"
	"strings"
	"time"
)

func GetMachineMetrics(prevDiskInfoMap *map[string]*int64, prevDiskTsMap *map[string]*int64,
	prevNetworkInfoMap *map[string]*int64, prevNetworkTsMap *map[string]*int64) *domain.MachineStat {
	data := new(domain.MachineStat)

	data.Name = GetHostName()
	data.Ip = GetIp()

	data.CpuUsage = GetCpuUsage()
	data.MemoryUsage = GetMemoryUsage()
	data.DiskUsages = GetDiskInfo(prevDiskInfoMap, prevDiskTsMap)
	data.NetworkUsages = GetNetworkInfo(prevNetworkInfoMap, prevNetworkTsMap)

	return data
}

func GetHostName() string {
	hostInfo, _ := host.Info()
	return hostInfo.Hostname
}

func GetIp() (ret string) {
	conn, err := gonet.Dial("udp", "8.8.8.8:53")
	if err != nil {
		logUtils.Errorf("udp error：%s", err.Error())
		return
	}

	localAddr := conn.LocalAddr().(*gonet.UDPAddr)
	ret = strings.Split(localAddr.String(), ":")[0]

	return
}

func GetCpuUsage() (ret float64) {
	percent, _ := cpu.Percent(time.Second, false) // false表示总使用率，true为单核

	ret = _floatUtils.PointNumb(percent[0], 2)

	return
}

func GetMemoryUsage() (ret float64) {
	memInfoVir, err := mem.VirtualMemory()
	if err != nil {
		return
	}

	ret = _floatUtils.PointNumb(memInfoVir.UsedPercent, 2)

	return
}

func GetNetworkInfo(prevInfoMap *map[string]*int64, prevTsMap *map[string]*int64) (ret map[string]float64) {
	ret = map[string]float64{}

	netIOs, _ := net.IOCounters(true)
	if netIOs == nil {
		return
	}

	for _, netIO := range netIOs {
		newTs := time.Now().UnixMilli()
		name := netIO.Name

		preBytes := (*prevInfoMap)[name]
		preTs := (*prevTsMap)[name]

		bytesNew := int64(netIO.BytesSent + netIO.BytesRecv)

		if preBytes != nil && preTs != nil {
			bytesDiff := bytesNew - *preBytes
			timeDiff := (newTs - *preTs) / 1000

			flow := float64(bytesDiff) / float64(timeDiff) / 1024
			flow = _floatUtils.PointNumb(flow, 2)

			if flow > 0 {
				ret[name] = _floatUtils.PointNumb(flow, 2)
			}
		}

		(*prevInfoMap)[name] = &bytesNew
		(*prevTsMap)[name] = &newTs
	}

	return
}

func GetDiskInfo(prevInfoMap *map[string]*int64, prevTsMap *map[string]*int64) (ret map[string]float64) {
	ret = map[string]float64{}

	infos, err := disk.IOCounters()
	if err != nil {
		return
	}

	for name, newInfo := range infos {
		newTs := time.Now().UnixMilli()

		preBytes := (*prevInfoMap)[name]
		preTs := (*prevTsMap)[name]

		bytesNew := int64(newInfo.ReadBytes + newInfo.WriteBytes)

		if preBytes != nil && preTs != nil {
			bytesDiff := bytesNew - *preBytes
			timeDiff := (newTs - *preTs) / 1000

			flow := float64(bytesDiff) / float64(timeDiff) / 1024 / 1024
			flow = _floatUtils.PointNumb(flow, 2)

			if flow > 0 {
				ret[name] = _floatUtils.PointNumb(flow, 2)
			}
		}

		(*prevInfoMap)[name] = &bytesNew
		(*prevTsMap)[name] = &newTs
	}

	return
}

func GetDiskTime(prevInfoMap *map[string]*disk.IOCountersStat, prevTsMap *map[string]*int64) (ret map[string]float64) {
	ret = map[string]float64{}

	infos, err := disk.IOCounters()
	if err != nil {
		return
	}

	for name, newInfo := range infos {
		newTs := time.Now().UnixMilli()

		preInfo := (*prevInfoMap)[name]
		preTs := (*prevTsMap)[name]

		if preInfo != nil && preTs != nil {
			tickDiff := newInfo.IoTime - preInfo.IoTime
			timeDiff := newTs - *preTs

			utilization := (float64(tickDiff) / float64(timeDiff)) * 100.0

			ret[newInfo.Name] = utilization
		}

		(*prevInfoMap)[name] = &newInfo
		(*prevTsMap)[name] = &newTs
	}

	return
}
