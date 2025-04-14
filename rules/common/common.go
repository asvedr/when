package common

import "github.com/asvedr/when/rules"

var All = []rules.Rule{
	SlashDMY(rules.Override),
}
