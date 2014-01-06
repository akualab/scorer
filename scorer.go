// Copyright 2013 AKUALAB INC. All Rights Reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package for scoring and evaluating discrete-time sequences. A sequence has samples at
times 0,1,...N-1. For each sample, we know the correct value, or reference (REF).
The scorer compares the REF to a hypothesized sequence (HYP).

The data is organized as a collection of sessions. A session is a sequence with a session ID.

The package includes various modules to score the inputs. The results can be computed
at the session level and overall for a group of sessions.
*/
package scorer

type Scorer interface {
	// The name of the scorer.
	Name() string
	// Description.
	Description() string
	// Results for session.
	Session(id string, ref, hyp []string) (result *Score, e error)
	// Cumulative results.
	Total() (result *Score)
}

type Score struct {
	// Score name.
	Name string
	// Results as formatted text.
	Text string
	// Results.
	Map map[string]interface{}
}
