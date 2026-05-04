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

var tasksGetMultipleCmd = &cobra.Command{
	Use:   "get-multiple",
	Short: "Get multiple tasks",
	RunE:  runTasksGetMultiple,
}

var tasksGetMultipleFlags struct {
	optPretty      bool
	limit          int
	offset         string
	assignee       string
	project        string
	section        string
	workspace      string
	completedSince string
	modifiedSince  string
	optFields      []string
}

func init() {
	tasksGetMultipleCmd.Flags().BoolVar(&tasksGetMultipleFlags.optPretty, "opt-pretty", false, "Provides “pretty” output. Provides the response in a “pretty” format. In the case of JSON this means doing proper line breaking and indentation to make it readable. This will take extra time and increase the response size so it is advisable only to use this during debugging.")
	tasksGetMultipleCmd.Flags().IntVar(&tasksGetMultipleFlags.limit, "limit", 0, "Results per page. The number of objects to return per page. The value must be between 1 and 100.")
	tasksGetMultipleCmd.Flags().StringVar(&tasksGetMultipleFlags.offset, "offset", "", "Offset token. An offset to the next page returned by the API. A pagination request will return an offset token, which can be used as an input parameter to the next request. If an offset is not passed in, the API will return the first page of results. *Note: You can only pass in an offset that was returned to you via a previously paginated request.*")
	tasksGetMultipleCmd.Flags().StringVar(&tasksGetMultipleFlags.assignee, "assignee", "", "The assignee to filter tasks on. If searching for unassigned tasks, assignee.any = null can be specified. *Note: If you specify `assignee`, you must also specify the `workspace` to filter on.*")
	tasksGetMultipleCmd.Flags().StringVar(&tasksGetMultipleFlags.project, "project", "", "The project to filter tasks on.")
	tasksGetMultipleCmd.Flags().StringVar(&tasksGetMultipleFlags.section, "section", "", "The section to filter tasks on.")
	tasksGetMultipleCmd.Flags().StringVar(&tasksGetMultipleFlags.workspace, "workspace", "", "The workspace to filter tasks on. *Note: If you specify `workspace`, you must also specify the `assignee` to filter on.*")
	tasksGetMultipleCmd.Flags().StringVar(&tasksGetMultipleFlags.completedSince, "completed-since", "", "Only return tasks that are either incomplete or that have been completed since this time.")
	tasksGetMultipleCmd.Flags().StringVar(&tasksGetMultipleFlags.modifiedSince, "modified-since", "", "Only return tasks that have been modified since the given time.  *Note: A task is considered “modified” if any of its properties change, or associations between it and other objects are modified (e.g.  a task being added to a project). A task is not considered modified just because another object it is associated with (e.g. a subtask) is modified. Actions that count as modifying the task include assigning, renaming, completing, and adding stories.*")
	tasksGetMultipleCmd.Flags().StringSliceVar(&tasksGetMultipleFlags.optFields, "opt-fields", nil, "This endpoint returns a compact resource, which excludes some properties by default. To include those optional properties, set this query parameter to a comma-separated list of the properties you wish to include.")

	tasksCmd.AddCommand(tasksGetMultipleCmd)
}

func runTasksGetMultiple(cmd *cobra.Command, args []string) error {
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
			Name:        "opt-pretty",
			Type:        "boolean",
			Required:    false,
			Location:    "query",
			Description: "Provides “pretty” output. Provides the response in a “pretty” format. In the case of JSON this means doing proper line breaking and indentation to make it readable. This will take extra time and increase the response size so it is advisable only to use this during debugging.",
		})
		flags = append(flags, flagSchema{
			Name:        "limit",
			Type:        "integer",
			Required:    false,
			Location:    "query",
			Description: "Results per page. The number of objects to return per page. The value must be between 1 and 100.",
		})
		flags = append(flags, flagSchema{
			Name:        "offset",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "Offset token. An offset to the next page returned by the API. A pagination request will return an offset token, which can be used as an input parameter to the next request. If an offset is not passed in, the API will return the first page of results. *Note: You can only pass in an offset that was returned to you via a previously paginated request.*",
		})
		flags = append(flags, flagSchema{
			Name:        "assignee",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "The assignee to filter tasks on. If searching for unassigned tasks, assignee.any = null can be specified. *Note: If you specify `assignee`, you must also specify the `workspace` to filter on.*",
		})
		flags = append(flags, flagSchema{
			Name:        "project",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "The project to filter tasks on.",
		})
		flags = append(flags, flagSchema{
			Name:        "section",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "The section to filter tasks on.",
		})
		flags = append(flags, flagSchema{
			Name:        "workspace",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "The workspace to filter tasks on. *Note: If you specify `workspace`, you must also specify the `assignee` to filter on.*",
		})
		flags = append(flags, flagSchema{
			Name:        "completed-since",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "Only return tasks that are either incomplete or that have been completed since this time.",
		})
		flags = append(flags, flagSchema{
			Name:        "modified-since",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "Only return tasks that have been modified since the given time.  *Note: A task is considered “modified” if any of its properties change, or associations between it and other objects are modified (e.g.  a task being added to a project). A task is not considered modified just because another object it is associated with (e.g. a subtask) is modified. Actions that count as modifying the task include assigning, renaming, completing, and adding stories.*",
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
			Description: "Successfully retrieved requested tasks.",
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
			"command":     "get-multiple",
			"description": "Get multiple tasks",
			"http": map[string]any{
				"method": "GET",
				"path":   "/tasks",
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

	req := &httpclient.Request{
		Method:      "GET",
		Path:        httpclient.SubstitutePath("/tasks", pathParams),
		QueryParams: map[string]string{},
		ArrayParams: map[string][]string{},
		Headers:     map[string]string{},
	}

	// Query parameters
	if cmd.Flags().Changed("opt-pretty") {
		req.QueryParams["opt_pretty"] = fmt.Sprintf("%v", tasksGetMultipleFlags.optPretty)
	}
	if cmd.Flags().Changed("limit") {
		req.QueryParams["limit"] = fmt.Sprintf("%v", tasksGetMultipleFlags.limit)
	}
	if cmd.Flags().Changed("offset") {
		req.QueryParams["offset"] = fmt.Sprintf("%v", tasksGetMultipleFlags.offset)
	}
	if cmd.Flags().Changed("assignee") {
		req.QueryParams["assignee"] = fmt.Sprintf("%v", tasksGetMultipleFlags.assignee)
	}
	if cmd.Flags().Changed("project") {
		req.QueryParams["project"] = fmt.Sprintf("%v", tasksGetMultipleFlags.project)
	}
	if cmd.Flags().Changed("section") {
		req.QueryParams["section"] = fmt.Sprintf("%v", tasksGetMultipleFlags.section)
	}
	if cmd.Flags().Changed("workspace") {
		req.QueryParams["workspace"] = fmt.Sprintf("%v", tasksGetMultipleFlags.workspace)
	}
	if cmd.Flags().Changed("completed-since") {
		req.QueryParams["completed_since"] = fmt.Sprintf("%v", tasksGetMultipleFlags.completedSince)
	}
	if cmd.Flags().Changed("modified-since") {
		req.QueryParams["modified_since"] = fmt.Sprintf("%v", tasksGetMultipleFlags.modifiedSince)
	}
	if cmd.Flags().Changed("opt-fields") {
		req.ArrayParams["opt_fields"] = tasksGetMultipleFlags.optFields
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
