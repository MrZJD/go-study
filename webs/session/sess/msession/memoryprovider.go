package msession

import (
	"container/list"
	"sync"
	"time"
)

// Provider session 底层存储结构 @interface
type Provider interface {
	SessionInit(sid string) (Session, error) // session初始化
	SessionRead(sid string) (Session, error) // 返回sid所代表的session 不存在则init
	SessionDestory(sid string) error
	SessionGC(maxLifeTime int64)
}

// Session session结构 @interface
type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}

// provider实现

// SessionStore @struct
type SessionStore struct {
	sid          string
	timeAccessed time.Time
	value        map[interface{}]interface{}
}

// Set SessionStore method@public @implement(Session)
func (st *SessionStore) Set(key, value interface{}) error {
	st.value[key] = value
	pder.SessionUpdate(st.sid)
	return nil
}

// Get SessionStore method@public @implement(Session)
func (st *SessionStore) Get(key interface{}) interface{} {
	pder.SessionUpdate(st.sid)
	if v, ok := st.value[key]; ok {
		return v
	} else {
		return nil
	}
}

// Delete SessionStore method@public @implement(Session)
func (st *SessionStore) Delete(key interface{}) error {
	delete(st.value, key)
	pder.SessionUpdate(st.sid)
	return nil
}

// SessionID SessionStore method@public @implement(Session)
func (st *SessionStore) SessionID() string {
	return st.sid
}

// MemoryProvider @struct @implement(Provider)
type MemoryProvider struct {
	lock     sync.Mutex
	sessions map[string]*list.Element
	list     *list.List
}

// SessionInit MemoryProvider method@public @implement(Provider)
func (pder *MemoryProvider) SessionInit(sid string) (Session, error) {
	pder.lock.Lock()
	defer pder.lock.Unlock()

	v := make(map[interface{}]interface{}, 0)
	newsess := &SessionStore{
		sid:          sid,
		timeAccessed: time.Now(),
		value:        v,
	}

	element := pder.list.PushBack(newsess)
	pder.sessions[sid] = element
	return newsess, nil
}

// SessionRead MemoryProvider method@public @implement(Provider)
func (pder *MemoryProvider) SessionRead(sid string) (Session, error) {
	if element, ok := pder.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	}
	sess, err := pder.SessionInit(sid)
	return sess, err
}

// SessionDestory MemoryProvider method@public @implement(Provider)
func (pder *MemoryProvider) SessionDestory(sid string) error {
	if element, ok := pder.sessions[sid]; ok {
		delete(pder.sessions, sid)
		pder.list.Remove(element)
		return nil
	}
	return nil
}

// SessionGC MemoryProvider method@public @implement(Provider)
func (pder *MemoryProvider) SessionGC(maxLifeTime int64) {
	pder.lock.Lock()
	defer pder.lock.Unlock()

	for {
		element := pder.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*SessionStore).timeAccessed.Unix() + maxLifeTime) < time.Now().Unix() {
			pder.list.Remove(element)
			delete(pder.sessions, element.Value.(*SessionStore).sid)
		} else {
			break
		}
	}
}

// SessionUpdate MemoryProvider method@public
func (pder *MemoryProvider) SessionUpdate(sid string) error {
	pder.lock.Lock()
	defer pder.lock.Unlock()

	if element, ok := pder.sessions[sid]; ok {
		element.Value.(*SessionStore).timeAccessed = time.Now()
		pder.list.MoveToFront(element)
		return nil
	}
	return nil
}

var pder = &MemoryProvider{
	list: list.New(),
}

func init() {
	pder.sessions = make(map[string]*list.Element, 0)

	Register("memory", pder)
}
