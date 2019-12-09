package controller

import (
	"dto"
	"encoding/json"
	"fmt"
	"net/http"
	"encoding/csv"
	"strconv"
	"strings"
)

func StatController(w http.ResponseWriter, r *http.Request) {
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	if origin := r.Header.Get("Origin"); origin != "" {
      w.Header().Set("Access-Control-Allow-Origin", "*")
      w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
      w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
      w.Header().Set("Access-Control-Expose-Headers", "Authorization")
  }
	pathVar := strings.Split(r.URL.Path, "/")
  url := dto.ServerProperties.CsvFile+pathVar[2]


	data, err := readCSVFromUrl(url)
	if err != nil {
		panic(err)
	}
	var dutys []dto.StatDto
	for idx, row := range data {
		// skip header
		if idx <= 0 {
			continue
		}
		//
		// if idx == 6 {
		// 	break
		// }

		if row[1] != ""{
			numberOfMembers, _ := strconv.Atoi(row[8])
			numberOfAttendance := strings.Split(strings.Replace(row[9], " ", "", -1), ",")
			var countOfAttendance = 0
			for i := 0 ; i < len(numberOfAttendance) ; i++ {
				if(!strings.Contains(numberOfAttendance[i], "_") && !strings.Contains(numberOfAttendance[i], " ") && numberOfAttendance[i] != ""){
					countOfAttendance++
				}
			}
			var isFull = countOfAttendance >= numberOfMembers
			if isFull {
				isFull = !strings.Contains(row[10], "*")
			}
			temp := dto.StatDto{
				Date :strings.Split(row[1], " ")[0],
				Time  :row[2],
				ActivityName :row[3],
				Location :row[4],
				Organiser :row[5],
				AnotherDept:row[6],
				Note	:row[7],
				NumberOfMember :numberOfMembers,
				NumberOfAttendance :countOfAttendance,
				ListOfMembers :row[9],
				Owed :row[10],
				DutyKey :row[11],
				ReceivedDate :row[12],
				IsFull : isFull,
			}
			dutys = append(dutys,temp)
		}else{
			if strings.Contains(row[0],"本月取消") {
				break
			}
		}
	}

	jsonResult, _ := json.Marshal(dutys)
	fmt.Fprintf(w, "%s", jsonResult)
}


func readCSVFromUrl(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ','
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}
