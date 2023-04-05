package main

import (
	"github.com/afrizal423/go-intervals/controllers"
	"github.com/gin-gonic/gin"
)

// // Payload adalah struktur untuk menyimpan data water dan wind.
// type Payload struct {
// 	Water float64 `json:"water"`
// 	Wind  float64 `json:"wind"`
// }

// // Response adalah struktur untuk menyimpan respons dari server.
// type Response struct {
// 	Water float64 `json:"water"`
// 	Wind  float64 `json:"wind"`
// }

// func main() {
// 	// Menjalankan fungsi kirimData setiap 15 detik.
// 	for range time.Tick(5 * time.Second) {
// 		kirimData()
// 		fmt.Println()
// 	}
// }

// func kirimData() {
// 	// Membuat payload dengan angka acak antara 1-100 untuk water dan wind.
// 	payload := Payload{
// 		Water: float64(rand.Intn(100) + 1),
// 		Wind:  float64(rand.Intn(100) + 1),
// 	}
// 	//fmt.Println("isi payload: ", payload)

// 	// Mengirim POST request dengan payload dalam format JSON.
// 	jsonPayload, _ := json.Marshal(payload)
// 	response, _ := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", strings.NewReader(string(jsonPayload)))

// 	// Membaca respons dari server dan menampilkan status water dan wind yang telah ditentukan.
// 	// opsi 1
// 	var result Response
// 	json.NewDecoder(response.Body).Decode(&result)
// 	// opsi 2 (dgn id juga)
// 	// body, err := ioutil.ReadAll(response.Body)
// 	// if err != nil {
// 	// 	log.Fatalln(err)
// 	// }
// 	// response.Body.Close()
// 	Pretty(result)
// 	statusWater, statusWind := CekStatus(result)
// 	println("Status water :", statusWater)
// 	println("Status wind :", statusWind)
// }

// func CekStatus(payload Response) (string, string) {
// 	// Menentukan status water dan wind sesuai dengan ketentuan yang telah diberikan.
// 	statusWater := "aman"
// 	if payload.Water > 8 {
// 		statusWater = "bahaya"
// 	} else if payload.Water >= 6 {
// 		statusWater = "siaga"
// 	}
// 	statusWind := "aman"
// 	if payload.Wind > 15 {
// 		statusWind = "bahaya"
// 	} else if payload.Wind >= 7 {
// 		statusWind = "siaga"
// 	}

// 	return statusWater, statusWind
// }

// // Pretty display the claims licely in the terminal
// func Pretty(data interface{}) {
// 	b, err := json.MarshalIndent(data, "", " ")
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	fmt.Println(string(b))
// }

var (
	router = gin.Default()
)

func main() {
	router.GET("/cuaca", controllers.Wheater)

	router.Run(":8080")
}
