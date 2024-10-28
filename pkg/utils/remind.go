package utils

import "time"

type Phase interface {
	GetPhase() int16
	Phase(phase int16) time.Duration
}

type phase struct {
	phase  int16
	phases []time.Duration
}

func (p *phase) GetPhase() int16 {
	return p.phase
}

func (p *phase) Phase(phase int16) time.Duration {
	if phase > BiggestPhase {
		phase = BiggestPhase
	}

	return p.phases[int(phase)]
}

var phaseObj = &phase{phases: []time.Duration{
	Phase0, Phase1, Phase2, Phase3, Phase4, Phase5, Phase6, Phase7, Phase8,
}}

const (
	Phase0 = 6 * time.Hour
	Phase1 = 12 * time.Hour
	Phase2 = 24 * time.Hour * 1
	Phase3 = 24 * time.Hour * 3
	Phase4 = 24 * time.Hour * 7
	Phase5 = 24 * time.Hour * 15
	Phase6 = 24 * time.Hour * 30
	Phase7 = 24 * time.Hour * 30 * 3
	Phase8 = 24 * time.Hour * 30 * 6

	BiggestPhase = 8
)

func GetRemind(phase int16, p ...Phase) int64 {
	var obj Phase
	obj = phaseObj
	if p != nil {
		obj = p[0]
	}
	return time.Now().Add(obj.Phase(phase)).Unix()
}
