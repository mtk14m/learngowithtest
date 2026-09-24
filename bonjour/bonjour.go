package bonjour

// const anglais = "Anglais"
const francais = "Français"
const espagnol = "Espagnol"

const prefixeSalutAnglais = "Hello, "
const prefixeSalutFrancais = "Bonjour, "
const prefixeSalutEspagnol = "Hola, "

func Bonjour(langue, name string) string {
	var prefixe, defaultName string

	switch langue {
	case espagnol:
		prefixe = prefixeSalutEspagnol
		defaultName = "mundo"
	case francais:
		prefixe = prefixeSalutFrancais
		defaultName = "monde"
	default:
		prefixe = prefixeSalutAnglais
		defaultName = "world"
	}

	if name == "" {
		name = defaultName
	}

	return prefixe + name
}
