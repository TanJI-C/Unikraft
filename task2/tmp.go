package main

import (
	"fmt"
	"io/ioutil"
	"libvirt.org/go/libvirt"
)

func main() {
	// vName := "test_UK"
	xmlFilePath := "/home/ubuntu/hw/UKraft/work/Unikraft/task2/00-hello-world/UK.xml"
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		fmt.Println("Failed to connect to libvirt:", err)
		return
	}
	defer conn.Close()

	// 创建新的虚拟机
	//通过文件路径读入XML文件
	xmlData, err := ioutil.ReadFile(xmlFilePath)
    if err != nil {
		fmt.Println("failed to read XML file: %v", err)
		return 
    }
	//使用defineXML函数定义虚拟机
    dom, err := conn.DomainDefineXML(string(xmlData))
    if err != nil {
		fmt.Println("failed to create domain: %v", err)
		return
    }

	// Create Domain
	err = dom.Create()
	if err != nil {
		fmt.Println("Failed to start domain:", err)
		return
	}

	// //Suspend Domain
	// err = dom.Suspend()
	// if err != nil {
	// 	fmt.Println("Failed to suspend domain:", err)
	// 	return
	// }
		
	// //ShutDown Domain
	// err = dom.ShutDown()
	// if err != nil {
	// 	fmt.Println("Failed to shutdown domain:", err)
	// 	return
	// }

	// //Resume Domain
	// err = dom.Resume()
	// if err != nil {
	// 	fmt.Println("Failed to resume domain:", err)
	// 	return
	// }

	// 获取所有开启的虚拟机
	// 	const (
	// 	CONNECT_LIST_DOMAINS_ACTIVE         = ConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_ACTIVE)
	// 	CONNECT_LIST_DOMAINS_INACTIVE       = ConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_INACTIVE)
	// 	CONNECT_LIST_DOMAINS_PERSISTENT     = ConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_PERSISTENT)
	// 	CONNECT_LIST_DOMAINS_TRANSIENT      = ConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_TRANSIENT)
	// 	CONNECT_LIST_DOMAINS_RUNNING        = ConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_RUNNING)
	// 	CONNECT_LIST_DOMAINS_PAUSED         = ConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_PAUSED)
	// 	CONNECT_LIST_DOMAINS_SHUTOFF        = ConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_SHUTOFF)
	// 	CONNECT_LIST_DOMAINS_OTHER          = ConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_OTHER)
	// 	CONNECT_LIST_DOMAINS_MANAGEDSAVE    = ConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_MANAGEDSAVE)
	// 	CONNECT_LIST_DOMAINS_NO_MANAGEDSAVE = ConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_NO_MANAGEDSAVE)
	// 	CONNECT_LIST_DOMAINS_AUTOSTART      = ConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_AUTOSTART)
	// 	CONNECT_LIST_DOMAINS_NO_AUTOSTART   = ConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_NO_AUTOSTART)
	// 	CONNECT_LIST_DOMAINS_HAS_SNAPSHOT   = ConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_HAS_SNAPSHOT)
	// 	CONNECT_LIST_DOMAINS_NO_SNAPSHOT    = ConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_NO_SNAPSHOT)
	// 	CONNECT_LIST_DOMAINS_HAS_CHECKPOINT = ConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_HAS_CHECKPOINT)
	// 	CONNECT_LIST_DOMAINS_NO_CHECKPOINT  = ConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_NO_CHECKPOINT)
	// 	)
	
	// doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE)
	// if err != nil {
	// 	fmt.Println("err", err)
	// }
	// fmt.Printf("%d running domains:\n", len(doms))
	// for _, dom := range doms {
		
	// }
	name, err := dom.GetName()
	fmt.Println("Domain name is", name)

	info, err := dom.GetInfo()	//return DomainInfo, err
	if err != nil {
		fmt.Println("Get Info of domain faild:", err)
	} else {
		// type DomainInfo struct {
		// 	State     DomainState		1 表示running
		// 	MaxMem    uint64
		// 	Memory    uint64
		// 	NrVirtCpu uint
		// 	CpuTime   uint64
		// }
		fmt.Println("GetInfo", info)
	}

	//内存利用率
	// get tag 4:剩余 & 5:总共
	meminfo, err := dom.MemoryStats(10, 0)	//return DomainMemoryStat[], err
	if err != nil {
		fmt.Println("Get memory states of domain faild:", err)
	} else {
		// type DomainMemoryStat struct {
		// 	Tag int32
		// 	Val uint64
		// }
		fmt.Println("MemoryStats", meminfo)
	}

	// func (d *Domain) GetFSInfo(flags uint32) ([]DomainFSInfo, error)
	fsInfo, err := dom.GetFSInfo(0)
	if err != nil {
		fmt.Println("Get File system information of domain faild:", err)
	} else {
		// type DomainFSInfo struct {
		// 	MountPoint string
		// 	Name       string
		// 	FSType     string
		// 	DevAlias   []string
		// }
		fmt.Println("FSInfo", fsInfo)
	}

	//磁盘利用率-错误值（都是最大值）
	blockinfo, err := dom.GetBlockInfo("hda", 0)
	if err != nil {
		fmt.Println("Get block information of domain faild:", err)
	} else {
		// type DomainBlockInfo struct {
		// 	Capacity   uint64
		// 	Allocation uint64
		// 	Physical   uint64
		// }
		fmt.Println("DiskStats", blockinfo)
	}


	// 获取虚拟机的CPU信息
	cpuStats, err := dom.GetCPUStats(0, 1, 0)
	if err != nil {
		fmt.Println("Failed to get domain CPU stats:", err)
	} else {
		// type DomainCPUStats struct {
		// 	CpuTimeSet    bool
		// 	CpuTime       uint64
		// 	UserTimeSet   bool
		// 	UserTime      uint64
		// 	SystemTimeSet bool
		// 	SystemTime    uint64
		// 	VcpuTimeSet   bool
		// 	VcpuTime      uint64
		// }
		fmt.Println("CPUStats", cpuStats)
	}



	//Destory Domain
	err = dom.Destroy()
	if err != nil {
		fmt.Println("Failed to stop domain:", err)
		return
	}	
	//Undefine Domain
	err = dom.Undefine()
	if err != nil {
		fmt.Println("Failed to suspend domain:", err)
		return
	}
	
}

