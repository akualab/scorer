// Copyright 2013 AKUALAB INC. All Rights Reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scorer

import (
	"testing"
)

func TestAccuracy(t *testing.T) {

	sc := NewAccuracyScore(false, false)
	t.Logf("Name: %s", sc.Name())
	t.Logf("Description: %s", sc.Description())

	result, e := sc.Session("0", ref0, hyp0)
	if e != nil {
		t.Fatal(e)
	}
	t.Logf("#%02d: %s", 0, result.Text)
	n := result.Map["n"].(int)
	nerr := result.Map["errors"].(int)
	if n != 11 {
		t.Fatalf("bad n %d", n)
	}
	if nerr != 4 {
		t.Fatalf("bad nerr %d", nerr)
	}

	result, e = sc.Session("1", ref1, hyp1)
	if e != nil {
		t.Fatal(e)
	}
	t.Logf("#%02d: %s", 1, result.Text)
	n = result.Map["n"].(int)
	nerr = result.Map["errors"].(int)
	if n != 6 {
		t.Fatalf("bad n %d", n)
	}
	if nerr != 3 {
		t.Fatalf("bad nerr %d", nerr)
	}

	result = sc.Total()
	t.Logf("TOTAL: %s", result.Text)
}

func TestAccuracyMerge(t *testing.T) {

	sc := NewAccuracyScore(false, true)
	t.Logf("Name: %s", sc.Name())
	t.Logf("Description: %s", sc.Description())

	result, e := sc.Session("0", ref0, hyp0)
	if e != nil {
		t.Fatal(e)
	}
	t.Logf("#%02d: %s", 0, result.Text)
	n := result.Map["n"].(int)
	nerr := result.Map["errors"].(int)
	if n != 8 {
		t.Fatalf("bad n %d", n)
	}
	if nerr != 4 {
		t.Fatalf("bad nerr %d", nerr)
	}

	result, e = sc.Session("1", ref1, hyp1)
	if e != nil {
		t.Fatal(e)
	}
	t.Logf("#%02d: %s", 1, result.Text)
	n = result.Map["n"].(int)
	nerr = result.Map["errors"].(int)
	if n != 5 {
		t.Fatalf("bad n %d", n)
	}
	if nerr != 4 {
		t.Fatalf("bad nerr %d", nerr)
	}

	result = sc.Total()
	t.Logf("TOTAL: %s", result.Text)
}

var ref0 = []string{"A", "B", "A", "D", "D", "D", "A", "XX", "A", "A", "C"}
var hyp0 = []string{"A", "B", "C", "D", "D", "D", "D", "D", "A", "A"}

var ref1 = []string{"A", "B", "A", "A", "C", "XX"}
var hyp1 = []string{"A", "B", "C", "A", "C", "A", "A"}
