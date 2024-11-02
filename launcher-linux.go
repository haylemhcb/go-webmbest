
package main

import (
  "os/exec"
)


func RunNav() {
  exec.Command("/usr/bin/pkill", "ffmpeg").Run();
  exec.Command("x-www-browser", "http://localhost:8072").Run();
  defer exec.Command("/usr/bin/pkill", "ffmpeg").Run();
  defer exec.Command("/usr/bin/pkill", "tomp4").Run();

}

func main() {
  go RunNav();
  exec.Command("/usr/bin/pkill","towebm").Run();
  exec.Command("./towebm").Run();
}
