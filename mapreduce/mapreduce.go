package mapreduce

type KeyValue struct    {
    Key string
    Value string
}

type Mapper interface {
    Map(key, value string) []KeyValue
}

type Reducer interface {
    Reduce(key string, values []string) string
}
