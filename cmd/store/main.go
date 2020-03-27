package main

import (
	"fmt"
	"gitlab.com/etherlabs/dist-store/pkg/dist_store"
)

func main(){
	numberOfStores := 10
	distStore := dist_store.NewDistStorer(int32(numberOfStores))
	err := distStore.Write("vamshi", "krishna")
	if err != nil {
		fmt.Println("error while writing", err)
	}
	err = distStore.Write("krishna", "vamshi")
	if err != nil {
		fmt.Println("error while writing", err)
	}

	err = distStore.Write("testing", "124")
	if err != nil {
		fmt.Println("error while writing", err)
	}

	v, err := distStore.Read("vamshi")
	if err != nil {
		fmt.Println("error while reading", err)
	} else{
		fmt.Println("value ", v)
	}

	v, err = distStore.Read("testing")
	if err != nil {
		fmt.Println("error while reading", err)
	} else{
		fmt.Println("value ", v)
	}

	distStore.Dump()

	cydistStore := dist_store.NewCyclicDistStorer(int32(numberOfStores))
	err = cydistStore.Write("vamshi", "krishna")
	if err != nil {
		fmt.Println("error while writing", err)
	}
	err = cydistStore.Write("krishna", "vamshi")
	if err != nil {
		fmt.Println("error while writing", err)
	}

	err = cydistStore.Write("testing", "124")
	if err != nil {
		fmt.Println("error while writing", err)
	}

	v, err = cydistStore.Read("vamshi")
	if err != nil {
		fmt.Println("error while reading", err)
	} else{
		fmt.Println("value ", v)
	}

	v, err = cydistStore.Read("testing")
	if err != nil {
		fmt.Println("error while reading", err)
	} else{
		fmt.Println("value ", v)
	}

	cydistStore.Dump()

}