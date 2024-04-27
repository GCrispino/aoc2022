package common

func isNotRepeating(s string) bool {
  setMap := make(map[rune]struct{})

  for _, c := range s {
    if _, ok := setMap[c]; ok {
      return false
    }

    setMap[c] = struct{}{}
  }

  return true
}

func FindMarker(s string, markerSize int) int {
  for i := markerSize; i < len(s); i++ {
    if isNotRepeating(s[i - markerSize:i + 1]) {
      return i + 1
    }
  }

  return -1
}
