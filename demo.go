package main
import ("github.com/jinzhu/gorm"
 				_ "github.com/jinzhu/gorm/dialects/postgres"
				"fmt"
				"strconv"
				)

type Product struct {
  gorm.Model
  Code string
  Price uint
}

type Category struct {
  gorm.Model
  Name string
}

func main() {
  db, err := gorm.Open("postgres", "postgres://foxzi:@localhost/app-demo?sslmode=disable")
	if err != nil {
    panic("failed to connect database")
  }
  defer db.Close()
	fmt.Println("Begining")
  if (db.HasTable(&Category{})==false) {
    db.AutoMigrate(&Category{})
    db.Create(&Category{Name: "Apple"})
    db.Create(&Category{Name: "LG"})
    db.Create(&Category{Name: "SAMSUNG"})
  }
  if (db.HasTable(&Product{})==false) {
    db.AutoMigrate(&Product{})
    db.Model(&Product{}).AddForeignKey("category_id", "categories(id)", "RESTRICT", "RESTRICT")
  	for i := 0; i < 100; i++ {
  		db.Create(&Product{Code: "LA65"+strconv.Itoa(i), Price: 999})
  		fmt.Println("LA65"+strconv.Itoa(i))
  	}
  }
	fmt.Println("Finished")
	// var product Product
	// db.First(Product, 1)
	// fmt.Println(db.First(Product, 1))
}
