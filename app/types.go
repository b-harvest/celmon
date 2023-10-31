package app

type Config struct {
	General struct {
		Network              string `toml:"network"`
		Period               int    `toml:"period"`
		ConsGrpc             string `toml:"cons_grpc"`
		GrpcSecureConnection bool   `toml:"grpc_secure_connection"`
		DaAPI                string `toml:"da_api"`
		APIToken             string `toml:"api_token"`
		ListenPort           int    `toml:"listen_port"`
	} `toml:"general"`
	Tg struct {
		Enable bool   `toml:"enable"`
		Token  string `toml:"token"`
		ChatID string `toml:"chat_id"`
	} `toml:"tg"`
	AlarmCriteria struct {
		HeightDiffer int `toml:"height_differ"`
	} `toml:"alarm-criteria"`
}
