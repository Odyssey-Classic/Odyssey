package game

import "github.com/Odyssey-Classic/Odyssey/server/internal/services/network"

type Player struct {
	client *network.Client
}

func (p *Player) Send(msg any) {

}
