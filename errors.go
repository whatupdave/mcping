package mcping

import (
    "errors"
)


var resolveErr = errors.New("mcping: Could not resolve address")
var connectErr = errors.New("mcping: Could not connect to host")
var varintErr = errors.New("mcping: Could not decode varint")
var smallPacketErr = errors.New("mcping: Response too small")
var bigPacketErr = errors.New("mcping: Response too large")
var packetTypeErr = errors.New("mcping: Response packet type incorrect")