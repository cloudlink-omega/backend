package structs

// JSON structure for creating/updating a save slot.
type Save struct {
	UGI      string `json:"ugi" validate:"required" label:"ugi"`
	Token    string `json:"token" validate:"required" label:"token"`
	SaveSlot uint8  `json:"save_slot" validate:"required,min=1,max=10" label:"save_slot"`
	SaveData any    `json:"save_data" validate:"required,max=10000" label:"save_data"`
}

// JSON structure for loading a save slot.
type Load struct {
	UGI      string `json:"ugi" validate:"required" label:"ugi"`
	Token    string `json:"token" validate:"required" label:"token"`
	SaveSlot uint8  `json:"save_slot" validate:"required,min=1,max=10" label:"save_slot"`
}
