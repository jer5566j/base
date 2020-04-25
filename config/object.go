package config

type Conf struct {
	App    AppConf    `yaml:"app"`
	DB     DBConf     `yaml:"DB"`
	Http   HttpConf   `yaml:"http"`
	Token  TokenConf  `yaml:"token"`
	WeChat WeChatConf `yaml:"wechat"`
}

type AppConf struct {
	Debug    bool   `yaml:"debug"`
	RootPath string `yaml:"root_path"`
}

type DBConf struct {
	Dsn            string `yaml:"dsn"`
	Type           string `yaml:"type"`
	MaxOpenConnect int    `yaml:"max_open_connect"`
	MaxIdleConnect int    `yaml:"max_idle_connect"`
	TablePrefix    string `yaml:"table_prefix"`
}

type HttpConf struct {
	Ip      string `yaml:"ip"`
	Port    string `yaml:"port"`
	Host    string `yaml:"host"`
	PicHost string `yaml:"pic_host"`
}

type TokenConf struct {
	Secret string `yaml:"secret"`
	Expire int64  `yaml:"expire"`
}

type WeChatConf struct {
	AppId  string    `yaml:"app_id"`
	Secret string    `yaml:"secret"`
	Api    WeChatApi `yaml:"api"`
}

type WeChatApi struct {
	CodeToSession string `yaml:"code_to_session"`
}
