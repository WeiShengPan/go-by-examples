package example7

import (
	"fmt"
	"testing"
)

// 我们将通过两个函数：`zeroval` 和 `zeroptr` 来比较指针和
// 值类型的不同。`zeroval` 有一个 `int` 型参数，所以使用值
// 传递。`zeroval` 将从调用它的那个函数中得到一个 `ival`
// 形参的拷贝。
func TestPointer(t *testing.T) {

	i := 1
	fmt.Println("initial:", i)

	zeroVal(i)
	fmt.Println("zeroVal:", i)

	// 通过&来取得i的内存地址
	zeroPtr(&i)
	fmt.Println("zeroPtr:", i)

	fmt.Println("pointer:", &i)
}

func zeroVal(ival int) {
	ival = 0
}

// `zeroptr` 有一和上面不同的 `*int` 参数，意味着它用了一
// 个 `int`指针。函数体内的 `*iptr` 接着_解引用_ 这个指针，
// 从它内存地址得到这个地址对应的当前值。对一个解引用的指
// 针赋值将会改变这个指针引用的真实地址的值。
func zeroPtr(iptr *int) {
	*iptr = 0
}
