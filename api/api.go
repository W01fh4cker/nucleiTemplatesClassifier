package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"nucleiTemplatesClassifier/utils"
	"strconv"
)

func WriteJSONResponse(w http.ResponseWriter, targetDir string, years string, severity string) {
	cves := utils.GetCVEsInfo(targetDir, years, severity)
	jsonData, err := json.MarshalIndent(cves, "", "    ")
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func ApiStart(targetDir string, years string, severity string, port int) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		WriteJSONResponse(w, targetDir, years, severity)
	})
	fmt.Printf("[+] Api Server is running on port %d...\n[+] Try:\n", port)
	ipstr := utils.GetLocalIPStr()
	for _, localip := range ipstr {
		apiUrl := "    - http://" + localip + ":" + strconv.Itoa(port)
		fmt.Println(apiUrl)
	}
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
