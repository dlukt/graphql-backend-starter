// Code generated by ent, DO NOT EDIT.

package ent

import (
	"errors"
	"fmt"
	"time"

	"github.com/dlukt/graphql-backend-starter/ent/predicate"
	"github.com/dlukt/graphql-backend-starter/ent/profile"
	"github.com/rs/xid"
)

// ProfileWhereInput represents a where input for filtering Profile queries.
type ProfileWhereInput struct {
	Predicates []predicate.Profile  `json:"-"`
	Not        *ProfileWhereInput   `json:"not,omitempty"`
	Or         []*ProfileWhereInput `json:"or,omitempty"`
	And        []*ProfileWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *xid.ID  `json:"id,omitempty"`
	IDNEQ   *xid.ID  `json:"idNEQ,omitempty"`
	IDIn    []xid.ID `json:"idIn,omitempty"`
	IDNotIn []xid.ID `json:"idNotIn,omitempty"`
	IDGT    *xid.ID  `json:"idGT,omitempty"`
	IDGTE   *xid.ID  `json:"idGTE,omitempty"`
	IDLT    *xid.ID  `json:"idLT,omitempty"`
	IDLTE   *xid.ID  `json:"idLTE,omitempty"`

	// "create_time" field predicates.
	CreateTime      *time.Time  `json:"createTime,omitempty"`
	CreateTimeNEQ   *time.Time  `json:"createTimeNEQ,omitempty"`
	CreateTimeIn    []time.Time `json:"createTimeIn,omitempty"`
	CreateTimeNotIn []time.Time `json:"createTimeNotIn,omitempty"`
	CreateTimeGT    *time.Time  `json:"createTimeGT,omitempty"`
	CreateTimeGTE   *time.Time  `json:"createTimeGTE,omitempty"`
	CreateTimeLT    *time.Time  `json:"createTimeLT,omitempty"`
	CreateTimeLTE   *time.Time  `json:"createTimeLTE,omitempty"`

	// "update_time" field predicates.
	UpdateTime       *time.Time  `json:"updateTime,omitempty"`
	UpdateTimeNEQ    *time.Time  `json:"updateTimeNEQ,omitempty"`
	UpdateTimeIn     []time.Time `json:"updateTimeIn,omitempty"`
	UpdateTimeNotIn  []time.Time `json:"updateTimeNotIn,omitempty"`
	UpdateTimeGT     *time.Time  `json:"updateTimeGT,omitempty"`
	UpdateTimeGTE    *time.Time  `json:"updateTimeGTE,omitempty"`
	UpdateTimeLT     *time.Time  `json:"updateTimeLT,omitempty"`
	UpdateTimeLTE    *time.Time  `json:"updateTimeLTE,omitempty"`
	UpdateTimeIsNil  bool        `json:"updateTimeIsNil,omitempty"`
	UpdateTimeNotNil bool        `json:"updateTimeNotNil,omitempty"`

	// "sub" field predicates.
	Sub             *string  `json:"sub,omitempty"`
	SubNEQ          *string  `json:"subNEQ,omitempty"`
	SubIn           []string `json:"subIn,omitempty"`
	SubNotIn        []string `json:"subNotIn,omitempty"`
	SubGT           *string  `json:"subGT,omitempty"`
	SubGTE          *string  `json:"subGTE,omitempty"`
	SubLT           *string  `json:"subLT,omitempty"`
	SubLTE          *string  `json:"subLTE,omitempty"`
	SubContains     *string  `json:"subContains,omitempty"`
	SubHasPrefix    *string  `json:"subHasPrefix,omitempty"`
	SubHasSuffix    *string  `json:"subHasSuffix,omitempty"`
	SubEqualFold    *string  `json:"subEqualFold,omitempty"`
	SubContainsFold *string  `json:"subContainsFold,omitempty"`

	// "name" field predicates.
	Name             *string  `json:"name,omitempty"`
	NameNEQ          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGT           *string  `json:"nameGT,omitempty"`
	NameGTE          *string  `json:"nameGTE,omitempty"`
	NameLT           *string  `json:"nameLT,omitempty"`
	NameLTE          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameIsNil        bool     `json:"nameIsNil,omitempty"`
	NameNotNil       bool     `json:"nameNotNil,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`

	// "gender" field predicates.
	Gender             *string  `json:"gender,omitempty"`
	GenderNEQ          *string  `json:"genderNEQ,omitempty"`
	GenderIn           []string `json:"genderIn,omitempty"`
	GenderNotIn        []string `json:"genderNotIn,omitempty"`
	GenderGT           *string  `json:"genderGT,omitempty"`
	GenderGTE          *string  `json:"genderGTE,omitempty"`
	GenderLT           *string  `json:"genderLT,omitempty"`
	GenderLTE          *string  `json:"genderLTE,omitempty"`
	GenderContains     *string  `json:"genderContains,omitempty"`
	GenderHasPrefix    *string  `json:"genderHasPrefix,omitempty"`
	GenderHasSuffix    *string  `json:"genderHasSuffix,omitempty"`
	GenderIsNil        bool     `json:"genderIsNil,omitempty"`
	GenderNotNil       bool     `json:"genderNotNil,omitempty"`
	GenderEqualFold    *string  `json:"genderEqualFold,omitempty"`
	GenderContainsFold *string  `json:"genderContainsFold,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *ProfileWhereInput) AddPredicates(predicates ...predicate.Profile) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the ProfileWhereInput filter on the ProfileQuery builder.
func (i *ProfileWhereInput) Filter(q *ProfileQuery) (*ProfileQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyProfileWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyProfileWhereInput is returned in case the ProfileWhereInput is empty.
var ErrEmptyProfileWhereInput = errors.New("ent: empty predicate ProfileWhereInput")

// P returns a predicate for filtering profiles.
// An error is returned if the input is empty or invalid.
func (i *ProfileWhereInput) P() (predicate.Profile, error) {
	var predicates []predicate.Profile
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, profile.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Profile, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, profile.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Profile, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, profile.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, profile.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, profile.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, profile.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, profile.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, profile.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, profile.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, profile.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, profile.IDLTE(*i.IDLTE))
	}
	if i.CreateTime != nil {
		predicates = append(predicates, profile.CreateTimeEQ(*i.CreateTime))
	}
	if i.CreateTimeNEQ != nil {
		predicates = append(predicates, profile.CreateTimeNEQ(*i.CreateTimeNEQ))
	}
	if len(i.CreateTimeIn) > 0 {
		predicates = append(predicates, profile.CreateTimeIn(i.CreateTimeIn...))
	}
	if len(i.CreateTimeNotIn) > 0 {
		predicates = append(predicates, profile.CreateTimeNotIn(i.CreateTimeNotIn...))
	}
	if i.CreateTimeGT != nil {
		predicates = append(predicates, profile.CreateTimeGT(*i.CreateTimeGT))
	}
	if i.CreateTimeGTE != nil {
		predicates = append(predicates, profile.CreateTimeGTE(*i.CreateTimeGTE))
	}
	if i.CreateTimeLT != nil {
		predicates = append(predicates, profile.CreateTimeLT(*i.CreateTimeLT))
	}
	if i.CreateTimeLTE != nil {
		predicates = append(predicates, profile.CreateTimeLTE(*i.CreateTimeLTE))
	}
	if i.UpdateTime != nil {
		predicates = append(predicates, profile.UpdateTimeEQ(*i.UpdateTime))
	}
	if i.UpdateTimeNEQ != nil {
		predicates = append(predicates, profile.UpdateTimeNEQ(*i.UpdateTimeNEQ))
	}
	if len(i.UpdateTimeIn) > 0 {
		predicates = append(predicates, profile.UpdateTimeIn(i.UpdateTimeIn...))
	}
	if len(i.UpdateTimeNotIn) > 0 {
		predicates = append(predicates, profile.UpdateTimeNotIn(i.UpdateTimeNotIn...))
	}
	if i.UpdateTimeGT != nil {
		predicates = append(predicates, profile.UpdateTimeGT(*i.UpdateTimeGT))
	}
	if i.UpdateTimeGTE != nil {
		predicates = append(predicates, profile.UpdateTimeGTE(*i.UpdateTimeGTE))
	}
	if i.UpdateTimeLT != nil {
		predicates = append(predicates, profile.UpdateTimeLT(*i.UpdateTimeLT))
	}
	if i.UpdateTimeLTE != nil {
		predicates = append(predicates, profile.UpdateTimeLTE(*i.UpdateTimeLTE))
	}
	if i.UpdateTimeIsNil {
		predicates = append(predicates, profile.UpdateTimeIsNil())
	}
	if i.UpdateTimeNotNil {
		predicates = append(predicates, profile.UpdateTimeNotNil())
	}
	if i.Sub != nil {
		predicates = append(predicates, profile.SubEQ(*i.Sub))
	}
	if i.SubNEQ != nil {
		predicates = append(predicates, profile.SubNEQ(*i.SubNEQ))
	}
	if len(i.SubIn) > 0 {
		predicates = append(predicates, profile.SubIn(i.SubIn...))
	}
	if len(i.SubNotIn) > 0 {
		predicates = append(predicates, profile.SubNotIn(i.SubNotIn...))
	}
	if i.SubGT != nil {
		predicates = append(predicates, profile.SubGT(*i.SubGT))
	}
	if i.SubGTE != nil {
		predicates = append(predicates, profile.SubGTE(*i.SubGTE))
	}
	if i.SubLT != nil {
		predicates = append(predicates, profile.SubLT(*i.SubLT))
	}
	if i.SubLTE != nil {
		predicates = append(predicates, profile.SubLTE(*i.SubLTE))
	}
	if i.SubContains != nil {
		predicates = append(predicates, profile.SubContains(*i.SubContains))
	}
	if i.SubHasPrefix != nil {
		predicates = append(predicates, profile.SubHasPrefix(*i.SubHasPrefix))
	}
	if i.SubHasSuffix != nil {
		predicates = append(predicates, profile.SubHasSuffix(*i.SubHasSuffix))
	}
	if i.SubEqualFold != nil {
		predicates = append(predicates, profile.SubEqualFold(*i.SubEqualFold))
	}
	if i.SubContainsFold != nil {
		predicates = append(predicates, profile.SubContainsFold(*i.SubContainsFold))
	}
	if i.Name != nil {
		predicates = append(predicates, profile.NameEQ(*i.Name))
	}
	if i.NameNEQ != nil {
		predicates = append(predicates, profile.NameNEQ(*i.NameNEQ))
	}
	if len(i.NameIn) > 0 {
		predicates = append(predicates, profile.NameIn(i.NameIn...))
	}
	if len(i.NameNotIn) > 0 {
		predicates = append(predicates, profile.NameNotIn(i.NameNotIn...))
	}
	if i.NameGT != nil {
		predicates = append(predicates, profile.NameGT(*i.NameGT))
	}
	if i.NameGTE != nil {
		predicates = append(predicates, profile.NameGTE(*i.NameGTE))
	}
	if i.NameLT != nil {
		predicates = append(predicates, profile.NameLT(*i.NameLT))
	}
	if i.NameLTE != nil {
		predicates = append(predicates, profile.NameLTE(*i.NameLTE))
	}
	if i.NameContains != nil {
		predicates = append(predicates, profile.NameContains(*i.NameContains))
	}
	if i.NameHasPrefix != nil {
		predicates = append(predicates, profile.NameHasPrefix(*i.NameHasPrefix))
	}
	if i.NameHasSuffix != nil {
		predicates = append(predicates, profile.NameHasSuffix(*i.NameHasSuffix))
	}
	if i.NameIsNil {
		predicates = append(predicates, profile.NameIsNil())
	}
	if i.NameNotNil {
		predicates = append(predicates, profile.NameNotNil())
	}
	if i.NameEqualFold != nil {
		predicates = append(predicates, profile.NameEqualFold(*i.NameEqualFold))
	}
	if i.NameContainsFold != nil {
		predicates = append(predicates, profile.NameContainsFold(*i.NameContainsFold))
	}
	if i.Gender != nil {
		predicates = append(predicates, profile.GenderEQ(*i.Gender))
	}
	if i.GenderNEQ != nil {
		predicates = append(predicates, profile.GenderNEQ(*i.GenderNEQ))
	}
	if len(i.GenderIn) > 0 {
		predicates = append(predicates, profile.GenderIn(i.GenderIn...))
	}
	if len(i.GenderNotIn) > 0 {
		predicates = append(predicates, profile.GenderNotIn(i.GenderNotIn...))
	}
	if i.GenderGT != nil {
		predicates = append(predicates, profile.GenderGT(*i.GenderGT))
	}
	if i.GenderGTE != nil {
		predicates = append(predicates, profile.GenderGTE(*i.GenderGTE))
	}
	if i.GenderLT != nil {
		predicates = append(predicates, profile.GenderLT(*i.GenderLT))
	}
	if i.GenderLTE != nil {
		predicates = append(predicates, profile.GenderLTE(*i.GenderLTE))
	}
	if i.GenderContains != nil {
		predicates = append(predicates, profile.GenderContains(*i.GenderContains))
	}
	if i.GenderHasPrefix != nil {
		predicates = append(predicates, profile.GenderHasPrefix(*i.GenderHasPrefix))
	}
	if i.GenderHasSuffix != nil {
		predicates = append(predicates, profile.GenderHasSuffix(*i.GenderHasSuffix))
	}
	if i.GenderIsNil {
		predicates = append(predicates, profile.GenderIsNil())
	}
	if i.GenderNotNil {
		predicates = append(predicates, profile.GenderNotNil())
	}
	if i.GenderEqualFold != nil {
		predicates = append(predicates, profile.GenderEqualFold(*i.GenderEqualFold))
	}
	if i.GenderContainsFold != nil {
		predicates = append(predicates, profile.GenderContainsFold(*i.GenderContainsFold))
	}

	switch len(predicates) {
	case 0:
		return nil, ErrEmptyProfileWhereInput
	case 1:
		return predicates[0], nil
	default:
		return profile.And(predicates...), nil
	}
}
