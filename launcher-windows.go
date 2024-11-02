
package main

import (
  "os/exec"
)


func RunNav() {
  // exec.Command("taskkill", "/IM", "ffmpeg.exe").Run();
  exec.Command("C:\\Program Files (x86)\\Microsoft\\Edge\\Application\\msedge.exe", "--new-window", "--app=http://localhost:8072").Run();
  defer exec.Command("taskkill", "/IM", "ffmpeg.exe").Run();

}

func main() {
  go RunNav();
  exec.Command(".\\towebm-windows.exe").Run();
}
