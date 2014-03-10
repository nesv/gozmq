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

import (
	"fmt"
)

func NewContext() *Context {
	var c = &Context{}
	return c
}

type Context struct{}

func (c *Context) NewSocket(t SocketType) (s *Socket, err error) {
	switch t {
	case PAIR, PUB, SUB, REQ, REP, DEALER, ROUTER, PULL, PUSH:
		// Set default socket options.
		opts := map[SocketOption]interface{}{
			SNDHWM:                  1000,
			RCVHWM:                  1000,
			AFFINITY:                0,
			RATE:                    100,
			RECOVERY_IVL:            10000,
			SNDBUF:                  0,
			RCVBUF:                  0,
			LINGER:                  -1,
			RECONNECT_IVL:           100,
			RECONNECT_IVL_MAX:       0,
			MAXMSGSIZE:              -1,
			MULTICAST_HOPS:          1,
			RCVTIMEO:                -1,
			SNDTIMEO:                -1,
			IPV4ONLY:                1,
			DELAY_ATTACH_ON_CONNECT: 0,
			ROUTER_MANDATORY:        0,
			XPUB_VERBOSE:            0,
			TCP_KEEPALIVE:           -1,
			TCP_KEEPALIVE_IDLE:      -1,
			TCP_KEEPALIVE_CNT:       -1,
			TCP_KEEPALIVE_INTVL:     -1}
		s = &Socket{Type: t, Options: opts}
	default:
		err = fmt.Errorf("unsupported socket type %v", t)
	}
	return
}
