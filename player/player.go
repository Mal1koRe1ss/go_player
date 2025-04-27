package player

import (
	"fmt"
	"os"
	"time"
	
	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

type MusicPlayer struct {
	Ctrl     *beep.Ctrl
	Streamer beep.StreamSeekCloser
	Done     chan bool
}

func NewMusicPlayer(filePath string) (*MusicPlayer, error) {
	file, err := openFile(filePath) // * reading file here...
	if err != nil {
		return nil, err
	}

	streamer, format, err := mp3.Decode(file) // * decoding file for opening ...
	if err != nil {
		return nil, fmt.Errorf("MP3 decode error: %w", err) 
	}

	err = initSpeaker(format.SampleRate) // * speaker method for alsa(pipewire)
	if err != nil {
		return nil, err
	}

	return &MusicPlayer{
		Ctrl:     &beep.Ctrl{Streamer: beep.Loop(1, streamer)},// * 1=1time only 0=inf
		Streamer: streamer,
		Done:     make(chan bool),
	}, nil
}

func (p *MusicPlayer) Play() {
	speaker.Play(beep.Seq(p.Ctrl, beep.Callback(func() {
		// * beep.sqe : ses akisi
		// * p.ctrl : ana ses akisi
		// * beep.callback : bitis callbacki
		// * p.Done : programa true gondererek sonlandiriyor.		
		p.Done <- true
	})))
}

func (p *MusicPlayer) Pause() {
	p.Ctrl.Paused = true // * controller
}

func (p *MusicPlayer) Resume() {
	p.Ctrl.Paused = false // * controller
}

func (p *MusicPlayer) Stop() {
	speaker.Clear()
	p.Streamer.Close()
	close(p.Done)
}

func openFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("File open error: %w", err)
	}
	return file, nil
}

func initSpeaker(sampleRate beep.SampleRate) error {
	return speaker.Init(sampleRate, sampleRate.N(time.Second/10)) // * speakeri tanimlama
}