package main

import (
	"net/http"
	"encoding/json"
	"net/url"
	"bytes"
	"net/http/httptest"
	"testing"
	"strconv"

)

func TestGetBirdHandler(t *testing.T){
	mockStore := InitMockStore()
	mockStore.On("GetBirds").Return([]*Bird{{"sparrow","small bird"}},nil).Once()
	req, err := http.NewRequest("GET", "", nil)
	if err!=nil{
		t.Fatal(err)
	}
	
	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(getBirdHandler)
	hf.ServeHTTP(recorder, req)
	
	if status := recorder.Code; status!=http.StatusOK{
		t.Errorf("got %v expected %v", recorder.Code, http.StatusOK)
	}
	
	expected := Bird{"sparrow","small bird"}
	b :=[]Bird{}
	err = json.NewDecoder(recorder.Body).Decode(&b)
	
	if err!=nil{
		t.Fatal(err)
	}
	
	actual := b[0]
	if actual != expected{
		t.Errorf("handler returned %v expected %v", actual,expected)
	}
	}


func TestBirdCreateHandler(t *testing.T){
	mockStore := InitMockStore()
	mockStore.On("CreateBird",&Bird{"pigeon","bird"}).Return(nil)
	form := newCreateBirdForm()
	
	req, err := http.NewRequest("POST","", bytes.NewBufferString(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length",strconv.Itoa(len(form.Encode())))
	if err!=nil{
		t.Fatal(err)
	}
	
	recorder :=httptest.NewRecorder()
	hf := http.HandlerFunc(createBirdHandler)
	hf.ServeHTTP(recorder, req)
	if status:= recorder.Code; status!=http.StatusFound{
		t.Errorf("%v expected %v found",http.StatusFound, recorder.Code)
	}

	
	
}

func newCreateBirdForm() *url.Values{
	form := url.Values{}
	form.Set("species", "pigeon")
	form.Set("description","bird")
	return &form
}


