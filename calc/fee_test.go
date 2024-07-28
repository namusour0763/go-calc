package calc_test

import (
	"calc/calc"
	"testing"
	"time"
)

func ToDate(t *testing.T, date string) time.Time {
	t.Helper()
	d, err := time.Parse("2006-01-02 15:04:05", date)
	if err != nil {
		t.Fatalf("toDate: %v", err)
	}
	return d
}

func TestFee(t *testing.T) {
	cases := map[string]struct {
		in        string
		want      int
		expectErr bool
	}{
		"daytime_10:00":      {"2022-01-02 10:00:00", 1000, false},
		"midnight_20:00":     {"2022-01-02 22:00:00", 1200, false},
		"early_morning_5:00": {"2022-01-02 05:00:00", 900, false},
		"err_2:00":           {"2022-01-02 02:00:00", 0, true},
	}

	for name, tt := range cases {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := calc.Fee(ToDate(t, tt.in))
			if tt.expectErr {
				if err == nil {
					t.Error("want err")
				}
			} else {
				if err != nil {
					t.Error("not want err")
				}
			}

			if got != tt.want {
				t.Errorf("want = %d, but got = %d", tt.want, got)
			}
		})
	}
}
