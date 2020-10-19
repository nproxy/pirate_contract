package logService

import (
	"context"
	"errors"
	"fmt"
	"github.com/kprc/nbsnetwork/tools"
	"log"
	"sync"
	"time"
)

/*
ls:=GetLogService()
go ls.Daemon()
ls.Start()
ls.Stop()
*/

type LogServiceItem struct {
	watch       func(doBatch *(chan *LogServiceItem)) error
	batch       func() error
	recover     func() error
	stop        chan struct{}
	name        string
	ootb        tools.OnlyOneThread
	ootw        tools.OnlyOneThread
	CurBlockNum func(cb uint64)
	evPos       EventPos
}

type LogService struct {
	watchErrChan chan string
	batchErrChan chan string
	logService   map[string]*LogServiceItem
	doWatchChan  chan *LogServiceItem
	doBatchChan  chan *LogServiceItem
	stop         chan struct{}
	tc           *time.Ticker
}

var logServiceMem *LogService
var logServiceOnce sync.Once

func NewLogService() *LogService {
	ls := &LogService{}

	ls.watchErrChan = make(chan string, 1024)
	ls.batchErrChan = make(chan string, 1024)
	ls.logService = make(map[string]*LogServiceItem)
	ls.doWatchChan = make(chan *LogServiceItem, 1024)
	ls.doBatchChan = make(chan *LogServiceItem, 1024)

	ls.stop = make(chan struct{})

	return ls
}

func GetLogService() *LogService {
	if logServiceMem == nil {
		logServiceOnce.Do(func() {
			logServiceMem = NewLogService()
		})
	}
	return logServiceMem
}

func (ls *LogService) RegLogSrv(lsi *LogServiceItem) error {
	ls.logService[lsi.name] = lsi
	return nil
}

func (ls *LogService) Daemon() error {
	for {
		select {
		case w := <-ls.watchErrChan:
			if w == "" {
				log.Println("watch chan error")
				return errors.New("watch chan error")
			}

			v, ok := ls.logService[w]
			if !ok {
				return errors.New("watch service name not correct " + w)
			}
			ls.doWatchChan <- v
		case b := <-ls.batchErrChan:
			if b == "" {
				log.Println("batch chan error")
				return errors.New("batch chan error")
			}

			v, ok := ls.logService[b]
			if !ok {
				return errors.New("batch service name not correct " + b)
			}
			ls.doBatchChan <- v
		case wlsi := <-ls.doWatchChan:
			if wlsi == nil {
				return errors.New("do watch chan closed")
			}

			fmt.Println("start watch ", wlsi.name)

			go func() {
				if !wlsi.ootw.Start() {
					return
				}
				defer wlsi.ootw.Release()
				err := wlsi.watch(&ls.doBatchChan)
				if err != nil {
					ls.watchErrChan <- wlsi.name
				}

			}()

		case blsi := <-ls.doBatchChan:
			if blsi == nil {
				return errors.New("do batch chan closed")
			}

			fmt.Println("start watch ", blsi.name)

			go func() {
				if !blsi.ootb.Start() {
					return
				}
				defer blsi.ootb.Release()
				err := blsi.batch()
				if err != nil {
					ls.batchErrChan <- blsi.name
				}

			}()

		case <-ls.stop:
			return nil
		}
	}
}

func (ls *LogService) Start() {
	fmt.Println("Log service starting...")

	for _, v := range ls.logService {
		if v.recover != nil {
			v.recover()
		}
	}

	for _, v := range ls.logService {
		ls.doWatchChan <- v
	}

	for _, v := range ls.logService {
		ls.doBatchChan <- v
	}

	ls.tc = time.NewTicker(time.Second * 17)

	go func() {
		defer ls.tc.Stop()
		mc,err := GetLogConf().NewMarketClient()
		if err!=nil{
			panic(err)
		}
		defer func() {
			mc.Close()
		}()
		for {
			select {
			case <-ls.stop:
				return
			case <-ls.tc.C:
				n, err := mc.GetClient().BlockNumber(context.TODO())
				if err != nil {
					mc.Close()
					mc,err = GetLogConf().NewMarketClient()
					if err!= nil{
						panic("can't create eth client")
					}
				}
				for _, v := range ls.logService {
					if v.CurBlockNum != nil {
						v.CurBlockNum(n)
					}
					ls.doBatchChan <- v
				}
			}
		}
	}()

}

func (ls *LogService) Stop() {
	fmt.Println("begin to stop log service...")
	close(ls.stop)
}
