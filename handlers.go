package main

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
)

func GetServerInfo(w http.ResponseWriter, r *http.Request, returnOriginal bool) {
    vars := mux.Vars(r)
    serverCode := vars["server_code"]

	if vanityCode, ok := vanityCodes.VanityCodes[serverCode]; ok {
        serverCode = vanityCode
    }

    response, err := http.Get("https://servers-frontend.fivem.net/api/servers/single/" + serverCode)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer response.Body.Close()

    var responseData map[string]interface{}
    if err := json.NewDecoder(response.Body).Decode(&responseData); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    selectedData := map[string]interface{}{
        "endpoint": responseData["EndPoint"],
    }

	if data, ok := responseData["Data"].(map[string]interface{}); ok {
		selectedData["hostname"] = data["hostname"]
		playerData := map[string]interface{}{
			"count":   data["clients"],
			"self-reported": data["selfReportedClients"],
			"list": data["players"],
		}
		selectedData["players"] = playerData
		selectedData["slots"] = data["sv_maxclients"]
		selectedData["last-seen"] = data["lastSeen"]
		if lastSeen, ok := data["lastSeen"].(string); ok {
			selectedData["online"] = ServerisOnline(lastSeen)
		} else {
			selectedData["online"] = nil
		}
		selectedData["private"] = data["private"]
		if returnOriginal {
			selectedData["original"] = responseData
		}
		selectedData["connect"] = "https://cfx.re/join/" + serverCode
	}

	jsonData, err := json.Marshal(selectedData)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Credit", "https://github.com/akatiggerx04")
    w.Write(jsonData)
}