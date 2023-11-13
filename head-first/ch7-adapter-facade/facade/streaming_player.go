package facade

import "log"

type IStreamingPlayer interface {
	On()
	Off()
	Play(movie string)
	Stop()
}

type StreamingPlayer struct {
	movie string
}

func NewStreamingPlayer() *StreamingPlayer {
	return &StreamingPlayer{}
}

func (s *StreamingPlayer) On() {
	log.Println("Streaming Player on")
}

func (s *StreamingPlayer) Off() {
	log.Println("Streaming Player off")
}

func (s *StreamingPlayer) Play(movie string) {
	s.movie = movie
	log.Printf(`Streaming Player playing "%s"%s`, s.movie, "\n")
}

func (s *StreamingPlayer) Stop() {
	log.Printf(`Streaming Player stopped "%s"%s`, s.movie, "\n")
}
