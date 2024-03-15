package config

import (
	"flag"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

// env Структура для хранения переменных среды
type env struct {
	Host       string
	Port       string
	DbUri      string
	DbName     string
	TableName  string
	Production bool
	SecretKey  string
}

// Env глобальная переменная для доступа к переменным среды
var Env env

// CheckFlagEnv Метод проверяющий флаги
func CheckFlagEnv() {

	var host string
	var port string
	var dbUri string
	var dbName string
	var tableName string
	var production bool
	var secretKey string

	// сканируем env файл
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	var flagHost = flag.String("h", "", "host")
	var flagPort = flag.String("p", "", "port")
	var flagDbName = flag.String("dn", "", "dbName")
	var flagDbUri = flag.String("du", "", "dbUri")
	var flagTableName = flag.String("tn", "", "table_name")
	var flagProduction = flag.Bool("pr", false, "production")
	var flagSecretKey = flag.String("ske", "", "secret key for jwt")

	flag.Parse()

	if os.Getenv("HOST") != "" {
		host = os.Getenv("HOST")
	} else {
		host = "localhost"
	}

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	} else {
		port = "8080"
	}

	if os.Getenv("TABLE_NAME") != "" {
		tableName = os.Getenv("TABLE_NAME")
	} else {
		tableName = ""
	}

	if os.Getenv("DB_NAME") != "" {
		dbName = os.Getenv("DB_NAME")
	} else {
		dbName = ""
	}

	if os.Getenv("PRODUCTION") != "" {
		production, _ = strconv.ParseBool(os.Getenv("PRODUCTION"))
	} else {
		production = false
	}

	if os.Getenv("SECRET_KEY") != "" {
		secretKey = os.Getenv("SECRET_KEY")
	} else {
		secretKey = ""
	}

	if os.Getenv("DB_URI") != "" {
		dbUri = os.Getenv("DB_URI")
	} else {
		dbName = ""
	}

	if *flagHost != "" {
		host = *flagHost
	}

	if *flagPort != "" {
		port = *flagPort
	}

	if *flagDbName != "" {
		dbName = *flagDbName
	}

	if *flagProduction != false {
		production = *flagProduction
	}

	if *flagSecretKey != "" {
		secretKey = *flagSecretKey
	}

	if *flagTableName != "" {
		tableName = *flagTableName
	}

	if *flagDbUri != "" {
		dbUri = *flagDbUri
	}

	Env = env{
		Host:       host,
		Port:       port,
		DbUri:      dbUri,
		TableName:  tableName,
		DbName:     dbName,
		Production: production,
		SecretKey:  secretKey,
	}
}
