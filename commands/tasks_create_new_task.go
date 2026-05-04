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

var tasksCreateNewTaskCmd = &cobra.Command{
	Use:   "create-new-task",
	Short: "Create a task",
	RunE:  runTasksCreateNewTask,
}

var tasksCreateNewTaskFlags struct {
	optPretty bool
	optFields []string
	data      string
	body      string
}

func init() {
	tasksCreateNewTaskCmd.Flags().BoolVar(&tasksCreateNewTaskFlags.optPretty, "opt-pretty", false, "Provides “pretty” output. Provides the response in a “pretty” format. In the case of JSON this means doing proper line breaking and indentation to make it readable. This will take extra time and increase the response size so it is advisable only to use this during debugging.")
	tasksCreateNewTaskCmd.Flags().StringSliceVar(&tasksCreateNewTaskFlags.optFields, "opt-fields", nil, "This endpoint returns a compact resource, which excludes some properties by default. To include those optional properties, set this query parameter to a comma-separated list of the properties you wish to include.")
	tasksCreateNewTaskCmd.Flags().StringVar(&tasksCreateNewTaskFlags.data, "data", "", "")
	// Note: body fields are not MarkFlagRequired — --body JSON satisfies them too.
	tasksCreateNewTaskCmd.Flags().StringVar(&tasksCreateNewTaskFlags.body, "body", "", "Full request body as JSON (overrides individual flags)")

	tasksCmd.AddCommand(tasksCreateNewTaskCmd)
}

func runTasksCreateNewTask(cmd *cobra.Command, args []string) error {
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
			Name:        "opt-fields",
			Type:        "array",
			Required:    false,
			Location:    "query",
			Description: "This endpoint returns a compact resource, which excludes some properties by default. To include those optional properties, set this query parameter to a comma-separated list of the properties you wish to include.",
		})
		flags = append(flags, flagSchema{
			Name:        "data",
			Type:        "string",
			Required:    false,
			Location:    "body",
			Description: "",
		})

		type responseSchema struct {
			Status      string `json:"status"`
			ContentType string `json:"content_type,omitempty"`
			Description string `json:"description,omitempty"`
		}
		var responses []responseSchema
		responses = append(responses, responseSchema{
			Status:      "201",
			ContentType: "application/json",
			Description: "Successfully created a new task.",
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
			"command":     "create-new-task",
			"description": "Create a task",
			"http": map[string]any{
				"method": "POST",
				"path":   "/tasks",
			},
			"input": map[string]any{
				"flags":         flags,
				"body_flag":     true,
				"body_required": true,
			},
			"output": map[string]any{
				"responses": responses,
			},
			"semantics": map[string]any{
				"safe":         false,
				"idempotent":   false,
				"reversible":   true,
				"side_effects": []string{"creates_resource"},
				"impact":       "medium",
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
		Method:      "POST",
		Path:        httpclient.SubstitutePath("/tasks", pathParams),
		QueryParams: map[string]string{},
		ArrayParams: map[string][]string{},
		Headers:     map[string]string{},
	}

	// Query parameters
	if cmd.Flags().Changed("opt-pretty") {
		req.QueryParams["opt_pretty"] = fmt.Sprintf("%v", tasksCreateNewTaskFlags.optPretty)
	}
	if cmd.Flags().Changed("opt-fields") {
		req.ArrayParams["opt_fields"] = tasksCreateNewTaskFlags.optFields
	}

	// Header parameters

	// Request body
	bodyMap := map[string]any{}
	if tasksCreateNewTaskFlags.body != "" {
		if err := json.Unmarshal([]byte(tasksCreateNewTaskFlags.body), &bodyMap); err != nil {
			_invState.errorType = "parse_error"
			cliErr := &output.CLIError{
				Error:    true,
				Code:     "validation_error",
				Message:  fmt.Sprintf("invalid JSON in --body: %v", err),
				ExitCode: output.ExitValidation,
			}
			cliErr.Write(os.Stderr)
			return output.NewExitError(cliErr)
		}
	}
	// Individual flags overlay onto body (flags take precedence over --body JSON)
	if cmd.Flags().Changed("data") {
		bodyMap["data"] = tasksCreateNewTaskFlags.data
	}
	req.Body = bodyMap

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
