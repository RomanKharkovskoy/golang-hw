package unicode_name

import "unicode"

// AnalyzeName считает байты, Unicode-кодовые точки, латинские и кириллические
// буквы и возвращает байтовые позиции начала каждой Unicode-кодовой точки.
// Предполагается, что входная строка содержит корректный UTF-8.
func AnalyzeName(name string) (byteCount, runeCount, latinCount, cyrillicCount int, positions []int) {
	byteCount = len(name)
	positions = make([]int, 0, byteCount)
	for position, letter := range name {
		positions = append(positions, position)
		runeCount++
		if unicode.IsLetter(letter) {
			if unicode.Is(unicode.Latin, letter) {
				latinCount++
			} else if unicode.Is(unicode.Cyrillic, letter) {
				cyrillicCount++
			}
		}
	}
	return
}
