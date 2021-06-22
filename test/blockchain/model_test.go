package blockchain

import (
	"reflect"
	"testing"
)

//BlockChain Model
func TestAddBlockModel(t *testing.T) {
	block := BlockChain{}
	block.MakeBlockWithNewHash(BlockChain{
		Index: 0,
		Body:  "Testt2",
	})
	got := block

	want := BlockChain{
		Index:    0,
		Body:     "Testt2",
		PrevHash: "",
		Hash:     "d05f9e1b8458e1159efbf3c2b5fb4961276cb0f32b02a7e68ea5861fe9347d98",
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestAddSecondaryBlockModel(t *testing.T) {
	block := BlockChain{
		Index:    0,
		Body:     "Testt",
		PrevHash: "",
	}
	block.MakeBlockWithNewHash(BlockChain{
		Index:    1,
		Body:     "Testt2",
		PrevHash: "d05f9e1b8458e1159efbf3c2b5fb4961276cb0f32b02a7e68ea5861fe9347d98",
	})
	got := block

	want := BlockChain{
		Index:    1,
		Body:     "Testt2",
		PrevHash: "d05f9e1b8458e1159efbf3c2b5fb4961276cb0f32b02a7e68ea5861fe9347d98",
		Hash:     "466d465d06b80b2ac2748f568b60dcb3084aa81fe806a97af0dc3e956efa21d9",
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestAddhBlockToChainModel(t *testing.T) {
	mockChain := BlockChain{
		Index:    0,
		Body:     "Testt2",
		PrevHash: "",
		Hash:     "9b638070939bf21b917ea5d4eecaf95d0b2bb818a23e613d527aeb52d1f229f3",
	}
	block := ChainArray{}
	block.Add(mockChain)
	got := block
	want := ChainArray{mockChain}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestSearchBlockInChainChainModel(t *testing.T) {
	mockChain := ChainArray{
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

	got := mockChain.Search("5fcb9413df11212039a1eddd613d475db3b8f50451246bf3e6cd056566ff13c4")
	want := BlockChain{
		Index:    2,
		Body:     "Testt2",
		PrevHash: "dba37682e87dd491e4a1b89f5b402397b96f65d33c104089950a47ff4bcfc0a7",
		Hash:     "5fcb9413df11212039a1eddd613d475db3b8f50451246bf3e6cd056566ff13c4",
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

// CreateBlockChainReq

func TestCreateBlockChainReq(t *testing.T) {
	mockBody := "Test123123"
	req := CreateBlockChainReq{}
	req.Create(mockBody)

	got := req
	want := CreateBlockChainReq{Body: mockBody}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}

}
