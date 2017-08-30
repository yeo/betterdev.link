package baja

type reporterFunc func(link string, previous, dupe int)

func CheckDupe(reporter reporterFunc) {
	// Loopover issue, Fetch links and find dupe
	page := buildPage()
	linkDB := make(map[string]int)
	for _, issue := range page.Issues {
		for thisIssue, link := range issue.Links {
			if previousIssue, ok := linkDB[link.URI]; ok {
				reporter(link.URI, previousIssue, thisIssue)
			} else {
				linkDB[link.URI] = thisIssue
			}
		}
	}
}
