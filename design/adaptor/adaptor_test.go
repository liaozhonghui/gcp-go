package adaptor

import "testing"

func TestAdaptor(t *testing.T) {
  adaptor := PlayAdaptor{}
  mp := []MusicPlayer{&adaptor}
  for _, player := range mp {
    player.play("mp3", "fantasy.")
    player.play("wma", "last dance")
    player.play("mp4", "megalobox")
  }
}
