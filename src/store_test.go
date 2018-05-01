package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"github.com/stretchr/testify/suite"
)

type StoreSuite struct{
	suite.Suite
	store *dbStore
	db *sql.DB	
}

func(s *StoreSuite) SetupSuite(){
	connString := "root:root@/bird_encyclopedia"
	db, err := sql.Open("mysql", connString)
	if err!=nil{
		s.T().Fatal(err)
	}
	s.db= db
	s.store=&dbStore{db: db}
}

func (s *StoreSuite) SetupTest(){
	_,err:= s.db.Query("DELETE FROM birds")
	if err!=nil{
		s.T().Fatal(err)
	}
}

func (s *StoreSuite) TearDownSuite(){
	s.db.Close()
}

func TestStoreSuite(t *testing.T){
	s := new(StoreSuite)
	suite.Run(t,s)
} 

func (s *StoreSuite) TestCreateBird(){
	s.store.CreateBird(&Bird{Description: "sample bird", Species: "Sample species", })
	res, err := s.db.Query(`SELECT COUNT(*) FROM birds where species='Sample species' and description='sample bird'`) 
	if err!=nil{
		s.T().Fatal(err)
	}
	
	var count int
	for res.Next(){
		err := res.Scan(&count)
		if err!=nil{
			s.T().Error(err)
		}
	}
	
	if count!=1{
		s.T().Errorf("incorrect count : %d",count)
	}
	
}

func (s *StoreSuite) TestGetBirds(){
	_,err := s.db.Query(`INSERT INTO birds(species, description) values('bird','descr')`)
	if err!=nil{
		s.T().Fatal(err)
	}
	
	birds, err := s.store.GetBirds()
	if err!=nil{
		s.T().Fatal(err)
	}
	
	nBirds := len(birds)
	if nBirds!=1{
		s.T().Errorf("incorrect count : %d",nBirds)
		
	}
	expectedBird := Bird{"bird","descr"}
	if *birds[0] != expectedBird{
		s.T().Errorf("expected %v got %v",expectedBird, *birds[0])
	}
	
}


