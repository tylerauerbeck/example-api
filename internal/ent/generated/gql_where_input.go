// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"errors"
	"fmt"
	"time"

	"go.infratographer.com/example-api/internal/ent/generated/example"
	"go.infratographer.com/example-api/internal/ent/generated/predicate"
	"go.infratographer.com/x/gidx"
)

// ExampleWhereInput represents a where input for filtering Example queries.
type ExampleWhereInput struct {
	Predicates []predicate.Example  `json:"-"`
	Not        *ExampleWhereInput   `json:"not,omitempty"`
	Or         []*ExampleWhereInput `json:"or,omitempty"`
	And        []*ExampleWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *gidx.PrefixedID  `json:"id,omitempty"`
	IDNEQ   *gidx.PrefixedID  `json:"idNEQ,omitempty"`
	IDIn    []gidx.PrefixedID `json:"idIn,omitempty"`
	IDNotIn []gidx.PrefixedID `json:"idNotIn,omitempty"`
	IDGT    *gidx.PrefixedID  `json:"idGT,omitempty"`
	IDGTE   *gidx.PrefixedID  `json:"idGTE,omitempty"`
	IDLT    *gidx.PrefixedID  `json:"idLT,omitempty"`
	IDLTE   *gidx.PrefixedID  `json:"idLTE,omitempty"`

	// "created_at" field predicates.
	CreatedAt      *time.Time  `json:"createdAt,omitempty"`
	CreatedAtNEQ   *time.Time  `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGT    *time.Time  `json:"createdAtGT,omitempty"`
	CreatedAtGTE   *time.Time  `json:"createdAtGTE,omitempty"`
	CreatedAtLT    *time.Time  `json:"createdAtLT,omitempty"`
	CreatedAtLTE   *time.Time  `json:"createdAtLTE,omitempty"`

	// "updated_at" field predicates.
	UpdatedAt      *time.Time  `json:"updatedAt,omitempty"`
	UpdatedAtNEQ   *time.Time  `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn    []time.Time `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn []time.Time `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGT    *time.Time  `json:"updatedAtGT,omitempty"`
	UpdatedAtGTE   *time.Time  `json:"updatedAtGTE,omitempty"`
	UpdatedAtLT    *time.Time  `json:"updatedAtLT,omitempty"`
	UpdatedAtLTE   *time.Time  `json:"updatedAtLTE,omitempty"`

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
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`

	// "description" field predicates.
	Description             *string  `json:"description,omitempty"`
	DescriptionNEQ          *string  `json:"descriptionNEQ,omitempty"`
	DescriptionIn           []string `json:"descriptionIn,omitempty"`
	DescriptionNotIn        []string `json:"descriptionNotIn,omitempty"`
	DescriptionGT           *string  `json:"descriptionGT,omitempty"`
	DescriptionGTE          *string  `json:"descriptionGTE,omitempty"`
	DescriptionLT           *string  `json:"descriptionLT,omitempty"`
	DescriptionLTE          *string  `json:"descriptionLTE,omitempty"`
	DescriptionContains     *string  `json:"descriptionContains,omitempty"`
	DescriptionHasPrefix    *string  `json:"descriptionHasPrefix,omitempty"`
	DescriptionHasSuffix    *string  `json:"descriptionHasSuffix,omitempty"`
	DescriptionEqualFold    *string  `json:"descriptionEqualFold,omitempty"`
	DescriptionContainsFold *string  `json:"descriptionContainsFold,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *ExampleWhereInput) AddPredicates(predicates ...predicate.Example) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the ExampleWhereInput filter on the ExampleQuery builder.
func (i *ExampleWhereInput) Filter(q *ExampleQuery) (*ExampleQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyExampleWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyExampleWhereInput is returned in case the ExampleWhereInput is empty.
var ErrEmptyExampleWhereInput = errors.New("generated: empty predicate ExampleWhereInput")

// P returns a predicate for filtering examples.
// An error is returned if the input is empty or invalid.
func (i *ExampleWhereInput) P() (predicate.Example, error) {
	var predicates []predicate.Example
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, example.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Example, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, example.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Example, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, example.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, example.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, example.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, example.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, example.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, example.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, example.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, example.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, example.IDLTE(*i.IDLTE))
	}
	if i.CreatedAt != nil {
		predicates = append(predicates, example.CreatedAtEQ(*i.CreatedAt))
	}
	if i.CreatedAtNEQ != nil {
		predicates = append(predicates, example.CreatedAtNEQ(*i.CreatedAtNEQ))
	}
	if len(i.CreatedAtIn) > 0 {
		predicates = append(predicates, example.CreatedAtIn(i.CreatedAtIn...))
	}
	if len(i.CreatedAtNotIn) > 0 {
		predicates = append(predicates, example.CreatedAtNotIn(i.CreatedAtNotIn...))
	}
	if i.CreatedAtGT != nil {
		predicates = append(predicates, example.CreatedAtGT(*i.CreatedAtGT))
	}
	if i.CreatedAtGTE != nil {
		predicates = append(predicates, example.CreatedAtGTE(*i.CreatedAtGTE))
	}
	if i.CreatedAtLT != nil {
		predicates = append(predicates, example.CreatedAtLT(*i.CreatedAtLT))
	}
	if i.CreatedAtLTE != nil {
		predicates = append(predicates, example.CreatedAtLTE(*i.CreatedAtLTE))
	}
	if i.UpdatedAt != nil {
		predicates = append(predicates, example.UpdatedAtEQ(*i.UpdatedAt))
	}
	if i.UpdatedAtNEQ != nil {
		predicates = append(predicates, example.UpdatedAtNEQ(*i.UpdatedAtNEQ))
	}
	if len(i.UpdatedAtIn) > 0 {
		predicates = append(predicates, example.UpdatedAtIn(i.UpdatedAtIn...))
	}
	if len(i.UpdatedAtNotIn) > 0 {
		predicates = append(predicates, example.UpdatedAtNotIn(i.UpdatedAtNotIn...))
	}
	if i.UpdatedAtGT != nil {
		predicates = append(predicates, example.UpdatedAtGT(*i.UpdatedAtGT))
	}
	if i.UpdatedAtGTE != nil {
		predicates = append(predicates, example.UpdatedAtGTE(*i.UpdatedAtGTE))
	}
	if i.UpdatedAtLT != nil {
		predicates = append(predicates, example.UpdatedAtLT(*i.UpdatedAtLT))
	}
	if i.UpdatedAtLTE != nil {
		predicates = append(predicates, example.UpdatedAtLTE(*i.UpdatedAtLTE))
	}
	if i.Name != nil {
		predicates = append(predicates, example.NameEQ(*i.Name))
	}
	if i.NameNEQ != nil {
		predicates = append(predicates, example.NameNEQ(*i.NameNEQ))
	}
	if len(i.NameIn) > 0 {
		predicates = append(predicates, example.NameIn(i.NameIn...))
	}
	if len(i.NameNotIn) > 0 {
		predicates = append(predicates, example.NameNotIn(i.NameNotIn...))
	}
	if i.NameGT != nil {
		predicates = append(predicates, example.NameGT(*i.NameGT))
	}
	if i.NameGTE != nil {
		predicates = append(predicates, example.NameGTE(*i.NameGTE))
	}
	if i.NameLT != nil {
		predicates = append(predicates, example.NameLT(*i.NameLT))
	}
	if i.NameLTE != nil {
		predicates = append(predicates, example.NameLTE(*i.NameLTE))
	}
	if i.NameContains != nil {
		predicates = append(predicates, example.NameContains(*i.NameContains))
	}
	if i.NameHasPrefix != nil {
		predicates = append(predicates, example.NameHasPrefix(*i.NameHasPrefix))
	}
	if i.NameHasSuffix != nil {
		predicates = append(predicates, example.NameHasSuffix(*i.NameHasSuffix))
	}
	if i.NameEqualFold != nil {
		predicates = append(predicates, example.NameEqualFold(*i.NameEqualFold))
	}
	if i.NameContainsFold != nil {
		predicates = append(predicates, example.NameContainsFold(*i.NameContainsFold))
	}
	if i.Description != nil {
		predicates = append(predicates, example.DescriptionEQ(*i.Description))
	}
	if i.DescriptionNEQ != nil {
		predicates = append(predicates, example.DescriptionNEQ(*i.DescriptionNEQ))
	}
	if len(i.DescriptionIn) > 0 {
		predicates = append(predicates, example.DescriptionIn(i.DescriptionIn...))
	}
	if len(i.DescriptionNotIn) > 0 {
		predicates = append(predicates, example.DescriptionNotIn(i.DescriptionNotIn...))
	}
	if i.DescriptionGT != nil {
		predicates = append(predicates, example.DescriptionGT(*i.DescriptionGT))
	}
	if i.DescriptionGTE != nil {
		predicates = append(predicates, example.DescriptionGTE(*i.DescriptionGTE))
	}
	if i.DescriptionLT != nil {
		predicates = append(predicates, example.DescriptionLT(*i.DescriptionLT))
	}
	if i.DescriptionLTE != nil {
		predicates = append(predicates, example.DescriptionLTE(*i.DescriptionLTE))
	}
	if i.DescriptionContains != nil {
		predicates = append(predicates, example.DescriptionContains(*i.DescriptionContains))
	}
	if i.DescriptionHasPrefix != nil {
		predicates = append(predicates, example.DescriptionHasPrefix(*i.DescriptionHasPrefix))
	}
	if i.DescriptionHasSuffix != nil {
		predicates = append(predicates, example.DescriptionHasSuffix(*i.DescriptionHasSuffix))
	}
	if i.DescriptionEqualFold != nil {
		predicates = append(predicates, example.DescriptionEqualFold(*i.DescriptionEqualFold))
	}
	if i.DescriptionContainsFold != nil {
		predicates = append(predicates, example.DescriptionContainsFold(*i.DescriptionContainsFold))
	}

	switch len(predicates) {
	case 0:
		return nil, ErrEmptyExampleWhereInput
	case 1:
		return predicates[0], nil
	default:
		return example.And(predicates...), nil
	}
}
