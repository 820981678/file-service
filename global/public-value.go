package global

type Config struct {
	ListenAddr string `json:"ListenAddr"`
	SaveRootPath string `json:"SaveRootPath"`
}
var ApplicationConfig Config