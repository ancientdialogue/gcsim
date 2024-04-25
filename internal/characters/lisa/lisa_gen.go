// Code generated by "pipeline"; DO NOT EDIT.
package lisa

import (
	_ "embed"

	"fmt"
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/keys"
	"github.com/genshinsim/gcsim/pkg/gcs/validation"
	"github.com/genshinsim/gcsim/pkg/model"
	"google.golang.org/protobuf/encoding/prototext"
	"slices"
)

//go:embed data_gen.textproto
var pbData []byte
var base *model.AvatarData
var paramKeysValidation = map[action.Action][]string{
	1: {"hold"},
}

func init() {
	base = &model.AvatarData{}
	err := prototext.Unmarshal(pbData, base)
	if err != nil {
		panic(err)
	}
	validation.RegisterCharParamValidationFunc(keys.Lisa, ValidateParamKeys)
}

func ValidateParamKeys(a action.Action, keys []string) error {
	valid, ok := paramKeysValidation[a]
	if !ok {
		return nil
	}
	for _, v := range keys {
		if !slices.Contains(valid, v) {
			return fmt.Errorf("key %v is invalid for action %v", v, a.String())
		}
	}
	return nil
}

func (x *char) Data() *model.AvatarData {
	return base
}

var (
	attack = [][]float64{
		attack_1,
		attack_2,
		attack_3,
		attack_4,
	}
)

var (
	// attack: attack_1 = [0]
	attack_1 = []float64{
		0.396,
		0.4257,
		0.4554,
		0.495,
		0.5247,
		0.5544,
		0.594,
		0.6336,
		0.6732,
		0.7128,
		0.753984,
		0.80784,
		0.861696,
		0.915552,
		0.969408,
	}
	// attack: attack_2 = [1]
	attack_2 = []float64{
		0.3592,
		0.38614,
		0.41308,
		0.449,
		0.47594,
		0.50288,
		0.5388,
		0.57472,
		0.61064,
		0.64656,
		0.683917,
		0.732768,
		0.781619,
		0.83047,
		0.879322,
	}
	// attack: attack_3 = [2]
	attack_3 = []float64{
		0.428,
		0.4601,
		0.4922,
		0.535,
		0.5671,
		0.5992,
		0.642,
		0.6848,
		0.7276,
		0.7704,
		0.814912,
		0.87312,
		0.931328,
		0.989536,
		1.047744,
	}
	// attack: attack_4 = [3]
	attack_4 = []float64{
		0.5496,
		0.59082,
		0.63204,
		0.687,
		0.72822,
		0.76944,
		0.8244,
		0.87936,
		0.93432,
		0.98928,
		1.046438,
		1.121184,
		1.19593,
		1.270675,
		1.345421,
	}
	// attack: charge = [4]
	charge = []float64{
		1.7712,
		1.90404,
		2.03688,
		2.214,
		2.34684,
		2.47968,
		2.6568,
		2.83392,
		3.01104,
		3.18816,
		3.372365,
		3.613248,
		3.854131,
		4.095014,
		4.335898,
	}
	// skill: skillHold = [0 1 2 3]
	skillHold = [][]float64{
		{
			3.2,
			3.44,
			3.68,
			4,
			4.24,
			4.48,
			4.8,
			5.12,
			5.44,
			5.76,
			6.08,
			6.4,
			6.8,
			7.2,
			7.6,
		},
		{
			3.68,
			3.956,
			4.232,
			4.6,
			4.876,
			5.152,
			5.52,
			5.888,
			6.256,
			6.624,
			6.992,
			7.36,
			7.82,
			8.28,
			8.74,
		},
		{
			4.24,
			4.558,
			4.876,
			5.3,
			5.618,
			5.936,
			6.36,
			6.784,
			7.208,
			7.632,
			8.056,
			8.48,
			9.01,
			9.54,
			10.07,
		},
		{
			4.872,
			5.2374,
			5.6028,
			6.09,
			6.4554,
			6.8208,
			7.308,
			7.7952,
			8.2824,
			8.7696,
			9.2568,
			9.744,
			10.353,
			10.962,
			11.571,
		},
	}
	// skill: skillPress = [5]
	skillPress = []float64{
		0.8,
		0.86,
		0.92,
		1,
		1.06,
		1.12,
		1.2,
		1.28,
		1.36,
		1.44,
		1.52,
		1.6,
		1.7,
		1.8,
		1.9,
	}
	// burst: burst = [0]
	burst = []float64{
		0.3656,
		0.39302,
		0.42044,
		0.457,
		0.48442,
		0.51184,
		0.5484,
		0.58496,
		0.62152,
		0.65808,
		0.69464,
		0.7312,
		0.7769,
		0.8226,
		0.8683,
	}
)
