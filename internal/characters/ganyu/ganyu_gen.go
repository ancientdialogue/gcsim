// Code generated by "pipeline"; DO NOT EDIT.
package ganyu

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
	3: {"travel"},
	7: {"hold", "travel", "weakspot"},
}

func init() {
	base = &model.AvatarData{}
	err := prototext.Unmarshal(pbData, base)
	if err != nil {
		panic(err)
	}
	validation.RegisterCharParamValidationFunc(keys.Ganyu, ValidateParamKeys)
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
		attack_5,
		attack_6,
	}
)

var (
	// attack: aim = [6]
	aim = []float64{
		0.4386,
		0.4743,
		0.51,
		0.561,
		0.5967,
		0.6375,
		0.6936,
		0.7497,
		0.8058,
		0.867,
		0.9282,
		0.9894,
		1.0506,
		1.1118,
		1.173,
	}
	// attack: attack_1 = [0]
	attack_1 = []float64{
		0.31734,
		0.34317,
		0.369,
		0.4059,
		0.43173,
		0.46125,
		0.50184,
		0.54243,
		0.58302,
		0.6273,
		0.678037,
		0.737705,
		0.797372,
		0.857039,
		0.922131,
	}
	// attack: attack_2 = [1]
	attack_2 = []float64{
		0.35604,
		0.38502,
		0.414,
		0.4554,
		0.48438,
		0.5175,
		0.56304,
		0.60858,
		0.65412,
		0.7038,
		0.760725,
		0.827669,
		0.894613,
		0.961556,
		1.034586,
	}
	// attack: attack_3 = [2]
	attack_3 = []float64{
		0.45494,
		0.49197,
		0.529,
		0.5819,
		0.61893,
		0.66125,
		0.71944,
		0.77763,
		0.83582,
		0.8993,
		0.972038,
		1.057577,
		1.143116,
		1.228655,
		1.321971,
	}
	// attack: attack_4 = [3]
	attack_4 = []float64{
		0.45494,
		0.49197,
		0.529,
		0.5819,
		0.61893,
		0.66125,
		0.71944,
		0.77763,
		0.83582,
		0.8993,
		0.972038,
		1.057577,
		1.143116,
		1.228655,
		1.321971,
	}
	// attack: attack_5 = [4]
	attack_5 = []float64{
		0.48246,
		0.52173,
		0.561,
		0.6171,
		0.65637,
		0.70125,
		0.76296,
		0.82467,
		0.88638,
		0.9537,
		1.030838,
		1.121551,
		1.212265,
		1.302979,
		1.401939,
	}
	// attack: attack_6 = [5]
	attack_6 = []float64{
		0.5762,
		0.6231,
		0.67,
		0.737,
		0.7839,
		0.8375,
		0.9112,
		0.9849,
		1.0586,
		1.139,
		1.231125,
		1.339464,
		1.447803,
		1.556142,
		1.67433,
	}
	// attack: ffa = [8]
	ffa = []float64{
		1.28,
		1.376,
		1.472,
		1.6,
		1.696,
		1.792,
		1.92,
		2.048,
		2.176,
		2.304,
		2.432,
		2.56,
		2.72,
		2.88,
		3.04,
	}
	// attack: ffb = [9]
	ffb = []float64{
		2.176,
		2.3392,
		2.5024,
		2.72,
		2.8832,
		3.0464,
		3.264,
		3.4816,
		3.6992,
		3.9168,
		4.1344,
		4.352,
		4.624,
		4.896,
		5.168,
	}
	// attack: fullaim = [7]
	fullaim = []float64{
		1.24,
		1.333,
		1.426,
		1.55,
		1.643,
		1.736,
		1.86,
		1.984,
		2.108,
		2.232,
		2.356,
		2.48,
		2.635,
		2.79,
		2.945,
	}
	// skill: lotus = [1]
	lotus = []float64{
		1.32,
		1.419,
		1.518,
		1.65,
		1.749,
		1.848,
		1.98,
		2.112,
		2.244,
		2.376,
		2.508,
		2.64,
		2.805,
		2.97,
		3.135,
	}
	// burst: shower = [0]
	shower = []float64{
		0.70272,
		0.755424,
		0.808128,
		0.8784,
		0.931104,
		0.983808,
		1.05408,
		1.124352,
		1.194624,
		1.264896,
		1.335168,
		1.40544,
		1.49328,
		1.58112,
		1.66896,
	}
)
