package auth

const (
	Email  Name = "email"
	Vk     Name = "vk"
	Ok     Name = "ok"
	Yandex Name = "yandex"
)

type Name string

func (n Name) String() string {
	return string(n)
}
