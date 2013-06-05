/*
Package shodan is a thin wrapper around shodanhq.com API.

You need a APIKEY to use this package.
Get one here: http://www.shodanhq.com/api_doc

A simple example using the package:

		api := shodan.NewWebAPI(APIKEY)
		res, err := api.Search("apache", 1)
		if err != nil {
			panic(err)
		}

		for x := range res.Matches {
			fmt.Printf("%s : %s\n", res.Matches[x].IP, res.Matches[x].RegionName)
		}
*/
package goshodan
