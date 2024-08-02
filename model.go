package main

type QuestResponse struct {
	Data Data `json:"data"`
}

type Data struct {
	Quests []Quest `json:"quests"`
}

type Quest struct {
	Article  Article   `json:"article"`
	Comments []Comment `json:"comments"`
}

type Article struct {
	BookmarkCount       int           `json:"bookmark_count"`
	BookmarkedAt        interface{}   `json:"bookmarked_at"`
	Card                interface{}   `json:"card"`
	Channel             interface{}   `json:"channel"`
	CommentCount        int           `json:"comment_count"`
	Content             string        `json:"content"`
	ContentTranslation  interface{}   `json:"content_translation"`
	CreateAt            float64       `json:"create_at"`
	Creator             Creator       `json:"creator"`
	CreatorID           string        `json:"creator_id"`
	DepositDeadline     int64         `json:"deposit_deadline"`
	DistrustCount       int           `json:"distrust_count"`
	Draw                interface{}   `json:"draw"`
	DrawID              interface{}   `json:"draw_id"`
	Entities            Entities      `json:"entities"`
	EntitiesTranslation interface{}   `json:"entities_translation"`
	HasPermission       bool          `json:"has_permission"`
	ID                  int           `json:"id"`
	Images              []interface{} `json:"images"`
	IsBookmarked        bool          `json:"is_bookmarked"`
	IsOriginal          bool          `json:"is_original"`
	IsPaid              bool          `json:"is_paid"`
	IsPinned            bool          `json:"is_pinned"`
	IsRead              bool          `json:"is_read"`
	IsReposted          bool          `json:"is_reposted"`
	IsRewardSettled     bool          `json:"is_reward_settled"`
	IsTrust             interface{}   `json:"is_trust"`
	IsVisible           bool          `json:"is_visible"`
	IsWithdrawn         bool          `json:"is_withdrawn"`
	NotVisibleReason    string        `json:"not_visible_reason"`
	OfficialList        []interface{} `json:"official_list"`
	PayCount            int           `json:"pay_count"`
	PayQA               interface{}   `json:"pay_qa"`
	Permissions         Permissions   `json:"permissions"`
	Poll                interface{}   `json:"poll"`
	PollID              interface{}   `json:"poll_id"`
	Price               interface{}   `json:"price"`
	Proposal            interface{}   `json:"proposal"`
	QualityDegree       interface{}   `json:"quality_degree"`
	Quest               QuestDetails  `json:"quest"`
	QuoteArticle        interface{}   `json:"quote_article"`
	QuoteCount          int           `json:"quote_count"`
	ReadCount           int           `json:"read_count"`
	RealizedValue       float64       `json:"realized_value"`
	RepostCount         int           `json:"repost_count"`
	RewardSettledAt     interface{}   `json:"reward_settled_at"`
	RewardUSDValue      float64       `json:"reward_usd_value"`
	Summary             string        `json:"summary"`
	SummaryTranslation  interface{}   `json:"summary_translation"`
	Template            interface{}   `json:"template"`
	ThreadCount         int           `json:"thread_count"`
	Threads             []interface{} `json:"threads"`
	TrustCount          int           `json:"trust_count"`
	TrustDegree         float64       `json:"trust_degree"`
	TVR                 float64       `json:"tvr"`
	Type                string        `json:"type"`
	Value               interface{}   `json:"value"`
	VRAvg               float64       `json:"vr_avg"`
	VRMedian            float64       `json:"vr_median"`
}

type Comment struct {
	// Assuming the structure of comments based on provided data
	// The structure can be added later if details are provided
}

type Creator struct {
	AvatarThumbnailURL  string      `json:"avatar_thumbnail_url"`
	AvatarURL           string      `json:"avatar_url"`
	BannerThumbnailURL  string      `json:"banner_thumbnail_url"`
	BannerURL           string      `json:"banner_url"`
	CreateAt            float64     `json:"create_at"`
	FollowerCount       int         `json:"follower_count"`
	ID                  int         `json:"id"`
	InitialPrice        int         `json:"initial_price"`
	Intro               string      `json:"intro"`
	IsFollowing         bool        `json:"is_following"`
	IsMuted             bool        `json:"is_muted"`
	IsPremium           bool        `json:"is_premium"`
	Name                string      `json:"name"`
	OfficialSiteID      interface{} `json:"official_site_id"`
	OfficialTwitterID   interface{} `json:"official_twitter_id"`
	RankAt              int         `json:"rank_at"`
	RepliedRate         float64     `json:"replied_rate"`
	SiteID              string      `json:"site_id"`
	Slug                string      `json:"slug"`
	TVF                 float64     `json:"tvf"`
	TwitterID           string      `json:"twitter_id"`
	TwitterIDVerified   bool        `json:"twitter_id_verified"`
	Type                string      `json:"type"`
	UnchargedOfferCount int         `json:"uncharged_offer_count"`
	UnchargedOfferValue int         `json:"uncharged_offer_value"`
	VerifyComment       interface{} `json:"verify_comment"`
	VerifyStatus        int         `json:"verify_status"`
}

type Entities struct {
	Channels []interface{} `json:"channels"`
	Mentions []interface{} `json:"mentions"`
	Tags     []interface{} `json:"tags"`
	URL      []URL         `json:"url"`
}

type URL struct {
	DisplayURL string `json:"display_url"`
	Indices    []int  `json:"indices"`
	RawString  string `json:"raw_string"`
	URL        string `json:"url"`
}

type Permissions struct {
	HasWeb3ID   bool `json:"has_web3_id"`
	MinNetWorth int  `json:"min_net_worth"`
}

type QuestDetails struct {
	Actions            []Action `json:"actions"`
	CreateAt           float64  `json:"create_at"`
	Draw               Draw     `json:"draw"`
	DrawID             int      `json:"draw_id"`
	Duration           int      `json:"duration"`
	EstimatedJoinCount int      `json:"estimated_join_count"`
	HasPermission      bool     `json:"has_permission"`
	ID                 int      `json:"id"`
	IsCompleted        bool     `json:"is_completed"`
	IsJoined           bool     `json:"is_joined"`
	IsPrivate          bool     `json:"is_private"`
	JoinCount          int      `json:"join_count"`
	Name               string   `json:"name"`
	Permissions        struct {
		HasWeb3ID bool `json:"has_web3_id"`
	} `json:"permissions"`
	StartAt int64  `json:"start_at"`
	Status  string `json:"status"`
	UnitXP  int    `json:"unit_xp"`
}

type Action struct {
	CreateAt float64 `json:"create_at"`
	Data     struct {
		ActionURL      string `json:"action_url"`
		TweetID        string `json:"tweet_id"`
		TwitterID      string `json:"twitter_id"`
		TwitterLink    string `json:"twitter_link"`
		DataURL        string `json:"data_url"`
		Desc           string `json:"desc"`
		RelatedLink    string `json:"related_link"`
		Title          string `json:"title"`
		ChainID        string `json:"chain_id"`
		ChainName      string `json:"chain_name"`
		CollectionID   string `json:"collection_id"`
		CollectionName string `json:"collection_name"`
		LogoURL        string `json:"logo_url"`
	} `json:"data"`
	ID         int    `json:"id"`
	IsVerified bool   `json:"is_verified"`
	QuestID    int    `json:"quest_id"`
	TaskID     int    `json:"task_id"`
	Type       string `json:"type"`
	UnitXP     int    `json:"unit_xp"`
}

type Draw struct {
	BlockID                 interface{} `json:"block_id"`
	CreateAt                float64     `json:"create_at"`
	Duration                int64       `json:"duration"`
	FinishAt                float64     `json:"finish_at"`
	HasPermission           bool        `json:"has_permission"`
	ID                      int         `json:"id"`
	IsSettled               bool        `json:"is_settled"`
	JoinCount               int         `json:"join_count"`
	Operations              interface{} `json:"operations"`
	Permissions             interface{} `json:"permissions"`
	PrizeCount              int         `json:"prize_count"`
	PrizeCustom             interface{} `json:"prize_custom"`
	PrizeCustomDistribution interface{} `json:"prize_custom_distribution"`
	PrizeValue              int         `json:"prize_value"`
	SettledAt               interface{} `json:"settled_at"`
	WinnerCount             interface{} `json:"winner_count"`
}

// TG
type InlineKeyboardButton struct {
	Text string `json:"text"`
	URL  string `json:"url"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}
