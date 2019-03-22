package cmd

import (
	"fmt"
	"os"

	"github.com/everywan/foundation-go/database"
	flog "github.com/everywan/foundation-go/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var rootCmd = &cobra.Command{
	Use:   "demo",
	Short: "Short Ins",
	Long:  `Long Ins`,
}

func init() {
	// 读取配置
	cobra.OnInitialize(initConfig)
	// 添加配置参数
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file")
}

// Execute is ..
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	if cfgFile == "" {
		return
	}
	viper.SetConfigFile(cfgFile)
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file: ", viper.ConfigFileUsed())
	}
}

type (
	// DatabaseOps is 数据库配置
	DatabaseOps struct {
		Mysql database.MysqlOptions `mapstructure:"mysql" yaml:"mysql"`
	}
	// ApplicationOps is ...
	ApplicationOps struct {
		Logger   flog.Options `mapstructure:"logger" yaml:"logger"`
		Database DatabaseOps  `mapstructure:"database" yaml:"database"`
		Server   ServerOps    `mapstructure:"server" yaml:"server"`
	}
)

// Load 使用viper加载配置文件
func (opts *ApplicationOps) Load() {
	err := viper.Unmarshal(opts)
	if err != nil {
		// 加入log组件, 改用log记录
		fmt.Printf("failed to parse config file: %s", err)
	}
}

func loadApplocationOps() *ApplicationOps {
	opts := &ApplicationOps{}
	opts.Load()
	return opts
}
