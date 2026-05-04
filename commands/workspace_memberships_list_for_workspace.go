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

var workspaceMembershipsListForWorkspaceCmd = &cobra.Command{
	Use:   "list-for-workspace",
	Short: "Get the workspace memberships for a workspace",
	RunE:  runWorkspaceMembershipsListForWorkspace,
}

var workspaceMembershipsListForWorkspaceFlags struct {
	workspaceGid string
	user         string
	optPretty    bool
	limit        int
	offset       string
	optFields    []string
}

func init() {
	workspaceMembershipsListForWorkspaceCmd.Flags().StringVar(&workspaceMembershipsListForWorkspaceFlags.workspaceGid, "workspace-gid", "", "Globally unique identifier for the workspace or organization.")
	workspaceMembershipsListForWorkspaceCmd.MarkFlagRequired("workspace-gid")
	workspaceMembershipsListForWorkspaceCmd.Flags().StringVar(&workspaceMembershipsListForWorkspaceFlags.user, "user", "", "A string identifying a user. This can either be the string \"me\", an email, or the gid of a user.")
	workspaceMembershipsListForWorkspaceCmd.Flags().BoolVar(&workspaceMembershipsListForWorkspaceFlags.optPretty, "opt-pretty", false, "Provides “pretty” output. Provides the response in a “pretty” format. In the case of JSON this means doing proper line breaking and indentation to make it readable. This will take extra time and increase the response size so it is advisable only to use this during debugging.")
	workspaceMembershipsListForWorkspaceCmd.Flags().IntVar(&workspaceMembershipsListForWorkspaceFlags.limit, "limit", 0, "Results per page. The number of objects to return per page. The value must be between 1 and 100.")
	workspaceMembershipsListForWorkspaceCmd.Flags().StringVar(&workspaceMembershipsListForWorkspaceFlags.offset, "offset", "", "Offset token. An offset to the next page returned by the API. A pagination request will return an offset token, which can be used as an input parameter to the next request. If an offset is not passed in, the API will return the first page of results. *Note: You can only pass in an offset that was returned to you via a previously paginated request.*")
	workspaceMembershipsListForWorkspaceCmd.Flags().StringSliceVar(&workspaceMembershipsListForWorkspaceFlags.optFields, "opt-fields", nil, "This endpoint returns a compact resource, which excludes some properties by default. To include those optional properties, set this query parameter to a comma-separated list of the properties you wish to include.")

	workspaceMembershipsCmd.AddCommand(workspaceMembershipsListForWorkspaceCmd)
}

func runWorkspaceMembershipsListForWorkspace(cmd *cobra.Command, args []string) error {
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
			Name:        "user",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "A string identifying a user. This can either be the string \"me\", an email, or the gid of a user.",
		})
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
			Description: "Successfully retrieved the requested workspace's memberships.",
		})

		schema := map[string]any{
			"command":     "list-for-workspace",
			"description": "Get the workspace memberships for a workspace",
			"http": map[string]any{
				"method": "GET",
				"path":   "/workspaces/{workspace_gid}/workspace_memberships",
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
	pathParams["workspace_gid"] = fmt.Sprintf("%v", workspaceMembershipsListForWorkspaceFlags.workspaceGid)

	req := &httpclient.Request{
		Method:      "GET",
		Path:        httpclient.SubstitutePath("/workspaces/{workspace_gid}/workspace_memberships", pathParams),
		QueryParams: map[string]string{},
		ArrayParams: map[string][]string{},
		Headers:     map[string]string{},
	}

	// Query parameters
	if cmd.Flags().Changed("user") {
		req.QueryParams["user"] = fmt.Sprintf("%v", workspaceMembershipsListForWorkspaceFlags.user)
	}
	if cmd.Flags().Changed("opt-pretty") {
		req.QueryParams["opt_pretty"] = fmt.Sprintf("%v", workspaceMembershipsListForWorkspaceFlags.optPretty)
	}
	if cmd.Flags().Changed("limit") {
		req.QueryParams["limit"] = fmt.Sprintf("%v", workspaceMembershipsListForWorkspaceFlags.limit)
	}
	if cmd.Flags().Changed("offset") {
		req.QueryParams["offset"] = fmt.Sprintf("%v", workspaceMembershipsListForWorkspaceFlags.offset)
	}
	if cmd.Flags().Changed("opt-fields") {
		req.ArrayParams["opt_fields"] = workspaceMembershipsListForWorkspaceFlags.optFields
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
