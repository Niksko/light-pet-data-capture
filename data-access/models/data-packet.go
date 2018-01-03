package models

import "time"

type DataPacket struct {
	Timestamp time.Time `sql:",pk"`
	ChipId uint32 `sql:",pk"`
	HumiditySampleRate uint32
	TemperatureSampleRate uint32
	AudioSampleRate uint32
	LightSampleRte uint32
	HumidityData []uint32
	TemperatureData []uint32
	AudioData []uint32
	LightData []uint32
}
