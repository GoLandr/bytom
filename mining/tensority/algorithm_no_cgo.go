// +build !cgo

package tensority

import (
	log "github.com/sirupsen/logrus"

	"github.com/bytom/protocol/bc"
)

func simdAlgorithm(bh, seed *bc.Hash) *bc.Hash {
	log.Warn("The SIMD-version is not supported on releases, please compile the lib according to README to enable this feature.")
	return legacyAlgorithm(bh, seed)
}