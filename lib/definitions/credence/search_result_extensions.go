package credence

func (searchResult *SearchResult) DeduplicateKeys() {
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

  newResults := []*SearchResult_KeyBreakdown{}

  for _, index := range dedupedKeyMap {
    newResults = append(newResults, searchResult.Results[index])
  }

  searchResult.Results = newResults
}