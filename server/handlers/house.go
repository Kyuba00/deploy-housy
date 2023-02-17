package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	housesdto "housy/dto/house"
	dto "housy/dto/result"
	"housy/models"
	"housy/repositories"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gorilla/mux"
	"gorm.io/datatypes"
)

type handlerHouse struct {
	HouseRepository repositories.HouseRepository
}

func HandlerHouse(HouseRepository repositories.HouseRepository) *handlerHouse {
	return &handlerHouse{HouseRepository}
}

func (h *handlerHouse) FindHouses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	houses, err := h.HouseRepository.FindHouses()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	for i, p := range houses {
		houses[i].Image = p.Image
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: houses}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerHouse) GetHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	house, err := h.HouseRepository.GetHouse(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseHouse(house)}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerHouse) CreateHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get image filepath
	dataContex := r.Context().Value("dataFile")
	filepath := dataContex.(string)

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")
	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "uploads"})
	if err != nil {
		fmt.Println(err.Error())
	}

	price, _ := strconv.Atoi(r.FormValue("price"))
	Bedroom, _ := strconv.Atoi(r.FormValue("Bedroom"))
	Bathroom, _ := strconv.Atoi(r.FormValue("Bathroom"))

	request := housesdto.HouseRequest{
		Name:        r.FormValue("name"),
		CityName:    r.FormValue("cityname"),
		Address:     r.FormValue("address"),
		TypeRent:    r.FormValue("type_rent"),
		Description: r.FormValue("description"),
		Area:        r.FormValue("area"),
		Amenities:   datatypes.JSON(r.FormValue("amenities")),
		Bedroom:     Bedroom,
		Price:       price,
		Bathroom:    Bathroom,
		Image:       resp.SecureURL,
	}

	// validation := validator.New()
	// err := validation.Struct(request)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	house := models.House{
		Name:        request.Name,
		CityName:    request.CityName,
		Address:     request.Address,
		Price:       request.Price,
		TypeRent:    request.TypeRent,
		Amenities:   request.Amenities,
		Bedroom:     request.Bedroom,
		Bathroom:    request.Bathroom,
		Description: request.Description,
		Area:        request.Area,
		Image:       resp.SecureURL,
	}

	// err := mysql.DB.Create(&product).Error
	house, err = h.HouseRepository.CreateHouse(house)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	house, _ = h.HouseRepository.GetHouse(house.ID)
	// house.Image = os.Getenv("PATH_FILE") + house.Image
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: house}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerHouse) DeleteHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	house, err := h.HouseRepository.GetHouse(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.HouseRepository.DeleteHouse(house)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerHouse) UpdateHouse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile") // add this code
	filename := dataContex.(string)             // add this code

	price, _ := strconv.Atoi(r.FormValue("price"))
	bedroom, _ := strconv.Atoi(r.FormValue("Bedroom"))
	bathroom, _ := strconv.Atoi(r.FormValue("Bathroom"))
	request := housesdto.HouseRequest{
		Name:        r.FormValue("name"),
		CityName:    r.FormValue("cityname"),
		Address:     r.FormValue("address"),
		TypeRent:    r.FormValue("type_rent"),
		Description: r.FormValue("description"),
		Area:        r.FormValue("area"),
		Amenities:   datatypes.JSON(r.FormValue("amenities")),
		Price:       price,
		Bedroom:     bedroom,
		Bathroom:    bathroom,
		Image:       filename,
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	house, err := h.HouseRepository.GetHouse(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Name != "" {
		house.Name = request.Name
	}

	if request.CityName != "" {
		house.CityName = request.CityName
	}

	if request.Address != "" {
		house.Address = request.Address
	}

	if request.Price != 0 {
		house.Price = request.Price
	}

	if request.TypeRent != "" {
		house.TypeRent = request.TypeRent
	}

	if request.Amenities != nil {
		house.Amenities = request.Amenities
	}

	if request.Bedroom != 0 {
		house.Bedroom = request.Bedroom
	}

	if request.Description != "" {
		house.Description = request.Description
	}

	if request.Area != "" {
		house.Area = request.Area
	}

	if request.Bathroom != 0 {
		house.Bathroom = request.Bathroom
	}

	if request.Image != "" {
		house.Image = request.Image
	}

	data, err := h.HouseRepository.UpdateHouse(house)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: data}
	json.NewEncoder(w).Encode(response)
}

func convertResponseHouse(u models.House) housesdto.ResponseHouse {
	return housesdto.ResponseHouse{
		ID:          u.ID,
		Name:        u.Name,
		CityName:    u.CityName,
		Address:     u.Address,
		Price:       u.Price,
		TypeRent:    u.TypeRent,
		Amenities:   u.Amenities,
		Bedroom:     u.Bedroom,
		Bathroom:    u.Bathroom,
		Image:       u.Image,
		Description: u.Description,
		Area:        u.Area,
	}
}
