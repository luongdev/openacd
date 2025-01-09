package types

import (
	"errors"
	"math"
	"reflect"
)

var ErrUnsupportedCriterionType = errors.New("unsupported criterion type")

type CriterionType int

const (
	StatusCriterion CriterionType = 0
	SkillCriterion  CriterionType = 1
)

type CriterionOption struct {
	Name        string        `json:"name"`
	DisplayName string        `json:"displayName"`
	Score       float64       `json:"score"`
	Weight      float64       `json:"weight"`
	MaxScore    float64       `json:"maxScore"`
	Type        CriterionType `json:"type"`
}

func WithName(name string) *CriterionOption {
	return &CriterionOption{Name: name}
}

func WithDisplayName(displayName string) *CriterionOption {
	return &CriterionOption{DisplayName: displayName}
}

func WithScore(score float64) *CriterionOption {
	return &CriterionOption{Score: score}
}

func WithWeight(weight float64) *CriterionOption {
	return &CriterionOption{Weight: weight}
}

func WithMaxScore(maxScore float64) *CriterionOption {
	return &CriterionOption{MaxScore: maxScore}
}

func WithType(t CriterionType) *CriterionOption {
	return &CriterionOption{Type: t}
}

func MergeOptions(opts ...*CriterionOption) *CriterionOption {
	if len(opts) == 1 {
		return opts[0]
	}

	c := &CriterionOption{}
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		optValue := reflect.ValueOf(opt).Elem()
		cValue := reflect.ValueOf(c).Elem()
		for i := 0; i < optValue.NumField(); i++ {
			field := optValue.Field(i)
			fieldType := optValue.Type().Field(i)
			if field.CanSet() && fieldType.PkgPath == "" && !field.IsZero() {
				cValue.Field(i).Set(field)
			}
		}
	}

	return c
}

type Criterion interface {
	Name() string
	DisplayName() string
	Type() CriterionType
	Weight() float64
	MaxScore() float64
	CalculateScore(i ...float64) float64
}

type criterion struct {
	name        string
	displayName string
	weight      float64
	score       float64
	maxScore    float64
}

func newCriterion(name, displayName string, score float64, others ...float64) criterion {
	c := criterion{
		name:        name,
		displayName: displayName,
		score:       score,
	}

	if len(others) > 0 {
		c.weight = others[0]
	}

	if len(others) > 1 {
		c.maxScore = others[1]
	}

	return c
}

func (c *criterion) Name() string {
	return c.name
}

func (c *criterion) DisplayName() string {
	return c.displayName
}

func (c *criterion) Weight() float64 {
	return c.weight
}

func (c *criterion) MaxScore() float64 {
	return c.maxScore
}

func (c *criterion) CalculateScore(i ...float64) float64 {
	applyMaxScore := func(score float64) float64 {
		if c.maxScore > 0 && score > c.maxScore {
			return c.maxScore
		}

		return score
	}

	if len(i) == 0 || i[0] <= 0 {
		return applyMaxScore(c.score)
	}

	return applyMaxScore(c.score * math.Floor(i[0]*100) / 100)
}
