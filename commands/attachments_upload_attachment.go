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

var attachmentsUploadAttachmentCmd = &cobra.Command{
	Use:   "upload-attachment",
	Short: "Upload an attachment",
	RunE:  runAttachmentsUploadAttachment,
}

var attachmentsUploadAttachmentFlags struct {
	optPretty       bool
	optFields       []string
	resourceSubtype string
	file            string
	parent          string
	url             string
	name            string
	connectToApp    bool
	body            string
}

func init() {
	attachmentsUploadAttachmentCmd.Flags().BoolVar(&attachmentsUploadAttachmentFlags.optPretty, "opt-pretty", false, "Provides “pretty” output. Provides the response in a “pretty” format. In the case of JSON this means doing proper line breaking and indentation to make it readable. This will take extra time and increase the response size so it is advisable only to use this during debugging.")
	attachmentsUploadAttachmentCmd.Flags().StringSliceVar(&attachmentsUploadAttachmentFlags.optFields, "opt-fields", nil, "This endpoint returns a compact resource, which excludes some properties by default. To include those optional properties, set this query parameter to a comma-separated list of the properties you wish to include.")
	attachmentsUploadAttachmentCmd.Flags().StringVar(&attachmentsUploadAttachmentFlags.resourceSubtype, "resource-subtype", "", "The type of the attachment. Must be one of the given values. If not specified, a file attachment of type `asana` will be assumed. Note that if the value of `resource_subtype` is `external`, a `parent`, `name`, and `url` must also be provided. ")
	// Note: body fields are not MarkFlagRequired — --body JSON satisfies them too.
	attachmentsUploadAttachmentCmd.Flags().StringVar(&attachmentsUploadAttachmentFlags.file, "file", "", "Required for `asana` attachments. ")
	// Note: body fields are not MarkFlagRequired — --body JSON satisfies them too.
	attachmentsUploadAttachmentCmd.Flags().StringVar(&attachmentsUploadAttachmentFlags.parent, "parent", "", "Required identifier of the parent task, project, or project_brief, as a string. ")
	// Note: body fields are not MarkFlagRequired — --body JSON satisfies them too.
	attachmentsUploadAttachmentCmd.Flags().StringVar(&attachmentsUploadAttachmentFlags.url, "url", "", "The URL of the external resource being attached. Required for attachments of type `external`. ")
	// Note: body fields are not MarkFlagRequired — --body JSON satisfies them too.
	attachmentsUploadAttachmentCmd.Flags().StringVar(&attachmentsUploadAttachmentFlags.name, "name", "", "The name of the external resource being attached. Required for attachments of type `external`. ")
	// Note: body fields are not MarkFlagRequired — --body JSON satisfies them too.
	attachmentsUploadAttachmentCmd.Flags().BoolVar(&attachmentsUploadAttachmentFlags.connectToApp, "connect-to-app", false, "*Optional*. Only relevant for external attachments with a parent task. A boolean indicating whether the current app should be connected with the attachment for the purposes of showing an app components widget. Requires the app to have been added to a project the parent task is in. ")
	// Note: body fields are not MarkFlagRequired — --body JSON satisfies them too.
	attachmentsUploadAttachmentCmd.Flags().StringVar(&attachmentsUploadAttachmentFlags.body, "body", "", "Full request body as JSON (overrides individual flags)")

	attachmentsCmd.AddCommand(attachmentsUploadAttachmentCmd)
}

func runAttachmentsUploadAttachment(cmd *cobra.Command, args []string) error {
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
			Name:        "resource-subtype",
			Type:        "string",
			Required:    false,
			Location:    "body",
			Description: "The type of the attachment. Must be one of the given values. If not specified, a file attachment of type `asana` will be assumed. Note that if the value of `resource_subtype` is `external`, a `parent`, `name`, and `url` must also be provided. ",
		})
		flags = append(flags, flagSchema{
			Name:        "file",
			Type:        "string",
			Required:    false,
			Location:    "body",
			Description: "Required for `asana` attachments. ",
		})
		flags = append(flags, flagSchema{
			Name:        "parent",
			Type:        "string",
			Required:    false,
			Location:    "body",
			Description: "Required identifier of the parent task, project, or project_brief, as a string. ",
		})
		flags = append(flags, flagSchema{
			Name:        "url",
			Type:        "string",
			Required:    false,
			Location:    "body",
			Description: "The URL of the external resource being attached. Required for attachments of type `external`. ",
		})
		flags = append(flags, flagSchema{
			Name:        "name",
			Type:        "string",
			Required:    false,
			Location:    "body",
			Description: "The name of the external resource being attached. Required for attachments of type `external`. ",
		})
		flags = append(flags, flagSchema{
			Name:        "connect-to-app",
			Type:        "boolean",
			Required:    false,
			Location:    "body",
			Description: "*Optional*. Only relevant for external attachments with a parent task. A boolean indicating whether the current app should be connected with the attachment for the purposes of showing an app components widget. Requires the app to have been added to a project the parent task is in. ",
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
			Description: "Successfully uploaded the attachment to the parent object.",
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
			"command":     "upload-attachment",
			"description": "Upload an attachment",
			"http": map[string]any{
				"method": "POST",
				"path":   "/attachments",
			},
			"input": map[string]any{
				"flags":         flags,
				"body_flag":     true,
				"body_required": false,
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
		Path:        httpclient.SubstitutePath("/attachments", pathParams),
		QueryParams: map[string]string{},
		ArrayParams: map[string][]string{},
		Headers:     map[string]string{},
	}

	// Query parameters
	if cmd.Flags().Changed("opt-pretty") {
		req.QueryParams["opt_pretty"] = fmt.Sprintf("%v", attachmentsUploadAttachmentFlags.optPretty)
	}
	if cmd.Flags().Changed("opt-fields") {
		req.ArrayParams["opt_fields"] = attachmentsUploadAttachmentFlags.optFields
	}

	// Header parameters

	// Request body
	bodyMap := map[string]any{}
	if attachmentsUploadAttachmentFlags.body != "" {
		if err := json.Unmarshal([]byte(attachmentsUploadAttachmentFlags.body), &bodyMap); err != nil {
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
	if cmd.Flags().Changed("resource-subtype") {
		bodyMap["resource_subtype"] = attachmentsUploadAttachmentFlags.resourceSubtype
	}
	if cmd.Flags().Changed("file") {
		bodyMap["file"] = attachmentsUploadAttachmentFlags.file
	}
	if cmd.Flags().Changed("parent") {
		bodyMap["parent"] = attachmentsUploadAttachmentFlags.parent
	}
	if cmd.Flags().Changed("url") {
		bodyMap["url"] = attachmentsUploadAttachmentFlags.url
	}
	if cmd.Flags().Changed("name") {
		bodyMap["name"] = attachmentsUploadAttachmentFlags.name
	}
	if cmd.Flags().Changed("connect-to-app") {
		bodyMap["connect_to_app"] = attachmentsUploadAttachmentFlags.connectToApp
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
