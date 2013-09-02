goshodan
========

[![Build Status](https://drone.io/github.com/stone/goshodan/status.png)](https://drone.io/github.com/stone/goshodan/latest)

goshodan is a work in progress thin wrapper around shodanhq.com API.

![img](http://i.imgur.com/IMUKRWQ.jpg)

You need a APIKEY to use this package.
Get one here: http://www.shodanhq.com/api_doc

A simple example using the package:

```Go
    api := goshodan.NewWebAPI(APIKEY)
    res, err := api.Search("apache", 1)
    if err != nil {
      panic(err)
    }

    for x := range res.Matches {
      fmt.Printf("%s : %s\n", res.Matches[x].IP, res.Matches[x].RegionName)
    }
```


Want to hack on goshodan? Awesome! Send me pull requests!

NOTE: The package are faaaar from perfect.
