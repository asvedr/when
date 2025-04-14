package br

import (
	"regexp"
	"strings"
	"time"

	"github.com/asvedr/when/rules"
)

func CasualTime(s rules.Strategy) rules.Rule {
	overwrite := s == rules.Override

	return &rules.F{
		RegExp: regexp.MustCompile(`(?i)(?:\W|^)(((?:nesta|esta|ao|à))?\s*(manhã|tarde|noite|meio[- ]dia))`),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			lower := strings.ToLower(strings.TrimSpace(m.String()))

			if (c.Hour != nil || c.Minute != nil) && !overwrite {
				return false, nil
			}

			switch {
			case strings.Contains(lower, "tarde"):
				if o.Afternoon != 0 {
					c.Hour = &o.Afternoon
				} else {
					c.Hour = rules.Ptr(15)
				}
				c.Minute = rules.Ptr(0)
			case strings.Contains(lower, "noite"):
				if o.Evening != 0 {
					c.Hour = &o.Evening
				} else {
					c.Hour = rules.Ptr(18)
				}
				c.Minute = rules.Ptr(0)
			case strings.Contains(lower, "manhã"):
				if o.Morning != 0 {
					c.Hour = &o.Morning
				} else {
					c.Hour = rules.Ptr(8)
				}
				c.Minute = rules.Ptr(0)
			case strings.Contains(lower, "meio-dia"), strings.Contains(lower, "meio dia"):
				if o.Noon != 0 {
					c.Hour = &o.Noon
				} else {
					c.Hour = rules.Ptr(12)
				}
				c.Minute = rules.Ptr(0)
			}

			return true, nil
		},
	}
}
