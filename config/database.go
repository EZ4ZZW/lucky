package config

func GetDbConfig() map[string]interface{} {

	// init db config
	dbConfig := make(map[string]interface{})

	dbConfig["hostname"] = "rm-bp1pkcz9y269h30fm5o.mysql.rds.aliyuncs.com"
	dbConfig["port"] = "3306"
	dbConfig["database"] = "lucky_2021"
	dbConfig["username"] = "lucky_2021"
	dbConfig["password"] = "Muxistudio304"
	dbConfig["charset"] = "utf8mb4"
	dbConfig["parseTime"] = "True"

	return dbConfig
}
