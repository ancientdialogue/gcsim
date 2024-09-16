package mualani

import (
	tmpl "github.com/genshinsim/gcsim/internal/template/character"
	"github.com/genshinsim/gcsim/internal/template/nightsoul"
	"github.com/genshinsim/gcsim/pkg/core"
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/info"
	"github.com/genshinsim/gcsim/pkg/core/keys"
	"github.com/genshinsim/gcsim/pkg/core/player/character"
	"github.com/genshinsim/gcsim/pkg/model"
)

func init() {
	core.RegisterCharFunc(keys.Mualani, NewChar)
}

type char struct {
	*tmpl.Character
	nightsoulState *nightsoul.State
	nightsoulSrc   int
	momentumStacks int
	lastStackFrame int
	momentumSrc    int
	a4Stacks       int
	c1Done         bool

	a1Count int
}

func NewChar(s *core.Core, w *character.CharWrapper, _ info.CharacterProfile) error {
	c := char{}
	c.Character = tmpl.NewWithWrapper(s, w)

	c.EnergyMax = 60
	c.NormalHitNum = normalHitNum
	c.SkillCon = 3
	c.BurstCon = 5
	c.HasArkhe = false

	w.Character = &c

	c.nightsoulState = nightsoul.New(s, w)

	return nil
}

func (c *char) Init() error {
	c.a4()

	if c.Base.Cons >= 1 {
		c.c1()
	}

	if c.Base.Cons >= 2 {
		c.c2()
	}

	if c.Base.Cons >= 4 {
		c.c4()
	}
	c.SetNumCharges(action.ActionAttack, 1)

	return nil
}

func (c *char) ActionReady(a action.Action, p map[string]int) (bool, action.Failure) {
	if a == action.ActionAttack && c.nightsoulState.HasBlessing() {
		if c.AvailableCDCharge[a] <= 0 {
			// TODO: Implement AttackCD warning
			return false, action.CharacterDeceased
		}
	}

	return c.Character.ActionReady(a, p)
}

func (c *char) Condition(fields []string) (any, error) {
	switch fields[0] {
	case "momentum":
		return c.momentumStacks, nil
	case "nightsoulpoints":
		return c.nightsoulState.Points(), nil
	default:
		return c.Character.Condition(fields)
	}
}

func (c *char) AnimationStartDelay(k model.AnimationDelayKey) int {
	if c.nightsoulState.HasBlessing() && c.momentumStacks >= 3 {
		switch k {
		case model.AnimationXingqiuN0StartDelay:
			return 33
		default:
			return 30
		}
	}
	return c.Character.AnimationStartDelay(k)
}