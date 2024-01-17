package translation

var (
	simpleTags = []string{
		"required",
		"alpha",
		"alphanum",
		"numeric",
		"number",
		"email",
		"latitude",
		"longitude",
	}
	withParamTags = []string{
		"contains",
		"containsany",
		"excludes",
		"excludesall",
		"excludesrune",
		"oneof",
		"eq",
		"ne",
		"time",
		"date",
	}
	fieldCompareTags = []string{
		"eqfield",
		"eqcsfield",
		"necsfield",
		"gtcsfield",
		"gtecsfield",
		"ltcsfield",
		"ltecsfield",
		"nefield",
		"gtfield",
		"gtefield",
		"ltfield",
		"ltefield",
	}
)
