package SerendipityResponse

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Media struct {
	Url       string `json:"url,omitempty"`
	MediaType string `json:"mediaType,omitempty"`
	Period    int    `json:"period,omitempty"`
}

type Post struct {
	ID            primitive.ObjectID `json:"id,omitempty"`
	ToolkitType   primitive.ObjectID `json:"toolkitType,omitempty"`
	Title         string             `json:"title,omitempty"`
	Description   string             `json:"description,omitempty"`
	Medias        []Media            `json:"medias,omitempty"`
	PostedAt      string             `json:"postedAt,omitempty"`
	SortTypeId    int                `json:"sortTypeId"`
	TodayActivity bool               `json:"todayActivity"`
}

type Comment struct {
	ID          primitive.ObjectID `json:"id,omitempty"`
	Description string             `json:"description,omitempty"`
	PostedAt    string             `json:"postedAt,omitempty"`
	PostedBy    primitive.ObjectID `json:"postedBy,omitempty"`
	PostId      primitive.ObjectID `json:"postId,omitEmpty"`
	Emotions    []int              `json:"emotions,omitempty"`
}

type ToolkitType struct {
	ID               primitive.ObjectID `json:"id,omitempty"`
	Title            string             `json:"title,omitempty"`            // -> Nutrition Resources, Foundation Resources, Recipe Resources, Movement Archive, Meditation Archive
	CoverLetterImage string             `json:"coverletterimage,omitempty"` //-> image url
	Description      string             `json:"description,omitempty"`
	SortType         []string           `json:"sortType"`
}

type ForumType struct {
	ID               primitive.ObjectID `jsdon:"title,omitempty"`
	Title            string             `json:"title,omitempty"`
	CoverLetterImage string             `json:"coverLetterImage,omitempty"`
}

type ForumPost struct {
	ID               primitive.ObjectID `json:"id,omitempty"`
	Title            string             `json:"title,omitempty"`
	CoverLetterImage string             `json:"coverLetterImage,omitempty"`
	Description      string             `json:"description,omitemtpy"`
	CreatedAt        string             `json:"createdAt"`
	CreatedBy        primitive.ObjectID `json:"createdBy,omitempty"`
	Comments         []Comment          `json:"comments"`
	VisitCount       int                `json:"visitCount"`
	Emotions         int                `json:"emotions"`
	ForumField       primitive.ObjectID `json:"forumField,omitempty"`
}

type UserSignUpResponse struct {
	ID               primitive.ObjectID   `json:"id,omitempty"`
	Email            string               `json:"email,omitempty"`
	FirstName        string               `json:"firstName,omitempty"`
	LastName         string               `json:"lastName,omitempty"`
	Password         string               `json:"password,omitempty"`
	SocialType       string               `json:"socialType"`
	SocialId         string               `json:"socialId"`
	PhoneNumber      string               `json:"phoneNumber,omitempty"`
	PushNotification bool                 `json:"pushNotification"`
	Avatar           string               `json:"avatar"`
	Follows          []primitive.ObjectID `json:"follows"`
	AccessToken      string               `json:"accessToken,omitempty"`
}

type UserLoginResponse struct {
	ID               primitive.ObjectID   `json:"id,omitempty"`
	Email            string               `json:"email,omitempty"`
	FirstName        string               `json:"firstName,omitempty"`
	LastName         string               `json:"lastName,omitempty"`
	Password         string               `json:"password,omitempty"`
	SocialType       string               `json:"socialType"`
	SocialId         string               `json:"socialId"`
	PhoneNumber      string               `json:"phoneNumber,omitempty"`
	PushNotification bool                 `json:"pushNotification"`
	Avatar           string               `json:"avatar"`
	Follows          []primitive.ObjectID `json:"follows"`
	AccessToken      string               `json:"accessToken,omitempty"`
}

type UserUpdateResponse struct {
	ID               primitive.ObjectID   `json:"id,omitempty"`
	Email            string               `json:"email,omitempty"`
	FirstName        string               `json:"firstName,omitempty"`
	LastName         string               `json:"lastName,omitempty"`
	Password         string               `json:"password,omitempty"`
	SocialType       string               `json:"socialType"`
	SocialId         string               `json:"socialId"`
	PhoneNumber      string               `json:"phoneNumber,omitempty"`
	PushNotification bool                 `json:"pushNotification"`
	Avatar           string               `json:"avatar"`
	Follows          []primitive.ObjectID `json:"follows"`
}

type AddNewToolkitResponse struct {
	ID               primitive.ObjectID `json:"id,omitempty"`
	Title            string             `json:"title,omitempty"`
	CoverLetterImage string             `json:"coverLetterImage,omitempty"`
	Description      string             `json:"description,omitempty"`
	SortType         []string           `json:"sortType"`
}

type AddNewToolkitPostResponse struct {
	ID          primitive.ObjectID `json:"id,omitempty"`
	ToolkitType primitive.ObjectID `json:"toolkittype,omitempty"`
	Medias      []Media            `json:"medias,omitempty"`
	PostedAt    string             `json:"postedAt,omitempty"`
	SortTypeId  int                `json:"sortTypeId,omitempty"`
}

type AddNewForumTypeResponse struct {
	ID               primitive.ObjectID `json:"id,omitempty"`
	Title            string             `json:"title,omitempty"`
	CoverLetterImage string             `json:"coverletterimage,omitempty"`
}

type AddNewForumPostResponse struct {
	ID               primitive.ObjectID `json:"id,omitempty"`
	Title            string             `json:"title,omitempty"`
	CoverLetterImage string             `json:"coverletterimage,omitempty"`
	Description      string             `json:"description,omitempty"`
	CreatedAt        string             `json:"createdAt,omitempty"`
	CreatedBy        primitive.ObjectID `json:"createdBy,omitempty"`
	Comments         []Comment          `json:"comment"`
	VisitCount       int                `json:"visitCount"`
	Emotions         int                `json:"emotions"`
	ForumField       primitive.ObjectID `json:"forumField,omitempty"`
}

type AddNewForumPostCommentResponse struct {
	ID               primitive.ObjectID `json:"id,omitempty"`
	Title            string             `json:"title,omitempty"`
	CoverLetterImage string             `json:"coverLetterImage,omitempty"`
	Description      string             `json:"description,omitempty"`
	CreatedAt        string             `json:"createdAt,omitempty"`
	CreatedBy        string             `json:"createdBy,omitempty"`
	Comments         []Comment          `json:"comments"`
	VisitCount       int                `json:"visitCount"`
	Emotions         int                `json:"emotions"`
	ForumField       string             `json:"forumField,omitempty"`
}

// ! Internal
type PostResponse struct {
	Title            string `json:"title,omitempty"`
	CoverLetterImage string `json:"coverLetterImage,omitempty"`
	TotalLength      int    `json:"totalLength,omitemtpy"`
	Description      string `json:"description,omitempty"`
}

// ! Internal

type GetToolkitPostResponse struct {
	// Posts       []PostResponse `json:"posts,omitempty"`
	// DisplayType int            `json:"displayType,omitempty"`

	ToolkitPosts []Post `json:"toolkistPosts,omitempty"`
}

type GetForumPostResponse struct {
	// 	ForumPosts  []PostResponseOfForum `json:"forumPosts,omitempty"`
	// 	DisplayType int                   `json:"displayType,omitEmpty"`

	ForumPosts []ForumPost `json:"forumPosts,omitempty"`
}

type CommentsOfPostOfForum struct {
	CommentId   primitive.ObjectID `json:"commentId,omitempty"`
	Description string             `json:"description,omitempty"`
	PostedAt    time.Time          `json:"postedAt,omitempty"`
	PostedBy    primitive.ObjectID `json:"postedBy,omitempty"`
	PostId      primitive.ObjectID `json:"postId,omitempty"`
	Emotions    []int              `json:"emotions,omitempty"`
}

// ! Internal
type PostResponseOfForum struct {
	Title            string                  `json:"title,omitempty"`
	CoverLetterImage string                  `json:"coverLetterImage,omitempty"`
	VisitCount       int                     `json:"visitCount,omitempty"`
	Emotions         int                     `json:"emotions,omitempty"`
	Comments         []CommentsOfPostOfForum `json:"comments,omitempty"`
}

type SaveForumPostResponse struct {
	ForumPost PostResponseOfForum `json:"forumPost,omitempty"`
}

type AddToolkitResponse struct {
	ID               primitive.ObjectID `json:"id,omitempty"`
	Title            string             `json:"title,omitempty"`
	Description      string             `json:"description,omitempty"`
	CoverLetterImage string             `json:"coverLetterImage,omitempty"`
	SortType         []string           `json:"sortType,omitempty"`
}

type AddToolkitPostResponse struct {
	ToolkitPostId primitive.ObjectID `json:"toolkitPostId,omitempty"`
	Title         string             `json:"title,omitempty"`
	Description   string             `json:"description,omitempty"`
	ToolkitType   primitive.ObjectID `json:"toolkitType,omitempty"`
	Medias        []Media            `json:"medias,omitempty"`
	SortTypeId    int                `json:"sortTypeId"`
	PostedAt      time.Time          `json:"postedAt,omitempty"`
}

type AddForumPostResponse struct {
	ForumPostId      primitive.ObjectID `json:"forumPostId,omitempty"`
	Title            string             `json:"title,omitempty"`
	CoverLetterImage string             `json:"coverLetterImage,omitempty"`
	Description      string             `json:"description,omitempty"`
	CreatedAt        string             `json:"createdAt,omitempty"`
	CreatedBy        string             `json:"createdBy,omitempty"`
	Comments         []Comment          `json:"comments,omitempty"`
	VisitCount       int                `json:"visitCount,omitempty"`
	Emotions         int                `json:"emotions,omitempty"`
	ForumField       primitive.ObjectID `json:"forumField,omitempty"`
}

type AddForumTypeResponse struct {
	ForumId          primitive.ObjectID `json:"forumId,omitempty"`
	Title            string             `json:"title,omitempty"`
	CoverLetterImage string             `json:"coverLetterImage,omitempty"`
}

type AddForumPostCommentResponse struct {
	CommentId primitive.ObjectID `json:"commentId,omitempty"`
	Comment   Comment            `json:"comment,omitempty"`
}

type AddEmotionForForumPostCommentResponse struct {
	ID               primitive.ObjectID `json:"id,omitempty"`
	Title            string             `json:"title,omitempty"`
	CoverLetterImage string             `json:"coverLetterImage,omitempty"`
	Description      string             `json:"description,omitempty"`
	CreatedAt        string             `json:"createdAt,omitempty"`
	CreatedBy        primitive.ObjectID `json:"createdBy,omitempty"`
	Comments         []Comment          `json:"comments,omitempty"`
	VisitCount       int                `json:"visitCount,omitempty"`
	Emotions         int                `json:"emotions,omitempty"`
	ForumField       primitive.ObjectID `json:"forumField,omitempty"`
}

type AddEmotionForForumPostResponse struct {
	ID               primitive.ObjectID `json:"id,omitempty"`
	Title            string             `json:"title,omitempty"`
	CoverLetterImage string             `json:"coverLetterImage,omitempty"`
	Description      string             `json:"description,omitempty"`
	CreatedAt        string             `json:"createdAt,omitempty"`
	CreatedBy        primitive.ObjectID `json:"createdBy,omitempty"`
	Comments         []Comment          `json:"comments"`
	VisitCount       int                `json:"visitCount"`
	Emotions         int                `json:"emotions,omitempty"`
	ForumField       primitive.ObjectID `json:"forumField,omitempty"`
}

type GetToolkitPostWithToolkitIdResponse struct {
	Posts []Post `json:"posts,omitempty"`
}

type UpdateToolkitResponse struct {
	ID               primitive.ObjectID `json:"id"`
	Title            string             `json:"title"`
	CoverLetterImage string             `json:"coverLetterImage"`
	Description      string             `json:"description"`
	SortType         []string           `json:"sortType"`
}

type GetAllToolkitTypesResponse struct {
	ToolkitTypes []ToolkitType `json:"toolkitTypes,omitempty"`
}

type GetAllForumTypesResponse struct {
	ForumType []ForumType `json:"forumType,omitempty"`
}

type GetForumPostWithPostId struct {
	ForumPost ForumPost `json:"forumPost,omitempty"`
}

type UpdateForumPostWithPostIdResponse struct {
	ForumPost ForumPost `json:"forumPost,omitempty"`
}

type DeleteForumPostWithPostIdResponse struct {
}

type GetTodayToolkitPostsResponse struct {
	TodaysActivities []Post `json:"todayActivity"`
}
