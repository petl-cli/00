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

var goalRelationshipsGetCompactRecordsCmd = &cobra.Command{
	Use:   "get-compact-records",
	Short: "Get goal relationships",
	RunE:  runGoalRelationshipsGetCompactRecords,
}

var goalRelationshipsGetCompactRecordsFlags struct {
	optPretty       bool
	limit           int
	offset          string
	supportedGoal   string
	resourceSubtype string
	optFields       []string
}

func init() {
	goalRelationshipsGetCompactRecordsCmd.Flags().BoolVar(&goalRelationshipsGetCompactRecordsFlags.optPretty, "opt-pretty", false, "Provides “pretty” output. Provides the response in a “pretty” format. In the case of JSON this means doing proper line breaking and indentation to make it readable. This will take extra time and increase the response size so it is advisable only to use this during debugging.")
	goalRelationshipsGetCompactRecordsCmd.Flags().IntVar(&goalRelationshipsGetCompactRecordsFlags.limit, "limit", 0, "Results per page. The number of objects to return per page. The value must be between 1 and 100.")
	goalRelationshipsGetCompactRecordsCmd.Flags().StringVar(&goalRelationshipsGetCompactRecordsFlags.offset, "offset", "", "Offset token. An offset to the next page returned by the API. A pagination request will return an offset token, which can be used as an input parameter to the next request. If an offset is not passed in, the API will return the first page of results. *Note: You can only pass in an offset that was returned to you via a previously paginated request.*")
	goalRelationshipsGetCompactRecordsCmd.Flags().StringVar(&goalRelationshipsGetCompactRecordsFlags.supportedGoal, "supported-goal", "", "Globally unique identifier for the supported goal in the goal relationship.")
	goalRelationshipsGetCompactRecordsCmd.MarkFlagRequired("supported-goal")
	goalRelationshipsGetCompactRecordsCmd.Flags().StringVar(&goalRelationshipsGetCompactRecordsFlags.resourceSubtype, "resource-subtype", "", "If provided, filter to goal relationships with a given resource_subtype.")
	goalRelationshipsGetCompactRecordsCmd.Flags().StringSliceVar(&goalRelationshipsGetCompactRecordsFlags.optFields, "opt-fields", nil, "This endpoint returns a compact resource, which excludes some properties by default. To include those optional properties, set this query parameter to a comma-separated list of the properties you wish to include.")

	goalRelationshipsCmd.AddCommand(goalRelationshipsGetCompactRecordsCmd)
}

func runGoalRelationshipsGetCompactRecords(cmd *cobra.Command, args []string) error {
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
			Name:        "supported-goal",
			Type:        "string",
			Required:    true,
			Location:    "query",
			Description: "Globally unique identifier for the supported goal in the goal relationship.",
		})
		flags = append(flags, flagSchema{
			Name:        "resource-subtype",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "If provided, filter to goal relationships with a given resource_subtype.",
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
			Description: "Successfully retrieved the requested goal relationships.",
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
			"command":     "get-compact-records",
			"description": "Get goal relationships",
			"http": map[string]any{
				"method": "GET",
				"path":   "/goal_relationships",
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
		Path:        httpclient.SubstitutePath("/goal_relationships", pathParams),
		QueryParams: map[string]string{},
		ArrayParams: map[string][]string{},
		Headers:     map[string]string{},
	}

	// Query parameters
	if cmd.Flags().Changed("opt-pretty") {
		req.QueryParams["opt_pretty"] = fmt.Sprintf("%v", goalRelationshipsGetCompactRecordsFlags.optPretty)
	}
	if cmd.Flags().Changed("limit") {
		req.QueryParams["limit"] = fmt.Sprintf("%v", goalRelationshipsGetCompactRecordsFlags.limit)
	}
	if cmd.Flags().Changed("offset") {
		req.QueryParams["offset"] = fmt.Sprintf("%v", goalRelationshipsGetCompactRecordsFlags.offset)
	}
	if cmd.Flags().Changed("supported-goal") {
		req.QueryParams["supported_goal"] = fmt.Sprintf("%v", goalRelationshipsGetCompactRecordsFlags.supportedGoal)
	}
	if cmd.Flags().Changed("resource-subtype") {
		req.QueryParams["resource_subtype"] = fmt.Sprintf("%v", goalRelationshipsGetCompactRecordsFlags.resourceSubtype)
	}
	if cmd.Flags().Changed("opt-fields") {
		req.ArrayParams["opt_fields"] = goalRelationshipsGetCompactRecordsFlags.optFields
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
