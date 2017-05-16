package view

import "testing"

func TestStage(t *testing.T) {
	test := []struct {
		from, to, test Stage
		enter, exit    bool
	}{
		{StageDead, StageDead, StagePrepreload, false, false},
		{StageDead, StageMounted, StagePrepreload, false, false},
		{StageDead, StagePrepreload, StagePrepreload, true, false},
		{StageDead, StagePreload, StagePrepreload, true, false},
		{StageDead, StageVisible, StagePrepreload, true, false},

		{StageMounted, StageDead, StagePrepreload, false, false},
		{StageMounted, StageMounted, StagePrepreload, false, false},
		{StageMounted, StagePrepreload, StagePrepreload, true, false},
		{StageMounted, StagePreload, StagePrepreload, true, false},
		{StageMounted, StageVisible, StagePrepreload, true, false},

		{StagePrepreload, StageDead, StagePrepreload, false, true},
		{StagePrepreload, StageMounted, StagePrepreload, false, true},
		{StagePrepreload, StagePrepreload, StagePrepreload, false, false},
		{StagePrepreload, StagePreload, StagePrepreload, false, false},
		{StagePrepreload, StageVisible, StagePrepreload, false, false},

		{StagePreload, StageDead, StagePrepreload, false, true},
		{StagePreload, StageMounted, StagePrepreload, false, true},
		{StagePreload, StagePrepreload, StagePrepreload, false, false},
		{StagePreload, StagePreload, StagePrepreload, false, false},
		{StagePreload, StageVisible, StagePrepreload, false, false},

		{StageVisible, StageDead, StagePrepreload, false, true},
		{StageVisible, StageMounted, StagePrepreload, false, true},
		{StageVisible, StagePrepreload, StagePrepreload, false, false},
		{StageVisible, StagePreload, StagePrepreload, false, false},
		{StageVisible, StageVisible, StagePrepreload, false, false},
	}

	for _, i := range test {
		if EntersStage(i.from, i.to, i.test) != i.enter {
			t.Error("enter", i.from, i.to, i.test)
		}
		if ExitsStage(i.from, i.to, i.test) != i.exit {
			t.Error("exit", i.from, i.to, i.test)
		}
	}
}
