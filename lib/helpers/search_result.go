package helpers

import (
  "github.com/jphastings/credence/lib/definitions/credence"
)

func DeduplicateKeys(searchResult *credence.SearchResult) {
  dedupedKeyMap := make(map[string]int)

  for i, keyBreakdown := range searchResult.Results {
    otherIndex, exists := dedupedKeyMap[keyBreakdown.Key]
    if exists {
      for _, assertion := range keyBreakdown.Assertions {
        searchResult.Results[otherIndex].Assertions = append(searchResult.Results[otherIndex].Assertions, assertion)
      }
    } else {
      dedupedKeyMap[keyBreakdown.Key] = i
    }
  }

  newResults := []*credence.SearchResult_KeyBreakdown{}

  for _, index := range dedupedKeyMap {
    newResults = append(newResults, searchResult.Results[index])
  }

  searchResult.Results = newResults
}