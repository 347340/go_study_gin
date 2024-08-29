package v1

import (
  "encoding/json"
  "net/http"
  "strconv"
	"github.com/347340/go_study_gin/models"

)

func GetStudent(w http.ResponseWriter, r *http.Request) {
  idStr := r.URL.Query().Get("id")
  id, _ := strconv.Atoi(idStr)

  student, _:= models.GetStudentByID(id)

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(student)
}
