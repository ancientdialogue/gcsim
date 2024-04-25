// Code generated by "pipeline"; DO NOT EDIT.
package jean

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
	2: {"enter", "enter_delay"},
}

func init() {
	base = &model.AvatarData{}
	err := prototext.Unmarshal(pbData, base)
	if err != nil {
		panic(err)
	}
	validation.RegisterCharParamValidationFunc(keys.Jean, ValidateParamKeys)
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
	auto = [][]float64{
		auto_1,
		auto_2,
		auto_3,
		auto_4,
		auto_5,
	}
)

var (
	// attack: auto_1 = [0]
	auto_1 = []float64{
		0.48332,
		0.52266,
		0.562,
		0.6182,
		0.65754,
		0.7025,
		0.76432,
		0.82614,
		0.88796,
		0.9554,
		1.032675,
		1.12355,
		1.214426,
		1.305301,
		1.404438,
	}
	// attack: auto_2 = [1]
	auto_2 = []float64{
		0.4558,
		0.4929,
		0.53,
		0.583,
		0.6201,
		0.6625,
		0.7208,
		0.7791,
		0.8374,
		0.901,
		0.973875,
		1.059576,
		1.145277,
		1.230978,
		1.32447,
	}
	// attack: auto_3 = [2]
	auto_3 = []float64{
		0.60286,
		0.65193,
		0.701,
		0.7711,
		0.82017,
		0.87625,
		0.95336,
		1.03047,
		1.10758,
		1.1917,
		1.288088,
		1.401439,
		1.514791,
		1.628143,
		1.751799,
	}
	// attack: auto_4 = [3]
	auto_4 = []float64{
		0.65876,
		0.71238,
		0.766,
		0.8426,
		0.89622,
		0.9575,
		1.04176,
		1.12602,
		1.21028,
		1.3022,
		1.407525,
		1.531387,
		1.655249,
		1.779112,
		1.914234,
	}
	// attack: auto_5 = [4]
	auto_5 = []float64{
		0.79206,
		0.85653,
		0.921,
		1.0131,
		1.07757,
		1.15125,
		1.25256,
		1.35387,
		1.45518,
		1.5657,
		1.692338,
		1.841263,
		1.990189,
		2.139115,
		2.301579,
	}
	// attack: charge = [5]
	charge = []float64{
		1.62024,
		1.75212,
		1.884,
		2.0724,
		2.20428,
		2.355,
		2.56224,
		2.76948,
		2.97672,
		3.2028,
		3.46185,
		3.766493,
		4.071136,
		4.375778,
		4.708116,
	}
	// skill: skill = [0]
	skill = []float64{
		2.92,
		3.139,
		3.358,
		3.65,
		3.869,
		4.088,
		4.38,
		4.672,
		4.964,
		5.256,
		5.548,
		5.84,
		6.205,
		6.57,
		6.935,
	}
	// burst: burst = [0]
	burst = []float64{
		4.248,
		4.5666,
		4.8852,
		5.31,
		5.6286,
		5.9472,
		6.372,
		6.7968,
		7.2216,
		7.6464,
		8.0712,
		8.496,
		9.027,
		9.558,
		10.089,
	}
	// burst: burstDotHealFlat = [5]
	burstDotHealFlat = []float64{
		154.03249,
		169.43788,
		186.12704,
		204.1,
		223.35674,
		243.89726,
		265.72156,
		288.82965,
		313.2215,
		338.89716,
		365.8566,
		394.0998,
		423.6268,
		454.4376,
		486.53214,
	}
	// burst: burstDotHealPer = [4]
	burstDotHealPer = []float64{
		0.2512,
		0.27004,
		0.28888,
		0.314,
		0.33284,
		0.35168,
		0.3768,
		0.40192,
		0.42704,
		0.45216,
		0.47728,
		0.5024,
		0.5338,
		0.5652,
		0.5966,
	}
	// burst: burstEnter = [1]
	burstEnter = []float64{
		0.784,
		0.8428,
		0.9016,
		0.98,
		1.0388,
		1.0976,
		1.176,
		1.2544,
		1.3328,
		1.4112,
		1.4896,
		1.568,
		1.666,
		1.764,
		1.862,
	}
	// burst: burstInitialHealFlat = [3]
	burstInitialHealFlat = []float64{
		1540.3248,
		1694.3788,
		1861.2705,
		2041,
		2233.5674,
		2438.9727,
		2657.2156,
		2888.2964,
		3132.215,
		3388.9717,
		3658.566,
		3940.998,
		4236.268,
		4544.376,
		4865.3213,
	}
	// burst: burstInitialHealPer = [2]
	burstInitialHealPer = []float64{
		2.512,
		2.7004,
		2.8888,
		3.14,
		3.3284,
		3.5168,
		3.768,
		4.0192,
		4.2704,
		4.5216,
		4.7728,
		5.024,
		5.338,
		5.652,
		5.966,
	}
)
