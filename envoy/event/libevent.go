package event

type LibeventScheduler struct {
}

func (s *LibeventScheduler) Run(rType RunType) {
	/*
	  int flag = 0;
	  switch (mode) {
	  case Dispatcher::RunType::NonBlock:
	    flag = LibeventScheduler::flagsBasedOnEventType();
	  case Dispatcher::RunType::Block:
	    // The default flags have 'block' behavior. See
	    // http://www.wangafu.net/~nickm/libevent-book/Ref3_eventloop.html
	    break;
	  case Dispatcher::RunType::RunUntilExit:
	    flag = EVLOOP_NO_EXIT_ON_EMPTY;
	    break;
	  }
	  event_base_loop(libevent_.get(), flag);
	*/
}
