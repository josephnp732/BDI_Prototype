package store

// Structure to store the JSON structure
type Plan struct {
	PlanCostShares     PlanCostShares      `json:"planCostShares"`
	LinkedPlanServices []LinkedPlanService `json:"linkedPlanServices"`
	Org                string              `json:"_org"`
	ObjectID           string              `json:"objectId"`
	ObjectType         string              `json:"objectType"`
	PlanType           string              `json:"planType"`
	CreationDate       string              `json:"creationDate"`
}

type LinkedPlanService struct {
	LinkedService         LinkedService  `json:"linkedService"`
	PlanserviceCostShares PlanCostShares `json:"planserviceCostShares"`
	Org                   string         `json:"_org"`
	ObjectID              string         `json:"objectId"`
	ObjectType            string         `json:"objectType"`
}

type LinkedService struct {
	Org        string `json:"_org"`
	ObjectID   string `json:"objectId"`
	ObjectType string `json:"objectType"`
	Name       string `json:"name"`
}

type PlanCostShares struct {
	Deductible int64  `json:"deductible"`
	Org        string `json:"_org"`
	Copay      int64  `json:"copay"`
	ObjectID   string `json:"objectId"`
	ObjectType string `json:"objectType"`
}
