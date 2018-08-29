package main
import (
        "bufio"
        "os"
        "log"
        "net/http"
        "io/ioutil"
				"fmt"
        "encoding/json"
        "time"
        "github.com/jinzhu/gorm"
   				_ "github.com/jinzhu/gorm/dialects/postgres"
				)

type Garages []struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Phone       string      `json:"phone"`
	Address     string      `json:"address"`
	OpenTime    time.Time   `json:"open_time"`
	CloseTime   time.Time   `json:"close_time"`
	Source      string      `json:"source"`
	IsPending   string      `json:"is_pending"`
	IsDeleted   string      `json:"is_deleted"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	CityID      int         `json:"city_id"`
	Longitude   float64     `json:"longitude"`
	Latitude    float64     `json:"latitude"`
}

type Garage struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Phone       string      `json:"phone"`
	Address     string      `json:"address"`
	OpenTime    time.Time   `json:"open_time"`
	CloseTime   time.Time   `json:"close_time"`
	Source      string      `json:"source"`
	IsPending   string      `json:"is_pending"`
	IsDeleted   string      `json:"is_deleted"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
	CityID      int         `json:"city_id"`
	Longitude   float64     `json:"longitude"`
	Latitude    float64     `json:"latitude"`
}

func fetchData() *http.Response {
  fmt.Print("Input key word search: ")
  reader := bufio.NewReader(os.Stdin)
  char, _, err_r := reader.ReadRune()
  if err_r != nil {
    fmt.Println(err_r)
  }

  var request_url = "http://fixmybike.herokuapp.com//api/v1/garages/find_garages?token=2e900f7419c3d358a28f48cc9ee5803a&name=" + string(char)
  fmt.Println("Signed URL:", request_url)
  resp, err := http.Get(request_url)
  if err != nil {
    log.Fatal(err)
  }
  return resp
}

func main() {

  db, e := gorm.Open("postgres", "postgres://foxzi:@localhost/app-demo?sslmode=disable")
  if e != nil {
    panic("failed to connect database")
  }
  defer db.Close()
  fmt.Println("Connect DB")
  db.DropTableIfExists(&Garage{})
  db.AutoMigrate(&Garage{})

  dataString, err := ioutil.ReadAll(fetchData().Body)
  if err==nil {
    garages := &Garages{}
    json.Unmarshal([]byte(dataString), &garages)
    fmt.Println(*garages)
    for i := 0; i < len(*garages); i++ {
      fmt.Println(*garages)
      db.Create(&Garage{
        ID: (*garages)[i].ID,
      	Name: (*garages)[i].Name,
      	Description: (*garages)[i].Description,
      	Phone: (*garages)[i].Phone,
      	Address: (*garages)[i].Address,
      	OpenTime: (*garages)[i].OpenTime,
      	CloseTime: (*garages)[i].CloseTime,
      	Source: (*garages)[i].Source,
      	IsPending: (*garages)[i].IsPending,
      	IsDeleted: (*garages)[i].IsDeleted,
      	CreatedAt: (*garages)[i].CreatedAt,
      	UpdatedAt: (*garages)[i].UpdatedAt,
      	CityID: (*garages)[i].CityID,
      	Longitude: (*garages)[i].Longitude,
      	Latitude: (*garages)[i].Latitude,
      })
    }
    count := 0
    db.Table("garages").Count(&count)
    fmt.Println(count)
  }
}
