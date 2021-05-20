/*
 * @lc app=leetcode.cn id=355 lang=golang
 *
 * [355] 设计推特
 *
 * https://leetcode-cn.com/problems/design-twitter/description/
 *
 * algorithms
 * Medium (40.63%)
 * Likes:    215
 * Dislikes: 0
 * Total Accepted:    21.9K
 * Total Submissions: 54K
 * Testcase Example:  '["Twitter","postTweet","getNewsFeed","follow","postTweet","getNewsFeed","unfollow","getNewsFeed"]\n' +
  '[[],[1,5],[1],[1,2],[2,6],[1],[1,2],[1]]'
 *
 *
 * 设计一个简化版的推特(Twitter)，可以让用户实现发送推文，关注/取消关注其他用户，能够看见关注人（包括自己）的最近十条推文。你的设计需要支持以下的几个功能：
 *
 *
 * postTweet(userId, tweetId): 创建一条新的推文
 * getNewsFeed(userId):
 * 检索最近的十条推文。每个推文都必须是由此用户关注的人或者是用户自己发出的。推文必须按照时间顺序由最近的开始排序。
 * follow(followerId, followeeId): 关注一个用户
 * unfollow(followerId, followeeId): 取消关注一个用户
 *
 *
 * 示例:
 *
 *
 * Twitter twitter = new Twitter();
 *
 * // 用户1发送了一条新推文 (用户id = 1, 推文id = 5).
 * twitter.postTweet(1, 5);
 *
 * // 用户1的获取推文应当返回一个列表，其中包含一个id为5的推文.
 * twitter.getNewsFeed(1);
 *
 * // 用户1关注了用户2.
 * twitter.follow(1, 2);
 *
 * // 用户2发送了一个新推文 (推文id = 6).
 * twitter.postTweet(2, 6);
 *
 * // 用户1的获取推文应当返回一个列表，其中包含两个推文，id分别为 -> [6, 5].
 * // 推文id6应当在推文id5之前，因为它是在5之后发送的.
 * twitter.getNewsFeed(1);
 *
 * // 用户1取消关注了用户2.
 * twitter.unfollow(1, 2);
 *
 * // 用户1的获取推文应当返回一个列表，其中包含一个id为5的推文.
 * // 因为用户1已经不再关注用户2.
 * twitter.getNewsFeed(1);
 *
 *
*/

// @lc code=start
// max heap for tweet
type MaxHeap []*

func (m MaxHeap) Len() int           { return len(m) }
func (m MaxHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m MaxHeap) Less(i, j int) bool { return m[i].time > m[j].time }
func (m *MaxHeap) Push(x interface{}) {
	*m = append(*m, x.(*Tweet))
}
func (m *MaxHeap) Pop() interface{} {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

var maxSize int = 10
var time int = 0

type Tweet struct {
	id   int
	time int
}
type Twitter struct {
	followee map[int]map[int]bool
	tweets   map[int]*list.List
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{map[int]map[int]bool{}, map[int]*list.List{}}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {

	if this.tweets[userId] == nil {
		this.tweets[userId] = list.New()
	}
	this.tweets[userId].PushFront(&Tweet{tweetId, time})

	// keep list max size at maxSize
	if this.tweets[userId].Len() > maxSize {
		this.tweets[userId].Remove(this.tweets[userId].Back())
	}
	// incrase the time
	time++
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {

	h := &MaxHeap{}
	heap.Init(h)
	if this.tweets[userId] != nil {
		for e := this.tweets[userId].Front(); e != nil; e = e.Next() {
			heap.Push(h, e.Value)
		}
	}

	for k := range this.followee[userId] {
		if this.tweets[k] != nil {
			for e := this.tweets[k].Front(); e != nil; e = e.Next() {
				heap.Push(h, e.Value)
			}
		}
	}
	count := maxSize
	if maxSize > h.Len() {
		count = h.Len()
	}
	ans := make([]int, 0, count)
	for count > 0 {
		ans = append(ans, heap.Pop(h).(*Tweet).id)
		count--
	}
	return ans
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {

	if this.followee[followerId] == nil {
		this.followee[followerId] = map[int]bool{}
	}
	this.followee[followerId][followeeId] = true
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if this.followee[followerId] != nil {
		delete(this.followee[followerId], followeeId)
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
// @lc code=end

