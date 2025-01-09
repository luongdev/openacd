package types

type Status struct {
	criterion
}

func (s *Status) Type() CriterionType {
	return StatusCriterion
}

func NewStatus(name, displayName string, score float64, others ...float64) *Status {
	s := &Status{
		criterion: newCriterion(name, displayName, score, others...),
	}

	return s
}

var _ Criterion = (*Status)(nil)
