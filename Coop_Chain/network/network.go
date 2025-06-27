package network

import "errors"

// Peer — интерфейс для представления пира в p2p-сети
// В MVP все методы возвращают ошибку 'not implemented'
type Peer interface {
	ID() string
	Send(msg []byte) error
	Close() error
}

// Broadcaster — интерфейс для широковещательной рассылки сообщений
// В MVP все методы возвращают ошибку 'not implemented'
type Broadcaster interface {
	Broadcast(msg []byte) error
}

// NotImplementedError — общая ошибка для не реализованных методов
var NotImplementedError = errors.New("not implemented")

// StubPeer — заглушка для Peer
type StubPeer struct{}

func (p *StubPeer) ID() string            { return "" }
func (p *StubPeer) Send(msg []byte) error { return NotImplementedError }
func (p *StubPeer) Close() error          { return NotImplementedError }

// StubBroadcaster — заглушка для Broadcaster
type StubBroadcaster struct{}

func (b *StubBroadcaster) Broadcast(msg []byte) error { return NotImplementedError }
