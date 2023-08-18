package hii

type Secret struct {
	password string
	key      string
	Name     string
}

func GetSecret(key string) Secret {
	return Secret{key: key, password: "default password"}
}

func (c Secret) GetPassword(key string) string {
	if key == c.key {
		return c.password
	}
	return ""
}
