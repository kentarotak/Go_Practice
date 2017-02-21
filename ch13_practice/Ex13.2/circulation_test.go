package circuChk

import "testing"

func Test(t *testing.T) {

	type Cycle struct {
		value int
		Tail  *Cycle
	}
	var c Cycle
	c = Cycle{42, &c}

	// 4変数リンク.
	var c1, c2, c3, c4 Cycle

	c1 = Cycle{100, &c2}
	c2 = Cycle{200, &c3}
	c3 = Cycle{300, &c4}
	c4 = Cycle{400, &c1}

	// 4変数の単方向リスト.
	var d1, d2, d3, d4 Cycle

	d1 = Cycle{500, &d2}
	d2 = Cycle{600, &d3}
	d3 = Cycle{700, &d4}

	type normalStruct struct {
		value  int
		value2 int
	}

	var d normalStruct
	d = normalStruct{100, 1000}

	type Test struct {
		value interface{}
		want  bool
	}

	var tests []Test

	tests = append(tests, Test{c, true})   // 循環構造体
	tests = append(tests, Test{d, false})  // 普通の構造体
	tests = append(tests, Test{c1, true})  // 4変数の循環リスト
	tests = append(tests, Test{d1, false}) // 4変数の単方向リスト
	tests = append(tests, Test{45, false}) // 普通のint型

	for _, test := range tests {
		if CirculationCheck(test.value) != test.want {
			t.Fatalf("err %v", test.value)
		}
	}

}
