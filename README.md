Library webhooks
================
[![Test](https://github.com/khulnasoft/webhooks/workflows/Test/badge.svg?branch=master)](https://github.com/khulnasoft/webhooks/actions)
[![Coverage Status](https://coveralls.io/repos/khulnasoft/webhooks/badge.svg?branch=master&service=github)](https://coveralls.io/github/khulnasoft/webhooks?branch=master)
[![Go Report Card](https://goreportcard.com/badge/khulnasoft/webhooks)](https://goreportcard.com/report/khulnasoft/webhooks)
[![GoDoc](https://godoc.org/github.com/khulnasoft/webhooks?status.svg)](https://godoc.org/github.com/khulnasoft/webhooks)

Library webhooks allows for easy receiving and parsing of GitHub, Bitbucket, GitLab, Docker Hub, Gogs and Azure DevOps Webhook Events

Features:

* Parses the entire payload, not just a few fields.
* Fields + Schema directly lines up with webhook posted json

Notes:

* Currently only accepting json payloads.

Installation
------------

Use go get.

```shell
go get -u github.com/khulnasoft/webhooks
```

Then import the package into your own code.

	import "github.com/khulnasoft/webhooks"

Usage and Documentation
------

Please see http://godoc.org/github.com/khulnasoft/webhooks for detailed usage docs.

##### Examples:
```go
package main

import (
	"fmt"

	"net/http"

	"github.com/khulnasoft/webhooks/github"
)

const (
	path = "/webhooks"
)

func main() {
	hook, _ := github.New(github.Options.Secret("MyGitHubSuperSecretSecret...?"))

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r, github.ReleaseEvent, github.PullRequestEvent)
		if err != nil {
			if err == github.ErrEventNotFound {
				// ok event wasn't one of the ones asked to be parsed
			}
		}
		switch payload.(type) {

		case github.ReleasePayload:
			release := payload.(github.ReleasePayload)
			// Do whatever you want from here...
			fmt.Printf("%+v", release)

		case github.PullRequestPayload:
			pullRequest := payload.(github.PullRequestPayload)
			// Do whatever you want from here...
			fmt.Printf("%+v", pullRequest)
		}
	})
	http.ListenAndServe(":3000", nil)
}

```

Contributing
------

Pull requests for other services are welcome!

If the changes being proposed or requested are breaking changes, please create an issue for discussion.

License
------

Distributed under MIT License, please see license file in code for more details.
