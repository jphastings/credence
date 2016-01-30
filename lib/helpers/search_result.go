package helpers

import (
  "github.com/jphastings/credence/lib/definitions/credence"
)

func DeduplicateKeys(searchResult *credence.SearchResult) {
  dedupedKeyMap := make(map[string]int)

  for i, sourceBreakdown := range searchResult.Results {
    otherIndex, exists := dedupedKeyMap[sourceBreakdown.SourceUri]
    if exists {
      for _, assertion := range sourceBreakdown.Assertions {
        searchResult.Results[otherIndex].Assertions = append(searchResult.Results[otherIndex].Assertions, assertion)
      }
    } else {
      dedupedKeyMap[sourceBreakdown.SourceUri] = i
    }
  }

  newResults := []*credence.SearchResult_SourceBreakdown{}

  for _, index := range dedupedKeyMap {
    newResults = append(newResults, searchResult.Results[index])
  }

  searchResult.Results = newResults
}