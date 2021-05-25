package adaptor

import "fmt"

type MusicPlayer interface {
  play(fileType string, fileName string)
}

type ExistPlayer struct{}

func (e *ExistPlayer) playMp3(fileName string) {
  fmt.Println("play mp3:", fileName)
}

func (e *ExistPlayer) playWma(fileName string) {
  fmt.Println("play wma :", fileName)
}

type PlayAdaptor struct {
  existPlayer ExistPlayer
}

func (player *PlayAdaptor) play(fileType string, fileName string) {
  switch fileType {
  case "mp3":
    player.existPlayer.playMp3(fileName)
  case "wma":
    player.existPlayer.playWma(fileName)
  default:
    fmt.Println("not supported.")
  }
}
