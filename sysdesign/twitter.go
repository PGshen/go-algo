package sysdesign

import (
	"container/heap"
	"time"
)

type Twitter struct {
	UserMap map[int]User // 用户列表
}

func Constructor() Twitter {
	return Twitter{
		UserMap: map[int]User{},
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	user, ok := this.UserMap[userId]
	if !ok {
		user = NewUser(userId)
		this.UserMap[userId] = user
	}
	user.Post(tweetId)
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	var res []int = []int{}
	user, ok := this.UserMap[userId]
	if !ok {
		return res
	}
	pq := make(TweetPq, 10)
	heap.Init(&pq)
	followees := user.Followees         // 关注列表
	for followeeId := range followees { // 初始化
		followee, ok := this.UserMap[followeeId]
		if ok && followee.Tweets != nil {
			heap.Push(&pq, followee.Tweets)
		}
	}
	for len(pq) > 0 {
		if len(res) > 10 {
			break
		}
		tweet := heap.Pop(&pq).(*Tweet) // 取队头加入到结果集合
		res = append(res, tweet.TweetId)
		if tweet.Next != nil { // 还有后续，添加到
			heap.Push(&pq, tweet.Next)
		}
	}
	return res
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	follower, ok := this.UserMap[followeeId]
	if ok {
		follower.Follow(followeeId)
	}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	follower, ok := this.UserMap[followeeId]
	if ok {
		follower.Unfollow(followeeId)
	}
}

// 用户
type User struct {
	UserId    int              // 用户ID
	Followees map[int]struct{} // 关注列表
	Followers map[int]struct{} // 粉丝列表
	Tweets    *Tweet           // 推文列表，按发布顺序
}

func NewUser(userId int) User {
	user := User{
		UserId:    userId,
		Followees: map[int]struct{}{userId: struct{}{}}, // 初始化，关注自己
		Followers: map[int]struct{}{},
		Tweets:    &Tweet{},
	}
	return user
}

// 关注
func (user *User) Follow(userId int) {
	user.Followees[userId] = struct{}{}
}

// 取消关注
func (user *User) Unfollow(userId int) {
	if userId != user.UserId {
		delete(user.Followees, userId)
	}
}

// 发送推文
func (user *User) Post(tweetId int) {
	tweet := Tweet{
		TweetId:  tweetId,
		PostTime: time.Now(),
	}
	tweet.Next = user.Tweets
	user.Tweets = &tweet
}

// 推文
type Tweet struct {
	TweetId  int       // 推文ID
	PostTime time.Time // 发布时间
	Next     *Tweet    // 链表形式保存
}

// 实现优先级队列
type TweetPq []*Tweet

func (pq TweetPq) Len() int {
	return len(pq)
}

// 发布时间越晚优先级越高
func (pq TweetPq) Less(i, j int) bool {
	return pq[i].PostTime.After(pq[j].PostTime)
}

func (pq TweetPq) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *TweetPq) Push(x interface{}) {
	item := x.(*Tweet)
	*pq = append(*pq, item)
}

func (pq *TweetPq) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	return item
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
