package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"kalkulatorWebService/model"
	"fmt"
	"strconv"
)

//Welcome Message
func Welcome(c *gin.Context) {
	result := gin.H{
		"message": "Assalamualaikum, semangat gais",
	}
	c.JSON(http.StatusOK, result)
	return
}


func Hitung(c *gin.Context){
	var(
		angka model.Format
	)

	if err:= c.Bind(&angka); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	angka1:=angka.Angka1
	angka2:=angka.Angka2
	_result := model.Hasil{Pengurangan:angka1-angka2,Pertambahan:angka1+angka2,Perkalian:angka1*angka2,Pembagian:angka1/angka2}
	fmt.Println(_result)
	//result = gin.H{
	//	"result": _result,
	//	"status":"mantap bosku",
	//}
	c.JSON(http.StatusOK, gin.H{
		"result":_result,
	})
}

func HitungQuery(c *gin.Context){



	query1 := c.DefaultQuery("angka1", "1")
	query2 := c.DefaultQuery("angka2","1")

	//c.String(http.StatusOK, "Hello %s %s", firstname, lastname)

	angka1,_:= strconv.ParseFloat(query1,64)
	angka2,_:= strconv.ParseFloat(query2,64)
	_result := model.Hasil{Pengurangan:angka1-angka2,Pertambahan:angka1+angka2,Perkalian:angka1*angka2,Pembagian:angka1/angka2}
	fmt.Println(_result)
	//result = gin.H{
	//	"result": _result,
	//	"status":"mantap bosku",
	//}
	c.JSON(http.StatusOK, gin.H{
		"result":_result,
	})
}
//func GetData(c *gin.Context){
//	url := "https://dteti.au-syd.mybluemix.net/api"
//	response, err := http.Get(url)
//	if err != nil {
//		fmt.Printf("The HTTP request failed with error %s\n", err)
//	} else {
//		data, _ := ioutil.ReadAll(response.Body)
//		hasil := &Data{}
//		err := json.Unmarshal([]byte(string(data)), hasil)
//		fmt.Println(err)
//		c.JSON(http.StatusOK, gin.H{
//			"result":hasil.Message,
//		})
//		return
//	}
//}

type Data struct {
	Message string
}