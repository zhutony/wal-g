package testtools

import (
	"bytes"

	"github.com/wal-g/wal-g/internal"
)

// FileTarBallMaker creates a new FileTarBall
// with the directory that files should be
// extracted to.
type FileTarBallMaker struct {
	number int
	Size   *int64
	Out    string
}

// Make creates a new FileTarBall.
func (tarBallMaker *FileTarBallMaker) Make(inheritState bool) internal.TarBall {
	tarBallMaker.number++
	return &FileTarBall{
		number:          tarBallMaker.number,
		allTarballsSize: tarBallMaker.Size,
		out:             tarBallMaker.Out,
	}
}

type BufferTarBallMaker struct {
	number        int
	Size          *int64
	BufferToWrite *bytes.Buffer
}

func (tarBallMaker *BufferTarBallMaker) Make(dedicatedUploader bool) internal.TarBall {
	tarBallMaker.number++
	return &BufferTarBall{
		number:          tarBallMaker.number,
		allTarballsSize: tarBallMaker.Size,
		underlying:      tarBallMaker.BufferToWrite,
	}
}
