package mcping

import (
	"errors"
)

//Could not parse address
var ErrAddress = errors.New("mcping: could not parse address")

//Could not resolve address
var ErrResolve = errors.New("mcping: Could not resolve address")

//Could not connect to host
var ErrConnect = errors.New("mcping: Could not connect to host")

//Could not decode varint
var ErrVarint = errors.New("mcping: Could not decode varint")

//Response is too small
var ErrSmallPacket = errors.New("mcping: Response too small")

//Response is too large
var ErrBigPacket = errors.New("mcping: Response too large")

//Response packet incorrect
var ErrPacketType = errors.New("mcping: Response packet type incorrect")

//Timeout error
var ErrTimeout = errors.New("mcping: Timeout occured")
