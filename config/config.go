package config

import (
	"go-wsadmin/utils"
)

type conf struct {
	DbDsn              string `json:"db_dsn"`
	DebugMode          bool   `json:"debug_mode"`
	SuperAdmin         string `json:"super_admin"`
	SuperAdminPassword string `json:"super_admin_password"`
}

var Conf conf

func Init(filename string) error {
	if filename == "" {
		filename = "./config"
	}
	return utils.JsonFromFile(filename, &Conf)
}
