package main

import (
	"fmt"
	"net/http"
	"encoding/json"

)

type Bird struct{
	Species string `json:"species"`
	Description string `json:"description"`
}



func getBirdHandler(w http.ResponseWriter, r *http.Request){
	birdRes, err := store.GetBirds()
	birdListBytes, err := json.Marshal(birdRes)
	if err!=nil{
		fmt.Println(fmt.Errorf("Error : %v",err))
		w.WriteHeader(http.StatusInternalServerError)
		return
			
		
	}
	w.Write(birdListBytes)
	
}

func createBirdHandler(w http.ResponseWriter, r *http.Request){
	bird:=Bird{}
	err :=r.ParseForm()
	if err!=nil{
		fmt.Println(fmt.Errorf("Error: %v",err))
			w.WriteHeader(http.StatusInternalServerError)
			return
	}
	bird.Species=r.Form.Get("species")
	bird.Description=r.Form.Get("description")
	err=store.CreateBird(&bird)
	if err!=nil{
		fmt.Println(err)
	}
//	http.Redirect(w, r, "/assets/", http.StatusFound)
}

