package stringreplace

import (
	"regexp"
)

func replace_map(original string, mapper map[string]string) (betterlog string) {
  betterlog=original
  for match_this, replacement := range mapper {
     if betterlog==original {		
       r, _ := regexp.Compile(match_this)
       betterlog = r.ReplaceAllString(original,replacement)
     }
  }
  return betterlog
}

