package bluesky

type createSessionResponse struct {
	AccessJwt       string `json:"accessJwt"`
	Active          bool   `json:"active"`
	Did             string `json:"did"`
	DidDoc          didDoc `json:"didDoc"`
	Email           string `json:"email"`
	EmailAuthFactor bool   `json:"emailAuthFactor"`
	EmailConfirmed  bool   `json:"emailConfirmed"`
	Handle          string `json:"handle"`
	RefreshJwt      string `json:"refreshJwt"`
}

type didDoc struct {
	Context            []string             `json:"@context"`
	AlsoKnownAs        []string             `json:"alsoKnownAs"`
	Id                 string               `json:"id"`
	Service            []service            `json:"service"`
	VerificationMethod []verificationMethod `json:"verificationMethod"`
}

type service struct {
	Id              string `json:"id"`
	ServiceEndpoint string `json:"serviceEndpoint"`
	Type            string `json:"type"`
}

type verificationMethod struct {
	Controller         string `json:"controller"`
	Id                 string `json:"id"`
	PublicKeyMultibase string `json:"publicKeyMultibase"`
	Type               string `json:"type"`
}

type createSessionRequest struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type createRecordRequest struct {
	Repo       string `json:"repo"`
	Collection string `json:"collection"`
	Record     record `json:"record"`
}

type record struct {
	Type      string `json:"$type"`
	CreatedAt string `json:"createdAt"`
	Text      string `json:"text"`
}
