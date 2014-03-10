/*
github.com/nesv/gozmq/zmtp -- A pure, ZeroMQ-compatible library for Go.
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

package zmtp

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
