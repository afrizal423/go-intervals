package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/afrizal423/go-intervals/models"
	"github.com/afrizal423/go-intervals/utils"
	"github.com/gin-gonic/gin"
)

func Wheater(c *gin.Context) {
	// Menjalankan fungsi kirimData setiap 15 detik.
	for range time.Tick(15 * time.Second) {
		// cuaca()
		payload := models.Cuaca{
			Water: float64(rand.Intn(100) + 1),
			Wind:  float64(rand.Intn(100) + 1),
		}
		//fmt.Println("isi payload: ", payload)

		// Mengirim POST request dengan payload dalam format JSON.
		jsonPayload, _ := json.Marshal(payload)
		response, _ := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", strings.NewReader(string(jsonPayload)))

		// Membaca respons dari server dan menampilkan status water dan wind yang telah ditentukan.
		// opsi 1
		var result models.Cuaca
		json.NewDecoder(response.Body).Decode(&result)
		// opsi 2 (dgn id juga)
		// body, err := ioutil.ReadAll(response.Body)
		// if err != nil {
		// 	log.Fatalln(err)
		// }
		// response.Body.Close()
		utils.Pretty(result)
		statusWater, statusWind := utils.CekStatus(result)
		println("Status water :", statusWater)
		println("Status wind :", statusWind)
		fmt.Println()
	}
}
