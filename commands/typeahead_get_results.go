package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/rishimantri795/CLICreator/runtime/httpclient"
	"github.com/rishimantri795/CLICreator/runtime/output"
	"github.com/spf13/cobra"
)

var typeaheadGetResultsCmd = &cobra.Command{
	Use:   "get-results",
	Short: "Get objects via typeahead",
	RunE:  runTypeaheadGetResults,
}

var typeaheadGetResultsFlags struct {
	workspaceGid string
	resourceType string
	type_        string
	query        string
	count        int
	optPretty    bool
	optFields    []string
}

func init() {
	typeaheadGetResultsCmd.Flags().StringVar(&typeaheadGetResultsFlags.workspaceGid, "workspace-gid", "", "Globally unique identifier for the workspace or organization.")
	typeaheadGetResultsCmd.MarkFlagRequired("workspace-gid")
	typeaheadGetResultsCmd.Flags().StringVar(&typeaheadGetResultsFlags.resourceType, "resource-type", "", "The type of values the typeahead should return. You can choose from one of the following: `custom_field`, `goal`, `project`, `project_template`, `portfolio`, `tag`, `task`, `team`, and `user`. Note that unlike in the names of endpoints, the types listed here are in singular form (e.g. `task`). Using multiple types is not yet supported.")
	typeaheadGetResultsCmd.MarkFlagRequired("resource-type")
	typeaheadGetResultsCmd.Flags().StringVar(&typeaheadGetResultsFlags.type_, "type", "", "*Deprecated: new integrations should prefer the resource_type field.*")
	typeaheadGetResultsCmd.Flags().StringVar(&typeaheadGetResultsFlags.query, "query", "", "The string that will be used to search for relevant objects. If an empty string is passed in, the API will return results.")
	typeaheadGetResultsCmd.Flags().IntVar(&typeaheadGetResultsFlags.count, "count", 0, "The number of results to return. The default is 20 if this parameter is omitted, with a minimum of 1 and a maximum of 100. If there are fewer results found than requested, all will be returned.")
	typeaheadGetResultsCmd.Flags().BoolVar(&typeaheadGetResultsFlags.optPretty, "opt-pretty", false, "Provides “pretty” output. Provides the response in a “pretty” format. In the case of JSON this means doing proper line breaking and indentation to make it readable. This will take extra time and increase the response size so it is advisable only to use this during debugging.")
	typeaheadGetResultsCmd.Flags().StringSliceVar(&typeaheadGetResultsFlags.optFields, "opt-fields", nil, "This endpoint returns a compact resource, which excludes some properties by default. To include those optional properties, set this query parameter to a comma-separated list of the properties you wish to include.")

	typeaheadCmd.AddCommand(typeaheadGetResultsCmd)
}

func runTypeaheadGetResults(cmd *cobra.Command, args []string) error {
	// --schema: print full input/output type contract without making any network call.
	if rootFlags.schema {
		type flagSchema struct {
			Name        string `json:"name"`
			Type        string `json:"type"`
			Required    bool   `json:"required"`
			Location    string `json:"location"`
			Description string `json:"description,omitempty"`
		}
		var flags []flagSchema
		flags = append(flags, flagSchema{
			Name:        "workspace-gid",
			Type:        "string",
			Required:    true,
			Location:    "path",
			Description: "Globally unique identifier for the workspace or organization.",
		})
		flags = append(flags, flagSchema{
			Name:        "resource-type",
			Type:        "string",
			Required:    true,
			Location:    "query",
			Description: "The type of values the typeahead should return. You can choose from one of the following: `custom_field`, `goal`, `project`, `project_template`, `portfolio`, `tag`, `task`, `team`, and `user`. Note that unlike in the names of endpoints, the types listed here are in singular form (e.g. `task`). Using multiple types is not yet supported.",
		})
		flags = append(flags, flagSchema{
			Name:        "type",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "*Deprecated: new integrations should prefer the resource_type field.*",
		})
		flags = append(flags, flagSchema{
			Name:        "query",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "The string that will be used to search for relevant objects. If an empty string is passed in, the API will return results.",
		})
		flags = append(flags, flagSchema{
			Name:        "count",
			Type:        "integer",
			Required:    false,
			Location:    "query",
			Description: "The number of results to return. The default is 20 if this parameter is omitted, with a minimum of 1 and a maximum of 100. If there are fewer results found than requested, all will be returned.",
		})
		flags = append(flags, flagSchema{
			Name:        "opt-pretty",
			Type:        "boolean",
			Required:    false,
			Location:    "query",
			Description: "Provides “pretty” output. Provides the response in a “pretty” format. In the case of JSON this means doing proper line breaking and indentation to make it readable. This will take extra time and increase the response size so it is advisable only to use this during debugging.",
		})
		flags = append(flags, flagSchema{
			Name:        "opt-fields",
			Type:        "array",
			Required:    false,
			Location:    "query",
			Description: "This endpoint returns a compact resource, which excludes some properties by default. To include those optional properties, set this query parameter to a comma-separated list of the properties you wish to include.",
		})

		type responseSchema struct {
			Status      string `json:"status"`
			ContentType string `json:"content_type,omitempty"`
			Description string `json:"description,omitempty"`
		}
		var responses []responseSchema
		responses = append(responses, responseSchema{
			Status:      "200",
			ContentType: "application/json",
			Description: "Successfully retrieved objects via a typeahead search algorithm.",
		})
		responses = append(responses, responseSchema{
			Status:      "400",
			ContentType: "application/json",
			Description: "This usually occurs because of a missing or malformed parameter. Check the documentation and the syntax of your request and try again.",
		})
		responses = append(responses, responseSchema{
			Status:      "401",
			ContentType: "application/json",
			Description: "A valid authentication token was not provided with the request, so the API could not associate a user with the request.",
		})
		responses = append(responses, responseSchema{
			Status:      "403",
			ContentType: "application/json",
			Description: "The authentication and request syntax was valid but the server is refusing to complete the request. This can happen if you try to read or write to objects or properties that the user does not have access to.",
		})
		responses = append(responses, responseSchema{
			Status:      "404",
			ContentType: "application/json",
			Description: "Either the request method and path supplied do not specify a known action in the API, or the object specified by the request does not exist.",
		})
		responses = append(responses, responseSchema{
			Status:      "500",
			ContentType: "application/json",
			Description: "There was a problem on Asana’s end. In the event of a server error the response body should contain an error phrase. These phrases can be used by Asana support to quickly look up the incident that caused the server error. Some errors are due to server load, and will not supply an error phrase.",
		})

		schema := map[string]any{
			"command":     "get-results",
			"description": "Get objects via typeahead",
			"http": map[string]any{
				"method": "GET",
				"path":   "/workspaces/{workspace_gid}/typeahead",
			},
			"input": map[string]any{
				"flags":         flags,
				"body_flag":     false,
				"body_required": false,
			},
			"output": map[string]any{
				"responses": responses,
			},
			"semantics": map[string]any{
				"safe":         true,
				"idempotent":   true,
				"reversible":   true,
				"side_effects": []string{},
				"impact":       "low",
			},
			"requires_auth": true,
		}
		data, _ := json.MarshalIndent(schema, "", "  ")
		fmt.Fprintln(_stdoutCounter, string(data))
		return nil
	}

	cfg, err := rootConfig()
	if err != nil {
		e := output.NetworkError(err)
		e.Write(os.Stderr)
		return output.NewExitError(e)
	}

	client := httpclient.New(cfg.BaseURL, cfg.AuthProvider())
	client.Debug = rootFlags.debug
	client.DryRun = rootFlags.dryRun
	if rootFlags.noRetries {
		client.RetryConfig.MaxRetries = 0
	}

	// Build path params
	pathParams := map[string]string{}
	pathParams["workspace_gid"] = fmt.Sprintf("%v", typeaheadGetResultsFlags.workspaceGid)

	req := &httpclient.Request{
		Method:      "GET",
		Path:        httpclient.SubstitutePath("/workspaces/{workspace_gid}/typeahead", pathParams),
		QueryParams: map[string]string{},
		ArrayParams: map[string][]string{},
		Headers:     map[string]string{},
	}

	// Query parameters
	if cmd.Flags().Changed("resource-type") {
		req.QueryParams["resource_type"] = fmt.Sprintf("%v", typeaheadGetResultsFlags.resourceType)
	}
	if cmd.Flags().Changed("type") {
		req.QueryParams["type"] = fmt.Sprintf("%v", typeaheadGetResultsFlags.type_)
	}
	if cmd.Flags().Changed("query") {
		req.QueryParams["query"] = fmt.Sprintf("%v", typeaheadGetResultsFlags.query)
	}
	if cmd.Flags().Changed("count") {
		req.QueryParams["count"] = fmt.Sprintf("%v", typeaheadGetResultsFlags.count)
	}
	if cmd.Flags().Changed("opt-pretty") {
		req.QueryParams["opt_pretty"] = fmt.Sprintf("%v", typeaheadGetResultsFlags.optPretty)
	}
	if cmd.Flags().Changed("opt-fields") {
		req.ArrayParams["opt_fields"] = typeaheadGetResultsFlags.optFields
	}

	// Header parameters

	resp, err := client.Do(req)
	if err != nil {
		if strings.Contains(err.Error(), "timeout") || strings.Contains(err.Error(), "deadline exceeded") {
			_invState.errorType = "timeout"
		} else {
			_invState.errorType = "network_error"
		}
		e := output.NetworkError(err)
		e.Write(os.Stderr)
		return output.NewExitError(e)
	}

	if resp.StatusCode >= 400 {
		if resp.StatusCode >= 500 {
			_invState.errorType = "http_5xx"
		} else {
			_invState.errorType = "http_4xx"
		}
		_invState.errorCode = resp.StatusCode
		e := output.HTTPError(resp.StatusCode, resp.Body)
		e.Write(os.Stderr)
		return output.NewExitError(e)
	}

	if rootFlags.jq != "" {
		return output.JQFilter(_stdoutCounter, resp.Body, rootFlags.jq)
	}
	return output.Print(_stdoutCounter, resp.Body, output.Format(cfg.OutputFormat))
}
