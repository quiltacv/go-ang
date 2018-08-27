package main
import (
        "log"
        "net/http"
        "io/ioutil"
				"fmt"
        // "github.com/jinzhu/gorm"
   			// 	_ "github.com/jinzhu/gorm/dialects/postgres"
				)

type Garage struct{
  id int `gorm:"primary_key";"AUTO_INCREMENT"`
  name string `gorm:"size:255"`
  address string `gorm:"type:varchar(100)"`
  description string `gorm:"type:varchar(100)"`
  phone string `gorm:"type:varchar(20)"`
}


func main() {
  var request_url = "http://fixmybike.herokuapp.com//api/v1/garages/find_garages?token=2e900f7419c3d358a28f48cc9ee5803a&name=Waino%20Sipes%20DVM"
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
    rep := string(robots)
    fmt.Println(rep)
  } else {
    fmt.Println("Broken")
  }
}
