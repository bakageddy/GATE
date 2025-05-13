package types

type Application struct {
	Bar Bar
}

type Bar struct {
	Links map[string]string
}
