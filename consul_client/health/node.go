package health

type NodeStatus struct {
	ID          string        `json:"ID"`
	Node        string        `json:"Node"`
	CheckID     string        `json:"CheckID"`
	Name        string        `json:"Name"`
	Status      string        `json:"Status"`
	Notes       string        `json:"Notes"`
	Output      string        `json:"Output"`
	ServiceID   string        `json:"ServiceID"`
	ServiceName string        `json:"ServiceName"`
	ServiceTags []interface{} `json:"ServiceTags"`
	Namespace   string        `json:"Namespace"`
}
