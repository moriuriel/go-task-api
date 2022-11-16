package domain

type (
	Owner struct {
		id   ID
		name string
	}

	TaskOwnerOutput struct {
		Id   string `json:"_id"`
		Name string `json:"name"`
	}
)

func NewOwner(id ID, name string) *Owner {
	return &Owner{
		id:   id,
		name: name,
	}
}

func (o Owner) ID() ID {
	return o.id
}

func (o Owner) Name() string {
	return o.name
}
