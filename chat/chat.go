package chat

import (
	"log"

	"golang.org/x/net/context"
)

// Server estructura del Servidor
type Server struct{}

// HolaPrro función que cumple con la interfase del msj
func (s *Server) HolaPrro(ctx context.Context, mssg *Message) (*Message, error) {

	log.Println("Ladrido recibido de un prro: ", mssg.Body)

	return &Message{Body: "Sssssakese pishi prro, oreleeee alv"}, nil
}
