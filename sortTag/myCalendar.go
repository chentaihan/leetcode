package sortTag

import "fmt"

type calItem struct {
	start int
	end   int
}

type MyCalendar struct {
	list []calItem
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (this *MyCalendar) Book(start int, end int) bool {
	item := calItem{start, end}
	if len(this.list) == 0 {
		this.list = append(this.list, item)
		return true
	}
	index, exist := binarySearch(this.list, start, end)
	if !exist {
		return false
	}
	if index == -1 {
		this.list = append([]calItem{item}, this.list...)
		return true
	}
	this.list = append(this.list[:index], append([]calItem{item}, this.list[index:]...)...)
	return true
}

//二分查找
func binarySearch(nums []calItem, min, max int) (int, bool) {
	start, middle, end := 0, 0, len(nums)-1
	for start <= end {
		middle = start + (end-start)/2
		if nums[middle].start == min {
			return middle, false
		} else if nums[middle].start > min {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}
	if nums[middle].start > min {
		middle--
	}
	if middle == -1 {
		return middle, true
	}
	if middle == len(nums)-1 {
		return middle, nums[middle].end <= min
	}
	return middle, nums[middle].end <= min && nums[middle+1].start >= max
}

func TestConstructor() {
	tests := []struct {
		start int
		end   int
		ok    bool
	}{
		{
			10, 20, true,
		},
		{
			15, 25, false,
		},
		{
			20, 30, true,
		},
		{
			5, 10, true,
		},
		{
			30, 40, true,
		},
		{
			16, 22, false,
		},
		{
			1, 5, true,
		},
		{
			100, 200, true,
		},
		{
			60, 100, true,
		},

	}
	mc := Constructor()
	for index, test := range tests {
		ok := mc.Book(test.start, test.end)
		if ok != test.ok {
			fmt.Println(index)
		}
		fmt.Println(ok == test.ok)
	}
}
