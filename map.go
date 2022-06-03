package unisendergo

type Map map[string]interface{}

func (m Map) Add(key string, value interface{}) {
	m[key] = value
}
