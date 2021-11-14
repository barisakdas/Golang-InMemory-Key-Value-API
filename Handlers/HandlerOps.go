package Handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	. "ys/Helpers"
	. "ys/Log"
	. "ys/Models"
)

var cache = NewMemoryCache()

// GetAllEndpoints : API için tanımlanmış tüm endpointleri ve açıklamalarını içeren json dosyasını son kullanıcıya dönecek.
func GetAllEndpoints(res http.ResponseWriter, req *http.Request) {
	ServerRequestLog("GetAllEndpoints", "/")
	file, _ := ioutil.ReadFile("endpoints.json")
	res.Write(file)
}

func GetAllKeys(res http.ResponseWriter, req *http.Request) {
	ServerRequestLog("GetAllKeys", "/get-all-keys")
	jsonResp, _ := json.Marshal(cache.GetAll())

	res.Write(jsonResp)
}

func GetKey(res http.ResponseWriter, req *http.Request) {
	ServerRequestLog("GetKey", "/get")
	key := req.URL.Query().Get("key")
	resp := make(map[string]string)
	data, _ := cache.Get(key)

	if data != "" {
		resp["key"] = key
		resp["value"] = data
	} else {
		resp["message"] = "Resource Not Found"
	}
	jsonResp, _ := json.Marshal(resp)

	res.Write(jsonResp)
}

func SetKeys(res http.ResponseWriter, req *http.Request) {
	ServerRequestLog("SetKeys", "/set")
	reqModel, err := ParseRequestToModel(req)
	resp := make(map[string]string)
	if err != nil {
		CreateLogJson("Error", "GetAllOperations", "There is an error while converting json input to golang model.", err.Error())
	}
	if reqModel != (DataModel{}) {
		err := cache.Set(reqModel.Key, reqModel.Value)
		if err != "" {
			CreateLogJson("Error", "SetKeys", "Set key to cache.", "A problem occurred while setting cached data.")
			resp["message"] = err
		} else {
			CreateLogJson("Success", "SetKeys", "Set key to cache.", "Setting key to cache is success")
			resp["message"] = "Setting key to cache is success."
		}
	} else {
		CreateLogJson("Error", "GetAllOperations", "There is no element in data. Key and value are requirement", err.Error())
		resp["message"] = "There is no element in data. Key and value are requirement."
	}

	jsonResp, _ := json.Marshal(resp)
	res.Write(jsonResp)
}

func DeleteAllKeys(res http.ResponseWriter, req *http.Request) {
	ServerRequestLog("DeleteAllKeys", "/delete-all")
	cache.DeleteAll() // Cache üzerindeki tüm kayıtları silecek olan fonksiyon
	resp := make(map[string]string)
	resp["message"] = "Deleted all data on cache."
	jsonResp, _ := json.Marshal(resp)
	res.Write(jsonResp)
	//CreateLogJson("Success", "DeleteAllKeys", "Delete all cache data", "Deleted all data on cache.")
}
