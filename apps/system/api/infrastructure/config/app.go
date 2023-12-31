package config

// fallback value
const _default Env = ""

var properties = map[Key]map[Env]string{
	timeLocation:    {_default: "Asia/Tokyo"},
	nowOrToday:      {_default: "", Local: "2020-07-10"},
	defaultLanguage: {_default: "ja"},
}

func property(env Env, key Key) string {
	v := properties[key]
	if result, ok := v[env]; ok {
		return result
	}
	return v[_default]
}
