package main

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

// func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
// 	// read the json into var
// 	var requestPayload JSONPayload
// 	_ := app.readJSON(w, r, &requestPayload)
// }
