package selector

import (
	"github.com/RekrutPoleks/GoNEWS/internal/internalType/RSSlayout"
)

func InitSelectorAndPrepare(startValCache int64, idurl int) func(news []RSSlayout.StructNews) []RSSlayout.StructNews {
	timeCashe := startValCache
	return func(news []RSSlayout.StructNews) []RSSlayout.StructNews {
		dr := 0
		t := timeCashe

		for i := 0; i < len(news); i++ {
			if news[i].PubDate.Unix() > timeCashe {
				if news[i].PubDate.Unix() > t { //Если время в каналах не попорядку.
					t = news[i].PubDate.Unix()
				}
				dr = i
				news[i].SetChannel(idurl)
			}
		}
		if dr == 0 {
			return nil
		}
		timeCashe = t
		return news[:dr]
	}
}
