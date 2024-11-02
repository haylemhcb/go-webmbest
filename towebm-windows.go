package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
        "io/ioutil"
        "os/exec"
)

func uploadFiles(w http.ResponseWriter, r *http.Request) {

        EndTask()

        DelFiles();

	r.ParseMultipartForm(10 * 1024 * 1024) // 500MB maximo

	files := r.MultipartForm.File["file-upload"]

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		dst, err := os.Create("uploads\\" + fileHeader.Filename)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer dst.Close()

		if _, err = io.Copy(dst, file); err != nil {
			fmt.Println(err)
			return
		}
	}

        ConvertVideos()

        http.Redirect(w, r, "/", http.StatusFound)

        defer EndTask()

}

func Stop(w http.ResponseWriter, r *http.Request) {
    EndTask()
}


func EndTask() {
        cmd := exec.Command("taskkill", "/IM", "ffmpeg.exe")
        cmd.Run()
}

func ConvertVideos() {
  files, err := ioutil.ReadDir("uploads")
      if err != nil {
          fmt.Println("Error al leer la carpeta:", err)
          return
      }

      fmt.Println("Archivos en la carpeta:")
      for _, file := range files {
          fmt.Println(file.Name())
          TransVid(file.Name())
      }
}


func DelFiles() {

  os.RemoveAll("static\\output")
  os.RemoveAll("uploads")
  os.MkdirAll("static\\output", 0755)
  os.MkdirAll("uploads", 0755)
}


func TransVid(nvid string) {

  cmd := exec.Command(".\\ffmpeg-7.0.1-amd64-static/ffmpeg", "-y", "-i", "uploads\\" + nvid, "-af", "afade,dynaudnorm,equalizer=f=1000:t=h:width_type=o:width=2:g=5,equalizer=f=100:t=1:width_type=o:width=2:g=-5,adeclip,afade,acompressor", "-vf", "unsharp,eq=brightness=0.03:saturation=1.3,fade", "-preset", "fast", "-b:a", "192k", "-s", "720x480", "static\\output\\" + nvid + ".webm")

  err := cmd.Run()
  if err != nil {
      fmt.Println("Error al ejecutar el comando:", err)
      return
  }
}

func main() {
        fs := http.FileServer(http.Dir("static"))
        http.Handle("/", fs)
	http.HandleFunc("/upload", uploadFiles)
	http.HandleFunc("/stop", Stop)
        fmt.Println("http://localhost:8070");
	http.ListenAndServe(":8072", nil)
}
