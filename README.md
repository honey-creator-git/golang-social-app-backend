# Serendipity BackEnd

## SignIn

- SignIn With User Email & Password
```
POST : /api/v1/user/signin
Request Body : 
        {
            "Email": "zhi@gad.ai",
            "Password": "IvanP.9899"
        }
Response :
        {
            "payload": {
                "data": {
                    "id": "62d8c194412b0114ac9e8af0",
                    "email": "zhi@gad.ai",
                    "firstName": "Zhi",
                    "lastName": "Huan",
                    "password": "IvanP.9899",
                    "phoneNumber": "+3534643523",
                    "socialType": "",
                    "socialId": "",
                    "pushNotification": true,
                    "avatar": "",
                    "follows": [
                        "62d96fa6dfae8e256f0cb2f5"
                    ]
                },
                "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTg1NDMxNDMsImlhdCI6MTY1ODU0MTM0MywibmJmIjoxNjU4NTQxMzQzLCJzdWIiOiJ6aGlAZ2FkLmFpIn0.3o2iPSFuL8ze0GDwImoixvZ5AV2DtMXy76ibE0uPhTxSIu7u9w5tu3nRWni-iqBaUrv64D_Ac9_DgXTF-OEm_w"
            },
            "result": true
        }
```

## SignUp

- SignUp With FirstName, LastName, Email, Password & Phone Number
```

POST: /api/v1/user/create
Request Body :
        {
            "Email": "lovricluka644@gmail.com",   
            "Password": "IvanP.9899",
            "FirstName": "Lovric",
            "LastName": "Luka",
            "PhoneNumber": "+235678673"
        }
Response :
        {
            "payload": {
                "data": {
                    "id": "62db5bafbff22618c2a37960",
                    "email": "lovricluka644@gmail.com",
                    "firstName": "Lovric",
                    "lastName": "Luka",
                    "password": "$2a$10$9PeXPO0S6Kh6gLNiHMuKO.7rR5pwhtqCn42N2MCj0ITUhQrfmRz5.",
                    "phoneNumber": "+235678673",
                    "socialType": "",
                    "socialId": "",
                    "pushNotification": false,
                    "avatar": "",
                    "follows": null
                },
                "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTg1NDQ4MjMsImlhdCI6MTY1ODU0MzAyMywibmJmIjoxNjU4NTQzMDIzLCJzdWIiOiJsb3ZyaWNsdWthNjQ0QGdtYWlsLmNvbSJ9.lHr6BV3KrT7_VDbr6vlAtaiYQdx4WNU05jVZHjBXmGC-aHz1NuJq8xt9SwE3BF2REi1h-50LxMwqygpLcFwlJg"
            },
            "status": true
        }
```

## Update User Profile

- Update User Profile with FirstName, LastName, Email, Password and PhoneNumber, Push Notification
```
PUT: /api/v1/user/:userId (i.e : /api/v1/user/62db5bafbff22618c2a37960)
Request Header :
        {
            "Authorization" : "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTg1NDczNDksImlhdCI6MTY1ODU0NTU0OSwibmJmIjoxNjU4NTQ1NTQ5LCJzdWIiOiJsb3ZyaWNsdWthNjQ0QGdtYWlsLmNvbSJ9.ajVnEazXIV__8lF5VtKP4-aKVeo8UaL9UrIJDRWAnlUhzjWuKK8ZBVh5J3k7W-wObuic7tGOPBTSCsTBpK2y3Q"
        }
Request Body :
        {
            "Email": "lovricluka644@gmail.com",   
            "Password": "IvanP.9899",
            "FirstName": "Lovoric",
            "LastName": "Lukas",
            "PhoneNumber": "+235678673",
            "PushNotification": true
        }
Response :
        {
            "result": {
                "payload": {
                    "email": "lovricluka644@gmail.com",
                    "firstname": "Lovoric",
                    "lastname": "Lukas",
                    "password": "IvanP.9899",
                    "phonenumber": "+235678673",
                    "pushnotification": true
                }
            },
            "status": true
        }
```

## Delete User

- Delete User with User ID
```
DELETE: /api/v1/user/:userId (i.e : /api/v1/user/62d8c50501304074c887da0d)
Response : {
    "status": true, 
    "result": "User Successfully deleted."
}
```

## Create New Toolkit

- Create Toolkit Type with Toolkie Title, Description, CoverLetterImage and SortType
```
POST: /api/v1/toolkit/create
Request Body: 
        {
            "Title": "Mediation Archive", 
            "CoverLetterImage": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQhzXfcXpMb4y1zZW3JWwcZEVX1-VpvB04J7KRlI9F2TWdIDkvjb7_hsLXKHiGzptsVRyY&usqp=CAU",
            "type": 5
        }
Response :
        {
            "payload": {
                "id": "62ee005aa9227bc0fc776050",
                "title": "Mediation Archive",
                "coverletterimage": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQhzXfcXpMb4y1zZW3JWwcZEVX1-VpvB04J7KRlI9F2TWdIDkvjb7_hsLXKHiGzptsVRyY&usqp=CAU",
                "sortType": null,
                "type": 5
            },
            "status": true
        }
```

## Create New Toolkit Post

- Create Toolkit Post with Toolkit Type, Post Title, Description, Medias and SortTypeId
```
POST: /api/v1/toolkit/post/create
Request Body:
        {
            "ToolkitType": "62d7cbce43b4f0f3bc45da46",
            "Title": "Traveling over the sea.",
            "Description": "Another way to help improve mood and outlook is through positive thinking. Members will receive daily affirmations and positive thoughts to help them focus on their strength and capabilities.",
            "Medias": [{
                "Url": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQa0uEA3CFNPhMSTvG4i8DL0OSLtiVkp3KXVQ&usqp=CAU",
                "MediaType": "image",
                "Period": 10
            }],
            "SortTypeId": 1
        }
Response :
        {
            "payload": {
                "id": "62dbbdbe013bc419b50b01fe",
                "toolkitType": "62d7cbce43b4f0f3bc45da46",
                "title": "Traveling over the SUN",
                "description": "Another way to help improve mood and outlook is through positive thinking. Members will receive daily affirmations and positive thoughts to help them focus on their strength and capabilities.",
                "medias": [
                    {
                        "url": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQa0uEA3CFNPhMSTvG4i8DL0OSLtiVkp3KXVQ&usqp=CAU",
                        "mediaType": "image",
                        "period": 10
                    }
                ],
                "postedAt": "2022-07-23 17:22:06",
                "sortTypeId": 1,
                "todayActivity": false
            },
            "status": true
        }
```

## Create New Forum

- Create Forum with Forum Title and CoverLetterImage
```
POST: /api/v1/forum/create
Request Body :
       {
            "Title": "How to keep healthy lifestyle forever ???",
            "CoverLetterImage":"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ95YqCNN-CvBoJlwflaXfeqIdMi7xlD-k7KQ&usqp=CAU"
        }
Response :
        {
            "payload": {
                "ID": "62dbc0ac9e20ee5498930f50",
                "title": "How to keep healthy lifestyle forever ???",
                "coverLetterImage": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ95YqCNN-CvBoJlwflaXfeqIdMi7xlD-k7KQ&usqp=CAU"
            },
            "status": true
        }
```

## Create Forum Post

- Create Forum Post with ForumField, Post Title, CoverLetterImage, Description and CreatedBy
```
POST: /api/v1/forum/post/create
Request Header : 
        {
            "Authorization" : "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTg0MTk1MDYsImlhdCI6MTY1ODQxNzcwNiwibmJmIjoxNjU4NDE3NzA2LCJzdWIiOiJ6aGlAZ2FkLmFpIn0.fp7vbKuFHQMxKIbnXdauYW8B1bg-X4_rG06N6eZbXHZLNIuQJMreON4nzgDgc9s-L7S-7MRy3SXpeImCb3He4g"
        }
Request Body :
        {
            "Title": "Keeping ideal mental condition",
            "CoverLetterImage": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTs7SgrB-dS-o113v3WfiQqDWPiNyAaoyuDeg&usqp=CAU",
            "Description": "Another way to help improve mood and outlook is through positive thinking. Members will receive daily affirmations and positive thoughts to help them focus on their strength and capabilities.",
            "CreatedBy": "62d81bcb835c45a59a53c14d",
            "ForumField": "62d819912fd8cf1fdc9d4fed"
        }
Response :
        {
            "payload": {
                "id": "62dbc2a1e9accb1ee07bc3f0",
                "title": "Keeping ideal mental condition",
                "coverLetterImage": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTs7SgrB-dS-o113v3WfiQqDWPiNyAaoyuDeg&usqp=CAU",
                "description": "Another way to help improve mood and outlook is through positive thinking. Members will receive daily affirmations and positive thoughts to help them focus on their strength and capabilities.",
                "createdAt": "2022-07-23 17:42:19",
                "createdBy": "62d81bcb835c45a59a53c14d",
                "comments": null,
                "visitCount": 0,
                "emotions": {},
                "forumField": "62d819912fd8cf1fdc9d4fed"
            },
            "status": true
        }
```

## Add Comment to Forum Post

- Add Comment to Forum Post with Description, PostedBy and PostId
```
POST: /api/v1/forum/post/comment/:postId (i.e : /api/v1/forum/post/comment/62d838b6730924cdf64ee532)
Request Header :
        {
            "Authorization" : "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTg0MTk1MDYsImlhdCI6MTY1ODQxNzcwNiwibmJmIjoxNjU4NDE3NzA2LCJzdWIiOiJ6aGlAZ2FkLmFpIn0.fp7vbKuFHQMxKIbnXdauYW8B1bg-X4_rG06N6eZbXHZLNIuQJMreON4nzgDgc9s-L7S-7MRy3SXpeImCb3He4g"
        }
Request Body :
        {
            "Description": "Fantastic !!!",
            "PostedBy": "62d81bcb835c45a59a53c14d",
            "PostId": "62d838b6730924cdf64ee532",
            "Emotions": {
                "Like": 0,
                "Dislike": 0
            }
        }
Response Body :
        {
            "payload": {
                "result": "successfully added comment to Forum Post.",
                "updated_comments": [
                    {
                        "id": "62dbd20c3022a8a3fc1c03ce",
                        "description": "This is really Good Post !!!!",
                        "postedAt": "2022-07-23 18:48:44",
                        "postedBy": "62d81bcb835c45a59a53c14d",
                        "postId": "62d838b6730924cdf64ee532",
                        "emotions": {
                            "like": 1
                        }
                    },
                    {
                        "id": "62dbd2243022a8a3fc1c03cf",
                        "description": "Fantastic !!!",
                        "postedAt": "2022-07-23 18:49:08",
                        "postedBy": "62d81bcb835c45a59a53c14d",
                        "postId": "62d838b6730924cdf64ee532",
                        "emotions": {}
                    }
                ]
            },
            "status": true
        }
```

## Add Emotion for Forum Post Comment

- Add Emotion for Forum Post Comment with Number of Emotions
```

PUT: /api/v1/forum/post/comment/emotion/:postId/:commentId (i.e : /api/v1/forum/post/comment/emotion/62dbd1ff3022a8a3fc1c03cc/62dbd20c3022a8a3fc1c03ce)
Request Header : 
        {
            "Authorization" : "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTg0MTk1MDYsImlhdCI6MTY1ODQxNzcwNiwibmJmIjoxNjU4NDE3NzA2LCJzdWIiOiJ6aGlAZ2FkLmFpIn0.fp7vbKuFHQMxKIbnXdauYW8B1bg-X4_rG06N6eZbXHZLNIuQJMreON4nzgDgc9s-L7S-7MRy3SXpeImCb3He4g"
        }
Request Body :
        {
            "Emotions": {
                "Like": 8,
                "Dislike": 5
            }
        }
Response :
        {
            "data": {
                "payload": {
                    "id": "62dbd20c3022a8a3fc1c03ce",
                    "description": "This is really Good Post !!!!",
                    "postedAt": "2022-07-23 18:48:44",
                    "postedBy": "62d81bcb835c45a59a53c14d",
                    "postId": "62d838b6730924cdf64ee532",
                    "emotions": {
                        "like": 8,
                        "dislike": 5
                    }
                },
                "result": "successfully updated forum post comment with Emotion."
            },
            "status": true
        }
```

## Add Emotion for Forum Post

- Add Emotion for Forum Post with Number of Emotions
```
PUT: /api/v1/forum/post/emotions/:postId (i.e: /api/v1/forum/post/emotions/62dbd1ff3022a8a3fc1c03cc)
Request Header :
        {
            "Authorization" : "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTg0MTk1MDYsImlhdCI6MTY1ODQxNzcwNiwibmJmIjoxNjU4NDE3NzA2LCJzdWIiOiJ6aGlAZ2FkLmFpIn0.fp7vbKuFHQMxKIbnXdauYW8B1bg-X4_rG06N6eZbXHZLNIuQJMreON4nzgDgc9s-L7S-7MRy3SXpeImCb3He4g"
        }
Request Body :
        {
            "Emotions": {
                "Like": 10,
                "Dislike": 3
            }
        }
Response :
        {
            "payload": {
                "id": "62dbd1ff3022a8a3fc1c03cc",
                "title": "Mental Works Illness",
                "description": "Another way to help improve mood and outlook is through positive thinking. Members will receive daily affirmations and positive thoughts to help them focus on their strength and capabilities. dddd dddd dddd",
                "createdAt": "2022-07-23 18:46:08",
                "createdBy": "62d81bcb835c45a59a53c14d",
                "comments": [
                    {
                        "id": "62dbd20c3022a8a3fc1c03ce",
                        "description": "This is really Good Post !!!!",
                        "postedAt": "2022-07-23 18:48:44",
                        "postedBy": "62d81bcb835c45a59a53c14d",
                        "postId": "62d838b6730924cdf64ee532",
                        "emotions": {
                            "like": 6,
                            "dislike": 3
                        }
                    },
                    {
                        "id": "62dbd2243022a8a3fc1c03cf",
                        "description": "Fantastic !!!",
                        "postedAt": "2022-07-23 18:49:08",
                        "postedBy": "62d81bcb835c45a59a53c14d",
                        "postId": "62d838b6730924cdf64ee532",
                        "emotions": {}
                    }
                ],
                "visitCount": 13,
                "emotions": {
                    "like": 5,
                    "dislike": 4
                },
                "forumField": "62d819912fd8cf1fdc9d4fed"
            },
            "status": true
        }
```

## Update VisitCount for Forum Post

- Add visit count for Forum Post with Number of VisitCount
```
PUT: /api/v1/forum/post/visitcount/:postId (i.e: /api/v1/forum/post/visitcount/62dbd1ff3022a8a3fc1c03cc)
Request Header :
        {
            "Authorization" : "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTg0MTk1MDYsImlhdCI6MTY1ODQxNzcwNiwibmJmIjoxNjU4NDE3NzA2LCJzdWIiOiJ6aGlAZ2FkLmFpIn0.fp7vbKuFHQMxKIbnXdauYW8B1bg-X4_rG06N6eZbXHZLNIuQJMreON4nzgDgc9s-L7S-7MRy3SXpeImCb3He4g"
        }
Request Body :
        {
            "VisitCount": 13
        }
Response :
        {
            "payload": {
                "id": "62dbd1ff3022a8a3fc1c03cc",
                "title": "Mental Works Illness",
                "description": "Another way to help improve mood and outlook is through positive thinking. Members will receive daily affirmations and positive thoughts to help them focus on their strength and capabilities. dddd dddd dddd",
                "createdAt": "2022-07-23 18:46:08",
                "createdBy": "62d81bcb835c45a59a53c14d",
                "comments": [
                    {
                        "id": "62dbd20c3022a8a3fc1c03ce",
                        "description": "This is really Good Post !!!!",
                        "postedAt": "2022-07-23 18:48:44",
                        "postedBy": "62d81bcb835c45a59a53c14d",
                        "postId": "62d838b6730924cdf64ee532",
                        "emotions": {
                            "like": 6,
                            "dislike": 3
                        }
                    },
                    {
                        "id": "62dbd2243022a8a3fc1c03cf",
                        "description": "Fantastic !!!",
                        "postedAt": "2022-07-23 18:49:08",
                        "postedBy": "62d81bcb835c45a59a53c14d",
                        "postId": "62d838b6730924cdf64ee532",
                        "emotions": {}
                    }
                ],
                "visitCount": 13,
                "emotions": {
                    "like": 4,
                    "dislike": 3
                },
                "forumField": "62d819912fd8cf1fdc9d4fed"
            },
            "status": true
        }
```

## Get Toolkit Posts with Toolkit ID

- Get all of the toolkit posts using Toolkit ID
```
GET: /api/v1/toolkit/posts/:toolkitType?results=3&page=1&sortField=id&sortOrder=descend (i.e : /api/v1/toolkit/posts/5?results=3&page=1&sortField=id&sortOrder=descend)
Request Header :
        {
            "Authorization" : "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTg0MTk1MDYsImlhdCI6MTY1ODQxNzcwNiwibmJmIjoxNjU4NDE3NzA2LCJzdWIiOiJ6aGlAZ2FkLmFpIn0.fp7vbKuFHQMxKIbnXdauYW8B1bg-X4_rG06N6eZbXHZLNIuQJMreON4nzgDgc9s-L7S-7MRy3SXpeImCb3He4g"
        }
Response :
        {
            "payload": [
                {
                    "id": "62ee182c93c90a8100a43f63",
                    "toolkitType": 5,
                    "title": "Hope, Strenth and Faith Meditation Foundation - 4",
                    "description": "",
                    "cookingPeriod": 0,
                    "preparation": 0,
                    "ingredients": null,
                    "instructions": null,
                    "coverLetterImage": "",
                    "medias": [
                        {
                            "url": "https://i.kfs.io/artist/global/34573525,0v1/fit/300x300.jpg",
                            "mediaType": "Video",
                            "period": 40
                        }
                    ],
                    "postedAt": "2022-08-06 15:28:44",
                    "sortTypeId": "",
                    "todayActivity": false
                },
                {
                    "id": "62ee181593c90a8100a43f61",
                    "toolkitType": 5,
                    "title": "Hope, Strenth and Faith Meditation Foundation - 3",
                    "description": "",
                    "cookingPeriod": 0,
                    "preparation": 0,
                    "ingredients": null,
                    "instructions": null,
                    "coverLetterImage": "",
                    "medias": [
                        {
                            "url": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR7PB2BVqzNBvwhd5I3jQvXTCxwXYgLRfBo_w&usqp=CAU",
                            "mediaType": "Video",
                            "period": 30
                        }
                    ],
                    "postedAt": "2022-08-06 15:28:21",
                    "sortTypeId": "",
                    "todayActivity": false
                },
                {
                    "id": "62ee17fd93c90a8100a43f5f",
                    "toolkitType": 5,
                    "title": "Hope, Strenth and Faith Meditation Foundation - 2",
                    "description": "",
                    "cookingPeriod": 0,
                    "preparation": 0,
                    "ingredients": null,
                    "instructions": null,
                    "coverLetterImage": "",
                    "medias": [
                        {
                            "url": "https://source.boomplaymusic.com/group10/M00/11/29/2511754c11fe49ffa90a1f6df5d8ee6e_464_464.jpg",
                            "mediaType": "Video",
                            "period": 20
                        }
                    ],
                    "postedAt": "2022-08-06 15:27:57",
                    "sortTypeId": "",
                    "todayActivity": false
                }
            ],
            "status": true
        }
```

## Get Forum Posts with Forum ID

- Get all of the Forum Posts using Forum ID
```
GET: /api/v1/forum/posts/:forumId (i.e : /api/v1/forum/posts/62d819912fd8cf1fdc9d4fed)
Request Header :
        {
            "Authorization" : "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTg0MTk1MDYsImlhdCI6MTY1ODQxNzcwNiwibmJmIjoxNjU4NDE3NzA2LCJzdWIiOiJ6aGlAZ2FkLmFpIn0.fp7vbKuFHQMxKIbnXdauYW8B1bg-X4_rG06N6eZbXHZLNIuQJMreON4nzgDgc9s-L7S-7MRy3SXpeImCb3He4g"
        }
Response :
        {
            "payload": [
                {
                    "id": "62da674d37ec0cc0aec4fff7",
                    "title": "Auto Heart Disease How to Prevent ??? ???",
                    "coverLetterImage": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTs7SgrB-dS-o113v3WfiQqDWPiNyAaoyuDeg&usqp=CAU",
                    "description": "Another way to help improve mood and outlook is through positive thinking. Members will receive daily affirmations and positive thoughts to help them focus on their strength and capabilities.",
                    "createdAt": "2022-07-22 17:00:40",
                    "createdBy": "62d81bcb835c45a59a53c14d",
                    "comments": [
                        {
                            "id": "62da67b637ec0cc0aec4fff9",
                            "description": "This is really fascinating !!! !!! !!! !!! !!!!!!!!!!!!!!! Really !",
                            "postedAt": "2022-07-22 17:02:46",
                            "postedBy": "62d81bcb835c45a59a53c14d",
                            "postId": "62d838b6730924cdf64ee532",
                            "emotions": {
                                "like": 8,
                                "dislike": 5
                            }
                        },
                        {
                            "id": "62da683a70096e2b055d4a2c",
                            "description": "In my case, it's really bad reputation.",
                            "postedAt": "2022-07-22 17:04:58",
                            "postedBy": "62d81bcb835c45a59a53c14d",
                            "postId": "62d838b6730924cdf64ee532",
                            "emotions": {
                                "like": 1
                            }
                        },
                        {
                            "id": "62da685b70096e2b055d4a2d",
                            "description": "This Post is really bad ... I don't like it",
                            "postedAt": "2022-07-22 17:05:31",
                            "postedBy": "62d81bcb835c45a59a53c14d",
                            "postId": "62d838b6730924cdf64ee532",
                            "emotions": {
                                "like": 2,
                                "dislike": 5
                            }
                        }
                    ],
                    "visitCount": 0,
                    "emotions": {
                        "like": 25,
                        "dislike": 17
                    },
                    "forumField": "62d819912fd8cf1fdc9d4fed"
                },
                {
                    "id": "62dbd1ff3022a8a3fc1c03cc",
                    "title": "Keeping ideal mental condition",
                    "coverLetterImage": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTs7SgrB-dS-o113v3WfiQqDWPiNyAaoyuDeg&usqp=CAU",
                    "description": "Another way to help improve mood and outlook is through positive thinking. Members will receive daily affirmations and positive thoughts to help them focus on their strength and capabilities.",
                    "createdAt": "2022-07-23 18:46:08",
                    "createdBy": "62d81bcb835c45a59a53c14d",
                    "comments": [
                        {
                            "id": "62dbd20c3022a8a3fc1c03ce",
                            "description": "This is really Good Post !!!!",
                            "postedAt": "2022-07-23 18:48:44",
                            "postedBy": "62d81bcb835c45a59a53c14d",
                            "postId": "62d838b6730924cdf64ee532",
                            "emotions": {
                                "like": 6,
                                "dislike": 3
                            }
                        },
                        {
                            "id": "62dbd2243022a8a3fc1c03cf",
                            "description": "Fantastic !!!",
                            "postedAt": "2022-07-23 18:49:08",
                            "postedBy": "62d81bcb835c45a59a53c14d",
                            "postId": "62d838b6730924cdf64ee532",
                            "emotions": {}
                        }
                    ],
                    "visitCount": 0,
                    "emotions": {
                        "like": 10,
                        "dislike": 3
                    },
                    "forumField": "62d819912fd8cf1fdc9d4fed"
                }
            ],
            "status": true
        }
```

## Update Toolkit

- Change Toolkit Properties with Toolkit ID
```
PUT: /api/v1/toolkit/:toolkitId (i.e : /toolkit/62d9207edb3e935ad265bc83)
Request Body :
        {
            "description": "This toolkit is for the usage of receipes to keep healthy lifestyle with latest version"
        }
Response :
        {
            "payload": {
                "id": "62d9207edb3e935ad265bc83",
                "title": "Receipe Archive Latest",
                "coverletterimage": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ9XkcY8BlukbxYxQJE5tbz9Nc7Ia_xEyM1vQ&usqp=CAU",
                "description": "This toolkit is for the usage of receipes to keep healthy lifestyle with latest version",
                "sortType": [
                    "Breakfast",
                    "Lunch",
                    "Dinner",
                    "Snack",
                    "Diet"
                ]
            },
            "status": true
        }
```

## Update Toolkit Post

- Change Toolkit Post Properties with Toolkit Post ID
```
PUT: /api/v1/toolkit/posts/:postId (i.e : /api/v1/toolkit/posts/62dbbdbe013bc419b50b01fe)
Request Body :
        {
            "ToolkitType": "62d7cbce43b4f0f3bc45da46",
            "Title": "Traveling over the sea.",
            "Description": "Another way to help improve mood and outlook is through positive thinking. Members will receive daily affirmations and positive thoughts to help them focus on their strength and capabilities.",
            "Medias": [{
                "Url": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQa0uEA3CFNPhMSTvG4i8DL0OSLtiVkp3KXVQ&usqp=CAU",
                "MediaType": "image",
                "Period": 10
            }],
            "SortTypeId": 1
        }
Response :
        {
            "payload": {
                "id": "62dbbdbe013bc419b50b01fe",
                "toolkitType": "62d7cbce43b4f0f3bc45da46",
                "title": "Biking over sea.",
                "description": "Another way to help improve mood and outlook is through positive thinking. Members will receive daily affirmations and positive thoughts to help them focus on their strength and capabilities.",
                "medias": [
                    {
                        "url": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQa0uEA3CFNPhMSTvG4i8DL0OSLtiVkp3KXVQ&usqp=CAU",
                        "mediaType": "image",
                        "period": 10
                    }
                ],
                "postedAt": "2022-07-23 17:22:06",
                "sortTypeId": 0,
                "todayActivity": false
            },
            "status": true
        }
```


- Get All Toolkit Types
```
GET: /api/v1/toolkits?results=2&page=4&sortField=id&sortOrder=ascend
Request Header :
        {
            "Authorization" : "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTg0MTk1MDYsImlhdCI6MTY1ODQxNzcwNiwibmJmIjoxNjU4NDE3NzA2LCJzdWIiOiJ6aGlAZ2FkLmFpIn0.fp7vbKuFHQMxKIbnXdauYW8B1bg-X4_rG06N6eZbXHZLNIuQJMreON4nzgDgc9s-L7S-7MRy3SXpeImCb3He4g"
        }
Response :
        {
            "payload": [
                {
                    "id": "62dba3cca10d33b86f1990a9",
                    "title": "Most comfortable way to keep lifecycle",
                    "coverletterimage": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ9XkcY8BlukbxYxQJE5tbz9Nc7Ia_xEyM1vQ&usqp=CAU",
                    "description": "This toolkit is for the usage of receipes to keep healthy lifestyle",
                    "sortType": null
                },
                {
                    "id": "62dba3ffa10d33b86f1990ad",
                    "title": "Working on mountain",
                    "coverletterimage": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ9XkcY8BlukbxYxQJE5tbz9Nc7Ia_xEyM1vQ&usqp=CAU",
                    "description": "This toolkit is for the usage of receipes to keep healthy lifestyle",
                    "sortType": [
                        "Seat",
                        "Location",
                        "Device"
                    ]
                }
            ],
            "status": true
        }
```

## Get Forum Types

- Get All Forum Types
```
GET: /api/v1/forums
Request Header :
        {
            "Authorization" : "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTg0MTk1MDYsImlhdCI6MTY1ODQxNzcwNiwibmJmIjoxNjU4NDE3NzA2LCJzdWIiOiJ6aGlAZ2FkLmFpIn0.fp7vbKuFHQMxKIbnXdauYW8B1bg-X4_rG06N6eZbXHZLNIuQJMreON4nzgDgc9s-L7S-7MRy3SXpeImCb3He4g"
        }
Response :
        {
            "payload": [
                {
                    "ID": "62d811cd89be62a8b442dd45",
                    "title": "Lupus",
                    "coverLetterImage": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRWgD7bE0GvOSC8rw4aW8bxe2eGq1-DnIuUhg&usqp=CAU"
                },
                {
                    "ID": "62d819912fd8cf1fdc9d4fed",
                    "title": "Women With Auto Immune Diseases",
                    "coverLetterImage": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ95YqCNN-CvBoJlwflaXfeqIdMi7xlD-k7KQ&usqp=CAU"
                },
                {
                    "ID": "62d978c8dfae8e256f0cb2ff",
                    "title": "Women With Auto Heart Diseases",
                    "coverLetterImage": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ95YqCNN-CvBoJlwflaXfeqIdMi7xlD-k7KQ&usqp=CAU"
                },
                {
                    "ID": "62dbc0ac9e20ee5498930f50",
                    "title": "How to keep healthy lifestyle forever ???",
                    "coverLetterImage": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ95YqCNN-CvBoJlwflaXfeqIdMi7xlD-k7KQ&usqp=CAU"
                },
                {
                    "ID": "62dcdb4a18d9d0a8a1678586",
                    "title": "How to earn money ???",
                    "coverLetterImage": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ95YqCNN-CvBoJlwflaXfeqIdMi7xlD-k7KQ&usqp=CAU"
                }
            ],
            "status": true
        }
```

## Get Forum Post with Post ID

- Get Forum Post with Post ID
```
GET: /api/v1/forum/post/:postId (i.e : /api/v1/forum/post/62dbd1ff3022a8a3fc1c03cc)
Request Header :
        {
            "Authorization" : "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTg0MTk1MDYsImlhdCI6MTY1ODQxNzcwNiwibmJmIjoxNjU4NDE3NzA2LCJzdWIiOiJ6aGlAZ2FkLmFpIn0.fp7vbKuFHQMxKIbnXdauYW8B1bg-X4_rG06N6eZbXHZLNIuQJMreON4nzgDgc9s-L7S-7MRy3SXpeImCb3He4g"
        }
Response :
        {
            "payload": {
                "id": "62dbd1ff3022a8a3fc1c03cc",
                "title": "Keeping ideal mental condition",
                "coverLetterImage": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTs7SgrB-dS-o113v3WfiQqDWPiNyAaoyuDeg&usqp=CAU",
                "description": "Another way to help improve mood and outlook is through positive thinking. Members will receive daily affirmations and positive thoughts to help them focus on their strength and capabilities.",
                "createdAt": "2022-07-23 18:46:08",
                "createdBy": "62d81bcb835c45a59a53c14d",
                "comments": [
                    {
                        "id": "62dbd20c3022a8a3fc1c03ce",
                        "description": "This is really Good Post !!!!",
                        "postedAt": "2022-07-23 18:48:44",
                        "postedBy": "62d81bcb835c45a59a53c14d",
                        "postId": "62d838b6730924cdf64ee532",
                        "emotions": {
                            "like": 6,
                            "dislike": 3
                        }
                    },
                    {
                        "id": "62dbd2243022a8a3fc1c03cf",
                        "description": "Fantastic !!!",
                        "postedAt": "2022-07-23 18:49:08",
                        "postedBy": "62d81bcb835c45a59a53c14d",
                        "postId": "62d838b6730924cdf64ee532",
                        "emotions": {}
                    }
                ],
                "visitCount": 0,
                "emotions": {
                    "like": 10,
                    "dislike": 3
                },
                "forumField": "62d819912fd8cf1fdc9d4fed"
            },
            "status": true
        }
```

## Update Forum Post with Post ID

- Update Forum Post Fields with Post ID
```
PUT: /api/v1/forum/post/:postId/:userId (i.e : api/v1/forum/post/62d838b6730924cdf64ee532)
Request Header :
        {
            "Authorization" : "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTg0MTk1MDYsImlhdCI6MTY1ODQxNzcwNiwibmJmIjoxNjU4NDE3NzA2LCJzdWIiOiJ6aGlAZ2FkLmFpIn0.fp7vbKuFHQMxKIbnXdauYW8B1bg-X4_rG06N6eZbXHZLNIuQJMreON4nzgDgc9s-L7S-7MRy3SXpeImCb3He4g"
        }
Request Body:
        {
            "description": "Another way to help improve mood and outlook is through positive thinking. Members will receive daily affirmations and positive thoughts to help them focus on their strength and capabilities. dddd dddd"
        }
Response :
        {
            "payload": {
                "id": "62dbd1ff3022a8a3fc1c03cc",
                "description": "Another way to help improve mood and outlook is through positive thinking. Members will receive daily affirmations and positive thoughts to help them focus on their strength and capabilities. dddd dddd",
                "createdAt": "2022-07-23 18:46:08",
                "createdBy": "62d81bcb835c45a59a53c14d",
                "comments": [
                    {
                        "id": "62dbd20c3022a8a3fc1c03ce",
                        "description": "This is really Good Post !!!!",
                        "postedAt": "2022-07-23 18:48:44",
                        "postedBy": "62d81bcb835c45a59a53c14d",
                        "postId": "62d838b6730924cdf64ee532",
                        "emotions": {
                            "like": 6,
                            "dislike": 3
                        }
                    },
                    {
                        "id": "62dbd2243022a8a3fc1c03cf",
                        "description": "Fantastic !!!",
                        "postedAt": "2022-07-23 18:49:08",
                        "postedBy": "62d81bcb835c45a59a53c14d",
                        "postId": "62d838b6730924cdf64ee532",
                        "emotions": {}
                    }
                ],
                "visitCount": 0,
                "emotions": {
                    "like": 10,
                    "dislike": 3
                },
                "forumField": "62d819912fd8cf1fdc9d4fed"
            },
            "status": true
        }
```

## Forum Post Delete With Post ID

- Forum Post Delete with Post ID as param
```
DELETE: /api/v1/forum/post/:postId (i.e : /aforum/post/62da3919754f1d0db99339e9)
Response :
        {
            "payload": "Forum Post is deleted successfully.",
            "status": true
        }
```

## Today's Activities for Toolkit post with Toolkit Post IDs

- Set Today's Activities with Toolkit Post Ids
```
POST: /api/v1/toolkit/posts/today_activities
Request Body :
        {
            "TodayActivities": ["62d9fb7e30319fa563f9bf92", "62d9a69b5b06a08cb8a3c7cb", "62d97848dfae8e256f0cb2fd"]
        }
Response :
        {
            "payload": "Successfully Set up for today's activities",
            "status": true
        }
```

## Today's Activities Get For Toolkit Posts

- Get Today's Activities
```
GET: /api/v1/toolkit/posts/today
Request Header :
        {
            "Authorization" : "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTg0MTk1MDYsImlhdCI6MTY1ODQxNzcwNiwibmJmIjoxNjU4NDE3NzA2LCJzdWIiOiJ6aGlAZ2FkLmFpIn0.fp7vbKuFHQMxKIbnXdauYW8B1bg-X4_rG06N6eZbXHZLNIuQJMreON4nzgDgc9s-L7S-7MRy3SXpeImCb3He4g"
        }
Response :
        {
            "payload": [
                {
                    "id": "62d97848dfae8e256f0cb2fd",
                    "toolkitType": "62d7cbce43b4f0f3bc45da46",
                    "title": "Movement Cycle",
                    "description": "Another way to help improve mood and outlook is through positive thinking. Members will receive daily affirmations and positive thoughts to help them focus on their strength and capabilities.",
                    "medias": [
                        {
                            "url": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQa0uEA3CFNPhMSTvG4i8DL0OSLtiVkp3KXVQ&usqp=CAU",
                            "mediaType": "image",
                            "period": 10
                        }
                    ],
                    "postedAt": "2022-07-22 00:01:12",
                    "sortTypeId": 0,
                    "todayActivity": true
                },
                {
                    "id": "62d9a69b5b06a08cb8a3c7cb",
                    "toolkitType": "62d7cbce43b4f0f3bc45da46",
                    "title": "Mountain Biking Overview",
                    "description": "Another way to help improve mood and outlook is through positive thinking. Members will receive daily affirmations and positive thoughts to help them focus on their strength and capabilities.",
                    "medias": [
                        {
                            "url": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQa0uEA3CFNPhMSTvG4i8DL0OSLtiVkp3KXVQ&usqp=CAU",
                            "mediaType": "image",
                            "period": 10
                        }
                    ],
                    "postedAt": "2022-07-22 03:18:50",
                    "sortTypeId": 0,
                    "todayActivity": true
                },
                {
                    "id": "62d9fb7e30319fa563f9bf92",
                    "toolkitType": "62d7cbce43b4f0f3bc45da46",
                    "title": "Mountain Biking Trains",
                    "description": "Another way to help improve mood and outlook is through positive thinking. Members will receive daily affirmations and positive thoughts to help them focus on their strength and capabilities.",
                    "medias": [
                        {
                            "url": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQa0uEA3CFNPhMSTvG4i8DL0OSLtiVkp3KXVQ&usqp=CAU",
                            "mediaType": "image",
                            "period": 10
                        }
                    ],
                    "postedAt": "2022-07-22 09:21:02",
                    "sortTypeId": 0,
                    "todayActivity": true
                }
            ],
            "status": true
        }
```

## Follow User

- Follow User
```
POST: /api/v1/user/follow/:userId (i.e : /user/follow/62d8c194412b0114ac9e8af0)
Request Header :
        {
            "Authorization" : "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTg0MTk1MDYsImlhdCI6MTY1ODQxNzcwNiwibmJmIjoxNjU4NDE3NzA2LCJzdWIiOiJ6aGlAZ2FkLmFpIn0.fp7vbKuFHQMxKIbnXdauYW8B1bg-X4_rG06N6eZbXHZLNIuQJMreON4nzgDgc9s-L7S-7MRy3SXpeImCb3He4g"
        }
Request Body :
        {
            "UserID": "62d96fa6dfae8e256f0cb2f5"
        }
Response :
        {
            "payload": {
                "id": "62d8c194412b0114ac9e8af0",
                "email": "zhi@gad.ai",
                "firstName": "Zhi",
                "lastName": "Huan",
                "password": "IvanP.9899",
                "phoneNumber": "+3534643523",
                "socialType": "",
                "socialId": "",
                "pushNotification": true,
                "avatar": "",
                "follows": [
                    "62d96fa6dfae8e256f0cb2f5"
                ]
            },
            "status": true
        }
```

## Delete Toolkit Type and Toolkit Posts involved in the Toolkit Type

- Delete Toolkit Type with Toolkit ID
```
DELETE: /api/v1/toolkit/:toolkitId (i.e: /api/v1/toolkit/62d9207edb3e935ad265bc83)
Response :
    {
        "result": "Toolkit Type Successfully Deleted.",
        "status": true
    }
```

## Delete Toolkit Post

- Delete Toolkit Post with Toolkit Post ID
```
DELETE: /api/v1/toolkit/posts/:postId (i.e: /api/v1/toolkit/posts/62dff3e24b450a734b3f27fb)
Response : 
        {
            "result": "Toolkit Post successfully deleted.",
            "status": true
        }
```

## Delete Forum Type and Forum Posts involved in the Forum Type

- Delete Forum Type With Forum ID
```
DELETE: /api/v1/forum/:forumId (i.e: /api/v1/forum/62d819912fd8cf1fdc9d4fed)
Response :
        {
            "result": "Forum Type Successfully Deleted.",
            "status": true
        }
```

## Create New Foundation

- Create a new Founction with Name
```
POST: /api/v1/user/updateFoundation/:userId (i.e: /api/v1/user/updateFoundation/62d8c194412b0114ac9e8af0)
Request Body :
    {
        "Name": "Crohn's And Colitis Foundation"
    }
Response :
    {
        "new_foundation": {
            "id": "62ebde0c5b8d413fa9ae0db1",
            "name": "Crohn's And Colitis Foundation"
        },
        "status": true
    }
```

## Set Foundation for User

- Set User's foundation with User ID
```
POST: /api/v1/user/updateFoundation/:userId (i.e: /api/v1/user/updateFoundation/62d8c194412b0114ac9e8af0)
Request Header : 
    {
        "Authorization" : "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTg1NDczNDksImlhdCI6MTY1ODU0NTU0OSwibmJmIjoxNjU4NTQ1NTQ5LCJzdWIiOiJsb3ZyaWNsdWthNjQ0QGdtYWlsLmNvbSJ9.ajVnEazXIV__8lF5VtKP4-aKVeo8UaL9UrIJDRWAnlUhzjWuKK8ZBVh5J3k7W-wObuic7tGOPBTSCsTBpK2y3Q"
    }

Request Body : 
    {
        "Foundation": "62ebde0c5b8d413fa9ae0db1"
    }

Response : 
    {
        "status": true,
        "userFoundation": {
            "id": "62ebde0c5b8d413fa9ae0db1",
            "name": "Crohn's And Colitis Foundation"
        }
    }
```

## Forum Type Update

- Update Forum Type With Forum ID
```
PUT: /api/v1/forum/:forumId (i.e : /api/v1/62e13fdd4050c2cd309615fa)
Request Body : 
    {
        "CoverLetterImage": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ95YqCNN-CvBoJlwflaXfeqIdMi7xlD-k7KQ&usqp=CAU",
        "Title": "Weekly Salary Limit"
    }
Response :
    {
        "payload": {
            "ID": "62e13fdd4050c2cd309615fa",
            "title": "Weekly Salary Limit",
            "coverLetterImage": "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQ95YqCNN-CvBoJlwflaXfeqIdMi7xlD-k7KQ&usqp=CAU"
        },
        "status": true
    }
```

## Get All Users

- Get All Users with results/page/sortField/sortOrder

```
GET: /api/v1/users?results=3&page=1&sortField=email&sortOrder=ascend
Response : 
        {
            "payload": [
                {
                    "id": "62db5bafbff22618c2a37960",
                    "email": "lovricluka644@gmail.com",
                    "firstName": "Lovoric",
                    "lastName": "Lukas",
                    "password": "IvanP.9899",
                    "phoneNumber": "+235678673",
                    "socialType": "",
                    "socialId": "",
                    "pushNotification": true,
                    "avatar": "",
                    "follows": null,
                    "foundation": "Lupus Foundation"
                },
                {
                    "id": "62d96fa6dfae8e256f0cb2f5",
                    "email": "michaelogboo@gmail.com",
                    "firstName": "Michale",
                    "lastName": "Logboo",
                    "password": "$2a$10$96p9en0OFgxBStVx116weOkTc8tslc/a/t4.F1RKgxt1sHAFAa92O",
                    "phoneNumber": "+3252352342",
                    "socialType": "",
                    "socialId": "",
                    "pushNotification": false,
                    "avatar": "",
                    "follows": null,
                    "foundation": "MS Society"
                },
                {
                    "id": "62ed2ae69d3c4b96000cc736",
                    "email": "ten1@gad.ai",
                    "firstName": "Ten",
                    "lastName": "Daamir",
                    "password": "$2a$10$tLnd/FbgWuY4Dmxs6nI0mOU9pO2NSjHdLYes2RjXj4hgk8jN6oAgq",
                    "phoneNumber": "+123456789",
                    "socialType": "",
                    "socialId": "",
                    "pushNotification": false,
                    "avatar": "",
                    "follows": null,
                    "foundation": "Lupus Foundation"
                }
            ],
            "status": true
        }
```

## Get All Foundations

- Get all of Foundations

```
GET: /api/v1/allFoundations
Response : 
        {
            "payload": [
                {
                    "id": "62ebde0c5b8d413fa9ae0db1",
                    "name": "Crohn's And Colitis Foundation"
                },
                {
                    "id": "62ed71e6b2385104ca07cc1a",
                    "name": "Arthritis Foundation"
                },
                {
                    "id": "62ed71f4b2385104ca07cc1c",
                    "name": "MS Society"
                },
                {
                    "id": "62ed71fcb2385104ca07cc1e",
                    "name": "Lupus Foundation"
                }
            ],
            "status": true
        }
```



## Update User Role

- Update User Role

```
PUT: /admin/api/v1/user/role/:userId (i.e: /admin/api/v1/user/role/62ef0b64ea476001be2d0bbd)
Request Header :
        {
            Authorization: "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTk4NDA3NDQsImlhdCI6MTY1OTgzNzE0NCwibmJmIjoxNjU5ODM3MTQ0LCJzdWIiOiJJdmFuUGV0cm92aWNoQGdtYWlsLmNvbSJ9.pJhi9NcjNdg-rpil79PLaaTWFXY446nLMX0t1DNXiKIrFS03-0FOghyCifr35bKgGg7RjemVUn56Tyt7w31RNg"
        }
Request Body :
        {
            "Role": 3
        }
Response :
        {
            "payload": {
                "id": "62ef0b64ea476001be2d0bbd",
                "email": "ten@gad.ai",
                "firstName": "Ten",
                "lastName": "Damir",
                "password": "$2a$10$AGNB0.6BzmIwcfUdppE3LOpRWa0PUGt8d/TLzOUgBn5c1BVDnbVJi",
                "phoneNumber": "+235678673",
                "socialType": "",
                "socialId": "",
                "pushNotification": false,
                "avatar": "",
                "follows": null,
                "foundation": "",
                "role": 3
            },
            "status": true
        }
```

## Test using Postman

BaseURL : http://localhost:6060/

- Serendipity SignUp with Email
{{BaseURL}}api/v1/user/create -> POST

- Serendipity SignIn User
{{BaseURL}}api/v1/user/signin -> POST

- Serendipity Update User
{{BaseURL}}api/v1/user/62db5bafbff22618c2a37960 -> PUT

- Serendipity User Delete
{{BaseURL}}api/v1/user/62d81bcb835c45a59a53c14d -> DELETE

- Serendipity Toolkit Create
{{BaseURL}}api/v1/toolkit/create -> POST

- Serendipity Toolkit Post Create
{{BaseURL}}api/v1/toolkit/post/create -> POST

- Serendipity Forum Create
{{BaseURL}}api/v1/forum/create -> POST

- Serendipity Forum Post Create
{{BaseURL}}api/v1/forum/post/create -> POST

- Serendipity Forum Post Update with New Comment
{{BaseURL}}api/v1/forum/post/comment/62d838b6730924cdf64ee532 -> POST

- Serendipity Forum Post Comment Update with Emotion
{{BaseURL}}api/v1/forum/post/comment/emotion/62d838b6730924cdf64ee532/62d8544fa2debd25d3e042d7 -> PUT

- Serendipity Forum Post Update with Emotion
{{BaseURL}}api/v1/forum/post/emotions/62d838b6730924cdf64ee532 -> PUT

- Serendipity Toolkit Post Get wtih Toolkit ID
{{BaseURL}}api/v1/toolkit/posts/5?results=3&page=1&sortField=id&sortOrder=descend -> GET

- Serendipity Forum Post Get With Forum ID
{{BaseURL}}api/v1/forum/post/62d819912fd8cf1fdc9d4fed/posts -> GET

- Serendipity Toolkit Type Update with Toolkit ID
{{BaseURL}}api/v1/toolkit/62d9207edb3e935ad265bc83 -> PUT

- Serendipity Toolkit Post Update Toolkit Post ID
{{BaseURL}}api/v1/toolkit/posts/62d8bd088f3524c7f0d08420 -> PUT

- Serendipity Get All Toolkit Types
{{BaseURL}}api/v1/toolkits -> GET

- Serendipity Forum Types Get
{{BaseURL}}api/v1/forums -> GET

- Serendipity Forum Post Get With Post Id
{{BaseURL}}api/v1/forum/posts/62d838b6730924cdf64ee532 -> GET

- Serendipity Forum Post Update with Post ID
{{BaseURL}}api/v1/forum/post/62d838b6730924cdf64ee532/62d81bcb835c45a59a53c14d -> PUT

- Serendipity Forum Post Delete with Post ID
{{BaseURL}}api/v1/forum/post/62da3919754f1d0db99339e9 -> DELETE

- Serendipity Set Today's Activities for Toolkit Posts with IDs
{{BaseURL}}api/v1/toolkit/posts/today_activities -> POST

- Serendipity Get Today's Activities for Toolkit Posts
{{BaseURL}}api/v1/toolkit/posts/today -> GET

- Serendipity DELETE Toolkit Type and Toolkit Posts in the Toolkit Type.
{{BaseURL}}api/v1/toolkit/62d9207edb3e935ad265bc83 -> DELETE

- Serendipity DELETE Toolkit Post with Toolkit Post ID
{{BaseURL}}api/v1/toolkit/posts/:postId -> DELETE

- Serendipity DELETE Forum Type and Forum Posts in the Forum Type.
{{BaseURL}}api/v1/forum/62d819912fd8cf1fdc9d4fed -> DELETE

- Serendipity Foundation Add New
{{BaseURL}}api/v1/foundation/create -> POST

- Serendipity User Set Foundation
{{BaseURL}}api/v1/user/updateFoundation/62d8c194412b0114ac9e8af0 -> POST

- Serendipity Forum Type Update
{{BaseURL}}api/v1/forum/62e13fdd4050c2cd309615fa -> PUT

- Serendipity Get All Users
{{BaseURL}}api/v1/users?results=20&page=1&sortField=email&sortOrder=ascend -> GET

- Serendipity Get All Foundations
{{BaseURL}}api/v1/allFoundations -> GET