package main

import "time"

type Tweet struct {
	ID               string `json:"id"`
	URL              string `json:"url"`
	Text             string `json:"text"`
	CreatedAt        string `json:"createdAt"`
	AuthorProfilePic string `json:"author.profilePicture"`
	RetweetCount     int    `json:"retweetCount"`
	ReplyCount       int    `json:"replyCount"`
	LikeCount        int    `json:"likeCount"`
	QuoteCount       int    `json:"quoteCount"`
	ViewCount        int    `json:"viewCount"`
	BookmarkCount    int    `json:"bookmarkCount"`
	Source           string `json:"source"`
	Lang             string `json:"lang"`
	IsReply          bool   `json:"isReply"`
	IsPinned         bool   `json:"isPinned"`
}

type ApifyRequest struct {
	FilterBlueVerified    bool   `json:"filter:blue_verified"`
	FilterConsumerVideo   bool   `json:"filter:consumer_video"`
	FilterHasEngagement   bool   `json:"filter:has_engagement"`
	FilterHashtags        bool   `json:"filter:hashtags"`
	FilterImages          bool   `json:"filter:images"`
	FilterLinks           bool   `json:"filter:links"`
	FilterMedia           bool   `json:"filter:media"`
	FilterMentions        bool   `json:"filter:mentions"`
	FilterNativeVideo     bool   `json:"filter:native_video"`
	FilterNativeRetweets  bool   `json:"filter:nativeretweets"`
	FilterNews            bool   `json:"filter:news"`
	FilterProVideo        bool   `json:"filter:pro_video"`
	FilterQuote           bool   `json:"filter:quote"`
	FilterReplies         bool   `json:"filter:replies"`
	FilterSafe            bool   `json:"filter:safe"`
	FilterSpaces          bool   `json:"filter:spaces"`
	FilterTwimg           bool   `json:"filter:twimg"`
	FilterVerified        bool   `json:"filter:verified"`
	FilterVideos          bool   `json:"filter:videos"`
	FilterVine            bool   `json:"filter:vine"`
	From                  string `json:"from"`
	IncludeNativeRetweets bool   `json:"include:nativeretweets"`
	MaxItems              int    `json:"maxItems"`
	Since                 string `json:"since"`
	Until                 string `json:"until"`
	QueryType             string `json:"queryType"`
	MinRetweets           int    `json:"min_retweets"`
	MinFaves              int    `json:"min_faves"`
	MinReplies            int    `json:"min_replies"`
	NegMinRetweets        int    `json:"-min_retweets"`
	NegMinFaves           int    `json:"-min_faves"`
	NegMinReplies         int    `json:"-min_replies"`
}

type ApifyResponse struct {
	Data struct {
		ID         string     `json:"id"`
		ActID      string     `json:"actId"`
		UserID     string     `json:"userId"`
		StartedAt  time.Time  `json:"startedAt"`
		FinishedAt *time.Time `json:"finishedAt"`
		Status     string     `json:"status"`
		Meta       struct {
			Origin    string `json:"origin"`
			UserAgent string `json:"userAgent"`
		} `json:"meta"`
		Stats struct {
			InputBodyLen   int `json:"inputBodyLen"`
			RebootCount    int `json:"rebootCount"`
			RestartCount   int `json:"restartCount"`
			ResurrectCount int `json:"resurrectCount"`
			ComputeUnits   int `json:"computeUnits"`
		} `json:"stats"`
		Options struct {
			Build        string `json:"build"`
			TimeoutSecs  int    `json:"timeoutSecs"`
			MemoryMbytes int    `json:"memoryMbytes"`
			MaxItems     int    `json:"maxItems"`
			DiskMbytes   int    `json:"diskMbytes"`
		} `json:"options"`
		BuildID                string `json:"buildId"`
		DefaultKeyValueStoreID string `json:"defaultKeyValueStoreId"`
		DefaultDatasetID       string `json:"defaultDatasetId"`
		DefaultRequestQueueID  string `json:"defaultRequestQueueId"`
		PricingInfo            struct {
			PricingModel          string    `json:"pricingModel"`
			ReasonForChange       string    `json:"reasonForChange"`
			PricePerUnitUsd       float64   `json:"pricePerUnitUsd"`
			UnitName              string    `json:"unitName"`
			CreatedAt             time.Time `json:"createdAt"`
			StartedAt             time.Time `json:"startedAt"`
			ApifyMarginPercentage float64   `json:"apifyMarginPercentage"`
			NotifiedAboutChangeAt time.Time `json:"notifiedAboutChangeAt"`
		} `json:"pricingInfo"`
		BuildNumber  string `json:"buildNumber"`
		ContainerURL string `json:"containerUrl"`
		Usage        struct {
			ActorComputeUnits              int `json:"ACTOR_COMPUTE_UNITS"`
			DatasetReads                   int `json:"DATASET_READS"`
			DatasetWrites                  int `json:"DATASET_WRITES"`
			KeyValueStoreReads             int `json:"KEY_VALUE_STORE_READS"`
			KeyValueStoreWrites            int `json:"KEY_VALUE_STORE_WRITES"`
			KeyValueStoreLists             int `json:"KEY_VALUE_STORE_LISTS"`
			RequestQueueReads              int `json:"REQUEST_QUEUE_READS"`
			RequestQueueWrites             int `json:"REQUEST_QUEUE_WRITES"`
			DataTransferInternalGbytes     int `json:"DATA_TRANSFER_INTERNAL_GBYTES"`
			DataTransferExternalGbytes     int `json:"DATA_TRANSFER_EXTERNAL_GBYTES"`
			ProxyResidentialTransferGbytes int `json:"PROXY_RESIDENTIAL_TRANSFER_GBYTES"`
			ProxySerps                     int `json:"PROXY_SERPS"`
		} `json:"usage"`
		UsageTotalUsd float64 `json:"usageTotalUsd"`
		UsageUsd      struct {
			ActorComputeUnits              float64 `json:"ACTOR_COMPUTE_UNITS"`
			DatasetReads                   float64 `json:"DATASET_READS"`
			DatasetWrites                  float64 `json:"DATASET_WRITES"`
			KeyValueStoreReads             float64 `json:"KEY_VALUE_STORE_READS"`
			KeyValueStoreWrites            float64 `json:"KEY_VALUE_STORE_WRITES"`
			KeyValueStoreLists             float64 `json:"KEY_VALUE_STORE_LISTS"`
			RequestQueueReads              float64 `json:"REQUEST_QUEUE_READS"`
			RequestQueueWrites             float64 `json:"REQUEST_QUEUE_WRITES"`
			DataTransferInternalGbytes     float64 `json:"DATA_TRANSFER_INTERNAL_GBYTES"`
			DataTransferExternalGbytes     float64 `json:"DATA_TRANSFER_EXTERNAL_GBYTES"`
			ProxyResidentialTransferGbytes float64 `json:"PROXY_RESIDENTIAL_TRANSFER_GBYTES"`
			ProxySerps                     float64 `json:"PROXY_SERPS"`
		} `json:"usageUsd"`
	} `json:"data"`
}
