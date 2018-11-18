// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// user HTTP client CLI support package
//
// Command:
// $ goa gen github.com/proepkes/speeddate/gamehostsvc/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	repositoryc "github.com/proepkes/speeddate/gamehostsvc/gen/http/repository/client"
	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `repository (insert|delete|get)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` repository insert --body '{
      "name": "3xq"
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		repositoryFlags = flag.NewFlagSet("repository", flag.ContinueOnError)

		repositoryInsertFlags    = flag.NewFlagSet("insert", flag.ExitOnError)
		repositoryInsertBodyFlag = repositoryInsertFlags.String("body", "REQUIRED", "")

		repositoryDeleteFlags  = flag.NewFlagSet("delete", flag.ExitOnError)
		repositoryDeleteIDFlag = repositoryDeleteFlags.String("id", "REQUIRED", "ID of user to remove")

		repositoryGetFlags     = flag.NewFlagSet("get", flag.ExitOnError)
		repositoryGetIDFlag    = repositoryGetFlags.String("id", "REQUIRED", "Get user by ID")
		repositoryGetViewFlag  = repositoryGetFlags.String("view", "", "")
		repositoryGetTokenFlag = repositoryGetFlags.String("token", "", "")
	)
	repositoryFlags.Usage = repositoryUsage
	repositoryInsertFlags.Usage = repositoryInsertUsage
	repositoryDeleteFlags.Usage = repositoryDeleteUsage
	repositoryGetFlags.Usage = repositoryGetUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if len(os.Args) < flag.NFlag()+3 {
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = os.Args[1+flag.NFlag()]
		switch svcn {
		case "repository":
			svcf = repositoryFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(os.Args[2+flag.NFlag():]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = os.Args[2+flag.NFlag()+svcf.NFlag()]
		switch svcn {
		case "repository":
			switch epn {
			case "insert":
				epf = repositoryInsertFlags

			case "delete":
				epf = repositoryDeleteFlags

			case "get":
				epf = repositoryGetFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if len(os.Args) > 2+flag.NFlag()+svcf.NFlag() {
		if err := epf.Parse(os.Args[3+flag.NFlag()+svcf.NFlag():]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "repository":
			c := repositoryc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "insert":
				endpoint = c.Insert()
				data, err = repositoryc.BuildInsertPayload(*repositoryInsertBodyFlag)
			case "delete":
				endpoint = c.Delete()
				data, err = repositoryc.BuildDeletePayload(*repositoryDeleteIDFlag)
			case "get":
				endpoint = c.Get()
				data, err = repositoryc.BuildGetPayload(*repositoryGetIDFlag, *repositoryGetViewFlag, *repositoryGetTokenFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// repositoryUsage displays the usage of the repository command and its
// subcommands.
func repositoryUsage() {
	fmt.Fprintf(os.Stderr, `The service makes it possible to insert, delete or get users.
Usage:
    %s [globalflags] repository COMMAND [flags]

COMMAND:
    insert: Add new user and return its ID.
    delete: Remove user from storage
    get: Get implements get.

Additional help:
    %s repository COMMAND --help
`, os.Args[0], os.Args[0])
}
func repositoryInsertUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] repository insert -body JSON

Add new user and return its ID.
    -body JSON: 

Example:
    `+os.Args[0]+` repository insert --body '{
      "name": "3xq"
   }'
`, os.Args[0])
}

func repositoryDeleteUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] repository delete -id STRING

Remove user from storage
    -id STRING: ID of user to remove

Example:
    `+os.Args[0]+` repository delete --id "Ut soluta eius quia."
`, os.Args[0])
}

func repositoryGetUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] repository get -id STRING -view STRING -token STRING

Get implements get.
    -id STRING: Get user by ID
    -view STRING: 
    -token STRING: 

Example:
    `+os.Args[0]+` repository get --id "Harum error est." --view "tiny" --token "Dicta laudantium iste ut cupiditate."
`, os.Args[0])
}
