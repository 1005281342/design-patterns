package facade

import "log"

type HomeTheater interface {
	WatchMovie(movie string)
	EndMovie()
}

type HomeTheaterFacade struct {
	amp       IAmplifier
	player    IStreamingPlayer
	popper    IPopcornPopper
	projector IProjector
	lights    ITheaterLights
	screen    IScreen
}

func NewHomeTheaterFacade(
	amp IAmplifier,
	player IStreamingPlayer,
	popper IPopcornPopper,
	projector IProjector,
	lights ITheaterLights,
	screen IScreen) *HomeTheaterFacade {
	return &HomeTheaterFacade{
		amp:       amp,
		player:    player,
		popper:    popper,
		projector: projector,
		lights:    lights,
		screen:    screen,
	}
}

func (h *HomeTheaterFacade) WatchMovie(movie string) {
	log.Println("Get ready to watch a movie...")
	h.popper.On()
	h.popper.Pop()
	h.lights.Dim(10)
	h.screen.Down()
	h.projector.On()
	h.projector.WideScreenMode()
	h.amp.On()
	h.amp.SetStreamingPlayer(h.player)
	h.amp.SetSurroundSound()
	h.amp.SetVolume(5)
	h.player.On()
	h.player.Play(movie)
}

func (h *HomeTheaterFacade) EndMovie() {
	log.Println("Shutting movie theater down...")
	h.popper.Off()
	h.lights.On()
	h.screen.Up()
	h.projector.Off()
	h.amp.Off()
	h.player.Stop()
	h.player.Off()
}
