package declarations

// Declarations model only gets the ARCHIVO attr which contain the filename for each document
type Declarations struct{
	_id		string `bson:"_id,omitempty"`
	ARCHIVO string `bson:"ARCHIVO"`
}