package gofilehead

import (
	"encoding/hex"
	"strings"
)

// Maximum headersize to determine filetype.
// var maxHeaderSize = 312

// Check the specified bytes array against hex signature.
func MatchFileHead(bstream []byte, hexsig string) bool {
	// Encode byte stream to hex with max length of provided hexsig.
	hBStream := encodeHex(bstream[:len(hexsig)])
	// Check if hex string starts with hex signature.
	return strings.HasPrefix(hBStream, strings.ToLower(hexsig))
}

// Encode string to hex.
func encodeHex(b []byte) string {
	hstr := hex.EncodeToString(b)
	return hstr
}
