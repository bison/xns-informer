package informers

import (
	"sync"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
)

type NewInformerFunc func(namespace string) informers.GenericInformer

// multiNamespaceGenericInformer satisfies the GenericInformer interface and
// provides cross-namespace informers and listers.
type multiNamespaceGenericInformer struct {
	informer *multiNamespaceInformer
	lister   cache.GenericLister
}

var _ informers.GenericInformer = &multiNamespaceGenericInformer{}

func (i *multiNamespaceGenericInformer) Informer() cache.SharedIndexInformer {
	return i.informer
}

func (i *multiNamespaceGenericInformer) Lister() cache.GenericLister {
	return i.lister
}

// informerData holds a single namespaced informer.
type informerData struct {
	informer cache.SharedIndexInformer
	lister   cache.GenericLister
	stopCh   chan struct{}
	started  bool
}

// eventHandlerData holds an event handler and its resync period.
type eventHandlerData struct {
	handler      cache.ResourceEventHandler
	resyncPeriod time.Duration
}

// multiNamespaceInformer satisfies the SharedIndexInformer interface and
// provides an informer that works across a set of namespaces -- though not all
// methods are actually usable.
type multiNamespaceInformer struct {
	informers     map[string]*informerData
	eventHandlers []eventHandlerData
	indexers      []cache.Indexers
	resyncPeriod  time.Duration
	namespaced    bool
	lock          sync.Mutex
	newInformer   NewInformerFunc
}

var _ cache.SharedIndexInformer = &multiNamespaceInformer{}

// NewMultiNamespaceInformer returns a new cross-namespace informer.  The given
// NewInformerFunc will be used to craft new single-namespace informers when
// adding namespaces.
func NewMultiNamespaceInformer(namespaced bool, resync time.Duration, newInformer NewInformerFunc) *multiNamespaceInformer {
	informer := &multiNamespaceInformer{
		informers:     make(map[string]*informerData),
		eventHandlers: make([]eventHandlerData, 0),
		indexers:      make([]cache.Indexers, 0),
		namespaced:    namespaced,
		resyncPeriod:  resync,
		newInformer:   newInformer,
	}

	// AddNamespace and RemoveNamespace are no-ops for cluster-scoped
	// informers.  They watch metav1.NamespaceAll only.
	if !namespaced {
		i := newInformer(metav1.NamespaceAll)

		informer.informers[metav1.NamespaceAll] = &informerData{
			informer: i.Informer(),
			lister:   i.Lister(),
			stopCh:   make(chan struct{}),
		}
	}

	return informer
}

func (i *multiNamespaceInformer) GetController() cache.Controller {
	panic("not implemented")
}

func (i *multiNamespaceInformer) GetStore() cache.Store {
	return NewCacheReader(i)
}

func (i *multiNamespaceInformer) GetIndexer() cache.Indexer {
	return NewCacheReader(i)
}

func (i *multiNamespaceInformer) LastSyncResourceVersion() string {
	panic("not implemented")
}

func (i *multiNamespaceInformer) SetWatchErrorHandler(handler cache.WatchErrorHandler) error {
	panic("not implemented") // TODO: This could probably be implemented.
}

// AddNamespace adds the given namespace to the informer.  This is a no-op if an
// informer for this namespace already exists.  You must call one of the run
// functions and wait for the caches to sync before the new informer is useful.
// This is usually done via a factory with Start() and WaitForCacheSync().
func (i *multiNamespaceInformer) AddNamespace(namespace string) {
	i.lock.Lock()
	defer i.lock.Unlock()

	// If an informer for this namespace already exists, or the
	// watched resource is cluster-scoped, this is a no-op.
	if _, ok := i.informers[namespace]; ok || !i.namespaced {
		return
	}

	informer := i.newInformer(namespace)

	i.informers[namespace] = &informerData{
		informer: informer.Informer(),
		lister:   informer.Lister(),
		stopCh:   make(chan struct{}),
	}

	// Add indexers to the new informer.
	for _, idx := range i.indexers {
		i.informers[namespace].informer.AddIndexers(idx)
	}

	// Add event handlers to the new informer.
	for _, handler := range i.eventHandlers {
		i.informers[namespace].informer.AddEventHandlerWithResyncPeriod(
			handler.handler,
			handler.resyncPeriod,
		)
	}
}

// RemoveNamespace stops and deletes the informer for the given namespace.
func (i *multiNamespaceInformer) RemoveNamespace(namespace string) {
	i.lock.Lock()
	defer i.lock.Unlock()

	// If there is no informer for this namespace, or the watched
	// resource is cluster-scoped, this is a no-op.
	if _, ok := i.informers[namespace]; !ok || !i.namespaced {
		return
	}

	close(i.informers[namespace].stopCh)
	delete(i.informers, namespace)
}

// WaitForStop waits for the channel to be closed, then stops all informers.
// TODO: This may be called multiple times, but should only wait once.
func (i *multiNamespaceInformer) WaitForStop(stopCh <-chan struct{}) {
	<-stopCh // Block until stopCh is closed.
	i.lock.Lock()
	defer i.lock.Unlock()

	for _, informer := range i.informers {
		if informer.started {
			close(informer.stopCh)
			informer.started = false
		}
	}
}

// NonBlockingRun starts all stopped informers and waits for the stop channel to
// close before stopping them.  This can be called safely multiple times.
func (i *multiNamespaceInformer) NonBlockingRun(stopCh <-chan struct{}) {
	i.lock.Lock()
	defer i.lock.Unlock()

	for _, informer := range i.informers {
		if !informer.started {
			go informer.informer.Run(informer.stopCh)
			informer.started = true
		}
	}

	go i.WaitForStop(stopCh)
}

// Run starts all stopped informers and waits for the stop channel to close
// before stopping them.  This can be called safely multiple times.  This
// version blocks until the stop channel is closed.
func (i *multiNamespaceInformer) Run(stopCh <-chan struct{}) {
	i.NonBlockingRun(stopCh)
	<-stopCh // Block until stopCh is closed.
}

// AddEventHandler adds the given handler to each namespaced informer.
func (i *multiNamespaceInformer) AddEventHandler(handler cache.ResourceEventHandler) {
	i.AddEventHandlerWithResyncPeriod(handler, i.resyncPeriod)
}

// AddEventHandlerWithResyncPeriod adds the given handler with a resync period
// to each namespaced informer.  The handler will also be added to any informers
// created later as namespaces are added.
func (i *multiNamespaceInformer) AddEventHandlerWithResyncPeriod(handler cache.ResourceEventHandler, resyncPeriod time.Duration) {
	i.lock.Lock()
	defer i.lock.Unlock()

	i.eventHandlers = append(i.eventHandlers, eventHandlerData{
		handler:      handler,
		resyncPeriod: resyncPeriod,
	})

	for _, informer := range i.informers {
		informer.informer.AddEventHandlerWithResyncPeriod(handler, resyncPeriod)
	}
}

// AddIndexers adds the given indexers to each namespaced informer.
func (i *multiNamespaceInformer) AddIndexers(indexers cache.Indexers) error {
	i.lock.Lock()
	defer i.lock.Unlock()

	i.indexers = append(i.indexers, indexers)

	for _, informer := range i.informers {
		err := informer.informer.AddIndexers(indexers)
		if err != nil {
			return err
		}
	}

	return nil
}

// HasSynced checks if each started namespaced informer has synced.
func (i *multiNamespaceInformer) HasSynced() bool {
	i.lock.Lock()
	defer i.lock.Unlock()

	for _, informer := range i.informers {
		if synced := informer.informer.HasSynced(); informer.started && !synced {
			return false
		}
	}

	return true
}

// getListers returns a map of namespaces to their GenericListers.
func (i *multiNamespaceInformer) getListers() map[string]cache.GenericLister {
	i.lock.Lock()
	defer i.lock.Unlock()

	res := make(map[string]cache.GenericLister, len(i.informers))
	for namespace, informer := range i.informers {
		res[namespace] = informer.lister
	}

	return res
}

// getIndexers returns a map of namespaces to their cache.Indexer.
func (i *multiNamespaceInformer) getIndexers() map[string]cache.Indexer {
	i.lock.Lock()
	defer i.lock.Unlock()

	res := make(map[string]cache.Indexer, len(i.informers))
	for namespace, informer := range i.informers {
		res[namespace] = informer.informer.GetIndexer()
	}

	return res
}