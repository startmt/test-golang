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
	expected := "61c78963dfc1541b9a2e462b92b792bb1ce5feec705d5545b739f4fbc42a12f8"
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected: \n %v\n got:\n%v", expected, got)
	}
}

func TestAddFirstChain(t *testing.T) {
	mockData := BlockChain{
		Index: 0,
		Body:  "TestData",
	}
	got := NewBlockBy(mockData)
	expected := BlockChain{
		Index:    0,
		Body:     "TestData",
		PrevHash: "",
		Hash:     "62bd487b2bed54c776108dca6062ae71da1bf191dc8ab9c7abd6ac4f015c7674",
	}

	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected: \n %v\n got:\n%v", expected, got)
	}
}

func TestValidateBlockChain(t *testing.T) {
	tableTest := map[string]struct {
		input []BlockChain
		expected  bool
	}{
		"Test Blockchain Valid": {
			input: []BlockChain{
				{
					Index: 0,
					Body:  "Testt2",
					Hash:  "d05f9e1b8458e1159efbf3c2b5fb4961276cb0f32b02a7e68ea5861fe9347d98",
				},
				{
					Index:    1,
					Body:     "Testt2",
					Hash:     "466d465d06b80b2ac2748f568b60dcb3084aa81fe806a97af0dc3e956efa21d9",
					PrevHash: "d05f9e1b8458e1159efbf3c2b5fb4961276cb0f32b02a7e68ea5861fe9347d98",
				},
				{
					Index:    2,
					Body:     "Testt2",
					Hash:     "6e0ffd6c5f7c7f3af9d5a14aaa7b712a685fb5eb8af0ebec4e446b3566f12d48",
					PrevHash: "466d465d06b80b2ac2748f568b60dcb3084aa81fe806a97af0dc3e956efa21d9",
				},
				{
					Index:    3,
					Body:     "Testt2",
					Hash:     "a3334b88eef16ddcdeb9368c02400fdda7fee585ab582b3d4ef8f5d37d54ca2a",
					PrevHash: "6e0ffd6c5f7c7f3af9d5a14aaa7b712a685fb5eb8af0ebec4e446b3566f12d48",
				},
			},
			expected: true,
		},
		"Test Blockchain InValid": {
			input: []BlockChain{
				{
					Index: 0,
					Body:  "Testt2",
					Hash:  "d05f9e1b8458e1159efbf3c2b5fb4961276cb0f32b02a7e68ea5861fe9347d98",
				},
				{
					Index:    1,
					Body:     "Testt2",
					Hash:     "466d465d06b80b2ac2748f568b60dcb3084aa81fe806a97af0dc3e956efa21d9",
					PrevHash: "d05f9e1b8458e1159efbf3c2b5fb4961276cb0f32b02a7e68ea5861fe9347d98",
				},
				{
					Index:    2,
					Body:     "Testt3",
					Hash:     "6e0ffd6c5f7c7f3af9d5a14aaa7b712a685fb5eb8af0ebec4e446b3566f12d48",
					PrevHash: "466d465d06b80b2ac2748f568b60dcb3084aa81fe806a97af0dc3e956efa21d9",
				},
				{
					Index:    3,
					Body:     "Testt2",
					Hash:     "a3334b88eef16ddcdeb9368c02400fdda7fee585ab582b3d4ef8f5d37d54ca2a",
					PrevHash: "6e0ffd6c5f7c7f3af9d5a14aaa7b712a685fb5eb8af0ebec4e446b3566f12d48",
				},
			},
			expected: false,
		},
	}
	for testName,testCaseData := range tableTest {
		got := ValidateBlockChain(testCaseData.input)
		if !reflect.DeepEqual(testCaseData.expected, got) {
			t.Fatalf("case %s is \n expected: \n %v\n got:\n%v", testName, testCaseData.expected, got)
		}
	}
}