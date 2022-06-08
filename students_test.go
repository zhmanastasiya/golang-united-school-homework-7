package coverage

import (
	"os"
	"testing"
	"time"
	
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

func TestLen(t *testing.T) {
	people := []Person{
    {"Bob", "Dilan",time.Date(1991, 6, 3, 0, 0, 0, 0, time.UTC)},
	{"John","Gray",time.Date(1991, 6, 3, 0, 0, 0, 0, time.UTC)},
	{"John", "Bush",time.Date(1991, 6, 3, 0, 0, 0, 0, time.UTC)},
	{"Michael", "Hostings",time.Date(2015, 3, 1, 0, 0, 0, 0, time.UTC)},
}
        var expects int
	for i := 0; i < len(people); i++ {
			expects +=1
	}
	
	if len(people) != expects{
		t.Errorf("The expected length %d, but actual %d", len(people),expects)
	}
	
}
func TestLess(t *testing.T) {
	pers1 := Person{"Bob", "Dilan",time.Date(1992, 8, 2, 0, 0, 0, 0, time.UTC)}
	pers2 := Person{"John","Gray",time.Date(1991, 6, 3, 0, 0, 0, 0, time.UTC)}
	pers3 := Person{"John", "Bush",time.Date(1991, 6, 3, 0, 0, 0, 0, time.UTC)}
	pers4 := Person{"Michael", "Hostings",time.Date(2015, 3, 1, 0, 0, 0, 0, time.UTC)}
	
		
	tests := []struct {
		input People
		expects bool}{
		{input: People{pers1,pers2},expects: false},
		{input: People{pers2,pers1},expects: true},
		{input: People{pers1,pers3},expects: false},
		{input: People{pers3,pers1},expects: true},
		{input: People{pers1,pers4},expects: true},
		{input: People{pers4,pers1},expects: false},
	}
	for _, v := range tests{
		if v.input.Less(1,0) != v.expects {
			t.Errorf("Expected %v, got %v \n", v.expects, v.input.Less(1,0))
		} 
	}

}

func TestSwap(t *testing.T){

	pers1 := Person{"Bob", "Dilan",time.Date(1992, 8, 2, 0, 0, 0, 0, time.UTC)}
	pers2 := Person{"John","Gray",time.Date(1991, 6, 3, 0, 0, 0, 0, time.UTC)}
	
	tests := People{pers1,pers2}

	tests.Swap(0,1)

	if tests[0] != tests[1]{
		t.Errorf("Person did not swap")
	}
	
		}

func TestNewMatrix(t *testing.T){

type testNew struct{
	input string
	cols,rows int
	data []int
	expectError error}
	
	test := []testNew{
		{input:"1 2 3\n 4 5 6",
		cols:3,
		rows:3,
		data:[]int{1,2,3,4,5,6},
		expectError: nil,},

		{input:"7 8\n 8 7",
		cols:2,
		rows:2,
		data:[]int{7,8,8,7},
		expectError: nil,},

		{input:"1 3 \n 5 7",
		cols:2,
		rows:2,
		data:[]int{1,3,5,7},
		expectError: nil,},
		}
	
		
		for _,p := range test{
			matr, err:= New(p.input)
				if len(matr.data) != len(p.data) && err != p.expectError{
					t.Errorf("Error in matrix %v", matr)
				}	
			}
		
	
		}
		
		
		