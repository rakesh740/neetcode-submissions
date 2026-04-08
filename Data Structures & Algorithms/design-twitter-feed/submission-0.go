type Twitter struct {
    followMap map[int]map[int]bool
	tweetMap map[int][]tweet
	time int
}

type tweet struct {
	tweetId int
	time int
}



func Constructor() Twitter {
    return Twitter{
		followMap: make(map[int]map[int]bool),
		tweetMap: make(map[int][]tweet),
		time: 0,
	}
}


func (this *Twitter) PostTweet(userId int, tweetId int)  {
    this.tweetMap[userId] = append(this.tweetMap[userId], tweet{tweetId, this.time} )
	this.time++
}


func (this *Twitter) GetNewsFeed(userId int) []int {
    // get all followers 
	// get all posts by each followers 
	// sort by time
	// return top 10 

	followeeMap := this.followMap[userId]
	var posts []tweet

	fmt.Println(" followers for ", userId, " followeeMap ", followeeMap)

	if followeeMap == nil {
		followeeMap = make(map[int]bool)
	}
	followeeMap[userId] = true

	for k := range followeeMap {
		posts = append(posts, this.tweetMap[k]...)
	}

	sort.Slice(posts, func(i, j int) bool {
		return posts[i].time > posts[j].time
	})

	fmt.Println(posts)

	var tweets []int
	 n := 10

	if len(posts) <= 10 {
		n = len(posts)
	}

	for i:= 0;i<n;i++ {
		tweets = append(tweets, posts[i].tweetId)
	}

	return tweets
}


func (this *Twitter) Follow(followerId int, followeeId int)  {
	// if followeeId == followerId {
	// 	return
	// }
	if this.followMap[followerId] == nil {
		this.followMap[followerId] = make(map[int]bool)
	}
    this.followMap[followerId][followeeId] = true
}


func (this *Twitter) Unfollow(followerId int, followeeId int)  {
	if _, ok := this.followMap[followerId]; ok {
		delete(this.followMap[followerId], followeeId)
	}
}
