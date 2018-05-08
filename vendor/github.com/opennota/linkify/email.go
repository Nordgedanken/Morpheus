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

import (
	"unicode"
	"unicode/utf8"
)

func findEmailStart(s string, start int) (_ int, _ bool) {
	end := start
	allowDot := false
	for end >= 0 {
		b := s[end]
		switch {
		case emailcs[b]:
			allowDot = true
		case b == '.':
			if !allowDot {
				return
			}
			allowDot = false
		default:
			if end == start {
				return
			}
			if s[end+1] == '.' {
				return
			}
			r, _ := utf8.DecodeLastRuneInString(s[:end+1])
			if r == utf8.RuneError {
				return
			}
			if !unicode.IsSpace(r) {
				return
			}
			return end + 1, true
		}
		end--
	}
	if end < start && s[end+1] == '.' {
		return
	}
	return end + 1, true
}

func findEmailEnd(s string, start int) (_ int, _ bool) {
	end := start
	allowDot := false
loop:
	for end < len(s) {
		b := s[end]
		switch {
		case emailcs[b]:
			allowDot = true
		case b == '.':
			if !allowDot {
				return
			}
			allowDot = false
		case b == '@':
			break loop
		default:
			return
		}
		end++
	}
	if end >= len(s)-5 {
		return
	}
	if end > start && s[end-1] == '.' {
		return
	}

	var dot int
	var ok bool
	end, dot, ok = findHostnameEnd(s, end+1)
	if !ok || dot == -1 {
		return
	}

	if dot+5 <= len(s) && s[dot+1:dot+5] == "xn--" {
		return end, true
	}

	if length := match(s[dot+1:]); dot+length+1 != end {
		return
	}

	return end, true
}
