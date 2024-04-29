// Code generated by "pipeline"; DO NOT EDIT.
package clorinde

import (
	_ "embed"

	"github.com/genshinsim/gcsim/pkg/model"
	"google.golang.org/protobuf/encoding/prototext"
)

//go:embed data_gen.textproto
var pbData []byte
var base *model.AvatarData

func init() {
	base = &model.AvatarData{}
	err := prototext.Unmarshal(pbData, base)
	if err != nil {
		panic(err)
	}
}

func (x *char) Data() *model.AvatarData {
	return base
}

var (
	attack = [][][]float64{
		{attack_1},
		{attack_2},
		attack_3,
		attack_4,
		{attack_5},
	}
)

var (
	// attack: attack_1 = [0]
	attack_1 = []float64{
		0.540596,
		0.584598,
		0.6286,
		0.69146,
		0.735462,
		0.78575,
		0.854896,
		0.924042,
		0.993188,
		1.06862,
		1.144052,
		1.219484,
		1.294916,
		1.370348,
		1.44578,
	}
	// attack: attack_2 = [1]
	attack_2 = []float64{
		0.516284,
		0.558307,
		0.60033,
		0.660363,
		0.702386,
		0.750413,
		0.816449,
		0.882485,
		0.948521,
		1.020561,
		1.092601,
		1.16464,
		1.23668,
		1.308719,
		1.380759,
	}
	// attack: attack_3 = [2 3]
	attack_3 = [][]float64{
		{
			0.34185,
			0.369675,
			0.3975,
			0.43725,
			0.465075,
			0.496875,
			0.5406,
			0.584325,
			0.62805,
			0.67575,
			0.72345,
			0.77115,
			0.81885,
			0.86655,
			0.91425,
		},
		{
			0.34185,
			0.369675,
			0.3975,
			0.43725,
			0.465075,
			0.496875,
			0.5406,
			0.584325,
			0.62805,
			0.67575,
			0.72345,
			0.77115,
			0.81885,
			0.86655,
			0.91425,
		},
	}
	// attack: attack_4 = [4 5 6]
	attack_4 = [][]float64{
		{
			0.23134,
			0.25017,
			0.269,
			0.2959,
			0.31473,
			0.33625,
			0.36584,
			0.39543,
			0.42502,
			0.4573,
			0.48958,
			0.52186,
			0.55414,
			0.58642,
			0.6187,
		},
		{
			0.23134,
			0.25017,
			0.269,
			0.2959,
			0.31473,
			0.33625,
			0.36584,
			0.39543,
			0.42502,
			0.4573,
			0.48958,
			0.52186,
			0.55414,
			0.58642,
			0.6187,
		},
		{
			0.23134,
			0.25017,
			0.269,
			0.2959,
			0.31473,
			0.33625,
			0.36584,
			0.39543,
			0.42502,
			0.4573,
			0.48958,
			0.52186,
			0.55414,
			0.58642,
			0.6187,
		},
	}
	// attack: attack_5 = [9]
	attack_5 = []float64{
		0.900102,
		0.973366,
		1.04663,
		1.151293,
		1.224557,
		1.308288,
		1.423417,
		1.538546,
		1.653675,
		1.779271,
		1.904867,
		2.030462,
		2.156058,
		2.281653,
		2.407249,
	}
	// attack: charge = [10]
	charge = []float64{
		1.2814,
		1.3857,
		1.49,
		1.639,
		1.7433,
		1.8625,
		2.0264,
		2.1903,
		2.3542,
		2.533,
		2.7118,
		2.8906,
		3.0694,
		3.2482,
		3.427,
	}
	// attack: collision = [12]
	collision = []float64{
		0.639324,
		0.691362,
		0.7434,
		0.81774,
		0.869778,
		0.92925,
		1.011024,
		1.092798,
		1.174572,
		1.26378,
		1.352988,
		1.442196,
		1.531404,
		1.620612,
		1.70982,
	}
	// attack: highPlunge = [14]
	highPlunge = []float64{
		1.596762,
		1.726731,
		1.8567,
		2.04237,
		2.172339,
		2.320875,
		2.525112,
		2.729349,
		2.933586,
		3.15639,
		3.379194,
		3.601998,
		3.824802,
		4.047606,
		4.27041,
	}
	// attack: lowPlunge = [13]
	lowPlunge = []float64{
		1.278377,
		1.382431,
		1.486485,
		1.635134,
		1.739187,
		1.858106,
		2.02162,
		2.185133,
		2.348646,
		2.527025,
		2.705403,
		2.883781,
		3.062159,
		3.240537,
		3.418915,
	}
	// skill: arkheCD = [10]
	arkheCD = []float64{
		10,
		10,
		10,
		10,
		10,
		10,
		10,
		10,
		10,
		10,
		10,
		10,
		10,
		10,
		10,
	}
	// skill: arkheDamage = [9]
	arkheDamage = []float64{
		0.432,
		0.4644,
		0.4968,
		0.54,
		0.5724,
		0.6048,
		0.648,
		0.6912,
		0.7344,
		0.7776,
		0.8208,
		0.864,
		0.918,
		0.972,
		1.026,
	}
	// skill: healingBOL = [8]
	healingBOL = []float64{
		0.8,
		0.8,
		0.8,
		0.8,
		0.8,
		0.8,
		0.8,
		0.8,
		0.8,
		0.8,
		0.8,
		0.8,
		0.8,
		0.8,
		0.8,
	}
	// skill: skillBOLGain = [2]
	skillBOLGain = []float64{
		0.35,
		0.35,
		0.35,
		0.35,
		0.35,
		0.35,
		0.35,
		0.35,
		0.35,
		0.35,
		0.35,
		0.35,
		0.35,
		0.35,
		0.35,
	}
	// skill: skillEnhancedNA = [1]
	skillEnhancedNA = []float64{
		0.38786,
		0.41943,
		0.451,
		0.4961,
		0.52767,
		0.56375,
		0.61336,
		0.66297,
		0.71258,
		0.7667,
		0.82082,
		0.87494,
		0.92906,
		0.98318,
		1.0373,
	}
	// skill: skillLungeFullBOL = [6]
	skillLungeFullBOL = []float64{
		0.23392,
		0.25296,
		0.272,
		0.2992,
		0.31824,
		0.34,
		0.36992,
		0.39984,
		0.42976,
		0.4624,
		0.49504,
		0.52768,
		0.56032,
		0.59296,
		0.6256,
	}
	// skill: skillLungeFullBOLHeal = [7]
	skillLungeFullBOLHeal = []float64{
		1.1,
		1.1,
		1.1,
		1.1,
		1.1,
		1.1,
		1.1,
		1.1,
		1.1,
		1.1,
		1.1,
		1.1,
		1.1,
		1.1,
		1.1,
	}
	// skill: skillLungeLowBOL = [4]
	skillLungeLowBOL = []float64{
		0.439632,
		0.475416,
		0.5112,
		0.56232,
		0.598104,
		0.639,
		0.695232,
		0.751464,
		0.807696,
		0.86904,
		0.930384,
		0.991728,
		1.053072,
		1.114416,
		1.17576,
	}
	// skill: skillLungeLowBOLHeal = [5]
	skillLungeLowBOLHeal = []float64{
		1.04,
		1.04,
		1.04,
		1.04,
		1.04,
		1.04,
		1.04,
		1.04,
		1.04,
		1.04,
		1.04,
		1.04,
		1.04,
		1.04,
		1.04,
	}
	// skill: skillLungeNoBOL = [3]
	skillLungeNoBOL = []float64{
		0.329724,
		0.356562,
		0.3834,
		0.42174,
		0.448578,
		0.47925,
		0.521424,
		0.563598,
		0.605772,
		0.65178,
		0.697788,
		0.743796,
		0.789804,
		0.835812,
		0.88182,
	}
	// skill: skillNA = [0]
	skillNA = []float64{
		0.267632,
		0.289416,
		0.3112,
		0.34232,
		0.364104,
		0.389,
		0.423232,
		0.457464,
		0.491696,
		0.52904,
		0.566384,
		0.603728,
		0.641072,
		0.678416,
		0.71576,
	}
	// skill: skillStateDuration = [11]
	skillStateDuration = []float64{
		9,
		9,
		9,
		9,
		9,
		9,
		9,
		9,
		9,
		9,
		9,
		9,
		9,
		9,
		9,
	}
	// burst: burstBOL = [1]
	burstBOL = []float64{
		0.66,
		0.72,
		0.78,
		0.84,
		0.9,
		0.96,
		1.02,
		1.08,
		1.14,
		1.2,
		1.26,
		1.32,
		1.38,
		1.44,
		1.5,
	}
	// burst: burstDamage = [0]
	burstDamage = []float64{
		1.2688,
		1.36396,
		1.45912,
		1.586,
		1.68116,
		1.77632,
		1.9032,
		2.03008,
		2.15696,
		2.28384,
		2.41072,
		2.5376,
		2.6962,
		2.8548,
		3.0134,
	}
)

const (
	// a1: a1Duration = [0]
	a1Duration float64 = 15
	// a1: a1FlatDmg = [1]
	a1FlatDmg float64 = 1800
	// a1: a1PercentBuff = [2]
	a1PercentBuff float64 = 0.2
	// a4: a4CritBuff = [2]
	a4CritBuff float64 = 0.1
	// a4: a4Duration = [3]
	a4Duration float64 = 15
)
