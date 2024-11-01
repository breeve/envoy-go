1. main thread 里面的 main_thread_guard_dog_ 在做什么？在work里面也有类似的dog
2. 服务生命周期这里的作用是什么？
dispatcher_->post([this] { notifyCallbacksForStage(Stage::Startup); });
3. warming 机制处理
4. createListenerFilterFactoryListImpl



ActiveTcpListener::onAccept
ActiveTcpListener::onAcceptWorker
onSocketAccepted
-> config_->filterChainFactory().createListenerFilterChain(*active_socket)
active_socket->startFilterChain();
{createListenerFilterFactoryListImpl}.onAccept()

void ActiveTcpSocket::continueFilterChain(bool success) {
.1 accept_filters_[x].onAccept(*this); // listener filer
.2 void ActiveStreamListenerBase::newConnection(Network::ConnectionSocketPtr&& socket
ListenerImpl::createNetworkFilterChain
ListenerFilterChainFactoryBuilder::buildFilterChain
buildFilterChainInternal
createNetworkFilterFactoryListImpl



bool ListenerImpl::createNetworkFilterChain(
    Network::Connection& connection, const Filter::NetworkFilterFactoriesList& filter_factories) {
  return Configuration::FilterChainUtility::buildFilterChain(connection, filter_factories);
}

bool ListenerImpl::createListenerFilterChain(Network::ListenerFilterManager& manager) {
  if (Configuration::FilterChainUtility::buildFilterChain(manager, listener_filter_factories_)) {
    return true;
  } else {
    ENVOY_LOG(debug, "New connection accepted while missing configuration. "
                     "Close socket and stop the iteration onAccept.");
    missing_listener_config_stats_.extension_config_missing_.inc();
    return false;
  }
}



void FilterChainManagerImpl::addFilterChains(


