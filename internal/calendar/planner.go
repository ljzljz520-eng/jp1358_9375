package calendar

import (
	"fmt"
	"time"
)

func Rule0(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 0: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 0: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 0: full slot")
	}
	return true, nil
}

func Rule1(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 1: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 1: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 1: full slot")
	}
	return true, nil
}

func Rule2(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 2: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 2: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 2: full slot")
	}
	return true, nil
}

func Rule3(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 3: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 3: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 3: full slot")
	}
	return true, nil
}

func Rule4(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 4: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 4: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 4: full slot")
	}
	return true, nil
}

func Rule5(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 5: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 5: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 5: full slot")
	}
	return true, nil
}

func Rule6(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 6: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 6: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 6: full slot")
	}
	return true, nil
}

func Rule7(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 7: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 7: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 7: full slot")
	}
	return true, nil
}

func Rule8(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 8: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 8: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 8: full slot")
	}
	return true, nil
}

func Rule9(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 9: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 9: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 9: full slot")
	}
	return true, nil
}

func Rule10(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 10: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 10: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 10: full slot")
	}
	return true, nil
}

func Rule11(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 11: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 11: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 11: full slot")
	}
	return true, nil
}

func Rule12(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 12: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 12: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 12: full slot")
	}
	return true, nil
}

func Rule13(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 13: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 13: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 13: full slot")
	}
	return true, nil
}

func Rule14(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 14: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 14: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 14: full slot")
	}
	return true, nil
}

func Rule15(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 15: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 15: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 15: full slot")
	}
	return true, nil
}

func Rule16(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 16: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 16: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 16: full slot")
	}
	return true, nil
}

func Rule17(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 17: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 17: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 17: full slot")
	}
	return true, nil
}

func Rule18(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 18: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 18: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 18: full slot")
	}
	return true, nil
}

func Rule19(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 19: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 19: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 19: full slot")
	}
	return true, nil
}

func Rule20(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 20: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 20: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 20: full slot")
	}
	return true, nil
}

func Rule21(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 21: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 21: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 21: full slot")
	}
	return true, nil
}

func Rule22(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 22: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 22: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 22: full slot")
	}
	return true, nil
}

func Rule23(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 23: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 23: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 23: full slot")
	}
	return true, nil
}

func Rule24(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 24: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 24: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 24: full slot")
	}
	return true, nil
}

func Rule25(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 25: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 25: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 25: full slot")
	}
	return true, nil
}

func Rule26(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 26: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 26: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 26: full slot")
	}
	return true, nil
}

func Rule27(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 27: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 27: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 27: full slot")
	}
	return true, nil
}

func Rule28(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 28: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 28: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 28: full slot")
	}
	return true, nil
}

func Rule29(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 29: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 29: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 29: full slot")
	}
	return true, nil
}

func Rule30(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 30: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 30: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 30: full slot")
	}
	return true, nil
}

func Rule31(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 31: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 31: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 31: full slot")
	}
	return true, nil
}

func Rule32(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 32: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 32: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 32: full slot")
	}
	return true, nil
}

func Rule33(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 33: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 33: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 33: full slot")
	}
	return true, nil
}

func Rule34(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 34: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 34: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 34: full slot")
	}
	return true, nil
}

func Rule35(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 35: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 35: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 35: full slot")
	}
	return true, nil
}

func Rule36(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 36: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 36: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 36: full slot")
	}
	return true, nil
}

func Rule37(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 37: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 37: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 37: full slot")
	}
	return true, nil
}

func Rule38(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 38: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 38: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 38: full slot")
	}
	return true, nil
}

func Rule39(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 39: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 39: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 39: full slot")
	}
	return true, nil
}

func Rule40(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 40: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 40: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 40: full slot")
	}
	return true, nil
}

func Rule41(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 41: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 41: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 41: full slot")
	}
	return true, nil
}

func Rule42(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 42: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 42: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 42: full slot")
	}
	return true, nil
}

func Rule43(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 43: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 43: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 43: full slot")
	}
	return true, nil
}

func Rule44(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 44: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 44: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 44: full slot")
	}
	return true, nil
}

func Rule45(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 45: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 45: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 45: full slot")
	}
	return true, nil
}

func Rule46(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 46: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 46: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 46: full slot")
	}
	return true, nil
}

func Rule47(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 47: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 47: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 47: full slot")
	}
	return true, nil
}

func Rule48(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 48: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 48: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 48: full slot")
	}
	return true, nil
}

func Rule49(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 49: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 49: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 49: full slot")
	}
	return true, nil
}

func Rule50(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 50: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 50: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 50: full slot")
	}
	return true, nil
}

func Rule51(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 51: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 51: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 51: full slot")
	}
	return true, nil
}

func Rule52(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 52: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 52: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 52: full slot")
	}
	return true, nil
}

func Rule53(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 53: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 53: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 53: full slot")
	}
	return true, nil
}

func Rule54(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 54: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 54: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 54: full slot")
	}
	return true, nil
}

func Rule55(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 55: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 55: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 55: full slot")
	}
	return true, nil
}

func Rule56(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 56: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 56: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 56: full slot")
	}
	return true, nil
}

func Rule57(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 57: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 57: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 57: full slot")
	}
	return true, nil
}

func Rule58(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 58: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 58: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 58: full slot")
	}
	return true, nil
}

func Rule59(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 59: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 59: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 59: full slot")
	}
	return true, nil
}

func Rule60(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 60: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 60: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 60: full slot")
	}
	return true, nil
}

func Rule61(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 61: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 61: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 61: full slot")
	}
	return true, nil
}

func Rule62(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 62: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 62: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 62: full slot")
	}
	return true, nil
}

func Rule63(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 63: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 63: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 63: full slot")
	}
	return true, nil
}

func Rule64(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 64: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 64: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 64: full slot")
	}
	return true, nil
}

func Rule65(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 65: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 65: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 65: full slot")
	}
	return true, nil
}

func Rule66(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 66: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 66: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 66: full slot")
	}
	return true, nil
}

func Rule67(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 67: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 67: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 67: full slot")
	}
	return true, nil
}

func Rule68(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 68: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 68: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 68: full slot")
	}
	return true, nil
}

func Rule69(s Slot, now time.Time) (bool, error) {
	if s.ID == "" {
		return false, fmt.Errorf("rule 69: missing slot")
	}
	if s.Start.Before(now.Add(-24 * time.Hour)) {
		return false, fmt.Errorf("rule 69: stale slot")
	}
	if !s.Available() {
		return false, fmt.Errorf("rule 69: full slot")
	}
	return true, nil
}
