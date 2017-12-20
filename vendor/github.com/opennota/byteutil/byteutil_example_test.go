// This program is free software: you can redistribute it and/or modify it
// under the terms of the GNU General Public License as published by the Free
// Software Foundation, either version 3 of the License, or (at your option)
// any later version.
//
// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General
// Public License for more details.
//
// You should have received a copy of the GNU General Public License along
// with this program.  If not, see <http://www.gnu.org/licenses/>.

package byteutil

import (
	"fmt"
)

func ExampleIsDigit() {
	fmt.Println(IsDigit('0'))
	fmt.Println(IsDigit('9'))
	fmt.Println(IsDigit('a'))

	// Output:
	// true
	// true
	// false
}

func ExampleIsHexDigit() {
	fmt.Println(IsHexDigit('0'))
	fmt.Println(IsHexDigit('9'))
	fmt.Println(IsHexDigit('a'))
	fmt.Println(IsHexDigit('F'))
	fmt.Println(IsHexDigit('G'))

	// Output:
	// true
	// true
	// true
	// true
	// false
}

func ExampleIsLetter() {
	fmt.Println(IsLetter('a'))
	fmt.Println(IsLetter('Z'))
	fmt.Println(IsLetter('0'))

	// Output:
	// true
	// true
	// false
}

func ExampleIsLowercaseLetter() {
	fmt.Println(IsLowercaseLetter('a'))
	fmt.Println(IsLowercaseLetter('z'))
	fmt.Println(IsLowercaseLetter('A'))

	// Output:
	// true
	// true
	// false
}

func ExampleIsUppercaseLetter() {
	fmt.Println(IsUppercaseLetter('A'))
	fmt.Println(IsUppercaseLetter('Z'))
	fmt.Println(IsUppercaseLetter('a'))

	// Output:
	// true
	// true
	// false
}

func ExampleIsAlphaNum() {
	fmt.Println(IsAlphaNum('0'))
	fmt.Println(IsAlphaNum('9'))
	fmt.Println(IsAlphaNum('a'))
	fmt.Println(IsAlphaNum('Z'))
	fmt.Println(IsAlphaNum('.'))

	// Output:
	// true
	// true
	// true
	// true
	// false
}

func ExampleToLower() {
	fmt.Println(ToLower(".012345abcdefABCDEFабвгд"))

	// Output:
	// .012345abcdefabcdefабвгд
}

func ExampleToUpper() {
	fmt.Println(ToUpper(".012345abcdefABCDEFабвгд"))

	// Output:
	// .012345ABCDEFABCDEFабвгд
}

func ExampleByteToLower() {
	fmt.Printf("%c\n", ByteToLower('a'))
	fmt.Printf("%c\n", ByteToLower('A'))
	fmt.Printf("%c\n", ByteToLower('0'))
	fmt.Printf("%c\n", ByteToLower('.'))

	// Output:
	// a
	// a
	// 0
	// .
}

func ExampleByteToUpper() {
	fmt.Printf("%c\n", ByteToUpper('a'))
	fmt.Printf("%c\n", ByteToUpper('A'))
	fmt.Printf("%c\n", ByteToUpper('0'))
	fmt.Printf("%c\n", ByteToUpper('.'))

	// Output:
	// A
	// A
	// 0
	// .
}

func ExampleIndexAny() {
	fmt.Println(IndexAny("abcdefghijklmnopqrstuvwxyz", "zyx"))

	// Output:
	// 23
}

func ExampleIndexAnyTable() {
	var t [256]bool
	t['z'], t['y'], t['x'] = true, true, true
	fmt.Println(IndexAnyTable("abcdefghijklmnopqrstuvwxyz", &t))

	// Output:
	// 23
}

func ExampleUnhex() {
	fmt.Println(Unhex('0'))
	fmt.Println(Unhex('9'))
	fmt.Println(Unhex('a'))
	fmt.Println(Unhex('A'))
	fmt.Println(Unhex('f'))
	fmt.Println(Unhex('F'))

	// Output:
	// 0
	// 9
	// 10
	// 10
	// 15
	// 15
}
