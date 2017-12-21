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

package linkify

import "unicode"

var (
	trigger    [256]bool
	unreserved [256]bool
	subdelims  [256]bool
	emailcs    [256]bool
	basicPunct [256]bool
)

func init() {
	for _, b := range "-._~" {
		unreserved[b] = true
	}
	for _, b := range "!$&'()*+,;=" {
		subdelims[b] = true
	}
	for _, b := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!#$%&'*+/=?^_`{|}~-" {
		emailcs[b] = true
	}
	for _, b := range ".,?!;:" {
		basicPunct[b] = true
	}
}

func isAllowedInEmail(r rune) bool {
	return r < 0x7f && emailcs[r]
}

func isLetterOrDigit(r rune) bool {
	return unicode.In(r, unicode.Letter, unicode.Digit)
}

func isPunctOrSpaceOrControl(r rune) bool {
	return r == '<' || r == '>' || unicode.In(r, unicode.Punct, unicode.Space, unicode.Cc)
}

func isUnreserved(r rune) bool {
	return (r < 0x7f && unreserved[r]) || isLetterOrDigit(r)
}

func isSubDelimiter(r rune) bool {
	return r < 0x7f && subdelims[r]
}
