package nl

import (
	"regexp"
	"strings"
	"time"

	"github.com/asvedr/when/rules"
)

func CasualDate(s rules.Strategy) rules.Rule {
	overwrite := s == rules.Override

	return &rules.F{
		RegExp: regexp.MustCompile("(?i)(?:\\W|^)(nu|vandaag|vanavond|vannacht|afgelopen\\s*nacht|morgen|gister|gisteren)(ochtend|morgen|middag|avond)?(?:\\W|$)"),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			lower := strings.ToLower(strings.TrimSpace(m.String()))

			if regexp.MustCompile("ochtend|\\s*morgen|middag|avond").MatchString(lower) {
				switch {
				case strings.Contains(lower, "ochtend"), regexp.MustCompile("(?i)(?:\\W|^)(\\s*morgen)(?:\\W|$)").MatchString(lower):
					if o.Morning != 0 {
						c.Hour = &o.Morning
					} else {
						c.Hour = rules.Ptr(8)
					}
				case strings.Contains(lower, "middag"):
					if o.Afternoon != 0 {
						c.Hour = &o.Afternoon
					} else {
						c.Hour = rules.Ptr(15)
					}
				case strings.Contains(lower, "avond"):
					if o.Evening != 0 {
						c.Hour = &o.Evening
					} else {
						c.Hour = rules.Ptr(18)
					}
				}
			}

			switch {
			case strings.Contains(lower, "vannacht"):
				if c.Hour == nil && c.Minute == nil || overwrite {
					c.Hour = rules.Ptr(23)
					c.Minute = rules.Ptr(0)
				}
			case strings.Contains(lower, "vandaag"):
				// c.Hour = rules.Ptr(18)
			case strings.Contains(lower, "morgen"):
				if c.Duration == 0 || overwrite {
					c.Duration += time.Hour * 24
				}
			case strings.Contains(lower, "gister"):
				if c.Duration == 0 || overwrite {
					c.Duration -= time.Hour * 24
				}
			case strings.Contains(lower, "afgelopen nacht"):
				if (c.Hour == nil && c.Duration == 0) || overwrite {
					c.Hour = rules.Ptr(23)
					c.Duration -= time.Hour * 24
				}
			}

			return true, nil
		},
	}
}
