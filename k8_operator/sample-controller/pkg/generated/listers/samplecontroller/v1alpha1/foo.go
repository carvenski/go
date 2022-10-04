package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "k8s.io/sample-controller/pkg/apis/samplecontroller/v1alpha1"
)

/////////////////////////////////////////////
/////// 重点使用cache/indexer来List资源 ///////
/////////////////////////////////////////////

// FooLister实现
// FooLister helps list Foos.
type FooLister interface {
	// List lists all Foos in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Foo, err error)
	// Foos returns an object that can list and get Foos.
	Foos(namespace string) FooNamespaceLister
	FooListerExpansion
}

// fooLister implements the FooLister interface.
type fooLister struct {
	indexer cache.Indexer
}

// NewFooLister returns a new FooLister.
func NewFooLister(indexer cache.Indexer) FooLister {
	return &fooLister{indexer: indexer}
}

// 从indexer里面ListAll,筛选出Foo类型,依靠标签选择器labels.Selector
// List lists all Foos in the indexer.
func (s *fooLister) List(selector labels.Selector) (ret []*v1alpha1.Foo, err error) {
	// cache.ListAll
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Foo))
	})
	return ret, err
}

// Foos returns an object that can list and get Foos.
func (s *fooLister) Foos(namespace string) FooNamespaceLister {
	return fooNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FooNamespaceLister helps list and get Foos.
type FooNamespaceLister interface {
	// List lists all Foos in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Foo, err error)
	// Get retrieves the Foo from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Foo, error)
	FooNamespaceListerExpansion
}

// fooNamespaceLister implements the FooNamespaceLister interface.
type fooNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// 还是从cache/indexer里面去ListAll,增加了namespace的筛选
// List lists all Foos in the indexer for a given namespace.
func (s fooNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Foo, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Foo))
	})
	return ret, err
}

// 还是从indexer里面来查询Foo, indexer.GetByKey(s.namespace + "/" + name)
// Get retrieves the Foo from the indexer for a given namespace and name.
func (s fooNamespaceLister) Get(name string) (*v1alpha1.Foo, error) {	
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("foo"), name)
	}
	return obj.(*v1alpha1.Foo), nil
}
