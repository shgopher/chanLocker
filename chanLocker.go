package chanLocker



type Locker struct {
	ch  chan struct{}
}
// 初始状态是已经有一个
func NewLocker()*Locker{
	ch := make(chan struct{},1)
	ch <- struct{}{}
	return &Locker{
		ch: ch,
	}
}
// lock 从缓存是1编程0，从chan中取出来数据
func (l *Locker) Lock()  {
	<- l.ch
}
// unlock 将缓存从0变成1，忘ch中放入数据。
func (l *Locker) Unlock() {
	l.ch <- struct{}{}
}
