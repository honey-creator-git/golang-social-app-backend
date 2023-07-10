package SerendipityRequest

import (
	"serendipity_backend/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginEmailRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignUpEmailRequest struct {
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	FirstName   string `json:"firstName" binding:"required"`
	LastName    string `json:"lastName" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
}

type GoogleAuthRequest struct {
	Email      string `json:"email" binding:"required"`
	FirstName  string `json:"firstName" binding:"required"`
	LastName   string `json:"lastName" binding:"required"`
	Avatar     string `json:"avatar" binding:"required"`
	SocialType string `json:"socialType" binding:"required"`
	SocialId   string `json:"socialId" binding:"required"`
}

type UpdateProfileRequest struct {
	Email            string `json:"email"`
	Password         string `json:"password"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	PhoneNumber      string `json:"phoneNumber"`
	PushNotification bool   `json:"pushNotification"`
	Avatar           string `json:"avatar"`
}

type UpdateUserRole struct {
	Role int `json:"role" binding:"required"`
}

type AddToolkitRequest struct {
	Title            string   `json:"title" binding:"required"`
	CoverLetterImage string   `json:"coverLetterImage" binding:"required"`
	SortType         []string `json:"sortType"`
	Type             int      `json:"type" binding:"required"`
}

type AddToolkitPostRequest struct {
	ToolkitType      int            `json:"toolkitType" binding:"required"`
	Title            string         `json:"title" binding:"required"`
	Description      string         `json:"description"`
	CoverLetterImage string         `json:"coverLetterImage"`
	Medias           []models.Media `json:"medias"`
	CookingPeriod    int            `json:"cookingPeriod"`
	Preparation      int            `json:"preparation"`
	Ingredients      []string       `json:"ingredients"`
	Instructions     []string       `json:"instructions"`
	SortTypeId       string         `json:"sortTypeId"`
	Link             string         `json:"link"`
}

type AddForumTypeRequest struct {
	Title            string `json:"title" binding:"required"`
	CoverLetterImage string `json:"coverLetterImage" binding:"required"`
	ForumType        int    `json:"forumType" binding:"required"`
}

type AddForumPostRequest struct {
	Title            string                  `json:"title" binding:"required"`
	CoverLetterImage string                  `json:"coverLetterImage" binding:"required"`
	Description      string                  `json:"description" binding:"required"`
	CreatedBy        primitive.ObjectID      `json:"createdBy" binding:"required"`
	Comments         []models.Comment        `json:"comments"`
	VisitCount       int                     `json:"visitCount"`
	Emotions         models.ForumPostEmotion `json:"emotions"`
	ForumType        int                     `json:"forumType" binding:"required"`
}

type NewCommentForForumPost struct {
	Description string             `json:"description" binding:"required"`
	PostedBy    primitive.ObjectID `json:"postedBy" binding:"required"`
	PostId      primitive.ObjectID `json:"postId" binding:"required"`
	Emotions    models.Emotion     `json:"emotions"`
}

type AddEmotionForForumPostComment struct {
	Emotions models.Emotion `json:"emotions" binding:"required"`
}

type AddNewForumPostCommentRequest struct {
	Description string             `json:"description" binding:"required"`
	PostedBy    primitive.ObjectID `json:"postedBy" binding:"required"`
	PostId      primitive.ObjectID `json:"postId" binding:"required"`
	Emotions    models.Emotion     `json:"emotions"`
}

type GetToolkitPostRequest struct {
	// ToolkitType string `json:"toolkitType" binding:"required"`
	// SortType    string `json:"sortType"`
	// Query       string `json:"query"`
	// Start       int    `json:"start" binding:"required"`
	// Limit       int    `json:"limit" binding:"required"`
}

type GetForumPostRequest struct {
	// SubjectType string `json:"subjectType" binding:"required"`
	// Query       string `json:"query" binding:"required"`
	// Start       int    `json:"start" binding:"required"`
	// Limit       int    `json:"limit" binding:"required"`
}

// type SaveForumPostRequest struct {
// 	PostId      primitive.ObjectID `json:"postId" binding:"required"` // ! Seems Unnecessary
// 	Title       string             `json:"title" binding:"required"`
// 	Description string             `json:"description" binding:"required"`
// 	Emotions    []int              `json:"emotions" binding:"required"`
// 	PostedAt    time.Time          `json:"postedAt" binding:"required"`
// 	PostedBy    primitive.ObjectID `json:"postedBy" binding:"required"`
// 	Medias      []models.Media
// }

type UpdateToolkitRequest struct {
	Title            string   `json:"title"`
	CoverLetterImage string   `json:"coverLetterImage"`
	Description      string   `json:"description"`
	SortType         []string `json:"sortType"`
}

type UpdateToolkitPostRequest struct {
	ToolkitType      int            `json:"toolkitType"`
	Title            string         `json:"title"`
	Description      string         `json:"description"`
	CookingPeriod    int            `json:"cookingPeriod"`
	Preparation      int            `json:"preparation"`
	Ingredients      []string       `json:"ingredients"`
	Instructions     []string       `json:"instructions"`
	Medias           []models.Media `json:"medias"`
	CoverLetterImage string         `json:"coverLetterImage"`
	SortTypeId       string         `json:"sortTypeId"`
	Link             string         `json:"link"`
}

type SetWeeklyPost struct {
	WeeklyPost bool `json:"weeklyPost"`
}

type UpdateForumPostWithPostIDUserID struct {
	Title            string `json:"title"`
	CoverLetterImage string `json:"coverLetterImage"`
	Description      string `json:"description"`
	ForumType        int    `json:"forumType" binding:"required"`
}

type UpdateForumPostWithEmotions struct {
	Emotions models.ForumPostEmotion `json:"emotions" binding:"required"`
}

type UpdateForumPostVisitCount struct {
	VisitCount int `json:"visitCount" binding:"required"`
}

type DeletePostRequest struct {
	PostId primitive.ObjectID `json:"postId" binding:"required"`
}

type FollowPostRequest struct {
	PostId primitive.ObjectID `json:"postId" binding:"required"`
	UserId primitive.ObjectID `json:"userId" binding:"required"` // ! Seems Unnecessary
}

type FollowForumPostwithUserId struct {
	UserID primitive.ObjectID `json:"userId" binding:"required"`
}

type LoginSocialRequest struct {
	SocialType string `json:"socialType" binding:"required"`
	SocialId   string `json:"socialId" binding:"required"`
}

type SignUpSocialRequest struct {
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	FirstName   string `json:"firstName" binding:"required"`
	LastName    string `json:"lastName" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
	SocialType  string `json:"socialType" binding:"required"`
	SocialId    string `json:"socialId" binding:"required"`
	Avatar      string `json:"avatar"`
}

type ForgetPasswordRequest struct {
	Email string `json:"email" binding:"required"`
}

type GetAllToolkitTypes struct {
}

type GetAllForumTypes struct {
}

type GetForumPostWithPostId struct {
}

type DeleteForumPostWithPostId struct {
}

type SetTodayActivity struct {
	ToolkitPostID  primitive.ObjectID `json:"toolkitPostID" binding:"required"`
	ActivityStatus bool               `json:"activityStatus"`
}

type GetTodayToolkitPosts struct {
}

type SetUserFoundation struct {
	Foundation string `json:"foundation" binding:"required"`
}

type SubmitFoundation struct {
	Name string `json:"name" binding:"required"`
}

type UpdateForumType struct {
	Title            string `json:"title"`
	CoverLetterImage string `json:"coverLetterImage"`
	ForumType        int    `json:"forumType"`
}

type AddNewMarketplace struct {
	Title            string `json:"title" binding:"required"`
	CoverLetterImage string `json:"coverLetterImage" binding:"required"`
	Type             int    `json:"type" binding:"required"`
}

type AddNewMarketplaceItem struct {
	Title           string `json:"title" binding:"required"`
	Description     string `json:"description" binding:"required"`
	Logo            string `json:"logo" binding:"required"`
	Link            string `json:"link" binding:"required"`
	MarketplaceType int    `json:"marketplaceType" binding:"required"`
}

type UpdateUserActivation struct {
	Activate bool `json:"activate"`
}

type UpdateMarketplaceRequest struct {
	Title            string `json:"title"`
	CoverLetterImage string `json:"coverLetterImage"`
}

type UpdateMarketplaceItemRequest struct {
	Title       string `json:"title"`
	Logo        string `json:"logo"`
	Link        string `json:"link"`
	Description string `json:"description"`
}

type CheckUserRequest struct {
	Email string `json:"email" binding:"required"`
}

type PasswordRestRequest struct {
	Email string `json:"email" binding:"required"`
}

type ResetWithNewPasswordRequest struct {
	Email       string `json:"email" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

type SetLogoutWithID struct {
	UserID primitive.ObjectID `json:"userId" binding:"required"`
}
