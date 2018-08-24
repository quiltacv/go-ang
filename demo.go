package main
import ("github.com/jinzhu/gorm"
 				_ "github.com/jinzhu/gorm/dialects/postgres"
				"fmt"
        "strconv"
        "math/rand"
				)

type Product struct {
  gorm.Model
  Name string `gorm:size:255`
  Code string
  Price uint
  Category Category `gorm:"foreignkey:CategoryID;association_foreignkey:Refer"`
  CategoryID uint
}

type Category struct {
  gorm.Model
  Name string `gorm:size:255`
}

func main() {
  db, err := gorm.Open("postgres", "postgres://foxzi:@localhost/go-lang?sslmode=disable")
	if err != nil {
    panic("failed to connect database")
  }
  defer db.Close()
	fmt.Println("Begining")
  db.DropTableIfExists(&Category{})
  db.DropTableIfExists(&Product{})

  db.AutoMigrate(&Category{}, &Product{})

  cate1 := &Category{Name: "Apple"}
  cate2 := &Category{Name: "LG"}
  cate3 := &Category{Name: "SAMSUNG"}
  db.Create(cate1)
  db.Create(cate2)
  db.Create(cate3)

	for i := 0; i < 100; i++ {
    var category = &Category{}
    switch i%3 {
    case 0: db.Where("Name = ?", "Apple").First(&category)
    case 2: db.Where("Name = ?", "LG").First(&category)
    default: db.Where("Name = ?", "SAMSUNG").First(&category)
    }
		db.Create(&Product{Code: "LA65"+strconv.Itoa(i), Name: category.Name+ "-"+ strconv.Itoa(i), Price: 999, CategoryID: category.Model.ID})
		fmt.Println("LA65"+strconv.Itoa(i))
	}
  products := []Product{}
  db.Debug().Find(&products)
  fmt.Println(len(products))
  for i := 0; i< len(products); i++ {
    fmt.Println(products[i].Code, "--" ,products[i].Name)
  }

	fmt.Println("Finished")
}
