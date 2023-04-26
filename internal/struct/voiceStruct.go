package _struct

type VoiceCallData struct {
	Country             string  `json:"country"`
	Bandwidth           int     `json:"bandwidth"`
	ResponseTime        int     `json:"response_time"`
	Provider            string  `json:"provider"`
	ConnectionStability float32 `json:"connection_stability"`
	TTFB                int     `json:"ttfb"`
	VoiceClarity        int     `json:"voice_purity"`
	MedianCallTime      int     `json:"median_of_call_time"`
}

var VoiceOperators = []string{"TransparentCalls", "E-Voice", "JustPhone"}
