package oob

import (
    "os/exec"
)

const tool = "md5"

type OobPlugin struct{}

// 获取OOB信息
func (o OobPlugin) OOB() (string, error) {
    cmd := exec.Command(tool, "/var/log/system.log")
    out, err := cmd.CombinedOutput()
    if err != nil {
        return "", err
    }
    return string(out), nil
}

// 重启服务器
func (o OobPlugin) PowerReset() (string, error) {
    cmd := exec.Command(tool, "/var/log/install.log")
    out, err := cmd.CombinedOutput()
    if err != nil {
        return "", err
    }
    return string(out), nil
}
