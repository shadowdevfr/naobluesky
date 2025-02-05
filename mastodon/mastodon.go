/*

Mastodon est un paquet pour poster sur Mastodon.
@author Maxim Lucas (maximlucas.dev)

*/

package mastodon

import "time"

type Mastodon struct {
	Instance string // instance url
	Token    string // Token
}

type MastodonPost struct {
	ID                 string     `json:"id"`
	CreatedAt          time.Time  `json:"created_at"`
	InReplyToID        *string    `json:"in_reply_to_id"`
	InReplyToAccountID *string    `json:"in_reply_to_account_id"`
	Sensitive          bool       `json:"sensitive"`
	SpoilerText        string     `json:"spoiler_text"`
	Visibility         string     `json:"visibility"`
	Language           string     `json:"language"`
	URI                string     `json:"uri"`
	URL                string     `json:"url"`
	RepliesCount       int        `json:"replies_count"`
	ReblogsCount       int        `json:"reblogs_count"`
	FavouritesCount    int        `json:"favourites_count"`
	EditedAt           *time.Time `json:"edited_at"`
	Favourited         bool       `json:"favourited"`
	Reblogged          bool       `json:"reblogged"`
	Muted              bool       `json:"muted"`
	Bookmarked         bool       `json:"bookmarked"`
	Pinned             bool       `json:"pinned"`
	Content            string     `json:"content"`
	Application        struct {
		Name    string  `json:"name"`
		Website *string `json:"website"`
	} `json:"application"`
	Account struct {
		ID              string    `json:"id"`
		Username        string    `json:"username"`
		Acct            string    `json:"acct"`
		DisplayName     string    `json:"display_name"`
		Locked          bool      `json:"locked"`
		Bot             bool      `json:"bot"`
		Discoverable    bool      `json:"discoverable"`
		Indexable       bool      `json:"indexable"`
		Group           bool      `json:"group"`
		CreatedAt       time.Time `json:"created_at"`
		Note            string    `json:"note"`
		URL             string    `json:"url"`
		URI             string    `json:"uri"`
		Avatar          string    `json:"avatar"`
		AvatarStatic    string    `json:"avatar_static"`
		Header          string    `json:"header"`
		HeaderStatic    string    `json:"header_static"`
		FollowersCount  int       `json:"followers_count"`
		FollowingCount  int       `json:"following_count"`
		StatusesCount   int       `json:"statuses_count"`
		LastStatusAt    string    `json:"last_status_at"`
		HideCollections *bool     `json:"hide_collections"`
		NoIndex         bool      `json:"noindex"`
		Emojis          []string  `json:"emojis"`
		Roles           []string  `json:"roles"`
		Fields          []struct {
			Name       string     `json:"name"`
			Value      string     `json:"value"`
			VerifiedAt *time.Time `json:"verified_at"`
		} `json:"fields"`
	} `json:"account"`
	MediaAttachments []string `json:"media_attachments"`
	Mentions         []string `json:"mentions"`
	Tags             []string `json:"tags"`
	Emojis           []string `json:"emojis"`
	Card             *string  `json:"card"`
	Poll             *string  `json:"poll"`
}
