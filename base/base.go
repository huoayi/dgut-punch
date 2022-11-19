package base

const YQFK_DATA_API_URL = "https://yqfk-daka-api.dgut.edu.cn/"
const LOGIN_URL = "https://auth.dgut.edu.cn/authserver/login?service=https%3A%2F%2Fauth.dgut.edu.cn%2Fauthserver%2Foauth2.0%2FcallbackAuthorize%3Fclient_id%3D1021534300621787136%26redirect_uri%3Dhttps%253A%252F%252Fyqfk-daka.dgut.edu.cn%252Fnew_login%252Fdgut%26response_type%3Dcode%26client_name%3DCasOAuthClient"
const AUTHORIZE_URL = "https://auth.dgut.edu.cn/authserver/oauth2.0/authorize?" +
	"response_type=code&client_id=1021534300621787136&redirect_uri=https://yqfk-daka.dgut.edu.cn/new_login/dgut&state=yqfk"

type UserAccount struct {
	Username string `yaml:"username" gorm:"username" json:"username"`
	Password string `yaml:"password" gorm:"password" json:"password"`
}

type DataForm struct {
	Data map[string]string
}

type AuthAccessToken struct {
	AccessToken string `json:"access_token"`
}

type ConfigApplication struct {
	UserAccount []UserAccount `yaml:"UserAccount"`
	DB          *DB           `yaml:"DB"`
	FuncMAX     int           `yaml:"func_max"`
	DataSource  int           `yaml:"data_source"`
	DebugModel  int           `yaml:"debug_model"`
}

type DB struct {
	DBHost     string `yaml:"DB_host"`
	DBPort     string `yaml:"DB_port"`
	DBName     string `yaml:"DB_name"`
	DBTable    string `yaml:"DB_table"`
	DBUsername string `yaml:"DB_username"`
	DBPassword string `yaml:"DB_password"`
}
