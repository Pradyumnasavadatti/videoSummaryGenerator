package internals

import "sync"

var (
	AudioName string = "audio.wav"
	SummaryName string = "summary.txt"
	Wg sync.WaitGroup
	VideoType = "video"
	AudioType = "audio"
);