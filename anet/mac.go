package anet

import (
    "net"
)

//获取本机mac地址
func GetMacList() (macList []string) {
    interfaces, err := net.Interfaces()
    if err != nil {
        return macList
    }

    for _, item := range interfaces {
        mac := item.HardwareAddr.String()
        if len(mac) == 0 {
            continue
        }
        macList = append(macList, mac)
    }
    return macList
}

//是否本机mac地址
func IsMac(mac string) bool {
    macList := GetMacList()
    for _, item := range macList {
        if item == mac {
            return true
        }
    }
    return false
}
