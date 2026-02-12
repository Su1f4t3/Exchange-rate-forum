package config

import (
	"log"

	"github.com/spf13/viper"
)

// 嵌套结构体 Config 包含 App 和 Database 两个字段，分别对应应用程序和数据库的配置
type Config struct {
	App struct {
		Name string // 应用程序名称
		Port string // 应用程序运行端口
	}
	Database struct {
		Dsn          string
		MaxIdleConns int
		MaxOpenConns int
	}
	Redis struct {
		Addr string
	}
}

// AppConfig 是一个全局变量，此时值为 nil（空）
var AppConfig *Config // *Config 是一个指向 Config 结构体的指针

func InitConfig() {
	viper.SetConfigName("config")   // 设置配置文件名（不带扩展名）
	viper.SetConfigType("yaml")     // 设置配置文件类型
	viper.AddConfigPath("./config") // 设置配置文件所在目录

	// 读取配置文件到 viper 的内存中
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// 创建一个新的 Config 结构体实例，然后把它的内存地址赋值给 AppConfig
	AppConfig = &Config{} // 现在 AppConfig 指向一个实际的 Config 结构体

	// config.yml 的数据是在 viper.Unmarshal() 调用时被"复制"到 AppConfig
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}

	InitDB()
	InitRedis()
}
