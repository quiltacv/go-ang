package main
import ("github.com/jinzhu/gorm"
 				_ "github.com/jinzhu/gorm/dialects/postgres"
				"fmt"
				)

type UserModel struct{
  Id int `gorm:"primary_key";"AUTO_INCREMENT"`
  Name string `gorm:"size:255"`
  Address string `gorm:"type:varchar(100)"`
}


func main() {
  db, err := gorm.Open("postgres", "postgres://foxzi:@localhost/app-demo?sslmode=disable")
  if err != nil {
    panic("failed to connect database")
  }
  defer db.Close()
	fmt.Println("Begining")
  db.DropTableIfExists(&UserModel{})
  db.AutoMigrate(&UserModel{})

  user := &UserModel{Name: "Foxzi", Address: "New City"}
  newUser :=&UserModel{Name: "Martin", Address: "Old City"}
  db.Create(user)
  db.Save(newUser)
  db.Find(&user).Update("address", "HCM CITY")

  users := &UserModel{}
  count := 0
  db.Debug().Find(&users)
  fmt.Println(len(users))
  db.Table("user_models").Count(&count)
  fmt.Println(count)
  fmt.Println("Finished")
  fmt.Println("Hello World"[1])
}
