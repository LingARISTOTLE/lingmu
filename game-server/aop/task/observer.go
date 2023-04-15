package task

type Observer struct {
	subs map[*Subscriber]struct{}
}

func (o *Observer) AddSubscriber() {

}

func (o *Observer) DELSubscriber() {

}

func (o *Observer) Notify(event interface{}) {
	for subscriber, _ := range o.subs {
		subscriber.OnNotify(event)
	}
}
