- 使用sync.Once包可以保证线程安全
- 懒汉式思想为时间换空间，没有调用，就不创建实例
- 饿汉式思想为空间换时间，只要包被加载，不管有没有调用，都会创建一个实例。