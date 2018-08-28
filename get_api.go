package main
import (
        "log"
        "net/http"
        "io/ioutil"
				"fmt"
        "encoding/json"
        "time"
        // "github.com/jinzhu/gorm"
   			// 	_ "github.com/jinzhu/gorm/dialects/postgres"
				)

  type Garage []struct {
  	ID          int         `json:"id"`
  	Name        string      `json:"name"`
  	Description string      `json:"description"`
  	Phone       string      `json:"phone"`
  	Address     string      `json:"address"`
  	OpenTime    time.Time   `json:"open_time"`
  	CloseTime   time.Time   `json:"close_time"`
  	Source      string      `json:"source"`
  	IsPending   interface{} `json:"is_pending"`
  	IsDeleted   interface{} `json:"is_deleted"`
  	CreatedAt   time.Time   `json:"created_at"`
  	UpdatedAt   time.Time   `json:"updated_at"`
  	CityID      int         `json:"city_id"`
  	Longitude   float64     `json:"longitude"`
  	Latitude    float64     `json:"latitude"`
  }


func main() {
  var request_url = "http://fixmybike.herokuapp.com//api/v1/garages/find_garages?token=2e900f7419c3d358a28f48cc9ee5803a&name=a"
  fmt.Println("Signed URL:", request_url)
  resp, err := http.Get(request_url)

  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
  if resp.StatusCode >= 200 && resp.StatusCode <=299 {
    fmt.Println("HTTP 2xx")
    robots, err := ioutil.ReadAll(resp.Body)
    if err != nil {
    		log.Fatal(err)
    	}
    garages := &Garage{}
    json.Unmarshal([]byte(robots), garages)
    fmt.Println(len(*garages))
    for i := 0; i < len(*garages); i++ {
      fmt.Println((*garages)[i])
    }
  } else {
    fmt.Println("Broken")
  }
}
