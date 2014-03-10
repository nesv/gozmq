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

import "errors"

// Various errors.
var (
	ErrAgain        = errors.New("message cannot be sent at the moment")
	ErrNotSupported = errors.New("operation not supported by this socket type")
	ErrFSM          = errors.New("socket is not in the appropriate state")
	ErrTerm         = errors.New("the context associated with this socket was terminated")
	ErrNotSock      = errors.New("the provided socket was invalid")
	ErrIntr         = errors.New("the operation was interrupted by delivery of a signal before the message was sent")
)
