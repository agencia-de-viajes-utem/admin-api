package models

type Usuario struct {
	ID              int    `json:"id"`
	Nombre          string `json:"nombre"`
	Apellido        string `json:"apellido"`
	Email           string `json:"email"`
	Rut             string `json:"rut"`
	Fono            string `json:"fono"`
	Fotodeperfil    string `json:"fotodeperfil"`
	FechaNacimiento string `json:"fecha_nacimiento"`
}

type UsuariosResponse struct {
	Usuarios []Usuario `json:"usuarios"`
}
