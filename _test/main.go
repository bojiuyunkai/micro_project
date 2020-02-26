package main

import (
    //"github.com/spf13/viper"
    "hz_project/internal/pkg/config"
    //"github.com/spf13/viper"
    "fmt"

   
)
type Db struct{
    Dbname string 
    Port int 
    Pass string
    Host string
    User string
}

func main() {
    viper, _ := config.New("/Users/gaozhanghu/Documents/hz_project/internal/pkg/conf/dev/db.yml")
   /* v   := viper.New()
    v.AddConfigPath(".")
    v.SetConfigFile(string("/Users/gaozhanghu/Documents/hz_project/internal/pkg/conf/dev/db.yml"))

    if err := v.ReadInConfig(); err == nil {
        fmt.Printf("use config file -> %s\n", v.ConfigFileUsed())
    } else {
       fmt.Println(err);
    }
    o := new(Db)
    fmt.Println(v.UnmarshalKey("db_kernel",o));
    fmt.Println(o.Dbname);*/

    o := new(Db)
    if err:= viper.UnmarshalKey("RES_DB_KERNEL", o); err != nil {
        fmt.Println(err);
    }
    fmt.Println(o);
}
