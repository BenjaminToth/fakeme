package Models

type Auth struct {
	//Username string `fake:"username"`
	Password string `fake:"{password}"`
}
