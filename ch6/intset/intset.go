package intset

import (
	"bytes"
	"fmt"
)

// IntSet представляет собой множество небольших неотрицательных
// целых чисел. Нулевое значение представляет пустое множество.
type IntSet struct {
	words []uint
}

const bitDepth = 32 << (uint(0) >> 63)

// Has указывает, содержит ли множество неотрицательное значение х.
func (s *IntSet) Has(x int) bool {
	word, bit := x/bitDepth, uint(x%bitDepth)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add добавляет неотрицательное значение x в множество.
func (s *IntSet) Add(x int) {
	word, bit := x/bitDepth, uint(x%bitDepth)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// Remove удаляет x из множества
func (s *IntSet) Remove(x int) {
	word, bit := x/bitDepth, uint(x%bitDepth)
	if word < len(s.words) {
		s.words[word] ^= (1 << bit)
	}
}

// Clear удаляет все элементы множества
func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

// UnionWith делает множество s равным объединению множеств s и t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith делает множество s равным пересечению множеств s и t.
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
}

// DifferenceWith делает множество s равным разнице множеств s и t.
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i := range s.words {
		if i < len(t.words) {
			s.words[i] &^= t.words[i]
		}
	}
}

// SymmetricDifferenceWith делает множество s равным симметричной разнице множеств s и t.
func (s *IntSet) SymmetricDifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Len возвращает количество элементов
func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bitDepth; j++ {
			if word&(1<<uint(j)) != 0 {
				count++
			}
		}
	}
	return count
}

// Elems возвращает элементы множества как срез
func (s *IntSet) Elems() []uint {
	var elements []uint
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bitDepth; j++ {
			if word&(1<<uint(j)) != 0 {
				elements = append(elements, uint(bitDepth*i+j))
			}
		}
	}
	return elements
}

// Copy возвращает копию множества
func (s *IntSet) Copy() *IntSet {
	var t IntSet
	t.words = make([]uint, len(s.words))
	copy(t.words, s.words)
	return &t
}

// String возвращает множество как строку вида "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bitDepth; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", bitDepth*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
