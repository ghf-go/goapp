package linuxservice

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ghf-go/goapp/base/gpath"
)

func installSystemd(path string) {
	bname := gpath.GetFileNameWithoutExt(path)
	content := fmt.Sprintf(`[Unit]
Description=%s
Documentation=%s
After=network.target
[Service]
Type=simple
ExecStart=%s
ExecStop=/bin/kill -SIGTERM $MAINPID
TimeoutStopSec=30
Restart=always
RestartSec=5s  
TimeoutStopSec=30  
[Install]
WantedBy=multi-user.target
Alias=%s.service`, bname, bname, path, bname)
	os.WriteFile(fmt.Sprintf("/etc/systemd/system/%s.service", bname), []byte(content), 0777)
	exec.Command("systemctl", "daemon-reload").Run()
	exec.Command("systemctl", "enable", "--now", bname).Run()
	exec.Command("systemctl", "start", bname).Run()
}

func uninstallSystemd(path string) {
	bname := gpath.GetFileNameWithoutExt(path)
	exec.Command("systemctl", "stop", bname).Run()
	os.Remove(fmt.Sprintf("/etc/systemd/system/%s.service", bname))
	exec.Command("systemctl", "daemon-reload").Run()
}
