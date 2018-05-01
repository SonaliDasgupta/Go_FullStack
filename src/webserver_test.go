package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"io/ioutil"
	_ "net/url"
	
	
	

)

func TestHandler(t *testing.T){
	req,err:=http.NewRequest("GET", "", nil)
	if err!=nil{
		t.Fatal(err)
	}
	recorder :=httptest.NewRecorder()
	hf:= http.HandlerFunc(handler)
	hf.ServeHTTP(recorder, req)
	if status:=recorder.Code; status!=http.StatusOK{
		t.Errorf("handler returned status %v want %v",status, http.StatusOK)
	}
	expected := `hello world`
	actual := recorder.Body.String()
	if actual!=expected{
		t.Errorf("expected content %q got %q",string(expected),string(actual))
	}
}

func TestCreateNewRouter(t *testing.T){
	r := createNewRouter()
	newServer := httptest.NewServer(r)
	resp, err:= http.Get(newServer.URL+"/hello")
	if err!=nil{
		t.Fatal(err)
	}
	if resp.StatusCode!=http.StatusOK{
		t.Errorf("%v is status , expected : %v", resp.StatusCode, http.StatusOK)
	}
	defer resp.Body.Close()
	b, err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		t.Fatal(err)
	}
	respString:=string(b)
	
	
	if respString!="hello world"{
		t.Errorf("hello world body expected , got %s",respString)
	}
}

func TestRouterForNonExistentRoute(t *testing.T){
	r:=createNewRouter()
	newServer := httptest.NewServer(r)
	resp, err:= http.Post(newServer.URL+"/hello", "", nil)
	if err!=nil{
		t.Fatal(err)
	}
	if resp.StatusCode!=http.StatusMethodNotAllowed{
		t.Errorf("%v status code, should be 405", resp.StatusCode)
	}
	defer resp.Body.Close()
	b, err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		t.Fatal(err)
	}
	respString:=string(b)
	if respString!=""{
		t.Errorf("got non-empty response %s", respString)
	}

}

func TestFileServer(t *testing.T){
	r := createNewRouter()
	testServer := httptest.NewServer(r)
	resp, err := http.Get(testServer.URL+"/assets/")
	if err!=nil{
		t.Fatal(err)
	}
	if resp.StatusCode!=http.StatusOK{
		t.Errorf("status is %d expected %v", resp.StatusCode,http.StatusOK)
	}
	defer resp.Body.Close()
	contentType:=resp.Header.Get("Content-Type")
	expectedType := "text/html; charset=utf-8"
	if expectedType!=contentType{
		t.Errorf("expected %v got %v content type", expectedType, contentType)
	} 
	
}



