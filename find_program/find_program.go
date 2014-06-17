package main
import (
    "log"
    "os"
    "os/exec"
    "strings"
    "fmt"
)

func main(){
// Look for ffmpeg
    cmd := exec.Command("which", "ffmpeg")
    buf, err := cmd.Output()
    if err != nil {
        log.Fatal("ffmpeg wasn't found on your system, it is required to convert to mkv.\n" +
            "Temp file left on your hardrive:\n")
        os.Exit(1)
    }
    ffmpegPath := strings.Trim(string(buf), "\n")
    fmt.Println("\n")
    fmt.Println(ffmpegPath)
}
