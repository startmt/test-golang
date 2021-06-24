package blockchain

import (
	"reflect"
	"testing"
)

func TestSearchBlockInChainChainModel(t *testing.T) {
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

	type TestTable struct {
		expected  BlockChain
		condition func(blockChain BlockChain) bool
	}
	testCase := map[string]TestTable{
		"search blockchain by hash (found)": {
			expected: BlockChain{
				Index:    2,
				Body:     "Testt2",
				PrevHash: "dba37682e87dd491e4a1b89f5b402397b96f65d33c104089950a47ff4bcfc0a7",
				Hash:     "5fcb9413df11212039a1eddd613d475db3b8f50451246bf3e6cd056566ff13c4",
			},
			condition: IsSameHash("5fcb9413df11212039a1eddd613d475db3b8f50451246bf3e6cd056566ff13c4"),
		},
		"search blockchain by hash (not found)": {
			expected:  BlockChain{},
			condition: IsSameHash("notfound"),
		},
		"search blockchain by index (found)": {
			expected: BlockChain{
				Index:    2,
				Body:     "Testt2",
				PrevHash: "dba37682e87dd491e4a1b89f5b402397b96f65d33c104089950a47ff4bcfc0a7",
				Hash:     "5fcb9413df11212039a1eddd613d475db3b8f50451246bf3e6cd056566ff13c4",
			},
			condition: IsSameIndex(2),
		},
		"search blockchain by index (not found)": {
			expected:  BlockChain{},
			condition: IsSameIndex(13),
		},
	}

	for testCase, caseTestTable := range testCase {
		got, err := SearchBlockChainBy(caseTestTable.condition)(mockData)
		expected := caseTestTable.expected

		switch {
		case !reflect.DeepEqual(got, expected):
			t.Fatalf("case %s is \n expected: \n %v\n got:\n%v", testCase, expected, got)
		case reflect.DeepEqual(got, BlockChain{}) && err == nil:
			t.Fatalf("case %s is \n should not found search \n expected: \n %v\n got:\n%v", testCase, expected, got)
		}
	}
}
