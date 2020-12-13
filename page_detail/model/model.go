package model

type (
	//PageDetailsData contains all data of a page
	PageDetailsData struct {
		Version       string      `json:"version"`
		Title         string      `json:"title"`
		HeadingLevels HLevelCount `json:"headings"`
		Links         Links       `json:"links"`
		HasLoginForm  bool        `json:"has_login_form"`
	}

	//HLevelCount contains all counts in heading levels
	HLevelCount struct {
		H1 int64 `json:"h1"`
		H2 int64 `json:"h2"`
		H3 int64 `json:"h3"`
		H4 int64 `json:"h4"`
		H5 int64 `json:"h5"`
		H6 int64 `json:"h6"`
	}

	//Links contains all links in a page
	Links struct {
		Internal     LinkDetail          `json:"internal"`
		External     LinkDetail          `json:"external"`
		Inaccessible LinkDetailWithError `json:"inaccessible"`
	}

	//LinkDetail contains count and links
	LinkDetail struct {
		Count int      `json:"count"`
		Data  []string `json:"data"`
	}

	//LinkDetailWithError contains count and links with error
	LinkDetailWithError struct {
		Count  int      `json:"count"`
		Data   []string `json:"data"`
		Errors []string `json:"errors"`
	}

	//CacheValue will be hold cache values
	CacheValue struct {
		Internal  bool
		Reachable bool
	}
)
