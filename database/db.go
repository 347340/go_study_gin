package database

import (
   "database/sql"
   "fmt"
   "github.com/joho/godotenv"
   _ "github.com/lib/pq"
   "os"
   "strconv"
)

// 在外部创建以使其全局
var Db *sql.DB

// 确保函数以大写开头以在目录外部调用
func ConnectDatabase() {

   // 默认是.env所以我们不用写
   err := godotenv.Load()
   if err != nil {
      fmt.Println("Error is occurred  on .env file please check")
   }
   // 读取.env 文件
   host := os.Getenv("HOST")
   // 不要忘记转换 int，因为 port 是 int 类型
   port, _ := strconv.Atoi(os.Getenv("PORT"))
   user := os.Getenv("USER")
   dbname := os.Getenv("DB_NAME")
   pass := os.Getenv("PASSWORD")

   // 设置 postgres sql 打开它
   psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
       host, port, user, dbname, pass)
   db, errSql := sql.Open("postgres", psqlSetup)
   if errSql != nil {
      fmt.Println("There is an error while connecting to the database ", err)
      panic(err)
   } else {
      Db = db
      fmt.Println("Successfully connected to database!")
   }
}
