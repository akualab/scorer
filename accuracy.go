// Copyright 2013 AKUALAB INC. All Rights Reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scorer

import "fmt"

// Computes number of errors.
type AccuracyScore struct {
	name        string
	description string
	errors      int
	n           int
	equalLength bool
	mergeTokens bool
}

// If equalLength is true, returns error if ref and hyp have different number of tokens.
// If equalLength is false, the difference in length is counted as errors.
// If mergeTokens is true, a set of identical consecutive tokens is reduced to a signle token.
func NewAccuracyScore(equlLength, mergeTokens bool) *AccuracyScore {

	return &AccuracyScore{equalLength: equlLength, mergeTokens: mergeTokens}
}

func (s *AccuracyScore) Name() string {
	return "Accuracy"
}

func (s *AccuracyScore) Description() string {
	return "Computes the number errors between the REF and the HYP."
}

func (s *AccuracyScore) Session(id string, ref, hyp []string) (result *Score, e error) {

	result = &Score{Name: s.Name()}
	n, miss, pct, e := s.session(id, ref, hyp)
	if e != nil {
		return
	}
	result.Text = fmt.Sprintf("Tokens: %4d, Errors: %4d (%6.2f%%)", n, miss, pct)

	result.Map = map[string]interface{}{
		"n":      n,
		"errors": miss,
	}

	return
}

func (s *AccuracyScore) session(id string, ref, hyp []string) (n, miss int, pct float64, e error) {

	if s.equalLength && len(ref) != len(hyp) {
		e = fmt.Errorf("ref and hyp must have the same length. [%d] vs. [%d]", len(ref), len(hyp))
		return
	}

	if s.mergeTokens {
		ref = merge(ref)
		hyp = merge(hyp)
	}

	n = len(ref)
	diff := n - len(hyp)
	minLen := n
	if diff > 0 {
		minLen = len(hyp)
	} else {
		diff = -diff
	}

	//	for k, v := range ref {
	for i := 0; i < minLen; i++ {
		if ref[i] != hyp[i] {
			miss++
		}
	}
	miss += diff

	// Accumulate stats.
	s.errors += miss
	s.n += n

	pct = float64(miss) / float64(n) * 100
	return
}

func (s *AccuracyScore) Total() *Score {

	result := &Score{Name: s.Name()}
	pct := float64(s.errors) / float64(s.n) * 100
	result.Text = fmt.Sprintf("Tokens: %4d, Errors: %4d (%6.2f%%)", s.n, s.errors, pct)
	result.Map = map[string]interface{}{
		"n":      s.n,
		"errors": s.errors,
	}

	return result
}

func merge(in []string) (out []string) {

	out = make([]string, 0)

	prev := ""
	for _, v := range in {
		if v == prev {
			continue
		}
		out = append(out, v)
		prev = v
	}
	return
}
