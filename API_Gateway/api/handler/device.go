package handler

import (
	proto "api_gateway/proto/device"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"google.golang.org/protobuf/encoding/protojson"
)

type DeviceClient struct {
	Client proto.DeviceControlServiceClient
	LG     *log.Logger
}

func NewDeviceClient(cl proto.DeviceControlServiceClient, lg *log.Logger) *DeviceClient {
	return &DeviceClient{Client: cl, LG: lg}
}

// @Router		/devices [post]
// @Summary		Add Devices
// @Description Bu endpoint yangi qurulma qo'shish uchun ishlatiladi
// @Security	BearerAuth
// @Tags		Device
// @Accept		json
// @Produce 	json
// @Param		body body proto.AddDeviceRequestSwag true "AddDeviceRequestSwag"
// @Success 	201 {object} map[string]interface{} "Device successfully added"
// @Failure 	400 {object} models.StandartError "Bad request error"
// @Failure 	403 {object} models.ForbiddenError "Forbidden error"
// @Failure 	500 {object} models.StandartError "Internal server error"
func (d *DeviceClient) AddDeviceApi(w http.ResponseWriter, r *http.Request) {
	var deviceReq proto.AddDeviceRequest

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		d.LG.Println("AddDeviceApi ERROR: POST: bodydan o'qib olishda xatolik...", err)
		http.Error(w, "POST: bodydan o'qib olishda xatolik...", http.StatusBadRequest)
		return
	}

	protojson.Unmarshal(bytes, &deviceReq)

	resp, err := d.Client.AddDevice(r.Context(), &deviceReq)
	if err != nil {
		log.Println(err)
		d.LG.Println("AddDeviceApi ERROR: POST: Serverdan ma'lumot olishda xatolik...", err)
		http.Error(w, "POST: Serverdan ma'lumot olishda xatolik...", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Println(err)
		d.LG.Println("AddDeviceApi ERROR: Ma'lumotni encode qilishda xatolik...", err)
		http.Error(w, "Ma'lumotni encode qilishda xatolik...", http.StatusInternalServerError)
		return
	}
}

// @Router		/devices/update [put]
// @Summary		Add Devices
// @Description Bu endpoint qurulmani yangilash uchun ishlatiladi
// @Security	BearerAuth
// @Tags		Device
// @Accept		json
// @Produce 	json
// @Param		body body proto.UpdateDeviceRequestSwag true "UpdateDeviceRequestSwag"
// @Success 	201 {object} map[string]interface{} "Device successfully updated"
// @Failure 	400 {object} models.StandartError "Bad request error"
// @Failure 	403 {object} models.ForbiddenError "Forbidden error"
// @Failure 	500 {object} models.StandartError "Internal server error"
func (d *DeviceClient) UpdateDeviceApi(w http.ResponseWriter, r *http.Request) {
	var deviceReq proto.UpdateDeviceRequest

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		d.LG.Println("UpdateDeviceApi ERROR: POST: bodydan o'qib olishda xatolik...", err)
		http.Error(w, "ERROR: %v", http.StatusBadRequest)
		return
	}

	protojson.Unmarshal(bytes, &deviceReq)

	resp, err := d.Client.UpdateDevice(r.Context(), &deviceReq)
	if err != nil {
		log.Println(err)
		d.LG.Println("UpdateDeviceApi ERROR:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Println(err)
		d.LG.Println("UpdateDeviceApi ERROR: Ma'lumotni encode qilishda xatolik...", err)
		http.Error(w, "Ma'lumotni encode qilishda xatolik...", http.StatusInternalServerError)
		return
	}
}

// @Router		/devices/delete [delete]
// @Summary		Add Devices
// @Description Bu endpoint qurulmani o'chirish uchun ishlatiladi
// @Security	BearerAuth
// @Tags		Device
// @Accept		json
// @Produce 	json
// @Param		body body proto.DeleteDeviceRequest true "DeleteDeviceRequest"
// @Success 	201 {object} map[string]interface{} "Device successfully deleted"
// @Failure 	400 {object} models.StandartError "Bad request error"
// @Failure 	403 {object} models.ForbiddenError "Forbidden error"
// @Failure 	500 {object} models.StandartError "Internal server error"
func (d *DeviceClient) DeleteDeviceApi(w http.ResponseWriter, r *http.Request) {
	var deviceReq proto.DeleteDeviceRequest

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		d.LG.Println("DeleteDeviceApi ERROR: POST: bodydan o'qib olishda xatolik...", err)
		http.Error(w, "POST: bodydan o'qib olishda xatolik...", http.StatusBadRequest)
		return
	}

	protojson.Unmarshal(bytes, &deviceReq)

	resp, err := d.Client.DeleteDevice(r.Context(), &deviceReq)
	if err != nil {
		log.Println(err)
		d.LG.Println("DeleteDeviceApi ERROR: POST: Serverdan ma'lumot olishda xatolik...", err)
		http.Error(w, "POST: Serverdan ma'lumot olishda xatolik...", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Println(err)
		d.LG.Println("DeleteDeviceApi ERROR: Ma'lumotni encode qilishda xatolik...", err)
		http.Error(w, "Ma'lumotni encode qilishda xatolik...", http.StatusInternalServerError)
		return
	}
}

// @Router		/control [post]
// @Summary		Create control
// @Description Bu endpoint qurulmani boshqarish uchun ishlatiladi
// @Security	BearerAuth
// @Tags		Device
// @Accept		json
// @Produce 	json
// @Param		body body proto.ControlRequest true "ControlRequest"
// @Success 	201 {object} map[string]interface{} "Commad sccessfully created"
// @Failure 	400 {object} models.StandartError "Bad request error"
// @Failure 	403 {object} models.ForbiddenError "Forbidden error"
// @Failure 	500 {object} models.StandartError "Internal server error"
func (d *DeviceClient) SendControlCommand(w http.ResponseWriter, r *http.Request) {
	var controlReq proto.ControlRequest

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		d.LG.Println("SendControlCommand ERROR: ", err)
		http.Error(w, "-1-", http.StatusBadRequest)
		return
	}

	protojson.Unmarshal(bytes, &controlReq)

	resp, err := d.Client.SendControlCommand(r.Context(), &controlReq)
	if err != nil {
		log.Println(err)
		d.LG.Println("SendControlCommand ERROR:", err)
		http.Error(w, "-2-", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Println(err)
		d.LG.Println("SendControlCommand ERROR:", err)
		http.Error(w, "Ma'lumotni encode qilishda xatolik...", http.StatusInternalServerError)
		return
	}
}
