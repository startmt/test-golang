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
	got := NewBlockBy(mockData)
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

func TestSearchBlockChainByPrevHashNotFound(t *testing.T) {
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

	_, err := SearchBlockChainByPrevHash(mockData, "notfoundthishash")
	want := "notfound"
	if !reflect.DeepEqual(want, err.Error()) || err == nil {
		t.Fatalf("expected: %v, error: %v", want, err)
	}
}

func TestValidateBlockChain(t *testing.T) {
	tableTest := map[string]struct{
		input []BlockChain
		want bool
	}{
		"Test Prev Hash InValid": {
			input: []BlockChain{
				{
					Index: 0,
					Body:  "Body1",
					Hash:  "aef9d52a83cc29df3c8a56192f11bf49a6c7a86c647579920ae26de7092bfc2c",
				},
				{
					Index:    1,
					Body:     "Body2",
					Hash:     "0b1c825c98f782f26dde3292b3116fbf734e00c3b3d23a1dc548e5e18db1c453",
					PrevHash: "aef9d52a83cc29df3c8a56192f11bf49a6c7a86c647579920ae26de7092bfc2a", // true prev hash lastchar is c
				},
				{
					Index:    2,
					Body:     "Body2",
					Hash:     "8126b51c590a3f25a2350409a582a1ffca58f987e25fb236e44c994ec04fe6f6",
					PrevHash: "0b1c825c98f782f26dde3292b3116fbf734e00c3b3d23a1dc548e5e18db1c453",
				},
				{
					Index:    3,
					Body:     "Body2",
					Hash:     "915b6e1399e83c0cc7b962cfb6db6db5a274efe595f56a36e9829f09bf686a18",
					PrevHash: "8126b51c590a3f25a2350409a582a1ffca58f987e25fb236e44c994ec04fe6f6",
				},
			},
			want: false,
		},
	}

	for name, testcase := range tableTest {
		got := ValidateBlockChain(testcase.input)
		if !reflect.DeepEqual(got, testcase.want) {
			t.Fatalf("case: %s expected: %v, got: %v", name, testcase.want, got)
		}
	}
}

func TestValidateBlockChainPrevHashInValid(t *testing.T) {
	mockFalse := []BlockChain{
		{
			Index: 0,
			Body:  "Body1",
			Hash:  "aef9d52a83cc29df3c8a56192f11bf49a6c7a86c647579920ae26de7092bfc2c",
		},
		{
			Index:    1,
			Body:     "Body2",
			Hash:     "0b1c825c98f782f26dde3292b3116fbf734e00c3b3d23a1dc548e5e18db1c453",
			PrevHash: "aef9d52a83cc29df3c8a56192f11bf49a6c7a86c647579920ae26de7092bfc2a", // true prev hash lastchar is c
		},
		{
			Index:    2,
			Body:     "Body2",
			Hash:     "8126b51c590a3f25a2350409a582a1ffca58f987e25fb236e44c994ec04fe6f6",
			PrevHash: "0b1c825c98f782f26dde3292b3116fbf734e00c3b3d23a1dc548e5e18db1c453",
		},
		{
			Index:    3,
			Body:     "Body2",
			Hash:     "915b6e1399e83c0cc7b962cfb6db6db5a274efe595f56a36e9829f09bf686a18",
			PrevHash: "8126b51c590a3f25a2350409a582a1ffca58f987e25fb236e44c994ec04fe6f6",
		},
	}
	got := ValidateBlockChain(mockFalse)
	want := false
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestValidateBlockChainHashInValid(t *testing.T) {
	mockFalse := []BlockChain{
		{
			Index: 0,
			Body:  "Body1",
			Hash:  "aef9d52a83cc29df3c8a56192f11bf49a6c7a86c647579920ae26de7092bfc2c",
		},
		{
			Index:    1,
			Body:     "Body2",
			Hash:     "0b1c825c98f782f26dde3292b3116fbf734e00c3b3d23a1dc548e5e18db1c453",
			PrevHash: "aef9d52a83cc29df3c8a56192f11bf49a6c7a86c647579920ae26de7092bfc2c",
		},
		{
			Index:    2,
			Body:     "Body2",
			Hash:     "8126b51c590a3f25a2350409a582a1ffca58f987e25fb236e44c994ec04fe6f1", // true hash lastchar is 6
			PrevHash: "0b1c825c98f782f26dde3292b3116fbf734e00c3b3d23a1dc548e5e18db1c453",
		},
		{
			Index:    3,
			Body:     "Body2",
			Hash:     "915b6e1399e83c0cc7b962cfb6db6db5a274efe595f56a36e9829f09bf686a18",
			PrevHash: "8126b51c590a3f25a2350409a582a1ffca58f987e25fb236e44c994ec04fe6f6",
		},
	}
	got := ValidateBlockChain(mockFalse)
	want := false
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestValidateBlockChainBodyInValid(t *testing.T) {
	mockFalse := []BlockChain{
		{
			Index: 0,
			Body:  "Body100",
			Hash:  "aef9d52a83cc29df3c8a56192f11bf49a6c7a86c647579920ae26de7092bfc2c",
		},
		{
			Index:    1,
			Body:     "Body2",
			Hash:     "0b1c825c98f782f26dde3292b3116fbf734e00c3b3d23a1dc548e5e18db1c453",
			PrevHash: "aef9d52a83cc29df3c8a56192f11bf49a6c7a86c647579920ae26de7092bfc2c",
		},
		{
			Index:    2,
			Body:     "Body2",
			Hash:     "8126b51c590a3f25a2350409a582a1ffca58f987e25fb236e44c994ec04fe6f1", // true hash lastchar is 6
			PrevHash: "0b1c825c98f782f26dde3292b3116fbf734e00c3b3d23a1dc548e5e18db1c453",
		},
		{
			Index:    3,
			Body:     "Body2",
			Hash:     "915b6e1399e83c0cc7b962cfb6db6db5a274efe595f56a36e9829f09bf686a18",
			PrevHash: "8126b51c590a3f25a2350409a582a1ffca58f987e25fb236e44c994ec04fe6f6",
		},
	}
	got := ValidateBlockChain(mockFalse)
	want := false
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestValidateBlockChainLess2(t *testing.T) {
	mockData := []BlockChain{
		{
			Index: 0,
			Body:  "Body1",
			Hash:  "aef9d52a83cc29df3c8a56192f11bf49a6c7a86c647579920ae26de7092bfc2c",
		},
	}
	got := ValidateBlockChain(mockData)
	want := true
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestValidateBlockChainLastInValid(t *testing.T) {
	mockFalse := []BlockChain{
		{
			Index: 0,
			Body:  "Body1",
			Hash:  "aef9d52a83cc29df3c8a56192f11bf49a6c7a86c647579920ae26de7092bfc2c",
		},
		{
			Index:    1,
			Body:     "Body2",
			Hash:     "0b1c825c98f782f26dde3292b3116fbf734e00c3b3d23a1dc548e5e18db1c453",
			PrevHash: "aef9d52a83cc29df3c8a56192f11bf49a6c7a86c647579920ae26de7092bfc2c",
		},
	}
	got := ValidateBlockChain(mockFalse)
	want := false
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestValidateBlockChainValid(t *testing.T) {
	mockData := []BlockChain{
		{
			Index: 0,
			Body:  "Body1",
			Hash:  "aef9d52a83cc29df3c8a56192f11bf49a6c7a86c647579920ae26de7092bfc2c",
		},
		{
			Index:    1,
			Body:     "Body2",
			Hash:     "0b1c825c98f782f26dde3292b3116fbf734e00c3b3d23a1dc548e5e18db1c453",
			PrevHash: "aef9d52a83cc29df3c8a56192f11bf49a6c7a86c647579920ae26de7092bfc2c",
		},
		{
			Index:    2,
			Body:     "Body2",
			Hash:     "8126b51c590a3f25a2350409a582a1ffca58f987e25fb236e44c994ec04fe6f6",
			PrevHash: "0b1c825c98f782f26dde3292b3116fbf734e00c3b3d23a1dc548e5e18db1c453",
		},
		{
			Index:    3,
			Body:     "Body2",
			Hash:     "915b6e1399e83c0cc7b962cfb6db6db5a274efe595f56a36e9829f09bf686a18",
			PrevHash: "8126b51c590a3f25a2350409a582a1ffca58f987e25fb236e44c994ec04fe6f6",
		},
	}
	got := ValidateBlockChain(mockData)
	want := true
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
