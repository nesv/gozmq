/*
github.com/nesv/gozmq -- A pure, ZeroMQ-compatible library for Go.
Copyright (C) 2014  Nick Saika

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package gozmq

type SendFlags int

const (
	DONTWAIT SendFlags = 1 << iota
	SNDMORE
)

type SocketOption int

const (
	SNDHWM SocketOption = 1 << iota
	RCVHWM
	AFFINITY
	SUBSCRIBE
	UNSUBSCRIBE
	IDENTITY
	RATE
	RECOVERY_IVL
	SNDBFU
	RCVBUF
	LINGER
	RECONNECT_IVL
	RECONNECT_IVL_MAX
	BACKLOG
	MAXMSGSIZE
	MULTICAST_HOPS
	RCVTIMEO
	SNDTIMEO
	IPV4ONLY
	DELAY_ATTACH_ON_CONNECT
	ROUTER_MANDATORY
	XPUB_VERBOSE
	TCP_KEEPALIVE
	TCP_KEEPALIVE_IDLE
	TCP_KEEPALIVE_CNT
	TCP_KEEPALIVE_INTVL
	TCP_ACCEPT_FILTER
)

type SocketType byte

const (
	PAIR   SocketType = 0x00
	PUB    SocketType = 0x01
	SUB    SocketType = 0x02
	REQ    SocketType = 0x03
	REP    SocketType = 0x04
	DEALER SocketType = 0x05
	ROUTER SocketType = 0x06
	PULL   SocketType = 0x07
	PUSH   SocketType = 0x08
)

type Socket struct {
	Type    SocketType
	Options map[SocketOption]interface{}
}

func Bind(net, addr string) error {
	return nil
}

func Connect(net, addr string) error {
	return nil
}

func Send(m []byte, flags SendFlags) (n int, err error) {
	return
}

func Recv() (m []byte, err error) {
	return
}

func SetSockOpt(op SocketOption, v interface{}) error {
	return nil
}

func GetSockOpt(op SocketOption) (v interface{}, err error) {
	return
}
