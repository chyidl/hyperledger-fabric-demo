package main

// Owmer contains the full name of an owner of a basic asset (资产)
type Owner struct {
	Forename string `json:"forename"`
	Surname  string `json:"surname"`
}

// BasicAsset a basic asset
type BasicAsset struct {
	ID        string `json:"id"`
	Owner     Owner  `json:"owner"`
	Value     int    `json:"value"`
	Condition int    `json:"condition"`
}

// SetConditionNew set the condition of a assert to mark as new
func (ba *BasicAsset) SetConditionNew() {
	ba.Condition = 0
}

// SetConditionUsed set the condition of the assert to mark as used
func (ba *BasicAsset) SetConditionUsed() {
	ba.Condition = 1
}
