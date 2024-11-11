package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// WorkflowJobTemplateNodeService implements awx job template node apis.
type WorkflowJobTemplateNodeService struct {
	client *Client
}

// ListWorkflowJobTemplateNodesResponse represents `ListWorkflowJobTemplateNodes` endpoint response.
type ListWorkflowJobTemplateNodesResponse struct {
	Pagination
	Results []*WorkflowJobTemplateNode `json:"results"`
}

const workflowJobTemplateNodeAPIEndpoint = "/api/v2/workflow_job_template_nodes/"

// GetWorkflowJobTemplateNodeByID shows the details of a job template node.
func (jt *WorkflowJobTemplateNodeService) GetWorkflowJobTemplateNodeByID(id int, params map[string]string) (*WorkflowJobTemplateNode, error) {
	result := new(WorkflowJobTemplateNode)
	endpoint := fmt.Sprintf("%s%d/", workflowJobTemplateNodeAPIEndpoint, id)
	resp, err := jt.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// ListWorkflowJobTemplateNodes shows a list of job templates nodes.
func (jt *WorkflowJobTemplateNodeService) ListWorkflowJobTemplateNodes(params map[string]string) ([]*WorkflowJobTemplateNode, *ListWorkflowJobTemplateNodesResponse, error) {
	result := new(ListWorkflowJobTemplateNodesResponse)

	resp, err := jt.client.Requester.GetJSON(workflowJobTemplateNodeAPIEndpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

// CreateWorkflowJobTemplateNode creates a job template node, without any pe exisiting nodes.
func (jt *WorkflowJobTemplateNodeService) CreateWorkflowJobTemplateNode(data map[string]interface{}, params map[string]string) (*WorkflowJobTemplateNode, error) {
	result := new(WorkflowJobTemplateNode)
	mandatoryFields = []string{"workflow_job_template", "unified_job_template", "identifier"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	resp, err := jt.client.Requester.PostJSON(workflowJobTemplateNodeAPIEndpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}
	if err := CheckResponse(resp); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateWorkflowJobTemplateNode updates a job template node.
func (jt *WorkflowJobTemplateNodeService) UpdateWorkflowJobTemplateNode(id int, data map[string]interface{}, params map[string]string) (*WorkflowJobTemplateNode, error) {
	result := new(WorkflowJobTemplateNode)
	endpoint := fmt.Sprintf("%s%d", workflowJobTemplateNodeAPIEndpoint, id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := jt.client.Requester.PatchJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}
	if err := CheckResponse(resp); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteWorkflowJobTemplateNode deletes a job template node.
func (jt *WorkflowJobTemplateNodeService) DeleteWorkflowJobTemplateNode(id int) (*WorkflowJobTemplateNode, error) {
	result := new(WorkflowJobTemplateNode)
	endpoint := fmt.Sprintf("%s%d", workflowJobTemplateNodeAPIEndpoint, id)

	resp, err := jt.client.Requester.Delete(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}

// AssociateNodeRelationship creates a relationship between nodes of the specified type
func (jt *WorkflowJobTemplateNodeService) AssociateNodeRelationship(sourceNodeID int, targetNodeID int, relationType string) error {
    endpoint := fmt.Sprintf("%s%d/%s/", workflowJobTemplateNodeAPIEndpoint, sourceNodeID, relationType)
    payload, err := json.Marshal(map[string]int{"id": targetNodeID})
    if err != nil {
        return err
    }

    resp, err := jt.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), nil, nil)
    if err != nil {
        return err
    }

    return CheckResponse(resp)
}

func (jt *WorkflowJobTemplateNodeService) DisassociateNodeRelationship(sourceNodeID int, targetNodeID int, relationType string) error {
    endpoint := fmt.Sprintf("%s%d/%s/", workflowJobTemplateNodeAPIEndpoint, sourceNodeID, relationType)
    payload, err := json.Marshal(map[string]interface{}{
        "id":           targetNodeID,
        "disassociate": true,
    })
    if err != nil {
        return err
    }

    resp, err := jt.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), nil, nil)
    if err != nil {
        return err
    }

    return CheckResponse(resp)
}

// GetNodeRelationships gets all related nodes of a specific type for a given node
func (jt *WorkflowJobTemplateNodeService) GetNodeRelationships(nodeID int, relationType string) ([]*WorkflowJobTemplateNode, error) {
    endpoint := fmt.Sprintf("%s%d/%s/", workflowJobTemplateNodeAPIEndpoint, nodeID, relationType)
    result := new(ListWorkflowJobTemplateNodesResponse)
    
    resp, err := jt.client.Requester.GetJSON(endpoint, result, nil)
    if err != nil {
        return nil, err
    }

    if err := CheckResponse(resp); err != nil {
        return nil, err
    }

    return result.Results, nil
}
