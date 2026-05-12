package levenshtein

// CalculateDistance find the levenshtein distance between the two string
// passed as arguments. The return value is their levenshtein distance
func CalculateDistance(s1, s2 string) int {
	if len(s1) == 0 {
		return len(s2)
	}

	if len(s2) == 0 {
		return len(s1)
	}

	if s1 == s2 {
		return 0
	}

	// swap to ensure s1 is always the shorter string
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}

	var row = make([]uint16, len(s1)+1)// uint16 use less bytes. Distances are small whole numbers, which will never exceed
	// row[0] = 0 is already initilized that way, so i can start at 1
	for i := 1; i < len(row); i++ {
		row[i] = uint16(i)// first row will always be [0, 1, ... len(x)-1]
	}

	for i := 0; i < len(s2); i++ {// i is the column
		prev := uint16(i+1)// prev is the diagonal
		for j := 0; j < len(s1); j++ {
			current := row[j]// in case of match, current save the diagonal value
			if s1[j] != s2[i] {// mismatch
				// the least costful operation among substitution, insertion and deletion is choosed
				current = min(	row[j], 	// substitution, diagonal cell
								prev,		// insertion, left cell
								row[j-1],	// deletion, top cell
				) +1
			}
			row[j] = prev
			prev = current
		}
		row[len(s1)] = prev
	}

	return int(row[len(s1)])
}