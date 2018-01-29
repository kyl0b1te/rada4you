package examples

import (
	"log"

	"github.com/zhikiri/rada4you/rada4you"
)

func SimpleRequestExample() {
	cli := rada4you.New("secret_key")
	res, err := cli.GetAllPeoples()
	if err != nil && err.IsOccur() {
		log.Fatal("Something with request")
		return
	}
	log.Print(res.Peoples)
}

func ComplexRequestExample() {
	cli := rada4you.New("secret_key")
	req := rada4you.GetAllDivisionsRequest{House: "rada"}
	res, err := cli.GetAllDivisions(req)
	if err != nil && err.IsOccur() {
		log.Fatal("Something with request")
		return
	}
	log.Print(res.Divisions)
}
