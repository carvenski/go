//评测题目: Map
// design and implement a map structure with the following functions.
// please also write unit test.
// Example:
// m.Put('a', 'aa')
// m.Put('b', 'bb')
// m.Get('a') = 'aa'
// m.PutAll('xx')
// m.Get('a') = 'xx'
// m.Put('b', 'bbb')
// m.Get('b') = 'bbb'               

type MyMap struct{
    has_putAll bool
    putAll_value string
    exceptions map[string]string
    gmap map[string]string
}

func (m *MyMap) Put(key, value string) {
  // TODO: implement the put function to set a specfic key to a specific value.
  // Requirement: need to implement in O(1) time complexity.
  if !m.has_putAll {
        m.gmap[key] = value
  }else {
        m.exceptions[key] = value
  }
}

func (m *MyMap) Get(key string) (string, bool) {
  // TODO: return the specific map value for the given key. if the key does not exit, return empty with false.
  // Requirement: need to implement in O(1) time complexity.
  if !m.has_putAll {
        v, b := m.gmap[key]
        return v, b
  }else {
        v, b := m.exceptions[key]
        if b {
            return v, b
        }else {
            return putAll_value, true
        }
  } 
}

func (m *MyMap) PutAll(value string) {
  // TODO: set all the existing keys in the map to the given value.
  // Requirement: need to implement in O(1) time complexity.
    m.has_putAll = true
    m.putAll_value = value
    m.exceptions = make(map[string]string)
}

