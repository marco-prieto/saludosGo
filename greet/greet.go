package greet

//ese var sirve en todo el paquete papu
//puedo compartir estructuras variables constantes etc

var emoji = "hola"

// no puedo hacer el emoji := ..
//si la primera letra es mayuscula go sabe q tiene q exportar y minuscula no tiene q exportar
//en este caso funciones estan iniciando con mayus y la variable con minuscula
func English() string {
	return "Hi " + emoji
}

func Italian() string {
	return "Ciao " + emoji
}
