package messaging

import "log"

type Dispatcher interface {
	GetChannel() (int, chan Item)
	Handle(item Item)
	Release(id int)
}

type ChannelDispatcher struct {
	actives      []bool
	itemChannels []chan Item
}

func NewDispatcher(poolSize int) *ChannelDispatcher {
	d := &ChannelDispatcher{}
	d.actives = make([]bool, poolSize)
	for i := range d.actives {
		d.actives[i] = false
		channel := make(chan Item)
		d.itemChannels = append(d.itemChannels, channel)
	}
	return d
}

func (d *ChannelDispatcher) Handle(item Item) {
	notified := 0
	for id, channel := range d.itemChannels {
		if d.actives[id] {
			channel <- item
			notified++
		}
	}
	log.Printf("notified %v channels", notified)
}

func (d *ChannelDispatcher) GetChannel() (int, chan Item) {
	for i, active := range d.actives {
		if !active {
			d.actives[i] = true
			log.Printf("assigned channel #%v", i)
			return i, d.itemChannels[i]
		}
	}
	return 0, nil
}

func (d *ChannelDispatcher) Release(id int) {
	d.actives[id] = false
	log.Printf("client connection #%v is closed", id)
}
