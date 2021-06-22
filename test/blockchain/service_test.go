package blockchain

import (
	"reflect"
	"testing"
)

func TestCreateNewHash(t *testing.T) {
	got := CreateNewHash(BlockChain{
		Index: 0,
		Body:  "Test Data",
	})
	want := "61c78963dfc1541b9a2e462b92b792bb1ce5feec705d5545b739f4fbc42a12f8"
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestAddFirstChain(t *testing.T) {
	mockData := BlockChain{
		Index: 0,
		Body:  "TestData",
	}
	got := AddOneChain(mockData)
	want := BlockChain{
		Index:    0,
		Body:     "TestData",
		PrevHash: "",
		Hash:     "62bd487b2bed54c776108dca6062ae71da1bf191dc8ab9c7abd6ac4f015c7674",
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestSearchChain(t *testing.T) {
	mockChain := ChainArray{
		{
			Index:    0,
			Body:     "Testt0",
			PrevHash: "",
			Hash:     "9b638070939bf21b917ea5d4eecaf95d0b2bb818a23e613d527aeb52d1f229f3",
		},
		{
			Index:    1,
			Body:     "Testt1",
			PrevHash: "9b638070939bf21b917ea5d4eecaf95d0b2bb818a23e613d527aeb52d1f229f3",
			Hash:     "dba37682e87dd491e4a1b89f5b402397b96f65d33c104089950a47ff4bcfc0a7",
		},
		{
			Index:    2,
			Body:     "Testt2",
			PrevHash: "dba37682e87dd491e4a1b89f5b402397b96f65d33c104089950a47ff4bcfc0a7",
			Hash:     "5fcb9413df11212039a1eddd613d475db3b8f50451246bf3e6cd056566ff13c4",
		},
	}

	got, err := mockChain.Search("5fcb9413df11212039a1eddd613d475db3b8f50451246bf3e6cd056566ff13c4")

	want := BlockChain{
		Index:    2,
		Body:     "Testt2",
		PrevHash: "dba37682e87dd491e4a1b89f5b402397b96f65d33c104089950a47ff4bcfc0a7",
		Hash:     "5fcb9413df11212039a1eddd613d475db3b8f50451246bf3e6cd056566ff13c4",
	}

	if !reflect.DeepEqual(want, got) || err != nil {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestSearchBlockChainNotFound(t *testing.T) {
	//create Mock data
	mockData := []BlockChain{
		{
			Index:    0,
			Body:     "Testt2",
			PrevHash: "",
			Hash:     "9b638070939bf21b917ea5d4eecaf95d0b2bb818a23e613d527aeb52d1f229f3",
		},
		{
			Index:    1,
			Body:     "Testt2",
			PrevHash: "9b638070939bf21b917ea5d4eecaf95d0b2bb818a23e613d527aeb52d1f229f3",
			Hash:     "dba37682e87dd491e4a1b89f5b402397b96f65d33c104089950a47ff4bcfc0a7",
		},
		{
			Index:    2,
			Body:     "Testt2",
			PrevHash: "dba37682e87dd491e4a1b89f5b402397b96f65d33c104089950a47ff4bcfc0a7",
			Hash:     "5fcb9413df11212039a1eddd613d475db3b8f50451246bf3e6cd056566ff13c4",
		},
	}

	_, err := SearchBlockChain(mockData, "notfoundthishash")
	want := "notfound"
	if !reflect.DeepEqual(want, err.Error()) || err == nil {
		t.Fatalf("expected: %v, error: %v", want, err)
	}
}
