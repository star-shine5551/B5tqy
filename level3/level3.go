package main

type Video struct {
	Author   string
	Selfname string
	Praise   int
	Stars    int
	Coin     int
}

func (v *Video) Praises() {
	v.Praise++
}
func (v *Video) GiveStar() {
	v.Stars++
}
func (v *Video) GiveCoin() {
	v.Coin++
}
func (v *Video) LongPress() {
	v.Praise++
	v.Stars++
	v.Coin++
}
func (v *Video) LongPress2() {
	v.Praises()
	v.GiveCoin()
	v.GiveStar()
}
func newVideo(author string, video string) Video {
	return Video{
		Author:   author,
		Selfname: video,
	}
}
