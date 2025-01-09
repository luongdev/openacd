package types

type CriterionFactory interface {
	New(opts ...*CriterionOption) (Criterion, error)
}

type criterionFactory struct {
}

func (c *criterionFactory) New(opts ...*CriterionOption) (Criterion, error) {
	o := MergeOptions(opts...)
	scores := make([]float64, 0)
	if o.Weight > 0 {
		scores = append(scores, o.Weight)
	}

	if o.MaxScore > 0 {
		scores = append(scores, o.MaxScore)
	}

	switch o.Type {
	case StatusCriterion:
		return NewStatus(o.Name, o.DisplayName, o.Score, scores...), nil
	case SkillCriterion:
		return NewSkill(o.Name, o.DisplayName, o.Score, scores...), nil
	default:
		return nil, ErrUnsupportedCriterionType
	}
}

func NewCriterionFactory() CriterionFactory {
	f := &criterionFactory{}

	return f
}

var _ CriterionFactory = (*criterionFactory)(nil)
