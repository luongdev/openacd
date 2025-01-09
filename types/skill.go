package types

type Skill struct {
	criterion
}

func (s *Skill) Type() CriterionType {
	return SkillCriterion
}

func NewSkill(name, displayName string, score float64, others ...float64) *Skill {
	s := &Skill{
		criterion: newCriterion(name, displayName, score, others...),
	}

	return s
}

var _ Criterion = (*Skill)(nil)
