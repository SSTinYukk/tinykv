package engine_util

/*
引擎是用于本地存储键/值对的低级系统（无分发或任何事务支持，
等等）。此包包含与此类引擎交互的代码。
CF表示“柱族”。中详细介绍了柱族https://github.com/facebook/rocksdb/wiki/Column-Families
（专门针对RocksDB，但一般概念是通用的）。简而言之，列族是键命名空间。
多列族通常实现为几乎独立的数据库。重要的是，每个柱族都可以
单独配置。写操作可以跨列族进行原子化，这对于单独的数据库是无法做到的。
engine_util包括以下软件包：
*引擎：用于保存unistore所需引擎的数据结构。
*write_batch：批处理代码写入单个原子“事务”。
*cf_迭代器：在badger中迭代整个列族的代码。
*/