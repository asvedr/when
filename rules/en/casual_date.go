package en

import (
	"regexp"
	"strings"
	"time"

	"github.com/asvedr/when/rules"
)

func CasualDate(s rules.Strategy) rules.Rule {
	overwrite := s == rules.Override

	return &rules.F{
		RegExp: regexp.MustCompile("(?i)(?:\\W|^)(now|today|tonight|last\\s*night|(?:tomorrow|tmr|yesterday)\\s*|tomorrow|tmr|yesterday)(?:\\W|$)"),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			lower := strings.ToLower(strings.TrimSpace(m.String()))

			switch {
			case strings.Contains(lower, "tonight"):
				if c.Hour == nil && c.Minute == nil || overwrite {
					c.Hour = rules.Ptr(23)
					c.Minute = rules.Ptr(0)
				}
			case strings.Contains(lower, "today"):
				// c.Hour = rules.Ptr(18)
			case strings.Contains(lower, "tomorrow"), strings.Contains(lower, "tmr"):
				if c.Duration == 0 || overwrite {
					c.Duration += time.Hour * 24
				}
			case strings.Contains(lower, "yesterday"):
				if c.Duration == 0 || overwrite {
					c.Duration -= time.Hour * 24
				}
			case strings.Contains(lower, "last night"):
				if (c.Hour == nil && c.Duration == 0) || overwrite {
					c.Hour = rules.Ptr(23)
					c.Duration -= time.Hour * 24
				}
			}

			return true, nil
		},
	}
}
