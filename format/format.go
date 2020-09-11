package format

import (
	"github.com/kukayyou/avcodec/av/avutil"
	"github.com/kukayyou/avcodec/format/aac"
	"github.com/kukayyou/avcodec/format/flv"
	"github.com/kukayyou/avcodec/format/mp4"
	"github.com/kukayyou/avcodec/format/rtmp"
	"github.com/kukayyou/avcodec/format/rtsp"
	"github.com/kukayyou/avcodec/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
