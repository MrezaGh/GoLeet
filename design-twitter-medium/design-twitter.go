package main

import (
	"container/heap"
	"fmt"
	"slices"
	"time"
)

type MinHeap []tweet

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].time.Compare(h[j].time) <= 0 }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(tweet))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type tweet struct {
	id        int
	time      time.Time
	timestamp int
}

func newTweet(id int) tweet {
	return tweet{id: id,
		time: time.Now(),
	}
}

type User struct {
	posts     []tweet
	following []int
}

type Twitter struct {
	users map[int]*User
}

func Constructor() Twitter {
	return Twitter{users: make(map[int]*User)}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	t := newTweet(tweetId)
	user := this.getUser(userId)

	user.posts = append(user.posts, t)
	if len(user.posts) > 10 {
		user.posts = user.posts[1:]
	}

}

func (this *Twitter) GetNewsFeed(userId int) []int {
	minHeap := MinHeap{}
	heap.Init(&minHeap)
	user := this.getUser(userId)

	for _, t := range user.posts {
		heap.Push(&minHeap, t)
	}

	for _, followee := range user.following {
		for _, t := range this.getUser(followee).posts {
			if minHeap.Len() < 10 {
				heap.Push(&minHeap, t)
			} else {
				head := minHeap[0]
				if head.time.Compare(t.time) < 0 {
					heap.Pop(&minHeap)
					heap.Push(&minHeap, t)
				} else {
					break
				}
			}
		}
	}

	newsFeed := make([]int, 0)
	for minHeap.Len() > 0 {
		newsFeed = append(newsFeed, heap.Pop(&minHeap).(tweet).id)
	}
	slices.Reverse(newsFeed)
	return newsFeed
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	user := this.getUser(followerId)
	for _, id := range user.following {
		if id == followeeId {
			return
		}
	}

	user.following = append(user.following, followeeId)
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	for i := 0; i < len(this.getUser(followerId).following); i++ {
		if this.users[followerId].following[i] == followeeId {
			this.users[followerId].following = append(this.users[followerId].following[:i], this.users[followerId].following[i+1:]...)
			return
		}
	}
}

func (this *Twitter) getUser(userId int) *User {
	user, exist := this.users[userId]
	if !exist {
		user = &User{
			posts:     []tweet{},
			following: []int{},
		}
	}
	this.users[userId] = user
	return user
}

/**
* Your Twitter object will be instantiated and called as such:
* obj := Constructor();
* obj.PostTweet(userId,tweetId);
* param_2 := obj.GetNewsFeed(userId);
* obj.Follow(followerId,followeeId);
* obj.Unfollow(followerId,followeeId);
 */

func main() {
	twitter := Constructor()
	twitter.PostTweet(1, 5)             // User 1 posts a new tweet (id = 5).
	fmt.Println(twitter.GetNewsFeed(1)) // User 1's news feed should return a list with 1 tweet id -> [5]. return [5]
	twitter.Follow(1, 2)                // User 1 follows user 2.
	twitter.PostTweet(2, 6)             // User 2 posts a new tweet (id = 6).
	fmt.Println(twitter.GetNewsFeed(1)) // User 1's news feed should return a list with 2 tweet ids -> [6, 5]. Tweet id 6 should precede tweet id 5 because it is posted after tweet id 5.
	twitter.Unfollow(1, 2)              // User 1 unfollows user 2.
	fmt.Println(twitter.GetNewsFeed(1)) // User 1's news feed should return a list with 1 tweet
}
