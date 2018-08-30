package main
import (
        "os"
        "bufio"
        "github.com/valyala/fasthttp"
				"fmt"
        "time"
        "encoding/json"
        "github.com/jinzhu/gorm"
          _ "github.com/jinzhu/gorm/dialects/postgres"
				)

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

type Garages []Garage
var db *gorm.DB

func fetchDate(url string) *Garages{
  req := fasthttp.AcquireRequest()
  req.SetRequestURI(url)

  resp := fasthttp.AcquireResponse()
  client := &fasthttp.Client{}
  client.Do(req, resp)
  bodyBytes := resp.Body()
  garages := &Garages{}

  json.Unmarshal([]byte(bodyBytes), &garages)
  return garages
}


func main() {
  db, e := gorm.Open("postgres", "postgres://foxzi:@localhost/app-demo?sslmode=disable")
  if e != nil {
    panic("failed to connect database")
  }
  defer db.Close()
  if !db.HasTable(&Garage{}) {
    fmt.Println("Create")
    db.DropTableIfExists(&Garage{})
    db.AutoMigrate(&Garage{})
  }
  fmt.Print("Fetch new data (Y/N...): ")
  reader := bufio.NewReader(os.Stdin)
  char, _, err_r := reader.ReadRune()
  if err_r != nil {
    fmt.Println(err_r)
  }
  if (string(char) == "Y" || string(char) == "y") {
    fmt.Print("Input Key: ")
    reader := bufio.NewReader(os.Stdin)
    key, _ := reader.ReadString('\n')
    var url = "http://fixmybike.herokuapp.com//api/v1/garages/find_garages?token=2e900f7419c3d358a28f48cc9ee5803a&name=" + string(key)
    fmt.Println("Begin: ", url)
    var garages = fetchDate(url)
    for i := 0; i < len(*garages); i++ {
      fmt.Println("======================")
      db.Create(&Garage{
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
  }
  count := 0
  db.Table("garages").Count(&count)
  fmt.Println(count)
  fmt.Println("===Query===")
  garages_qr := []Garage{}
  db.Debug().Where("name LIKE ?", "%Al%").Find(&garages_qr)
  fmt.Println(len(garages_qr))
  for i:=0; i < len(garages_qr); i++ {
      fmt.Println(garages_qr[i].Name)
  }
}
