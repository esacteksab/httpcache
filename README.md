## Fork

This is a fork of [https://github.com/bored-engineer/httpcache](https://github.com/bored-engineer/httpcache) which is a fork of [https://github.com/gregjones/httpcache](https://github.com/gregjones/httpcache). `gregjones/httpcache` is archived, that's how I went down the path of finding a "maintained" fork. I stumbled upon `bored-engineer/httpcache`. I chose to fork that repo because 1) the source and 2) I use `httpcache` on a project, and perhaps unwarranted out of anxiety, I want to preserve this package as-is for now "just in case". I won't likely be accepting any PR's or doing any further maintenance beyond general updates and what not. This is likely not the fork you're looking for. You more than likely want something from [gregjones/httpcache/forks](https://github.com/gregjones/httpcache/forks) which are almost certainly maintained by someone who knows what they're doing.

## httpcache

Package httpcache provides a http.RoundTripper implementation that works as a mostly [RFC 7234](https://tools.ietf.org/html/rfc7234) compliant cache for http responses.

It is only suitable for use as a 'private' cache (i.e. for a web-browser or an API-client and not for a shared proxy).

This project isn't actively maintained; it works for what I, and seemingly others, want to do with it, and I consider it "done". That said, if you find any issues, please open a Pull Request and I will try to review it. Any changes now that change the public API won't be considered.

## Cache Backends

- The built-in 'memory' cache stores responses in an in-memory map.
- [`github.com/gregjones/httpcache/diskcache`](https://github.com/gregjones/httpcache/tree/master/diskcache) provides a filesystem-backed cache using the [diskv](https://github.com/peterbourgon/diskv) library.
- [`github.com/gregjones/httpcache/memcache`](https://github.com/gregjones/httpcache/tree/master/memcache) provides memcache implementations, for both App Engine and 'normal' memcache servers.
- [`sourcegraph.com/sourcegraph/s3cache`](https://sourcegraph.com/github.com/sourcegraph/s3cache) uses Amazon S3 for storage.
- [`github.com/gregjones/httpcache/leveldbcache`](https://github.com/gregjones/httpcache/tree/master/leveldbcache) provides a filesystem-backed cache using [leveldb](https://github.com/syndtr/goleveldb/leveldb).
- [`github.com/die-net/lrucache`](https://github.com/die-net/lrucache) provides an in-memory cache that will evict least-recently used entries.
- [`github.com/die-net/lrucache/twotier`](https://github.com/die-net/lrucache/tree/master/twotier) allows caches to be combined, for example to use lrucache above with a persistent disk-cache.
- [`github.com/birkelund/boltdbcache`](https://github.com/birkelund/boltdbcache) provides a BoltDB implementation (based on the [bbolt](https://github.com/coreos/bbolt) fork).

If you implement any other backend and wish it to be linked here, please send a PR editing this file.

## License

- [MIT License](LICENSE.txt)
