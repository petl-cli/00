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

var tasksSearchInWorkspaceCmd = &cobra.Command{
	Use:   "search-in-workspace",
	Short: "Search tasks in a workspace",
	RunE:  runTasksSearchInWorkspace,
}

var tasksSearchInWorkspaceFlags struct {
	workspaceGid      string
	optPretty         bool
	text              string
	resourceSubtype   string
	assigneeAny       string
	assigneeNot       string
	portfoliosAny     string
	projectsAny       string
	projectsNot       string
	projectsAll       string
	sectionsAny       string
	sectionsNot       string
	sectionsAll       string
	tagsAny           string
	tagsNot           string
	tagsAll           string
	teamsAny          string
	followersNot      string
	createdByAny      string
	createdByNot      string
	assignedByAny     string
	assignedByNot     string
	likedByNot        string
	commentedOnByNot  string
	dueOnBefore       string
	dueOnAfter        string
	dueOn             string
	dueAtBefore       string
	dueAtAfter        string
	startOnBefore     string
	startOnAfter      string
	startOn           string
	createdOnBefore   string
	createdOnAfter    string
	createdOn         string
	createdAtBefore   string
	createdAtAfter    string
	completedOnBefore string
	completedOnAfter  string
	completedOn       string
	completedAtBefore string
	completedAtAfter  string
	modifiedOnBefore  string
	modifiedOnAfter   string
	modifiedOn        string
	modifiedAtBefore  string
	modifiedAtAfter   string
	isBlocking        bool
	isBlocked         bool
	hasAttachment     bool
	completed         bool
	isSubtask         bool
	sortBy            string
	sortAscending     bool
	optFields         []string
}

func init() {
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.workspaceGid, "workspace-gid", "", "Globally unique identifier for the workspace or organization.")
	tasksSearchInWorkspaceCmd.MarkFlagRequired("workspace-gid")
	tasksSearchInWorkspaceCmd.Flags().BoolVar(&tasksSearchInWorkspaceFlags.optPretty, "opt-pretty", false, "Provides “pretty” output. Provides the response in a “pretty” format. In the case of JSON this means doing proper line breaking and indentation to make it readable. This will take extra time and increase the response size so it is advisable only to use this during debugging.")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.text, "text", "", "Performs full-text search on both task name and description")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.resourceSubtype, "resource-subtype", "", "Filters results by the task's resource_subtype")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.assigneeAny, "assignee-any", "", "Comma-separated list of user identifiers")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.assigneeNot, "assignee-not", "", "Comma-separated list of user identifiers")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.portfoliosAny, "portfolios-any", "", "Comma-separated list of portfolio IDs")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.projectsAny, "projects-any", "", "Comma-separated list of project IDs")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.projectsNot, "projects-not", "", "Comma-separated list of project IDs")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.projectsAll, "projects-all", "", "Comma-separated list of project IDs")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.sectionsAny, "sections-any", "", "Comma-separated list of section or column IDs")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.sectionsNot, "sections-not", "", "Comma-separated list of section or column IDs")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.sectionsAll, "sections-all", "", "Comma-separated list of section or column IDs")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.tagsAny, "tags-any", "", "Comma-separated list of tag IDs")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.tagsNot, "tags-not", "", "Comma-separated list of tag IDs")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.tagsAll, "tags-all", "", "Comma-separated list of tag IDs")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.teamsAny, "teams-any", "", "Comma-separated list of team IDs")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.followersNot, "followers-not", "", "Comma-separated list of user identifiers")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.createdByAny, "created-by-any", "", "Comma-separated list of user identifiers")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.createdByNot, "created-by-not", "", "Comma-separated list of user identifiers")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.assignedByAny, "assigned-by-any", "", "Comma-separated list of user identifiers")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.assignedByNot, "assigned-by-not", "", "Comma-separated list of user identifiers")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.likedByNot, "liked-by-not", "", "Comma-separated list of user identifiers")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.commentedOnByNot, "commented-on-by-not", "", "Comma-separated list of user identifiers")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.dueOnBefore, "due-on-before", "", "ISO 8601 date string")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.dueOnAfter, "due-on-after", "", "ISO 8601 date string")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.dueOn, "due-on", "", "ISO 8601 date string or `null`")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.dueAtBefore, "due-at-before", "", "ISO 8601 datetime string")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.dueAtAfter, "due-at-after", "", "ISO 8601 datetime string")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.startOnBefore, "start-on-before", "", "ISO 8601 date string")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.startOnAfter, "start-on-after", "", "ISO 8601 date string")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.startOn, "start-on", "", "ISO 8601 date string or `null`")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.createdOnBefore, "created-on-before", "", "ISO 8601 date string")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.createdOnAfter, "created-on-after", "", "ISO 8601 date string")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.createdOn, "created-on", "", "ISO 8601 date string or `null`")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.createdAtBefore, "created-at-before", "", "ISO 8601 datetime string")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.createdAtAfter, "created-at-after", "", "ISO 8601 datetime string")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.completedOnBefore, "completed-on-before", "", "ISO 8601 date string")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.completedOnAfter, "completed-on-after", "", "ISO 8601 date string")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.completedOn, "completed-on", "", "ISO 8601 date string or `null`")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.completedAtBefore, "completed-at-before", "", "ISO 8601 datetime string")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.completedAtAfter, "completed-at-after", "", "ISO 8601 datetime string")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.modifiedOnBefore, "modified-on-before", "", "ISO 8601 date string")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.modifiedOnAfter, "modified-on-after", "", "ISO 8601 date string")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.modifiedOn, "modified-on", "", "ISO 8601 date string or `null`")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.modifiedAtBefore, "modified-at-before", "", "ISO 8601 datetime string")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.modifiedAtAfter, "modified-at-after", "", "ISO 8601 datetime string")
	tasksSearchInWorkspaceCmd.Flags().BoolVar(&tasksSearchInWorkspaceFlags.isBlocking, "is-blocking", false, "Filter to incomplete tasks with dependents")
	tasksSearchInWorkspaceCmd.Flags().BoolVar(&tasksSearchInWorkspaceFlags.isBlocked, "is-blocked", false, "Filter to tasks with incomplete dependencies")
	tasksSearchInWorkspaceCmd.Flags().BoolVar(&tasksSearchInWorkspaceFlags.hasAttachment, "has-attachment", false, "Filter to tasks with attachments")
	tasksSearchInWorkspaceCmd.Flags().BoolVar(&tasksSearchInWorkspaceFlags.completed, "completed", false, "Filter to completed tasks")
	tasksSearchInWorkspaceCmd.Flags().BoolVar(&tasksSearchInWorkspaceFlags.isSubtask, "is-subtask", false, "Filter to subtasks")
	tasksSearchInWorkspaceCmd.Flags().StringVar(&tasksSearchInWorkspaceFlags.sortBy, "sort-by", "", "One of `due_date`, `created_at`, `completed_at`, `likes`, or `modified_at`, defaults to `modified_at`")
	tasksSearchInWorkspaceCmd.Flags().BoolVar(&tasksSearchInWorkspaceFlags.sortAscending, "sort-ascending", false, "Default `false`")
	tasksSearchInWorkspaceCmd.Flags().StringSliceVar(&tasksSearchInWorkspaceFlags.optFields, "opt-fields", nil, "This endpoint returns a compact resource, which excludes some properties by default. To include those optional properties, set this query parameter to a comma-separated list of the properties you wish to include.")

	tasksCmd.AddCommand(tasksSearchInWorkspaceCmd)
}

func runTasksSearchInWorkspace(cmd *cobra.Command, args []string) error {
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
			Name:        "opt-pretty",
			Type:        "boolean",
			Required:    false,
			Location:    "query",
			Description: "Provides “pretty” output. Provides the response in a “pretty” format. In the case of JSON this means doing proper line breaking and indentation to make it readable. This will take extra time and increase the response size so it is advisable only to use this during debugging.",
		})
		flags = append(flags, flagSchema{
			Name:        "text",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "Performs full-text search on both task name and description",
		})
		flags = append(flags, flagSchema{
			Name:        "resource-subtype",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "Filters results by the task's resource_subtype",
		})
		flags = append(flags, flagSchema{
			Name:        "assignee-any",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "Comma-separated list of user identifiers",
		})
		flags = append(flags, flagSchema{
			Name:        "assignee-not",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "Comma-separated list of user identifiers",
		})
		flags = append(flags, flagSchema{
			Name:        "portfolios-any",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "Comma-separated list of portfolio IDs",
		})
		flags = append(flags, flagSchema{
			Name:        "projects-any",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "Comma-separated list of project IDs",
		})
		flags = append(flags, flagSchema{
			Name:        "projects-not",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "Comma-separated list of project IDs",
		})
		flags = append(flags, flagSchema{
			Name:        "projects-all",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "Comma-separated list of project IDs",
		})
		flags = append(flags, flagSchema{
			Name:        "sections-any",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "Comma-separated list of section or column IDs",
		})
		flags = append(flags, flagSchema{
			Name:        "sections-not",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "Comma-separated list of section or column IDs",
		})
		flags = append(flags, flagSchema{
			Name:        "sections-all",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "Comma-separated list of section or column IDs",
		})
		flags = append(flags, flagSchema{
			Name:        "tags-any",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "Comma-separated list of tag IDs",
		})
		flags = append(flags, flagSchema{
			Name:        "tags-not",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "Comma-separated list of tag IDs",
		})
		flags = append(flags, flagSchema{
			Name:        "tags-all",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "Comma-separated list of tag IDs",
		})
		flags = append(flags, flagSchema{
			Name:        "teams-any",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "Comma-separated list of team IDs",
		})
		flags = append(flags, flagSchema{
			Name:        "followers-not",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "Comma-separated list of user identifiers",
		})
		flags = append(flags, flagSchema{
			Name:        "created-by-any",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "Comma-separated list of user identifiers",
		})
		flags = append(flags, flagSchema{
			Name:        "created-by-not",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "Comma-separated list of user identifiers",
		})
		flags = append(flags, flagSchema{
			Name:        "assigned-by-any",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "Comma-separated list of user identifiers",
		})
		flags = append(flags, flagSchema{
			Name:        "assigned-by-not",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "Comma-separated list of user identifiers",
		})
		flags = append(flags, flagSchema{
			Name:        "liked-by-not",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "Comma-separated list of user identifiers",
		})
		flags = append(flags, flagSchema{
			Name:        "commented-on-by-not",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "Comma-separated list of user identifiers",
		})
		flags = append(flags, flagSchema{
			Name:        "due-on-before",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "ISO 8601 date string",
		})
		flags = append(flags, flagSchema{
			Name:        "due-on-after",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "ISO 8601 date string",
		})
		flags = append(flags, flagSchema{
			Name:        "due-on",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "ISO 8601 date string or `null`",
		})
		flags = append(flags, flagSchema{
			Name:        "due-at-before",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "ISO 8601 datetime string",
		})
		flags = append(flags, flagSchema{
			Name:        "due-at-after",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "ISO 8601 datetime string",
		})
		flags = append(flags, flagSchema{
			Name:        "start-on-before",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "ISO 8601 date string",
		})
		flags = append(flags, flagSchema{
			Name:        "start-on-after",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "ISO 8601 date string",
		})
		flags = append(flags, flagSchema{
			Name:        "start-on",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "ISO 8601 date string or `null`",
		})
		flags = append(flags, flagSchema{
			Name:        "created-on-before",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "ISO 8601 date string",
		})
		flags = append(flags, flagSchema{
			Name:        "created-on-after",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "ISO 8601 date string",
		})
		flags = append(flags, flagSchema{
			Name:        "created-on",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "ISO 8601 date string or `null`",
		})
		flags = append(flags, flagSchema{
			Name:        "created-at-before",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "ISO 8601 datetime string",
		})
		flags = append(flags, flagSchema{
			Name:        "created-at-after",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "ISO 8601 datetime string",
		})
		flags = append(flags, flagSchema{
			Name:        "completed-on-before",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "ISO 8601 date string",
		})
		flags = append(flags, flagSchema{
			Name:        "completed-on-after",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "ISO 8601 date string",
		})
		flags = append(flags, flagSchema{
			Name:        "completed-on",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "ISO 8601 date string or `null`",
		})
		flags = append(flags, flagSchema{
			Name:        "completed-at-before",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "ISO 8601 datetime string",
		})
		flags = append(flags, flagSchema{
			Name:        "completed-at-after",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "ISO 8601 datetime string",
		})
		flags = append(flags, flagSchema{
			Name:        "modified-on-before",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "ISO 8601 date string",
		})
		flags = append(flags, flagSchema{
			Name:        "modified-on-after",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "ISO 8601 date string",
		})
		flags = append(flags, flagSchema{
			Name:        "modified-on",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "ISO 8601 date string or `null`",
		})
		flags = append(flags, flagSchema{
			Name:        "modified-at-before",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "ISO 8601 datetime string",
		})
		flags = append(flags, flagSchema{
			Name:        "modified-at-after",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "ISO 8601 datetime string",
		})
		flags = append(flags, flagSchema{
			Name:        "is-blocking",
			Type:        "boolean",
			Required:    false,
			Location:    "query",
			Description: "Filter to incomplete tasks with dependents",
		})
		flags = append(flags, flagSchema{
			Name:        "is-blocked",
			Type:        "boolean",
			Required:    false,
			Location:    "query",
			Description: "Filter to tasks with incomplete dependencies",
		})
		flags = append(flags, flagSchema{
			Name:        "has-attachment",
			Type:        "boolean",
			Required:    false,
			Location:    "query",
			Description: "Filter to tasks with attachments",
		})
		flags = append(flags, flagSchema{
			Name:        "completed",
			Type:        "boolean",
			Required:    false,
			Location:    "query",
			Description: "Filter to completed tasks",
		})
		flags = append(flags, flagSchema{
			Name:        "is-subtask",
			Type:        "boolean",
			Required:    false,
			Location:    "query",
			Description: "Filter to subtasks",
		})
		flags = append(flags, flagSchema{
			Name:        "sort-by",
			Type:        "string",
			Required:    false,
			Location:    "query",
			Description: "One of `due_date`, `created_at`, `completed_at`, `likes`, or `modified_at`, defaults to `modified_at`",
		})
		flags = append(flags, flagSchema{
			Name:        "sort-ascending",
			Type:        "boolean",
			Required:    false,
			Location:    "query",
			Description: "Default `false`",
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
			Description: "Successfully retrieved the section's tasks.",
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
			"command":     "search-in-workspace",
			"description": "Search tasks in a workspace",
			"http": map[string]any{
				"method": "GET",
				"path":   "/workspaces/{workspace_gid}/tasks/search",
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
	pathParams["workspace_gid"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.workspaceGid)

	req := &httpclient.Request{
		Method:      "GET",
		Path:        httpclient.SubstitutePath("/workspaces/{workspace_gid}/tasks/search", pathParams),
		QueryParams: map[string]string{},
		ArrayParams: map[string][]string{},
		Headers:     map[string]string{},
	}

	// Query parameters
	if cmd.Flags().Changed("opt-pretty") {
		req.QueryParams["opt_pretty"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.optPretty)
	}
	if cmd.Flags().Changed("text") {
		req.QueryParams["text"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.text)
	}
	if cmd.Flags().Changed("resource-subtype") {
		req.QueryParams["resource_subtype"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.resourceSubtype)
	}
	if cmd.Flags().Changed("assignee-any") {
		req.QueryParams["assignee.any"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.assigneeAny)
	}
	if cmd.Flags().Changed("assignee-not") {
		req.QueryParams["assignee.not"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.assigneeNot)
	}
	if cmd.Flags().Changed("portfolios-any") {
		req.QueryParams["portfolios.any"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.portfoliosAny)
	}
	if cmd.Flags().Changed("projects-any") {
		req.QueryParams["projects.any"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.projectsAny)
	}
	if cmd.Flags().Changed("projects-not") {
		req.QueryParams["projects.not"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.projectsNot)
	}
	if cmd.Flags().Changed("projects-all") {
		req.QueryParams["projects.all"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.projectsAll)
	}
	if cmd.Flags().Changed("sections-any") {
		req.QueryParams["sections.any"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.sectionsAny)
	}
	if cmd.Flags().Changed("sections-not") {
		req.QueryParams["sections.not"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.sectionsNot)
	}
	if cmd.Flags().Changed("sections-all") {
		req.QueryParams["sections.all"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.sectionsAll)
	}
	if cmd.Flags().Changed("tags-any") {
		req.QueryParams["tags.any"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.tagsAny)
	}
	if cmd.Flags().Changed("tags-not") {
		req.QueryParams["tags.not"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.tagsNot)
	}
	if cmd.Flags().Changed("tags-all") {
		req.QueryParams["tags.all"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.tagsAll)
	}
	if cmd.Flags().Changed("teams-any") {
		req.QueryParams["teams.any"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.teamsAny)
	}
	if cmd.Flags().Changed("followers-not") {
		req.QueryParams["followers.not"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.followersNot)
	}
	if cmd.Flags().Changed("created-by-any") {
		req.QueryParams["created_by.any"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.createdByAny)
	}
	if cmd.Flags().Changed("created-by-not") {
		req.QueryParams["created_by.not"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.createdByNot)
	}
	if cmd.Flags().Changed("assigned-by-any") {
		req.QueryParams["assigned_by.any"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.assignedByAny)
	}
	if cmd.Flags().Changed("assigned-by-not") {
		req.QueryParams["assigned_by.not"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.assignedByNot)
	}
	if cmd.Flags().Changed("liked-by-not") {
		req.QueryParams["liked_by.not"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.likedByNot)
	}
	if cmd.Flags().Changed("commented-on-by-not") {
		req.QueryParams["commented_on_by.not"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.commentedOnByNot)
	}
	if cmd.Flags().Changed("due-on-before") {
		req.QueryParams["due_on.before"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.dueOnBefore)
	}
	if cmd.Flags().Changed("due-on-after") {
		req.QueryParams["due_on.after"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.dueOnAfter)
	}
	if cmd.Flags().Changed("due-on") {
		req.QueryParams["due_on"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.dueOn)
	}
	if cmd.Flags().Changed("due-at-before") {
		req.QueryParams["due_at.before"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.dueAtBefore)
	}
	if cmd.Flags().Changed("due-at-after") {
		req.QueryParams["due_at.after"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.dueAtAfter)
	}
	if cmd.Flags().Changed("start-on-before") {
		req.QueryParams["start_on.before"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.startOnBefore)
	}
	if cmd.Flags().Changed("start-on-after") {
		req.QueryParams["start_on.after"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.startOnAfter)
	}
	if cmd.Flags().Changed("start-on") {
		req.QueryParams["start_on"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.startOn)
	}
	if cmd.Flags().Changed("created-on-before") {
		req.QueryParams["created_on.before"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.createdOnBefore)
	}
	if cmd.Flags().Changed("created-on-after") {
		req.QueryParams["created_on.after"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.createdOnAfter)
	}
	if cmd.Flags().Changed("created-on") {
		req.QueryParams["created_on"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.createdOn)
	}
	if cmd.Flags().Changed("created-at-before") {
		req.QueryParams["created_at.before"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.createdAtBefore)
	}
	if cmd.Flags().Changed("created-at-after") {
		req.QueryParams["created_at.after"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.createdAtAfter)
	}
	if cmd.Flags().Changed("completed-on-before") {
		req.QueryParams["completed_on.before"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.completedOnBefore)
	}
	if cmd.Flags().Changed("completed-on-after") {
		req.QueryParams["completed_on.after"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.completedOnAfter)
	}
	if cmd.Flags().Changed("completed-on") {
		req.QueryParams["completed_on"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.completedOn)
	}
	if cmd.Flags().Changed("completed-at-before") {
		req.QueryParams["completed_at.before"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.completedAtBefore)
	}
	if cmd.Flags().Changed("completed-at-after") {
		req.QueryParams["completed_at.after"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.completedAtAfter)
	}
	if cmd.Flags().Changed("modified-on-before") {
		req.QueryParams["modified_on.before"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.modifiedOnBefore)
	}
	if cmd.Flags().Changed("modified-on-after") {
		req.QueryParams["modified_on.after"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.modifiedOnAfter)
	}
	if cmd.Flags().Changed("modified-on") {
		req.QueryParams["modified_on"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.modifiedOn)
	}
	if cmd.Flags().Changed("modified-at-before") {
		req.QueryParams["modified_at.before"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.modifiedAtBefore)
	}
	if cmd.Flags().Changed("modified-at-after") {
		req.QueryParams["modified_at.after"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.modifiedAtAfter)
	}
	if cmd.Flags().Changed("is-blocking") {
		req.QueryParams["is_blocking"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.isBlocking)
	}
	if cmd.Flags().Changed("is-blocked") {
		req.QueryParams["is_blocked"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.isBlocked)
	}
	if cmd.Flags().Changed("has-attachment") {
		req.QueryParams["has_attachment"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.hasAttachment)
	}
	if cmd.Flags().Changed("completed") {
		req.QueryParams["completed"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.completed)
	}
	if cmd.Flags().Changed("is-subtask") {
		req.QueryParams["is_subtask"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.isSubtask)
	}
	if cmd.Flags().Changed("sort-by") {
		req.QueryParams["sort_by"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.sortBy)
	}
	if cmd.Flags().Changed("sort-ascending") {
		req.QueryParams["sort_ascending"] = fmt.Sprintf("%v", tasksSearchInWorkspaceFlags.sortAscending)
	}
	if cmd.Flags().Changed("opt-fields") {
		req.ArrayParams["opt_fields"] = tasksSearchInWorkspaceFlags.optFields
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
